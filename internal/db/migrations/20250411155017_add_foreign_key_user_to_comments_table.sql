-- +goose Up
-- +goose StatementBegin
ALTER TABLE comments
ADD CONSTRAINT fk_comments_user
FOREIGN KEY (user_id) REFERENCES users(id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE comments
DROP FOREIGN KEY fk_comments_user;
-- +goose StatementEnd
