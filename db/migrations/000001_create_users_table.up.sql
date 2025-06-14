CREATE TABLE user_account(
    id serial primary key,
    username VARCHAR(72) UNIQUE NOT NULL,
    email VARCHAR(140) UNIQUE,
    nano_id VARCHAR(21) UNIQUE NOT NULL,
    password VARCHAR(140) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE OR REPLACE FUNCTION update_last_updated_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_account_update_lasted_updated BEFORE UPDATE
on user_account FOR EACH ROW EXECUTE PROCEDURE
update_last_updated_column();

