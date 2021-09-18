package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var db *sql.DB

func TestMain(m *testing.M) {
	// Only run if RUN_INTEGRATION_TESTS is YES
	if os.Getenv("RUN_INTEGRATION_TESTS") != "YES" {
		os.Exit(0)
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	user := "myexampleuser"
	password := "myexample"
	dbName := "listele"

	resource, err := pool.Run("postgres", "13.4-alpine", []string{
		fmt.Sprintf("POSTGRES_PASSWORD=%s", password),
		fmt.Sprintf("POSTGRES_USER=%s", user),
		fmt.Sprintf("POSTGRES_DB=%s", dbName),
	})
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	if err = pool.Retry(func() error {
		var errOpenConn error

		db, errOpenConn = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", user, password, resource.GetPort("5432/tcp"), dbName))
		if errOpenConn != nil {
			return errOpenConn
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("unable to initialize driver due to=%s", err)
	}

	mi, err := migrate.NewWithDatabaseInstance(
		"file://../db/migrations/", dbName, driver,
	)
	if err != nil {
		log.Fatalf("unable to initialize migrator due to=%s", err)
	}

	err = mi.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("unable to run migrations due to=%s", err)
	}

	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestPeopleRepository_FindAll(t *testing.T) {
	repo := PeopleRepository{Database: db}

	people, err := repo.FindAll()

	require.Nil(t, err, "unexpected to get an error at this step")
	assert.Equal(t, 4, len(people))

	var names []string

	for _, person := range people {
		names = append(names, person.FullName)
	}

	assert.Equal(t, []string{"John Doe", "Aida Bugg", "Maureen Biologist", "Allie Grater"}, names)
}
