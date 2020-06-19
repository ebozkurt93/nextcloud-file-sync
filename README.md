# nextcloud-file-sync

This tool uploads/downloads files from/to a given nextcloud instance.

## Setup / Configuration
The application requires a `config.env` file for reading nextcloud parameters. Download/upload paths could be defined in the config file, or passed as flags.

| Environment Variable | Description | Could be used via env file | Could be passed as a flag | Notes |
|-----|------|-----|-----|-----|
| `HOST_URL` | Nextcloud host url | :heavy_check_mark:| :x: | |
| `USERNAME` | Nextcloud username | :heavy_check_mark:| :x: | |
| `PASSWORD` | Nextcloud password | :heavy_check_mark:| :x: | |
| `NC_FOLDER_PATH` | Nextcloud upload/download folder path | :heavy_check_mark:| :heavy_check_mark: | |
| `NC_FILE_NAME` | Nextcloud upload/download file name | :heavy_check_mark:| :heavy_check_mark: | |
| `LOCAL_FOLDER_PATH` | Local upload/download folder path | :heavy_check_mark:| :heavy_check_mark: | |
| `LOCAL_FILE_NAME` | Local upload/download file name | :heavy_check_mark:| :heavy_check_mark: | |
| `d` | Download | :x:| :heavy_check_mark: | Only one of download/upload flags should be passed |
| `u` | Upload | :x:| :heavy_check_mark: | Only one of download/upload flags should be passed |

* If you are passing variables via flags, flags should be named in lowercase.
* Passing none or both of the download/upload flags will result in execution being stopped

Examples: 
* `go run . --d`
* `go run . --nc_folder_path <nc_folder_path> --nc_file_name <nc_file_name> --local_folder_path <local_folder_path> --local_file_name <local_file_name> --d`