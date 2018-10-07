package storage

import (
	"database/sql"
	"errors"
	"kne.st/models"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// FolderStorage is the interface through which methods will access the database
// in order to operate on folder objects.
type FolderStorage interface {
	List(...FolderFilter) ([]models.Folder, error)
	Get(...FolderFilter) (models.Folder, error)
	Create(models.Folder) error
	Update(models.Folder) error
	Delete(...FolderFilter) error
}

// FolderFilter is the set of criteria that will be used to select certain
// folders
type FolderFilter func(*FolderFilterConfig) error

// FolderFilterConfig is the struct that will be edited and then called by the
// FolderFilter interface for searching.
type FolderFilterConfig struct {
	OwnerID    int
	FolderName string
	Created    time.Time
	Completed  bool
	Downloaded bool
}

type FolderStore struct {
	db *sql.DB
}

// FolderOwnerIDFilter sets the ownerID field
func FolderOwnerIDFilter(ownerID int) FolderFilter {
	return func(fc *FolderFilterConfig) error {
		fc.OwnerID = ownerID
		return nil
	}
}

// FolderFolderNameFilter sets the folderName field
func FolderFolderNameFilter(folderName string) FolderFilter {
	return func(fc *FolderFilterConfig) error {
		fc.FolderName = folderName
		return nil
	}
}

// FolderCreatedFilter sets the created field
func FolderCreatedFilter(created time.Time) FolderFilter {
	return func(fc *FolderFilterConfig) error {
		fc.Created = created
		return nil
	}
}

// FolderCompletedFilter sets the completed field
func FolderCompletedFilter(completed bool) FolderFilter {
	return func(fc *FolderFilterConfig) error {
		fc.Completed = completed
		return nil
	}
}

// FolderDownloadedFilter sets the downloaded field
func FolderDownloadedFilter(downloaded bool) FolderFilter {
	return func(fc *FolderFilterConfig) error {
		fc.Downloaded = downloaded
		return nil
	}
}

func (fs *FolderStore) List(filters ...FolderFilter) ([]models.Folder, error) {
	return nil, nil
}

func (fs *FolderStore) Get(filters ...FolderFilter) (models.Folder, error) {
	return models.Folder{}, nil
}

func (fs *FolderStore) Create(folder models.Folder) error {
	return nil
}

func (fs *FolderStore) List(folder models.Folder) error {
	return nil
}

func (fs *FolderStore) List(filters ...FolderFilter) error {
	return nil
}

func NewFolderStore(db *sql.DB) FolderStorage {
	return &FolderStore(db)
}
