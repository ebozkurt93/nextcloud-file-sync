package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	hostURL         string
	username        string
	password        string
	ncFolderPath    string
	ncFileName      string
	localFolderPath string
	localFileName   string
	download        bool
	upload          bool
}

func main() {
	c, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#+v\n", c)
	if c.upload == c.download {
		log.Fatal("Only upload or download should be selected")
	}
	localFilePath := filepath.Join(c.localFolderPath, c.localFileName)
	fmt.Println(localFilePath)

	if c.download {
		log.Println("Downloading...")
	} else if c.upload {
		log.Println("Uploading...")
	}
}

// getConfig read parameters from env file and as cmd params and returns configuration in a struct
func getConfig() (config, error) {
	v := viper.New()
	// Support reading file & folder paths for nextcloud and local from flags
	pflag.String("nc_folder_path", "asda", "nextcloud folder path")
	v.BindPFlag("nc_folder_path", pflag.Lookup("nc_folder_path"))
	pflag.String("nc_file_path", "asda", "nextcloud file path")
	v.BindPFlag("nc_file_path", pflag.Lookup("nc_file_path"))
	pflag.String("local_folder_path", "asda", "local folder path")
	v.BindPFlag("local_folder_path", pflag.Lookup("local_folder_path"))
	pflag.String("local_file_path", "asda", "local file path")
	v.BindPFlag("local_file_path", pflag.Lookup("local_file_path"))

	pflag.Bool("d", false, "download from nextcloud server")
	v.BindPFlag("d", pflag.Lookup("d"))
	pflag.Bool("u", false, "upload to nextcloud server")
	v.BindPFlag("u", pflag.Lookup("u"))

	pflag.Parse()
	v.SetConfigName("config")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		return config{}, errors.Wrap(err, "Cannot read env config, env file should be named config.env")
	}
	v.AutomaticEnv()

	// viper cannot unmarshal when using AutomaticEnv
	var c config = config{
		hostURL:         v.GetString("host_url"),
		username:        v.GetString("username"),
		password:        v.GetString("password"),
		ncFolderPath:    v.GetString("nc_folder_path"),
		ncFileName:      v.GetString("nc_file_name"),
		localFolderPath: v.GetString("local_folder_path"),
		localFileName:   v.GetString("local_file_name"),
		download:        v.GetBool("d"),
		upload:          v.GetBool("u"),
	}
	return c, nil
}
