-- +goose Up
-- +goose StatementBegin
CREATE TABLE todo
(
    id          serial PRIMARY KEY       NOT NULL,
    title       varchar                  NOT NULL,
    description text                     NOT NULL,
    completed   boolean                  NOT NULL DEFAULT FALSE,
    created_at  timestamp WITH time zone NOT NULL DEFAULT (NOW()),
    updated_at  timestamp WITH time zone NOT NULL DEFAULT (NOW())
);
-- Added closing parenthesis here
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todo;
-- +goose StatementEnd
