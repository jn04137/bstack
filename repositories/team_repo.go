package repositories

import (
	"database/sql"

	"com/bstack/dependencies"
)

type TeamRepository struct {
	dbConn *sql.DB
}

func NewTeamRepository(env *dependencies.Environment) *TeamRepository {
	return &TeamRepository{
		dbConn: env.DBConn,
	}
}

func (repo TeamRepository) CreateTeam(teamName string, ownerNanoId string) (error) {
	db := repo.dbConn

	query := `INSERT INTO Team (team_name, owner) VALUES ($1, SELECT user_account.id FROM user_account WHERE nano_id=$2)`
	_, err := db.Exec(query, teamName, ownerNanoId)
	return err
}
