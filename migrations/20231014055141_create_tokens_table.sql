-- +goose Up
-- +goose StatementBegin
CREATE TYPE TokenType AS ENUM(
    'VALIDATION',
    'AUTHENTICATION',
    'OTP',
    'SECRET',
    'API_KEY',
    'REFRESH'
);

CREATE TABLE IF NOT EXISTS Tokens (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    type TokenType DEFAULT 'VALIDATION',
    token TEXT NOT NULL,
    account_id UUID  REFERENCES Accounts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Tokens;
DROP TYPE IF EXISTS TokenType;
-- +goose StatementEnd
