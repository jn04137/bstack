
CREATE TABLE Team_ESEA_Division (
    id smallserial PRIMARY KEY,
    division VARCHAR(30) NOT NULL
);

INSERT INTO Team_ESEA_Division (division) VALUES ('Open');
INSERT INTO Team_ESEA_Division (division) VALUES ('Intermediate');
INSERT INTO Team_ESEA_Division (division) VALUES ('Main');
INSERT INTO Team_ESEA_Division (division) VALUES ('Advanced');
INSERT INTO Team_ESEA_Division (division) VALUES ('ESL Challenger League');

CREATE TABLE Division_of_Team (
    id SERIAL PRIMARY KEY,
    division_id int,
    team_id int,
    FOREIGN KEY (division_id) REFERENCES Team_ESEA_Division(id),
    FOREIGN KEY (team_id) REFERENCES team(id),
    UNIQUE(division_id, team_id)
);

