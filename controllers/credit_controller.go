package controllers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/cholid97/go-kredit/models"
	"github.com/cholid97/go-kredit/services"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

type UserController struct {
	service services.UserService
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetContracts(ctx *gin.Context) {

	users, err := c.service.GetAllCredits()

	if err != nil || len(users) == 0 {
		resp := Response{
			Status:  "failed",
			Message: "TRX failed",
		}

		if len(users) == 0 {
			resp.Message = "No data found"
		}

		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp := Response{
		Status: "success",
		Data:   users,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := c.service.GetUserByID(uint(id))

	resp := Response{
		Status: "Success",
	}

	if err != nil {
		resp.Message = "User not found"
		resp.Status = "failed"

		ctx.JSON(http.StatusNotFound, resp)
		return
	}

	resp.Data = user

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	resp := Response{
		Status: "Success",
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		resp.Status = "failed"
		resp.Message = "Invalid request"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := c.service.CreateUser(&user); err != nil {
		resp.Status = "failed"
		resp.Message = "TRX failed"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Data = user
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) CreateContract(ctx *gin.Context) {
	var contract models.Contract
	wg.Add(1)
	errChan := make(chan error, 1)
	result := make(chan *models.Contract)

	resp := Response{
		Status: "Success",
	}

	if err := ctx.ShouldBindJSON(&contract); err != nil {
		resp.Status = "failed"
		resp.Message = "Invalid request"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	go func() {
		ctr, err := c.service.CreateContract(&contract)
		errChan <- err
		result <- ctr
	}()

	errs := <-errChan
	ch := <-result

	if errs != nil {
		resp.Status = "failed"
		resp.Message = "TRX failed"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Data = ch
	wg.Done()

	wg.Wait()

	ctx.JSON(http.StatusCreated, resp)
}
