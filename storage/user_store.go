package storage

import (
	"database/sql"
	"errors"
	"kne.st/models"
	//"strings"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// UserStorage is the interface through which methods will access the database
// in order to operate on user objects.
type UserStorage interface {
	List(...UserFilter) ([]models.User, error)
	Get(...UserFilter) (models.User, error)
	Create(models.User) (int, error)
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

// Right now, this only filters by subscription type
func (us *UserStore) List(filters ...UserFilter) ([]*models.User, error) {
	var users []*models.User

	stmt, err := us.db.Prepare(`select id, username, email, sub_type from users where sub_type = $N`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(filters.SubscriptionType)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.SubscriptionType)

		if err != nil {
			return nil, err
		}

		// Don't need information like this being passed for this operation
		u.Password = ""
		u.FullName = ""

		users = append(users, &u)
	}

	if err = rows.Err; err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (us *UserStore) Get(filters ...UserFilter) (models.User, error) {
	var u models.User

	stmt, err := us.db.Prepare(`select id, username, email, sub_type from users where id = $N`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(filters.ID).Scan(&u.ID, &u.Username, &u.Email, &u.SubscriptionType)
	if err != nil {
		return nil, errors.New("Unable to complete Get query!")
	}

	u.Password = ""
	u.FullName = ""

	return u, nil
}

func (us *UserStore) Create(u models.User) (int, error) {
	var userID int

	// Salt and hash the password for storage in database
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return nil, err.Error()
	}
	tx, err := us.db.Begin()
	if err != nil {
		return nil, errors.New("Unable to begin user creation transaction!")
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`insert into users(username, password, full_name, email, sub_type)
					values($1, $2, $3, $4, $5) RETURNING id`)
	if err != nil {
		return nil, errors.New("Unable to complete user creation statement preparation!")
	}
	defer stmt.Close()

	err = stmt.QueryRow(u.Username, u.Password, u.FullName, u.Email, u.SubscriptionType).Scan(&userID)
	if err != nil {
		return nil, errors.New("Unable to complete user creation query operation!")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("Unable to commit user creation transaction!")
	}

	return userID, nil
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
