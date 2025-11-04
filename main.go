package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/posts", func(c *gin.Context) {
		// Получить список постов для отображения в ленте.
	})

	r.POST("/posts", func(c *gin.Context) {
		// Опубликовать новый пост в ленте
	})

	r.DELETE("/posts/:id", func(c *gin.Context) {
		// Удалить указанный пост
	})
	r.PATCH("/posts/:id", func(c *gin.Context) {
		// частично изменить указанный пост
	})
	r.GET("/posts/:id/comments", func(c *gin.Context) {
		// получить список комментариев к указанному посту
	})
	r.DELETE("/posts/:id/comments/:id", func(c *gin.Context) {
		// удалить указанный комментарий указанного поста
	})
	r.PUT("/posts/:id/comments/:id", func(c *gin.Context) {
		// изменить указанный комментарий указанного поста
	})
	r.POST("/posts/:id/like", func(c *gin.Context) {
		// поставить лайк на указанный пост
	})
	r.POST("/posts/:id/comments/:id/like", func(c *gin.Context) {
		// поставить лайк на указанный комментарий указанного поста
	})
	r.POST("/direct/User/:id/messsage", func(c *gin.Context) {
		// отправить сообщение в директ указанному юзеру
	})
	r.DELETE("/direct/messages/:id", func(c *gin.Context) {
		// удалить указанное сообщение из директа
	})
	r.DELETE("/direct/chats/:id", func(c *gin.Context) {
		// удалить указанный чат из директа
	})

	// тут должны быть заготовки роутов как это сделано выше
	r.Run()
}
