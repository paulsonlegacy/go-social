-- +goose Up
-- +goose StatementBegin
ALTER TABLE comments
ADD CONSTRAINT fk_comments_post
FOREIGN KEY (post_id) REFERENCES posts(id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE comments
DROP FOREIGN KEY fk_comments_post;
-- +goose StatementEnd
