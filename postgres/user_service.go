package postgres

import (
	"database/sql"
	"github.com/adcrn/webknest"

	_ "github.com/lib/pq" // Driver for database/sql
	"golang.org/x/crypto/bcrypt"
)

// UserService allows us to interact with the Postgres database
type UserService struct {
	db *sql.DB
}

// ListBySubscriptionType returns all users with a certain subscription type
func (us *UserService) ListBySubscriptionType(subType int) ([]*webknest.User, error) {
	var users []*webknest.User

	stmt, err := us.db.Prepare(`select id, username, email, sub_type from users where sub_type = $N`)
	if err != nil {
		return []*webknest.User{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(subType)

	if err != nil {
		return []*webknest.User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var u webknest.User
		err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.SubscriptionType)

		if err != nil {
			return []*webknest.User{}, err
		}

		// Don't need information like this being passed for this operation
		u.Password = ""
		u.FirstName = ""
		u.LastName = ""

		users = append(users, &u)
	}

	if err = rows.Err(); err != nil {
		return []*webknest.User{}, err
	}

	return users, nil
}

// GetByID returns a user database record given a user ID
func (us *UserService) GetByID(userID int) (webknest.User, error) {
	var u webknest.User

	stmt, err := us.db.Prepare(`select * from users where id = $N`)
	if err != nil {
		return webknest.User{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID).Scan(&u.ID, &u.Username, &u.Password, &u.FirstName, &u.LastName, &u.Email, &u.SubscriptionType)
	if err != nil {
		return webknest.User{}, err
	}

	return u, nil
}

// GetByUsername returns a user database record given a username
func (us *UserService) GetByUsername(username string) (webknest.User, error) {
	var u webknest.User

	stmt, err := us.db.Prepare(`select * from users where username = $N`)
	if err != nil {
		return webknest.User{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&u.ID, &u.Username, &u.Password, &u.FirstName, &u.LastName,
		&u.Email, &u.SubscriptionType)
	if err != nil {
		return webknest.User{}, err
	}

	return u, nil
}

// Create takes a user object and creates a corresponding database record
func (us *UserService) Create(u webknest.User) (int, error) {
	var userID int

	// Salt and hash the password for storage in database
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return -1, err
	}

	// Start transaction
	tx, err := us.db.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`insert into users(username, password, first_name, last_name, email, sub_type)
					values($1, $2, $3, $4, $5, $6) RETURNING id`)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	// Make sure that the hashed password is stored instead of the actual one
	err = stmt.QueryRow(u.Username, string(hashedPass), u.FirstName, u.LastName, u.Email, u.SubscriptionType).Scan(&userID)
	if err != nil {
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return userID, nil
}

// Update uses the credential update struct and modies the user record
func (us *UserService) Update(u webknest.User, c webknest.CredentialUpdate) error {

	newPass, err := bcrypt.GenerateFromPassword([]byte(c.Password), 10)
	if err != nil {
		return err
	}

	tx, err := us.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`update users set password = $1, full_name = $2, email = $3, sub_type = $4
					where id = $5`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(string(newPass), c.FirstName, c.LastName, c.Email, c.SubscriptionType, u.ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Delete removes the corresponding user record from the database
func (us *UserService) Delete(userID int) error {

	// Start transaction
	tx, err := us.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`delete from users where id = $N`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Returns no rows so we just execute
	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
