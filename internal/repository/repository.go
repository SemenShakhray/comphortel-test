package repository

import (
	"database/sql"
	"fmt"
	"log/slog"

	"comphortel-test/internal/config"
	"comphortel-test/internal/models"
	"comphortel-test/internal/service"
	"comphortel-test/utils/logger/sl"

	_ "github.com/lib/pq"
)

type Repository struct {
	DB  *sql.DB
	log *slog.Logger
}

func NewRepository(cfg *config.Config, log *slog.Logger) (service.Repositorer, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error("failed to create database connection pool", sl.Err(err))

		return nil, fmt.Errorf("failed to create database connection pool: %w", err)
	}

	if err := db.Ping(); err != nil {
		log.Error("failed to ping database", sl.Err(err))

		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Repository{
		DB:  db,
		log: log,
	}, nil
}

func (r *Repository) Close() error {
	return r.DB.Close()
}

func (r *Repository) GetUser(id int) (*models.User, error) {
	var user models.User

	query := `SELECT id, 
	login, 
	full_name, 
	gender, 
	age, 
	contacts,
	avatar_url,
	created_at,
	active 
	FROM users WHERE id = $1`

	err := r.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Login,
		&user.FullName,
		&user.Gender,
		&user.Age,
		&user.Contacts,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.Active,
	)
	if err != nil {
		r.log.Error("failed to get user", sl.Err(err))

		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}
