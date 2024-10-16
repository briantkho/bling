package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func GetUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, User{
		UserID:    uuid.UUID{},
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@gmail.com",
		Password:  "SALT_PW",
		CreatedAt: time.Time{},
		LevelID:   1,
	})
}
