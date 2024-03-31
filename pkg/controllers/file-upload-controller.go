package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	// "github.com/joho/godotenv"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form data
	err := r.ParseMultipartForm(32 << 20) // adjust memory limit as needed
	if err != nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
		return
	}

	// Get uploaded file
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Fprintf(w, "Error retrieving file: %v", err)
		return
	}
	defer file.Close()

	// Validate file type
	if err := validateFileType(file); err != nil {
		fmt.Fprintf(w, "Error: Invalid file type. Only JPEG, JPG, and PNG files allowed: %v", err)
		return
	}

	folderPath := os.Getenv("ANNEX_FILE_UPLOADS_PATH")
	// Create folder if it doesn't exist
	_, err = os.Stat(folderPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(folderPath, os.ModePerm) // adjust permissions as needed
		if err != nil {
			fmt.Fprintf(w, "Error creating uploads folder: %v", err)
			return
		}
	}

	filename := fmt.Sprintf(
		"%s/%s_%d%s",
		folderPath,
		strings.TrimSuffix(handler.Filename, filepath.Ext(handler.Filename)),
		time.Now().Unix(),
		filepath.Ext(handler.Filename),
	)

	// Create new file and write uploaded data
	dst, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(w, "Error creating file: %v", err)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Fprintf(w, "Error saving file: %v", err)
		return
	}

	message := fmt.Sprintf("Successfully uploaded file: %s", filename)
	response := map[string]string{"message": message}
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func validateFileType(file io.Reader) error {

	buf := make([]byte, 512)
	file.Read(buf)

	contentType := http.DetectContentType(buf)
	fmt.Println(contentType)
	defer file.(io.Seeker).Seek(0, io.SeekStart)

	allowedTypes := []string{"image/jpeg", "image/png", "image/jpg"}
	for _, allowedType := range allowedTypes {
		if contentType == allowedType {
			return nil
		}
	}

	return errors.New("invalid file type")
}
