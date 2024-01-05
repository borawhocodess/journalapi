package handlers

import (
    "journalapi/database"
    "gorm.io/gorm"
)

// Handler struct holds required services for handler to function
type Handler struct {
    DB *gorm.DB
}

// NewHandler returns a new Handler with the given services
func NewHandler() *Handler {
    db := database.ConnectDB()
    return &Handler{
        DB: db,
    }
}

// NewHandlerWithDB returns a new Handler with the given database connection
func NewHandlerWithDB(db *gorm.DB) *Handler {
    return &Handler{
        DB: db,
    }
}
