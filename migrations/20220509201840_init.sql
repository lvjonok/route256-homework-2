-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    new.updated_at = NOW();
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE categories
(
    id          SERIAL PRIMARY KEY,
    category_id INT NOT NULL,
    task_number INT NOT NULL,
    title       TEXT NOT NULL DEFAULT '',
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE problems
(
    id          SERIAL PRIMARY KEY,
    problem_id  INT       NOT NULL,
    category_id INT       NOT NULL REFERENCES categories (id),
    image       TEXT      NOT NULL DEFAULT '',
    parts       TEXT[]    NOT NULL DEFAULT '{}',
    answer      TEXT      NOT NULL DEFAULT '',
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);



CREATE TABLE submissions
(
    id         SERIAL PRIMARY KEY,
    chat_id    SERIAL    NOT NULL,
    problem_id SERIAL REFERENCES problems (id),
    result     TEXT               DEFAULT 'pending' CHECK ( result IN ('correct', 'wrong', 'pending', 'aborted') ),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp_categories
    BEFORE UPDATE
    ON categories
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_problems
    BEFORE UPDATE
    ON problems
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_submissions
    BEFORE UPDATE
    ON submissions
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE categories CASCADE;
DROP TABLE images CASCADE;
DROP TABLE problems CASCADE;
DROP TABLE submissions CASCADE;
-- +goose StatementEnd
