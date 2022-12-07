package dbHelper

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"todo/database"
	"todo/models"
	"todo/utils"
)

func IsUserExists(email string) (bool, error) {
	// language=SQL
	SQL := `SELECT id FROM users WHERE email = TRIM(LOWER($1)) AND archived_at IS NULL`
	var id string
	err := database.Todo.Get(&id, SQL, email)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}

func CreateUser(db sqlx.Ext, name, email, password string) (string, error) {
	// language=SQL
	SQL := `INSERT INTO users(name, email, password) VALUES ($1, TRIM(LOWER($2)), $3) RETURNING id`
	var userID string
	if err := db.QueryRowx(SQL, name, email, password).Scan(&userID); err != nil {
		return "", err
	}
	return userID, nil
}

func CreateUserRole(db sqlx.Ext, userID string, role models.Role) error {
	// language=SQL
	SQL := `INSERT INTO user_roles(user_id, role) VALUES ($1, $2)`
	_, err := db.Exec(SQL, userID, role)
	return err
}

func CreateUserSession(db sqlx.Ext, userID, sessionToken string) error {
	// language=SQL
	SQL := `INSERT INTO user_session(user_id, session_token) VALUES ($1, $2)`
	_, err := db.Exec(SQL, userID, sessionToken)
	return err
}

func GetUserIDByPassword(email, password string) (string, error) {
	// language=SQL
	SQL := `SELECT
				u.id,
       			u.password
       		FROM
				users u
			WHERE
				u.archived_at IS NULL
				AND u.email = TRIM(LOWER($1))`
	var userID string
	var passwordHash string
	err := database.Todo.QueryRowx(SQL, email).Scan(&userID, &passwordHash)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	if err == sql.ErrNoRows {
		return "", nil
	}
	// compare password
	if passwordErr := utils.CheckPassword(password, passwordHash); passwordErr != nil {
		return "", passwordErr
	}
	return userID, nil
}

func GetUserBySession(sessionToken string) (*models.User, error) {
	// language=SQL
	SQL := `SELECT 
       			u.id, 
       			u.name, 
       			u.email, 
       			u.created_at 
			FROM users u
			JOIN user_session us on u.id = us.user_id
			WHERE u.archived_at IS NULL AND us.session_token = $1`
	var user models.User
	err := database.Todo.Get(&user, SQL, sessionToken)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	roles, roleErr := GetUserRoles(user.ID)
	if roleErr != nil {
		return nil, roleErr
	}
	user.Roles = roles
	return &user, nil
}

func GetUserRoles(userID string) ([]models.UserRole, error) {
	// language=SQL
	SQL := `SELECT id, user_id, role FROM user_roles WHERE user_id = $1 AND archived_at IS NULL`
	roles := make([]models.UserRole, 0)
	err := database.Todo.Select(&roles, SQL, userID)
	return roles, err
}

func DeleteSessionToken(userID, token string) error {
	// language=SQL
	SQL := `DELETE FROM user_session WHERE user_id = $1 AND session_token = $2`
	_, err := database.Todo.Exec(SQL, userID, token)
	return err
}
