
CREATE TABLE team (
    id SERIAL primary key,
    team_name VARCHAR(72) UNIQUE NOT NULL,
    owner INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    FOREIGN KEY (owner) REFERENCES user_account(id)
);

CREATE TABLE user_on_team (
    id serial primary key,
    team INT NOT NULL,
    player INT NOT NULL,
    accepted BOOL NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    FOREIGN KEY (team) REFERENCES team(id),
    FOREIGN KEY (player) REFERENCES user_account(id)
);

CREATE TRIGGER update_team_update_lasted_updated BEFORE UPDATE
    on team FOR EACH ROW EXECUTE PROCEDURE
    update_last_updated_column();

CREATE TRIGGER update_user_on_team_update_lasted_updated BEFORE UPDATE
    on user_on_team FOR EACH ROW EXECUTE PROCEDURE
    update_last_updated_column();
