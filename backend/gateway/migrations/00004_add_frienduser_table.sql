-- +goose Up
-- +goose StatementBegin

CREATE TABLE frienduser (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user1_id INTEGER NOT NULL REFERENCES userprofile(id) ON DELETE CASCADE,
    user2_id INTEGER NOT NULL REFERENCES userprofile(id) ON DELETE CASCADE,
    friend_user1 BOOLEAN NOT NULL DEFAULT false,
    friend_user2 BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

-- Уникальность пары и проверка порядка
    UNIQUE (user1_id, user2_id),
    CHECK (user1_id < user2_id)
);

-- Индексы
CREATE INDEX idx_frienduser_user1 ON frienduser(user1_id);
CREATE INDEX idx_frienduser_user2 ON frienduser(user2_id);
CREATE INDEX idx_frienduser_created ON frienduser(created_at);
CREATE INDEX idx_frienduser_friendship ON frienduser(user1_id, user2_id)
    WHERE friend_user1 AND friend_user2;


COMMENT ON TABLE frienduser IS 'Friendship relationships between users';
COMMENT ON COLUMN frienduser.user1_id IS 'Reference to first user in relationship';
COMMENT ON COLUMN frienduser.user2_id IS 'Reference to second user in relationship';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS frienduser;
-- +goose StatementEnd