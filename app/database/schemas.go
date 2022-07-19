package database

const createScheme = `
CREATE TABLE IF NOT EXISTS posts
(
	id SERIAL PRIMARY KEY,
	title TEXT,
	content TEXT,
	author TEXT
)
`

var insertPostSchema = `
INSERT INTO posts(title, content, author) VALUES ($1, $2, $3) RETURNING id
`
