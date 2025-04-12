-- +goose Up
-- +goose StatementBegin
ALTER TABLE comments
ADD CONSTRAINT fk_comments_parent
FOREIGN KEY (parent_id) REFERENCES comments(id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE comments
DROP FOREIGN KEY fk_comments_parent;
-- +goose StatementEnd
