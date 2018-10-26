package postgres

import (
	"database/sql"
	"github.com/adcrn/webknest"
	"time"

	_ "github.com/lib/pq" // Driver for database/sql
)

// FolderService allows for interaction with the database
type FolderService struct {
	DB *sql.DB
}

// ListByUser lists all folders that are tied to a particular user
func (fs *FolderService) ListByUser(ownerID int) ([]webknest.Folder, error) {
	var folders []webknest.Folder

	stmt, err := fs.DB.Prepare(`select * from folders where owner_id = $1`)
	if err != nil {
		return []webknest.Folder{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(ownerID)

	if err != nil {
		return []webknest.Folder{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var f webknest.Folder
		err = rows.Scan(&f.OwnerID, &f.FolderName, &f.S3Path, &f.UploadTime, &f.NumElements, &f.Completed, &f.Downloaded)

		if err != nil {
			return []webknest.Folder{}, err
		}

		folders = append(folders, f)
	}

	if err = rows.Err(); err != nil {
		return []webknest.Folder{}, err
	}

	return folders, nil
}

// GetByName returns a single folder object given a user ID and a folder name
func (fs *FolderService) GetByName(ownerID int, name string) (webknest.Folder, error) {
	var f webknest.Folder

	stmt, err := fs.DB.Prepare(`select * from folders where owner_id = $1 and name = $2`)
	if err != nil {
		return webknest.Folder{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(ownerID, name).Scan(&f.OwnerID, &f.FolderName,
		&f.S3Path, &f.UploadTime, &f.NumElements, &f.Completed, &f.Downloaded)
	if err != nil {
		return webknest.Folder{}, err
	}

	return f, nil
}

// Create takes in a folder object and creates a database record
func (fs *FolderService) Create(f webknest.Folder) (int, error) {
	var folderID int
	const sqlTimeFormat = "1993-01-05 23:45:00"

	tx, err := fs.DB.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`insert into folders (folder_name, path, upload_time, 
					num_elements, completed, downloaded) values($1, $2, $3, $4, $5, $6 RETURNING id`)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(f.FolderName, f.S3Path, time.Now().Format(sqlTimeFormat),
		f.NumElements, f.Completed, f.Downloaded).Scan(&folderID)
	if err != nil {
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, nil
	}

	return folderID, nil
}

// Update should only be used to update the completed and downloaded fields
func (fs *FolderService) Update(f webknest.Folder, up webknest.FolderUpdate) error {

	tx, err := fs.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`update folders set completed = $1, downloaded = $2 where id = $3`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(up.Completed, up.Downloaded, f.OwnerID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Delete takes in a user Id and folder name and removes the folder record
func (fs *FolderService) Delete(ownerID int, name string) error {

	tx, err := fs.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`delete from folders where id = $1 and folder_name = $2`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ownerID, name)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
