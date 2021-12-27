package database

const createSchema = `
CREATE TABLE IF NOT EXISTS posts
(
	id SERIAL PRIMARY KEY,
	title TEXT,
	content TEXT,
	author TEXT
)
`

const insertPostSchema = `
INSERT INTO posts(title, content, author) VALUES($1, $2, $3) RETURNING id
`

const updatePostSchema = `
UPDATE posts SET title = $2, content = $3, author = $4 WHERE id = $1 RETURNING id
`

const deletePostSchema = `
DELETE FROM posts WHERE id = $1
`
