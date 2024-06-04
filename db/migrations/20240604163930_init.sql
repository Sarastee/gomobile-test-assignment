-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS currencies (
    id BIGSERIAL PRIMARY KEY,
    valute_id VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    numeric_code INT NOT NULL,
    character_code VARCHAR(10) NOT NULL,
    nominal INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    value NUMERIC(20, 4) NOT NULL,
    vunit_rate NUMERIC(20, 4) NOT NULL,
    CONSTRAINT currency_date_unqiue UNIQUE (date, character_code)

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
