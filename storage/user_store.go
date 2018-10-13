package storage

import (
	"database/sql"
	"errors"
	"kne.st/models"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// UserStorage is the interface through which methods will access the database
// in order to operate on user objects.
type UserStorage interface {
	ListBySubscriptionType(int) ([]*models.User, error)
	GetByID(int) (models.User, error)
	GetByUsername(string) (models.User, error)
	Create(models.User) (int, error)
	Update(models.User) error
	Delete(int, string) error
}

type UserStore struct {
	db *sql.DB
}

// Right now, this only filters by subscription type
func (us *UserStore) ListBySubscriptionType(subType int) ([]*models.User, error) {
	var users []*models.User

	stmt, err := us.db.Prepare(`select id, username, email, sub_type from users where sub_type = $N`)
	if err != nil {
		return []*models.User{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(subType)

	if err != nil {
		return []*models.User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.SubscriptionType)

		if err != nil {
			return []*models.User{}, err
		}

		// Don't need information like this being passed for this operation
		u.Password = ""
		u.FullName = ""

		users = append(users, &u)
	}

	if err = rows.Err(); err != nil {
		return []*models.User{}, err
	} else {
		return users, nil
	}
}

func (us *UserStore) GetByID(userID int) (models.User, error) {
	var u models.User

	stmt, err := us.db.Prepare(`select id, username, email, sub_type from users where id = $N`)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID).Scan(&u.ID, &u.Username, &u.Email, &u.SubscriptionType)
	if err != nil {
		return models.User{}, err
	}

	u.Password = ""
	u.FullName = ""

	return u, nil
}

func (us *UserStore) GetByUsername(username string) (models.User, error) {
	var u models.User

	stmt, err := us.db.Prepare(`select id, username, email, sub_type from users where username = $N`)
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&u.ID, &u.Username, &u.Email, &u.SubscriptionType)
	if err != nil {
		return models.User{}, err
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
		return -1, err
	}

	// Start transaction
	tx, err := us.db.Begin()
	if err != nil {
		return -1, errors.New("Unable to begin user creation transaction!")
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`insert into users(username, password, full_name, email, sub_type)
					values($1, $2, $3, $4, $5) RETURNING id`)
	if err != nil {
		return -1, errors.New("Unable to complete user creation statement preparation!")
	}
	defer stmt.Close()

	err = stmt.QueryRow(u.Username, u.Password, u.FullName, u.Email, u.SubscriptionType).Scan(&userID)
	if err != nil {
		return -1, errors.New("Unable to complete user creation query operation!")
	}

	err = tx.Commit()
	if err != nil {
		return -1, errors.New("Unable to commit user creation transaction!")
	}

	return userID, nil
}

func (us *UserStore) Update(u models.User) error {
	return nil
}

// Final check of password before deleting an account
func (us *UserStore) Delete(userID int, password string) error {
	return nil
}

func NewUserStore(db *sql.DB) UserStorage {
	return &UserStore{db}
}
