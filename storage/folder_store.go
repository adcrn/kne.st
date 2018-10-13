package storage

import (
	"database/sql"
	"errors"
	"kne.st/models"
	//"strings"
	//"time"

	_ "github.com/lib/pq"
)

// FolderStorage is the interface through which methods will access the database
// in order to operate on folder objects.
type FolderStorage interface {
	ListByUser(int) ([]models.Folder, error)
	GetByName(int, string) (models.Folder, error)
	Create(models.Folder) error
	Update(models.Folder) error
	Delete(int, string) error
}

type FolderStore struct {
	db *sql.DB
}

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

func (fs *FolderStore) Create(folder models.Folder) error {
	return nil
}

func (fs *FolderStore) Update(folder models.Folder) error {
	return nil
}

func (fs *FolderStore) Delete(ownerID int, name string) error {
	return nil
}

func NewFolderStore(db *sql.DB) FolderStorage {
	return &FolderStore(db)
}
