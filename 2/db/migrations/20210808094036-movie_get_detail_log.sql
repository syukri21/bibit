
-- +migrate Up
CREATE TABLE movie_get_detail_logs
(
    id char(36) NOT NULL,
    deleted_at TIME DEFAULT NULL,
    created_at TIME NOT NULL,
    updated_at TIME NOT NULL,
    data jsonb DEFAULT NULL,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE movie_get_detail_logs;
