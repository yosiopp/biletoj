package data

const (
	// 初期化 (DDL)
	_SQL_INIT = `CREATE TABLE tickets (
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
		is_group INTEGER NOT NULL DEFAULT 0,
		is_range INTEGER NOT NULL DEFAULT 0
	);
	
	CREATE VIRTUAL TABLE tickets_fts USING fts5 (
		ticket_id UNINDEXED,
		title,
		content,
		tags,
		comments,
		tokenize="unicode61 remove_diacritics 2"
	);
	`

	// 初期データ投入（タグカタログ）
	_SQL_INIT_TAG_CATALOG = `INSERT INTO tag_catalog (tag, note, color, is_group, is_range) VALUES 
		('status:OPEN', '未処理', '#ff0000', 1, 0),
		('status:WIP', '処理中', '#ff0000', 1, 0),
		('status:DONE', '処理済', '#ff0000', 1, 0),
		('status:CLOSE', '完了', '#ff0000', 1, 0),
		('type:ISSUE', '課題', '#00ff00', 1, 0),
		('type:TASK', 'タスク', '#00ff00', 1, 0),
		('type:BUG', 'バグ', '#00ff00', 1, 0),
		('type:QUESTION', '質問', '#00ff00', 1, 0),
		('type:NOTE', 'メモ', '#00ff00', 1, 0),
		('due-date@:', '期限', NULL, 1, 1);`

	// タグカタログ追加
	_SQL_ADD_TAG_CATALOG = `INSERT INTO tag_catalog (tag, note, color, is_group, is_range) VALUES (?, ?, ?, ?, ?)`

	// チケット追加
	_SQL_ADD_TICKET = `INSERT INTO tickets (title, content, tags, created_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	_SQL_ADD_TICKET_HISTORY = `INSERT INTO ticket_histories (ticket_id, title, content, tags, created_by, created_at) VALUES (?, ?, ?, ?, ?, ?)`
	_SQL_ADD_TICKET_FTS = `INSERT INTO tickets_fts (ticket_id, title, content, tags, comments) VALUES (?, ?, ?, ?, ?)`

	// チケット編集
	_SQL_EDIT_TICKET = `UPDATE tickets SET title = ?, content = ?, tags = ?, updated_at = ? WHERE id = ?`
	_SQL_EDIT_TICKET_HISTORY = `INSERT INTO ticket_histories (ticket_id, title, content, tags, created_by, created_at) VALUES (?, ?, ?, ?, ?, ?)`
	_SQL_EDIT_TICKET_FTS = `UPDATE tickets_fts SET title = ?, content = ?, tags = ?, comments = ? WHERE ticket_id = ?`

	// コメント追加
	_SQL_ADD_COMMENT = `INSERT INTO comments (ticket_id, content, created_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	_SQL_ADD_COMMENT_HISTORY = `INSERT INTO comment_histories (comment_id, content, created_at) VALUES (?, ?, ?)`
	_SQL_ADD_COMMENT_FTS = `UPDATE tickets_fts SET comments = ? WHERE ticket_id = ?`

	// コメント編集
	_SQL_EDIT_COMMENT = `UPDATE comments SET content = ?, updated_at = ? WHERE id = ?`
	_SQL_EDIT_COMMENT_HISTORY = `INSERT INTO comment_histories (comment_id, content, created_at) VALUES (?, ?, ?)`
	_SQL_EDIT_COMMENT_FTS = `UPDATE tickets_fts SET comments = ? WHERE ticket_id = ?`

	// チケット検索
	_SQL_QUERY_TICKETS = `SELECT t.* FROM tickets t JOIN tickets_fts ft ON t.id = ft.ticket_id WHERE ft.title MATCH ? OR ft.content MATCH ? OR ft.comments MATCH ? OR ft.tags MATCH ? ORDER BY t.created_at DESC`

	// コメント取得
	_SQL_QUERY_COMMENTS = `SELECT * FROM comments WHERE ticket_id = ? ORDER BY created_at ASC`

	// タグ取得
	_SQL_QUERY_TAGS = `SELECT * FROM tag_catalog`
)