package api

import (
	"shafra-task-1/internal/database/db"
	"shafra-task-1/internal/database/interfaces"
	"shafra-task-1/internal/validation"
)

type Handler struct {
	DB db.PostgresDB
}

type UserHandler struct {
	Handler
	interfaces.IUserInterface
}

var V validation.Validation
