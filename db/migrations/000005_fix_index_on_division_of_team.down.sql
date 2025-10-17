
DROP INDEX IF EXISTS division_of_team_team_id_unique;

CREATE UNIQUE INDEX division_of_team_division_id_team_id_key ON division_of_team (team_id);
