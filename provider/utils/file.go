package utils

import (
	"archive/zip"
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type ContentType interface{}

func WriteToFile(
	folderPath string,
	fileName string,
	content ContentType,
	createPathIfNeeded bool,
) error {
	fullPath := filepath.Join(folderPath, fileName)

	if createPathIfNeeded {
		dirPath := filepath.Dir(fullPath)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			err := os.MkdirAll(dirPath, 0755)
			
			if err != nil {
				return err
			}
		}
	} 

	var data []byte
	switch v := content.(type) {
	case string:
		data = []byte(v)
	case *bytes.Buffer:
		data = v.Bytes()
	default:
		data = []byte(content.(string))
	}

	return os.WriteFile(fullPath, data, 0644)
}

func CreateTemporaryFolder(name *string, shouldDeleteContents *bool) (string, error) {
	tmpDir := os.TempDir()
	// create a folder name variable with the pid at the end 
	// to avoid conflicts with other temporary folders
	folderName := fmt.Sprintf("genezio-%d", os.Getpid())
	tmpParentFolder := filepath.Join(tmpDir, folderName)
	// if the folder doesn't exist, create it
	if _, err := os.Stat(tmpParentFolder); os.IsNotExist(err) {
		err := os.MkdirAll(tmpParentFolder, 0755)
		
		if err != nil {
			return "", err
		}
	}

	// generate a random name of 6 characters
	if name == nil {
		randomName := make([]byte, 6)
		_, err := rand.Read(randomName)
		
		if err != nil {
			return "", err
		}
		tmpName := fmt.Sprintf("%x", randomName)
		name = &tmpName
	}
	tmpFolder := filepath.Join(tmpParentFolder, *name)

	if _, err := os.Stat(tmpFolder); os.IsNotExist(err) {
		if shouldDeleteContents != nil && *shouldDeleteContents{
			err := os.RemoveAll(tmpFolder)
			if err != nil {
				return "", err
			}
		} else {
			return tmpFolder, nil
		}
	}

	err := os.Mkdir(tmpFolder, 0755)
	if err != nil {
		return "", err
	}

	return tmpFolder, nil
} 

func CopyFile(source string, dest string) error {
	sourceFile, err := os.Open(source)
	if err != nil{
		return err
	}
	defer sourceFile.Close()
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}
	return destFile.Sync()}

func CopyFolder(source string, dest string) error {
	srcInfo, err := os.Stat(source)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dest, srcInfo.Mode()); err != nil {
		return err
	}
	directory, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	fmt.Println("Copying folder 2")
	for _, entry := range directory {
		srcFilePath := filepath.Join(source, entry.Name())
		destFilePath := filepath.Join(dest, entry.Name())
		if entry.IsDir() {
			if err := CopyFolder(srcFilePath, destFilePath); err != nil {
				fmt.Printf("Error copying folder %s", err)
				return err
			}
		} else {
			if err := CopyFile(srcFilePath, destFilePath); err != nil {
				fmt.Printf("Error copying file %s", err)
				return err
			}
		}
	}
	return nil
}

func CopyFileOrFolder(source string, dest string) error {
	srcInfo, err := os.Stat(source)
	if err != nil {
		return err
	}
	if srcInfo.IsDir() {
		fmt.Println("Copying folder")
		return CopyFolder(source, dest)
	}
	fmt.Println("Copying file")
	return CopyFile(source, dest)
}

func ZipDirectory(source string, destination string, exclussion []string) error {
	zipFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		

		for _, exclussionPath := range exclussion {
			if relPath == exclussionPath {
				return nil
			}
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(relPath)

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

return err

}