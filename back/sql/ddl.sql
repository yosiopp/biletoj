CREATE TABLE tickets (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    tags TEXT,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL
);
CREATE INDEX tickets_at_idx ON tickets (created_at);
CREATE INDEX tickets_by_idx ON tickets (created_by);

CREATE TABLE ticket_histories (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    ticket_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    tags TEXT,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    ticket_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL
);
CREATE INDEX comments_idx ON comments (ticket_id, created_at);

CREATE TABLE comment_histories (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    comment_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tag_catalog (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    tag VARCHAR(255) NOT NULL UNIQUE,
    note VARCHAR(255),
    color VARCHAR(40),
    is_group INTEGER NOT NULL DEFAULT 0
    is_range INTEGER NOT NULL DEFAULT 0
);

create virtual table tickets_fts using fts5 (
    ticket_id unindexed,
    title, -- bigram
    content, -- bigram
    tags,
    comments, -- bigram
    tokenize="unicode61 remove_diacritics 2"
);
