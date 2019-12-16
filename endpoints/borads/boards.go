package borads

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"time"
	"trello_clone_backend/models"
)

func List(c echo.Context) error {
	model := []models.Board{
		{
			Model:     gorm.Model{ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: nil},
			UserId:    1,
			Title:     "a",
			DispOrder: 1,
		},
		{
			Model:     gorm.Model{ID: 2, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: nil},
			UserId:    1,
			Title:     "b",
			DispOrder: 2,
		},
		{
			Model:     gorm.Model{ID: 3, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: nil},
			UserId:    1,
			Title:     "c",
			DispOrder: 3,
		},
	}

	return c.JSON(http.StatusOK, model)
}
