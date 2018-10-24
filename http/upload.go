package http

import (
	"encoding/hex"
	"github.com/adcrn/webknest"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/sha3"
	"path/filepath"
	"time"
)

// UploadHandler handles the actual folder upload process
func UploadHandler(c *gin.Context, u webknest.User) (webknest.Folder, error) {
	var f webknest.Folder

	// take username and hash it
	h := make([]byte, 64)
	sha3.ShakeSum256(h, []byte(u.Username))

	// Take the user-supplied folder name; if there isn't one,
	// then take the current system time
	folderName := c.PostForm("name")
	if folderName == "" {
		folderName = time.Now().String()
	}

	// This is only for testing purposes, will be adjusted later
	rootStoragePath := "/var/folders"
	s := hex.EncodeToString(h)

	userFolderPath := filepath.Join(rootStoragePath, s)

	form, err := c.MultipartForm()
	if err != nil {
		return webknest.Folder{}, err
	}

	files := form.File["files"]
	for _, file := range files {
		err = c.SaveUploadedFile(file, filepath.Join(userFolderPath, file.Filename))
		if err != nil {
			return webknest.Folder{}, err
		}
	}

	f = webknest.Folder{
		OwnerID:     u.ID,
		FolderName:  folderName,
		S3Path:      userFolderPath,
		UploadTime:  time.Now(),
		NumElements: len(files),
		Completed:   false,
		Downloaded:  false,
	}

	return f, nil
}
