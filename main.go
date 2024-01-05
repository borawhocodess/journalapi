package main

import (
	"journalapi/database"
    "journalapi/handlers"
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
    _ "journalapi/docs" // Swagger documentation
)

// @title JournalAPI Documentation
// @description This is a sample server for a journal application.
// @version 1.0
// @host localhost:8080
// @BasePath /

func main() {
    router := gin.Default()
	db := database.ConnectDB()
    handler := handlers.NewHandlerWithDB(db)

    api := router.Group("/api")
    {
        users := api.Group("/users")
        {
            users.POST("/", handler.CreateUser)
            users.GET("/", handler.GetUsers)
            users.GET("/:id", handler.GetUser)
            users.PUT("/:id", handler.UpdateUser)
            users.DELETE("/:id", handler.DeleteUser)
        }

        entries := api.Group("/entries")
        {
            entries.POST("/", handler.CreateEntry)
            entries.GET("/", handler.GetAllEntries)
            entries.GET("/:id", handler.GetEntry)
            entries.PUT("/:id", handler.UpdateEntry)
            entries.DELETE("/:id", handler.DeleteEntry)
        }

        journals := api.Group("/journals")
        {
            journals.POST("/", handler.CreateJournal)
            journals.GET("/", handler.GetAllJournals)
            journals.GET("/:userid/:entryid", handler.GetJournal)
            journals.DELETE("/:userid/:entryid", handler.DeleteJournal)
        }
    }

    // Swagger route
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Start server
    router.Run(":8080")
}
