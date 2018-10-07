package storage

import (
	"database/sql"
	"errors"
	"kne.st/models"
	"strings"

	_ "github.com/lib/pq"
)

// UserStorage is the interface through which methods will access the database
// in order to operate on user objects.
type UserStorage interface {
	List(...UserFilter) ([]models.User, error)
	Get(...UserFilter) (models.User, error)
	Create(models.User) error
	Update(models.User) error
	Delete(...UserFilter) error
}

// UserFilter is the set of criteria that will be used to select certain users
type UserFilter func(*UserFilterConfig) error

// UserFilterConfig is the struct that will be edited and then called by the
// UserFilter interface for searching.
type UserFilterConfig struct {
	ID               int
	Username         string
	FullName         string
	Email            string
	SubscriptionType int
}

// UserIDFilter sets the ID field
func UserIDFilter(id int) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.ID = id
		return nil
	}
}

// UserUsernameFilter sets the username field
func UserUsernameFilter(username string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.Username = username
		return nil
	}
}

// UserFullNameFilter sets the full name field
func UserFullNameFilter(fullName string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.FullName = fullName
		return nil
	}
}

// UserEmailFilter sets the email field
func UserEmailFilter(email string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.Email = email
		return nil
	}
}

// UserSubscriptionTypeFilter sets the subscription type field
func UserSubscriptionTypeFilter(subType int) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.SubscriptionType = subType
		return nil
	}
}

type UserStore struct {
	db *sql.DB
}

func (us *UserStore) List(filters ...UserFilter) ([]*models.User, error) {
	return nil, nil
}

func (us *UserStore) Get(filters ...UserFilter) (models.User, error) {
	return models.User{}, nil
}

func (us *UserStore) Create(u models.User) error {
	return nil
}

func (us *UserStore) Update(u models.User) error {
	return nil
}

func (us *UserStore) Delete(filters ...UserFilter) error {
	return nil
}

func NewUserStore(db *sql.DB) UserStorage {
	return &UserStore{db}
}
