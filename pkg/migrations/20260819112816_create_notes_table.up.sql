CREATE TABLE IF NOT EXISTS notes
(
    id serial Primary key,
    title varchar(128) NOT NULL DEFAULT '',
    content text,
    tags text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz
);