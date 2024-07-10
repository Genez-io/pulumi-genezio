package requests

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func UploadContentToS3(
	presignedUrl *string,
	archivePath string,
	userId *string,
) error {
	if presignedUrl == nil || archivePath=="" {
		return fmt.Errorf("presignedUrl, archivePath are required")
	}

	fmt.Printf("Uploading to S3 2.1 %s ", *presignedUrl)
	_, err := url.Parse(*presignedUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Uploading to S3 2.2 %s", *presignedUrl)

	file, err:= os.Open(archivePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("Uploading to S3 2.3")

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	fmt.Println("Uploading to S3 2.4")

	req, err := http.NewRequest("PUT", *presignedUrl, file)
	if err != nil {
		return err
	}

	fmt.Println("Uploading to S3 2.5")
	
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	fmt.Println("Uploading to S3 2.6")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println("Uploading to S3 2.7")
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload content to S3: %d", resp.StatusCode)
	}
	return nil

}