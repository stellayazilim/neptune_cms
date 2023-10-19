-- +goose Up
-- +goose StatementBegin

ALTER TABLE IF EXISTS accounts ADD COLUMN deleted TIMESTAMP NULL;
ALTER TABLE IF EXISTS accounts ADD COLUMN created_at TIMESTAMP DEFAULT NOW();
ALTER TABLE IF EXISTS accounts ADD COLUMN updated_at TIMESTAMP DEFAULT NOW();
ALTER TABLE IF EXISTS tokens ADD COLUMN deleted TIMESTAMP NULL;
ALTER TABLE IF EXISTS tokens ADD COLUMN created_at TIMESTAMP DEFAULT NOW();
ALTER TABLE IF EXISTS tokens ADD COLUMN updated_at TIMESTAMP DEFAULT NOW();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS accounts DROP COLUMN  deleted;
ALTER TABLE IF EXISTS accounts DROP COLUMN  created_at;
ALTER TABLE IF EXISTS accounts DROP COLUMN  updated_at;
ALTER TABLE IF EXISTS tokens DROP  COLUMN  deleted;
ALTER TABLE IF EXISTS tokens DROP COLUMN  created_at;
ALTER TABLE IF EXISTS tokens DROP COLUMN updated_at;

-- +goose StatementEnd
