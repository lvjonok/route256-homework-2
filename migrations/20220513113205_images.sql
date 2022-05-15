-- +goose Up
-- +goose StatementBegin

CREATE TABLE images
(
    id    SERIAL      PRIMARY KEY,
    image BYTEA,
    href  text,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE images;
-- +goose StatementEnd
