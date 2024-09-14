package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"shafra-task-1/internal/database/db"
	"shafra-task-1/internal/models"
	"shafra-task-1/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserHandler_CreateUser(t *testing.T) {
	config, err := utils.LoadConfig("../../..")
	if err != nil {
		t.Fatal(err)
	}
	conn := db.ConnectToPostgres(config)
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewUserRepo(conn)

	testserver := newTestServer(t, store)

	user := &models.User{
		ID:          utils.RandomInt(1, 100),
		NameSurname: utils.RandomString(),
		Age:         utils.RandomInt(20, 100),
	}

	jsonUser, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(jsonUser))
	w := httptest.NewRecorder()
	testserver.Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

}
