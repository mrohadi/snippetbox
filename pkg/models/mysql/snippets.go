package mysql

import (
	"database/sql"

	"github.com/mrohadi/snippetbox/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB	
}

// This will be insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// This will return specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// This will return the most 10 recently created snippet
func (m *SnippetModel) Latest() (*models.Snippet, error) {
	return nil, nil
}