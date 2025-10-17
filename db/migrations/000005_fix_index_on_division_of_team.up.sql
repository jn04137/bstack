ALTER TABLE division_of_team DROP CONSTRAINT division_of_team_division_id_team_id_key;

ALTER TABLE division_of_team ADD CONSTRAINT unique_team_id UNIQUE (team_id);
ALTER TABLE division_of_team ALTER COLUMN team_id SET NOT NULL;
