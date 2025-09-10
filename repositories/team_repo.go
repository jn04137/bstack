package repositories

import (
	"database/sql"

	"com/bstack/dependencies"
	"com/bstack/models"
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

func (repo TeamRepository) GetAllTeams() ([]models.Team, error) {
	db := repo.dbConn

	query := `SELECT team.id,team_name,details,user_account.nano_id,team.nano_id from team INNER JOIN user_account
		ON team.owner=user_account.id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		if err := rows.Scan(&team.Id, &team.TeamName, &team.Details, &team.OwnerNanoId, &team.TeamNanoId); err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return teams, nil
}
