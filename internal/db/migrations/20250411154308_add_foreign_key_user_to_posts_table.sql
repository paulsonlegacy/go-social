-- +goose Up
-- +goose StatementBegin
ALTER TABLE posts
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id) REFERENCES users(id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE posts
DROP FOREIGN KEY fk_user_id;
-- +goose StatementEnd
