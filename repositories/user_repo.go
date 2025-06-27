package repositories

import (
	"database/sql"

	"com/bstack/dependencies"
	"com/bstack/models"
)

type UserRepository struct {
	dbConn *sql.DB
}

func NewUserRepository(env *dependencies.Environment) *UserRepository {
	return &UserRepository{
		dbConn: env.DBConn,
	}
}

func (repo UserRepository) CreateUser(user models.UserAccount) error {
	db := repo.dbConn

	query := `INSERT INTO user_account (username, email, nano_id, password) VALUES ($1,$2,$3,$4)`
	_, err := db.Exec(query, user.Username, user.Email, user.NanoId, user.Password)
	return err
}

func (repo UserRepository) GetUser(username string) (models.UserAccount, error) {
	db := repo.dbConn
	query := `SELECT username,password,nano_id from user_account WHERE username=$1`
	row := db.QueryRow(query, username)

	user := models.UserAccount{}
	err := row.Scan(&user.Username, &user.Password, &user.NanoId)

	return user, err
}

func (repo UserRepository) GetAllUsers() ([]models.UserAccount, error) {
	db := repo.dbConn
	query := `SELECT id,username,nano_id from user_account`

	users := []models.UserAccount{}
	rows, err := db.Query(query)

	for rows.Next() {
		var user models.UserAccount
		if err := rows.Scan(&user.Id, &user.Username, &user.NanoId); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, err

}

