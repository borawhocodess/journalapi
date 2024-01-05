package handlers

import (
    "journalapi/models"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// CreateEntry godoc
// @Summary Add a new entry
// @Description add by json entry
// @Tags entries
// @Accept  json
// @Produce  json
// @Param   entry   body    models.Entry   true  "Add Entry"
// @Success 200 {object} models.Entry
// @Router /entries [post]
func (h *Handler) CreateEntry(c *gin.Context) {
    var entry models.Entry
    if err := c.ShouldBindJSON(&entry); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if result := h.DB.Create(&entry); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, entry)
}

// GetEntry godoc
// @Summary Get an entry
// @Description get an entry by ID
// @Tags entries
// @Accept  json
// @Produce  json
// @Param   id   path    int  true  "Entry ID"
// @Success 200 {object} models.Entry
// @Router /entries/{id} [get]
func (h *Handler) GetEntry(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var entry models.Entry

    if result := h.DB.First(&entry, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
        return
    }

    c.JSON(http.StatusOK, entry)
}

// GetAllEntries godoc
// @Summary Get all entries
// @Description get all entries
// @Tags entries
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Entry
// @Router /entries [get]
func (h *Handler) GetAllEntries(c *gin.Context) {
    var entries []models.Entry
    if result := h.DB.Find(&entries); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving entries"})
        return
    }

    c.JSON(http.StatusOK, entries)
}

// UpdateEntry godoc
// @Summary Update an entry
// @Description update entry's details by ID
// @Tags entries
// @Accept  json
// @Produce  json
// @Param   id   path    int  true  "Entry ID"
// @Param   entry   body    models.Entry   true  "Update Entry"
// @Success 200 {object} models.Entry
// @Router /entries/{id} [put]
func (h *Handler) UpdateEntry(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var entry models.Entry

    if result := h.DB.First(&entry, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
        return
    }

    if err := c.ShouldBindJSON(&entry); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    h.DB.Save(&entry)
    c.JSON(http.StatusOK, entry)
}

// DeleteEntry godoc
// @Summary Delete an entry
// @Description delete an entry by ID
// @Tags entries
// @Accept  json
// @Produce  json
// @Param   id   path    int  true  "Entry ID"
// @Success 204 "No Content"
// @Router /entries/{id} [delete]
func (h *Handler) DeleteEntry(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var entry models.Entry

    if result := h.DB.Delete(&entry, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting entry"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Entry deleted successfully"})
}
