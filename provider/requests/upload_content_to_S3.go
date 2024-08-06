package requests

import (
	"bytes"
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
	if presignedUrl == nil || archivePath == "" {
		return fmt.Errorf("presignedUrl, archivePath are required")
	}
	_, err := url.Parse(*presignedUrl)
	if err != nil {
		return err
	}

	file, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	bufferFile := make([]byte, fileInfo.Size())
	_, err = file.Read(bufferFile)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, *presignedUrl, bytes.NewBuffer(bufferFile))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	if userId != nil {
		req.Header.Set("x-amz-meta-userid", *userId)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload content to S3: %v", resp)
	}
	return nil

}
