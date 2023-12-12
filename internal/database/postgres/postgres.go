package postgres

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // driver for postgres
	"github.com/jmoiron/sqlx"
	"github.com/max-bazarov/chat/internal/config"
)

const (
	UsersTable      = "users"
	TodoListsTable  = "todo_lists"
	UsersListsTable = "users_lists"
	TodoItemsTable  = "todo_items"
	ListsItemsTable = "lists_items"
)

func NewPostgresDB(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.DB_port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	return db, nil
}
