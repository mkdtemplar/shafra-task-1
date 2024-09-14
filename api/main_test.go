package api

import (
	"os"
	"shafra-task-1/internal/database/interfaces"
	"shafra-task-1/utils"
	"testing"

	"github.com/gin-gonic/gin"
)

func newTestServer(t *testing.T, store interfaces.IUserInterface) *Server {
	config, err := utils.LoadConfig("../../..")
	if err != nil {
		t.Fatal(err)
	}

	server, err := NewServer(config, store)
	if err != nil {
		t.Fatal(err)
	}
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
