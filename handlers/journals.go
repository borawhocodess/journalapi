package handlers

import (
    "journalapi/models"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// CreateJournal godoc
// @Summary Add a new journal entry
// @Description add by json journal
// @Tags journals
// @Accept  json
// @Produce  json
// @Param   journal   body    models.Journal   true  "Add Journal Entry"
// @Success 200 {object} models.Journal
// @Router /journals [post]
func (h *Handler) CreateJournal(c *gin.Context) {
    var journal models.Journal
    if err := c.ShouldBindJSON(&journal); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if result := h.DB.Create(&journal); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, journal)
}

// GetJournal godoc
// @Summary Get a journal entry
// @Description get a journal entry by UserID and EntryID
// @Tags journals
// @Accept  json
// @Produce  json
// @Param   userid   path    int  true  "User ID"
// @Param   entryid  path    int  true  "Entry ID"
// @Success 200 {object} models.Journal
// @Router /journals/{userid}/{entryid} [get]
func (h *Handler) GetJournal(c *gin.Context) {
    userID, _ := strconv.Atoi(c.Param("userid"))
    entryID, _ := strconv.Atoi(c.Param("entryid"))
    var journal models.Journal

    if result := h.DB.Where("user_id = ? AND entry_id = ?", userID, entryID).First(&journal); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Journal not found"})
        return
    }

    c.JSON(http.StatusOK, journal)
}

// GetAllJournals godoc
// @Summary Get all journal entries
// @Description get all journal entries
// @Tags journals
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Journal
// @Router /journals [get]
func (h *Handler) GetAllJournals(c *gin.Context) {
    var journals []models.Journal
    if result := h.DB.Find(&journals); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving journals"})
        return
    }

    c.JSON(http.StatusOK, journals)
}

// DeleteJournal godoc
// @Summary Delete a journal entry
// @Description delete a journal entry by UserID and EntryID
// @Tags journals
// @Accept  json
// @Produce  json
// @Param   userid   path    int  true  "User ID"
// @Param   entryid  path    int  true  "Entry ID"
// @Success 204 "No Content"
// @Router /journals/{userid}/{entryid} [delete]
func (h *Handler) DeleteJournal(c *gin.Context) {
    userID, _ := strconv.Atoi(c.Param("userid"))
    entryID, _ := strconv.Atoi(c.Param("entryid"))

    if result := h.DB.Where("user_id = ? AND entry_id = ?", userID, entryID).Delete(&models.Journal{}); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting journal"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Journal deleted successfully"})
}
