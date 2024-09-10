package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	db "github.com/tetrex/golang-project-template/db/sqlc"
	_ "github.com/tetrex/golang-project-template/docs" // docs is generated by Swag CLI
	"github.com/tetrex/golang-project-template/utils/config"
	"github.com/tetrex/golang-project-template/utils/logger"
)

// @title			server api
// @version			1.0
// @description		This is a backend api server
// @contact.name	github.com/tetrex
// @license.name	MIT License
//
// @host			localhost:8000
// @basePath		/
func main() {
	l := logger.New(true)

	// load config
	config, err := config.LoadConfig()
	if err != nil {
		l.Fatal().Err(errors.Errorf("cannot load config "))
		l.Fatal().Err(err)
	}

	// pg connection
	db_config, err := pgxpool.ParseConfig(config.PgConnStr)
	if err != nil {
		l.Fatal().Err(errors.Errorf("cannot connect to to db"))
	}

	db_config.MaxConns = 20                     // Maximum number of connections in the pool
	db_config.MaxConnLifetime = 5 * time.Minute // Maximum lifetime of a connection
	db_config.MaxConnIdleTime = 2 * time.Minute // Maximum time a connection can remain idle

	// ------------
	db_pool, err := pgxpool.NewWithConfig(context.Background(), db_config)
	if err != nil {
		l.Fatal().Err(errors.Errorf("cannot connect to db db_pool"))
	}
	defer db_pool.Close()
	queries := db.New(db_pool)

	s, err := NewServer(ServerParams{
		Config:  config,
		Logger:  l,
		Queries: queries,
	})
	if err != nil {
		l.Fatal().Err(errors.Errorf("cannot create new server"))
	}

}
