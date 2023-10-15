-- +goose Up
-- +goose StatementBegin
CREATE TYPE TokenStatus AS ENUM ('VALID', 'INVALID');
ALTER TABLE tokens RENAME COLUMN token TO value;
ALTER TABLE tokens ADD COLUMN  status TokenStatus DEFAULT 'VALID';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tokens RENAME  value TO token;
ALTER TABLE tokens DROP COLUMN status;
DROP TYPE IF EXISTS TokenStatus;
-- +goose StatementEnd
