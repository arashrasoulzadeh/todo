package providers

import (
	"awesomeProject/internal/repositories"
	"awesomeProject/internal/requests"
	"awesomeProject/internal/responses"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func NewRouter(engine *gin.Engine, db *repositories.TaskDatabase, log *zap.Logger) {
	engine.GET("/", func(c *gin.Context) {

		list := db.TaskList()
		tasks := make([]responses.Task, len(list))

		for index, t := range list {
			tasks[index].Fill(t)
		}
		c.JSON(http.StatusOK, gin.H{
			"items": tasks,
		})
	})
	engine.POST("/", func(c *gin.Context) {
		var task requests.Task
		if err := c.ShouldBindJSON(&task); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
			log.Error(err.Error())
			return
		}

		createdTask := db.CreateTask(task.Title, task.Desc)
		var response responses.Task
		response.Fill(createdTask)
		c.JSON(http.StatusOK, gin.H{
			"item": response,
		})
	})
	engine.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		var task requests.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
			return
		}

		updatedTask, err := db.UpdateTask(id, task.Title, task.Desc, task.Status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Error(err.Error())
			return
		}
		var response responses.Task
		response.Fill(*updatedTask)

		c.JSON(http.StatusBadRequest, gin.H{"item": response})

	})
	go func() {
		err := engine.Run()
		if err != nil {
			panic("Can't run web server," + err.Error())
		}
	}()
}
