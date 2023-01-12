package main

import (
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"
		"fmt"
)

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

func uploadHandler(config Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseMultipartForm(config.MaxFileSize);
			if err != nil {
				http.Error(w, "File size exceeded the maximum allowed", http.StatusBadRequest);
				return;
			}
			file, handler, err := r.FormFile("uploadfile");
			if err != nil {
				http.Error(w, "Error retrieving file", http.StatusBadRequest);
				return;
			}
			defer file.Close();
			filetype := strings.ToLower(filepath.Ext(handler.Filename));
			if !contains(config.AllowedFileTypes, filetype) {
					http.Error(w, "File type not allowed", http.StatusBadRequest);
					return;
			}
			filename := handler.Filename;
			if config.GenerateUniqueFilenames {
					filename = strings.TrimSuffix(filename, filepath.Ext(filename)) + "-" + time.Now().Format("20060102150405") + filepath.Ext(filename);
			}
			f, err := os.OpenFile(filepath.Join(config.UploadDir, filename), os.O_WRONLY|os.O_CREATE, 0666);
			if err != nil {
					http.Error(w, "Error creating file", http.StatusInternalServerError);
					return;
			}
			defer f.Close();
			io.Copy(f, file);
			w.Write([]byte("File uploaded successfully."));
    }
}

func contains(s []string, e string) bool {
    for _, a := range s {
			if a == e {
				return (true);
			}
    }
    return (false);
}

func main() {
    config := Config{
			MaxFileSize:           		32 << 20, // 32MB
			AllowedFileTypes:					[]string{".jpg", ".jpeg", ".png", ".gif"},
			UploadDir:								"uploads/",
			GenerateUniqueFilenames:	true,
    }
		// Create the upload directory if it doesn't exist
		if _, err := os.Stat(config.UploadDir); os.IsNotExist(err) {
			os.Mkdir(config.UploadDir, 0755);
		}
		// http://localhost:8080/upload
		fmt.Println("Server running on http://localhost:8080");
    http.HandleFunc("/upload", uploadHandler(config));
    http.ListenAndServe(":8080", nil);
}
