package main

import (
	"go-api-structure/data"
	"go-api-structure/infra"
	"go-api-structure/inputs"
	"go-api-structure/model"
	"go-api-structure/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type API struct {
	CreateService *service.CreateTodoService
	RemoveService *service.RemoveTodoService
	ToggleService *service.ToggleTodoService
	Query         data.TodoQueryRepository
}

func (api *API) Create(c *gin.Context) {
	var todoInput inputs.CreateTodoInput

	if err := c.BindJSON(&todoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if todo, err := api.CreateService.Execute(todoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id": todo.ID})
	}
}

func (api *API) GetAll(c *gin.Context) {
	if list, err := api.Query.GetAll(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, list)
	}
}

func (api *API) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if todo, err := api.Query.GetByID(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func (api *API) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := api.RemoveService.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}

func (api *API) Toggle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := api.ToggleService.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}

func main() {
	conn := make(map[int]model.Todo)
	repo := infra.NewInMemoryTodoRepository(&conn)
	createService := service.NewCreateTodoService(repo)
	removeService := service.NewRemoveTodoService(repo)
	toggleService := service.NewToggleTodoService(repo)
	query := infra.NewQueryRepository(&conn)

	api := &API{
		CreateService: createService,
		RemoveService: removeService,
		ToggleService: toggleService,
		Query:         query,
	}

	router := gin.Default()

	router.POST("/todos", api.Create)
	router.GET("/todos", api.GetAll)
	router.GET("/todos/:id", api.GetById)
	router.DELETE("/todos/:id", api.Delete)
	router.PUT("/todos/:id", api.Toggle)

	router.Run("0.0.0.0:8080")
}
