package repository

import (
	"context"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/andrersp/test-containers/customer"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestCustomerRepository(t *testing.T) {

	ctx := context.Background()
	pgContainer, err := postgres.Run(
		ctx, "postgres:10-alpine",
		postgres.WithInitScripts(filepath.Join("..", "testdata", "init-db.sql")),
		postgres.WithDatabase("test-data"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		log.Fatal(err)
	}
	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	conStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)
	customerRepo, err := NewRepository(ctx, conStr)

	assert.NoError(t, err)

	customer := customer.Customer{
		Name:  "Henry",
		Email: "henry@mail.com",
	}

	c, err := customerRepo.CreateCustomer(ctx, customer)
	assert.NoError(t, err)
	assert.NotNil(t, c)

	cms, err := customerRepo.GetCustomerByEmail(ctx, "henry@mail.com")
	assert.NoError(t, err)
	assert.NotNil(t, cms)
	assert.Equal(t, cms.Email, "henry@mail.com")

}
