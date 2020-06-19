package nextcloud

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// Config contains the required data to connect to a nextcloud instance
type Config struct {
	HostURL  string
	Username string
	Password string
}

// UploadFile uploads file to given remoteFolderPath & fileName from current directory
func (c *Config) UploadFile(remoteFolderPath, remoteFileName, localFilePath string) error {
	dat, err := ioutil.ReadFile(localFilePath)
	if err != nil {
		return err
	}

	uploadURL := fmt.Sprintf("%s/remote.php/dav/files/%s/%s/%s", c.HostURL, c.Username, remoteFolderPath, remoteFileName)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", uploadURL, bytes.NewBuffer(dat))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// DownloadFile downloads file from given remoteFolderPath & remoteFileName to current directory
func (c *Config) DownloadFile(remoteFolderPath, remoteFileName, localDownloadPath string) error {
	downloadURL := fmt.Sprintf("%s/remote.php/dav/files/%s/%s/%s", c.HostURL, c.Username, remoteFolderPath, remoteFileName)
	client := &http.Client{}
	req, err := http.NewRequest("GET", downloadURL, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	downloadPath := filepath.Join(localDownloadPath, remoteFileName)

	out, err := os.Create(downloadPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
