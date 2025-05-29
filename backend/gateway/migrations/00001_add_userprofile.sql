-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS goose_db_version (
id SERIAL PRIMARY KEY,
version_id BIGINT NOT NULL,
is_applied BOOLEAN NOT NULL,
tstamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE EXTENSION IF NOT EXISTS "pg_trgm";

CREATE TABLE userprofile (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    bio TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_userprofile_email ON userprofile(email);
CREATE INDEX idx_userprofile_name ON userprofile(name) WHERE name IS NOT NULL;

COMMENT ON TABLE userprofile IS 'User profiles table';
COMMENT ON COLUMN userprofile.id IS 'Primary key';
COMMENT ON COLUMN userprofile.name IS 'User full name (optional)';
COMMENT ON COLUMN userprofile.email IS 'Unique email address';
COMMENT ON COLUMN userprofile.password_hash IS 'Hashed password';
COMMENT ON COLUMN userprofile.bio IS 'User biography (optional)';
COMMENT ON COLUMN userprofile.created_at IS 'Timestamp of user creation';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS userprofile;
DROP EXTENSION IF EXISTS "pg_trgm";
-- +goose StatementEnd
