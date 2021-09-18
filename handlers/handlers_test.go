package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/listele/database"
	"net/http"
	"testing"
)

type testRepository struct {
	ReturnErrorOnFindAll bool
}

func (g *testRepository) FindAll() ([]database.Person, error) {
	if g.ReturnErrorOnFindAll {
		return nil, errors.New("some error occurred")
	}

	return []database.Person{
		{
			FullName: "John Doe",
		},
	}, nil
}

func TestHandleList(t *testing.T) {
	t.Run("it should list records", func(t *testing.T) {
		testRepo := &testRepository{ReturnErrorOnFindAll: false}

		app := fiber.New()
		app.Get("/", HandleList(testRepo))

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		require.Nil(t, err, "unexpected to get an error")

		res, err := app.Test(req, -1)
		assert.Nil(t, err, "unexpected to get an error")

		assert.Equalf(t, http.StatusOK, res.StatusCode, "expected to get status ok but got=%d", res.StatusCode)
		assert.Equal(t, fiber.MIMEApplicationJSON, res.Header.Get("Content-Type"))
	})

	t.Run("it should return internal server error if anything goes wrong", func(t *testing.T) {
		testRepo := &testRepository{ReturnErrorOnFindAll: true}

		app := fiber.New()
		app.Get("/", HandleList(testRepo))

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		require.Nil(t, err, "unexpected to get an error")

		res, err := app.Test(req, -1)
		assert.Nil(t, err, "unexpected to get an error")

		assert.Equalf(t, http.StatusInternalServerError, res.StatusCode, "expected to get internal server error but got=%d", res.StatusCode)
	})
}
