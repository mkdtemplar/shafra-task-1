package api

import (
	"log"
	"net/http"
	"shafra-task-1/internal/models"
	"shafra-task-1/utils"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

var wg sync.WaitGroup

func (s *Server) CreateUser(ctx *gin.Context) {
	var req *models.User
	start := time.Now()

	wg.Add(1)

	chResponse := make(chan *models.User)

	go func() {
		defer wg.Done()
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		checkRequest := V.ValidateFullName(req.NameSurname).ValidateAge(strconv.FormatInt(req.Age, 10))

		if checkRequest.Err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(checkRequest.Err))
			return
		}
		newUser, err := s.store.CreateUser(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		chResponse <- newUser

	}()

	defer close(chResponse)
	ctx.JSON(http.StatusCreated, <-chResponse)
	elapsed := time.Now().Sub(start)
	log.Printf("Request create user took %s", elapsed)

	wg.Wait()
}

func (s *Server) GetUserByID(ctx *gin.Context) {
	var req getUserRequest
	start := time.Now()

	wg.Add(1)

	chResponse := make(chan *models.User)

	go func() {
		defer wg.Done()
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		userById, err := s.store.GetUserById(ctx, req.ID)
		if err != nil {
			ctx.JSON(http.StatusNoContent, errorResponse(err))
			return
		}
		chResponse <- userById
	}()
	defer close(chResponse)
	ctx.JSON(http.StatusOK, <-chResponse)
	elapsed := time.Now().Sub(start)
	log.Printf("Request get user took %s", elapsed)
	wg.Wait()
}

func (s *Server) UpdateUser(ctx *gin.Context) {
	var req getUserRequest
	start := time.Now()
	wg.Add(1)
	chResponse := make(chan *models.User)

	go func() {
		defer wg.Done()
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		_, err := s.store.GetUserById(ctx, req.ID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		userForEdit, err := utils.ParseUserPrefRequestBody(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		userUpdated, err := s.store.UpdateUser(ctx, req.ID, userForEdit.NameSurname, int(userForEdit.Age))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		chResponse <- userUpdated
	}()

	defer close(chResponse)
	ctx.JSON(http.StatusOK, <-chResponse)
	elapsed := time.Now().Sub(start)
	log.Printf("Request get user took %s", elapsed)
	wg.Wait()
}

func (s *Server) DeleteUser(ctx *gin.Context) {
	var req getUserRequest
	start := time.Now()
	wg.Add(1)
	chResponse := make(chan error)

	go func() {
		defer wg.Done()
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		err := s.store.DeleteUser(ctx, req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		chResponse <- err

	}()
	defer close(chResponse)
	ctx.JSON(http.StatusAccepted, gin.H{"error": <-chResponse, "message": "user deleted"})
	elapsed := time.Now().Sub(start)
	log.Printf("Request delete user took %s", elapsed)
	wg.Wait()
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
