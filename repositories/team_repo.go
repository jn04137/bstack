package repositories

import (
	"log"
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

func (repo TeamRepository) GetTeamByNanoId(teamNanoId string) (models.Team, error) {
	db := repo.dbConn

	query := `SELECT team.id,team.team_name,team.details,user_account.username
		FROM 
			team
		INNER JOIN 
			user_account
		ON 
			team.owner=user_account.id
		WHERE 
			team.nano_id=$1`

	row := db.QueryRow(query, teamNanoId)

	var team models.Team
	err := row.Scan(&team.Id, &team.TeamName, &team.Details, &team.OwnerName)
	if err != nil {
		return team, err
	}
	team.TeamNanoId = teamNanoId

	return team, err
}

func (repo TeamRepository) GetTeamESEADivision() string {
	return ""
}

func (repo TeamRepository) SetESEADivision(teamNanoId string, eseaDivision string) error {
	db := repo.dbConn

	query := `INSERT INTO (division_id,team_id) VALUES ($1, $2)`

	_, err := db.Exec(query, teamNanoId, eseaDivision)
	return err
}

func (repo TeamRepository) GetESEADivision(teamId int) (string, error) {
	db := repo.dbConn
	query := `SELECT 
			division 
		FROM 
			team_esea_division 
		WHERE 
			team_esea_division.id=(SELECT division_id FROM division_of_team WHERE division_of_team=$1)`

		var eseaDivision string
		row := db.QueryRow(query, teamId)
		err := row.Scan(&eseaDivision)
		if err != nil {
			return eseaDivision, err
		}
		return eseaDivision, err
}

func (repo TeamRepository) GetESEADivisions() ([]string, error) {
	db := repo.dbConn
	query := `SELECT id,division FROM team_esea_division`

	var eseaDivisions []string
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("This is the error: %v", err)
		return eseaDivisions, err
	}
	for rows.Next() {
		var division string
		if err := rows.Scan(&division); err != nil {
			log.Printf("This is the error: %v", err)
			return eseaDivisions, err
		}
		eseaDivisions = append(eseaDivisions, division)
	}

	return eseaDivisions, err
}

func (repo TeamRepository) GetPlayersOnTeam(teamNanoId string) {

}

