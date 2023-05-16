# GoLang File Upload Service

A simple GoLang service that allows file uploads with configurable options.

## Features

- Maximum file size limit
- Allowed file types
- Upload directory configuration
- Option to generate unique filenames

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Anas-Qadil/Filebeam
```
2. Change into the project directory:

```bash
cd Filebeam
```
3. Build the executable:

```bash
go build
```
4. Run the service:

```bash
./Filebeam
```

## Configuration

The service can be configured by modifying the Config struct in the main.go file:

```golang
type Config struct {
    // Max file size allowed for uploads
    MaxFileSize int64;
    // Allowed file types for uploads
    AllowedFileTypes []string;
    // Upload directory
    UploadDir string;
    // Generate unique filenames
    GenerateUniqueFilenames bool;
}

func main() {
    config := Config{
        MaxFileSize:            32 << 20, // 32MB
        AllowedFileTypes:       []string{".jpg", ".jpeg", ".png", ".gif"},
        UploadDir:              "uploads/",
        GenerateUniqueFilenames: true,
    }
    // ...
}
```

Make sure to adjust the configuration according to your requirements before running the service.

## Usage

Once the service is running, you can upload files by sending a POST request to http://localhost:8080/upload.</br>
The uploaded files will be stored in the configured upload directory.</br>

To change the upload directory, modify the UploadDir field in the Config struct.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
