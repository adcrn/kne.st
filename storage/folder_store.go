package storage

import (
	"database/sql"
	"kne.st/models"
	"time"

	_ "github.com/lib/pq" // Driver for database/sql
)

// FolderStorage is the interface through which methods will access the database
// in order to operate on folder objects.
type FolderStorage interface {
	ListByUser(int) ([]models.Folder, error)
	GetByName(int, string) (models.Folder, error)
	Create(models.Folder) (int, error)
	Update(models.Folder, models.FolderUpdate) error
	Delete(int, string) error
}

// FolderStore allows for interaction with the database
type FolderStore struct {
	db *sql.DB
}

// ListByUser lists all folders that are tied to a particular user
func (fs *FolderStore) ListByUser(ownerID int) ([]models.Folder, error) {
	var folders []models.Folder

	stmt, err := fs.db.Prepare(`select * from folders where owner_id = $N`)
	if err != nil {
		return []models.Folder{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(ownerID)

	if err != nil {
		return []models.Folder{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var f models.Folder
		err = rows.Scan(&f.OwnerID, &f.FolderName, &f.FolderNameURL, &f.Created, &f.NumElements, &f.Completed, &f.Downloaded)

		if err != nil {
			return []models.Folder{}, err
		}

		folders = append(folders, f)
	}

	if err = rows.Err(); err != nil {
		return []models.Folder{}, err
	}

	return folders, nil
}

// GetByName returns a single folder object given a user ID and a folder name
func (fs *FolderStore) GetByName(ownerID int, name string) (models.Folder, error) {
	var f models.Folder

	stmt, err := fs.db.Prepare(`select * from folders where owner_id = $1 and name = $2`)
	if err != nil {
		return models.Folder{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(ownerID, name).Scan(&f.OwnerID, &f.FolderName,
		&f.FolderNameURL, &f.Created, &f.NumElements, &f.Completed, &f.Downloaded)
	if err != nil {
		return models.Folder{}, err
	}

	return f, nil
}

// Create takes in a folder object and creates a database record
func (fs *FolderStore) Create(f models.Folder) (int, error) {
	var folderID int
	const sqlTimeFormat = "1993-01-05 23:45:00"

	tx, err := fs.db.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`insert into folders(folder_name, upload_time, 
					num_elements, completed, downloaded) values($1, $2, $3, $4, $5 RETURNING id`)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(f.FolderName, time.Now().Format(sqlTimeFormat),
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
func (fs *FolderStore) Update(f models.Folder, up models.FolderUpdate) error {

	tx, err := fs.db.Begin()
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
func (fs *FolderStore) Delete(ownerID int, name string) error {

	tx, err := fs.db.Begin()
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

// NewFolderStore returns a struct that implements the FolderStorage interface
func NewFolderStore(db *sql.DB) FolderStorage {
	return &FolderStore{db}
}
