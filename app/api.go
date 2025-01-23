package main

import (
	"database/sql"
	"go-api-structure/data"
	"go-api-structure/infra"
	"go-api-structure/inputs"
	"go-api-structure/service"
	"log"
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

func Transactional(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx, err := db.Begin()
		if err != nil {
			log.Println("Failed to start transaction:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
			return
		}
		c.Next()
		if len(c.Errors) > 0 {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Println("Failed to rollback transaction:", rollbackErr)
			}
		} else {
			if commitErr := tx.Commit(); commitErr != nil {
				log.Println("Failed to commit transaction:", commitErr)
			}
		}
	}
}

func main() {
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	repo := infra.NewSQLTodoRepository(db)
	query := infra.NewSQLQueryRepository(db)

	createService := service.NewCreateTodoService(repo)
	removeService := service.NewRemoveTodoService(repo)
	toggleService := service.NewToggleTodoService(repo)

	api := &API{
		CreateService: createService,
		RemoveService: removeService,
		ToggleService: toggleService,
		Query:         query,
	}

	router := gin.Default()
	router.Use(Transactional(db))

	router.POST("/todos", api.Create)
	router.GET("/todos", api.GetAll)
	router.GET("/todos/:id", api.GetById)
	router.DELETE("/todos/:id", api.Delete)
	router.PATCH("/todos/:id/toggle", api.Toggle)

	router.Run("0.0.0.0:8080")
}
