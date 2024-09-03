package scripts

func CreateTables() string {
	sql := `
		CREATE TABLE IF NOT EXISTS posts(
			post_key serial primary key,
			title TEXT,
			content TEXT,
			created_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP 
		)`

	return sql
}

func SeedTables() string {
	return `
		INSERT INTO posts(title, content)
		VALUES ('First post', 'First post content')`
}

func DropTables() string {
	sql := `
		DROP TABLE posts
	`

	return sql
}
