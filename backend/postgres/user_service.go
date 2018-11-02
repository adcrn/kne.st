package postgres

import (
	"database/sql"
	"github.com/adcrn/webknest/backend"
	"github.com/adcrn/webknest/backend/errors"

	_ "github.com/lib/pq" // Driver for database/sql
	"golang.org/x/crypto/bcrypt"
)

// UserService allows us to interact with the Postgres database
type UserService struct {
	DB *sql.DB
}

// ListBySubscriptionType returns all users with a certain subscription type
func (us *UserService) ListBySubscriptionType(subType int) ([]*webknest.User, error) {
	var users []*webknest.User

	stmt, err := us.DB.Prepare(`select id, username, email, sub_type from users where sub_type = $1`)
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

	stmt, err := us.DB.Prepare(`select * from users where id = $1`)
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

	stmt, err := us.DB.Prepare(`select * from users where username = $1`)
	if err != nil {
		return webknest.User{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&u.ID, &u.Username, &u.Password, &u.FirstName, &u.LastName,
		&u.Email, &u.SubscriptionType)
	if err != nil {
		return webknest.User{}, errors.ErrUsernameNotFound
	}

	return u, nil
}

// Create takes a user object and creates a corresponding database record
func (us *UserService) Create(u webknest.User) (int, error) {
	var userID int

	// Salt and hash the password for storage in database
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	// Start transaction
	tx, err := us.DB.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`insert into users (username, password, first_name, last_name, email, sub_type)
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

// UpdateDetails uses the detail update struct and modifies the user record
func (us *UserService) UpdateDetails(u webknest.User, du webknest.DetailUpdate) error {

	tx, err := us.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`update users set first_name = $1, last_name = $2, sub_type = $3
					where id = $4`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(du.FirstName, du.LastName, du.SubscriptionType, u.ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// ChangePassword changes the password by checking the supplied current
// password and if it passes, then the password is changed to the new one
func (us *UserService) ChangePassword(userID int, pu webknest.PasswordUpdate) error {
	var passFromDB string
	stmt, err := us.DB.Prepare(`select password from users where id = $1`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Retrieve hash of current password from the database
	err = stmt.QueryRow(userID).Scan(&passFromDB)
	if err != nil {
		return err
	}

	// Compare the hash to the confirmation password supplied by user
	err = bcrypt.CompareHashAndPassword([]byte(passFromDB), []byte(pu.CurrentPassword))
	if err != nil {
		return errors.ErrPassDoesNotMatch
	}

	// Generate hash for new user-supplied password
	newHash, err := bcrypt.GenerateFromPassword([]byte(pu.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	tx, err := us.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt2, err := tx.Prepare(`update users set password = $1 where id = $2`)
	if err != nil {
		return err
	}
	defer stmt2.Close()

	// Update user password with new hash
	_, err = stmt2.Exec(string(newHash), userID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// ChangeEmail allows for the easy emodication of email records
func (us *UserService) ChangeEmail(userID int, newEmail string) error {
	tx, err := us.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`update users set email = $1 where id = $2`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newEmail, userID)
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
	tx, err := us.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`delete from users where id = $1`)
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
