package utils

import (
	"encoding/json"
	"io"
	"shafra-task-1/internal/models"

	"github.com/gin-gonic/gin"
)

func ParseUserPrefRequestBody(ctx *gin.Context) (*models.User, error) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}
