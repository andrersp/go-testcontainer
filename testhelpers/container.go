package testhelpers

import (
	"context"
	"path/filepath"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

func CratePostgresContainer(ctx context.Context) (*PostgresContainer, error) {

	pgContainer, err := postgres.Run(
		ctx, "postgres:10-alpine",
		postgres.WithInitScripts(filepath.Join("..", "testhelpers", "sql", "init-db.sql")),
		postgres.WithDatabase("test-data"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		return nil, err
	}

	conStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  conStr,
	}, nil

}
