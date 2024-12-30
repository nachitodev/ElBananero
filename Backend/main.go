package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ElBananero/schema"
)

func main() {
	var db *gorm.DB

	r := gin.Default()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&schema.Video{}, &schema.Comentario{})

	r.GET("/videos", func(c *gin.Context) {
		var videos []schema.Video
		result := db.Find(&videos)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  result.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   videos,
		})
	})

	r.GET("/videos/:id", func(c *gin.Context) {
		var videos schema.Video
		result := db.First(&videos, c.Param("id"))

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   videos,
		})
	})

	r.GET("/videos/:id/comentarios", func(c *gin.Context) {
		var comentarios []schema.Comentario
		result := db.Where("video_id = ?", c.Param("id")).Find(&comentarios)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   comentarios,
		})
	})

	r.POST("/videos", func(c *gin.Context) {
		db.Create(&schema.Video{})
		c.JSON(200, gin.H{
			"status": 200,
		})
	})
	r.Run()
}
