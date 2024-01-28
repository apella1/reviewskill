package config

import "reviewskill/internal/database"

type ApiConfig struct {
	DB *database.Queries
}
