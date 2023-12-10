package controllers

import (
	"epseed/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type NoteController struct {
	NoteService *services.NoteService
}

// GET /notes
func (controller *NoteController) GetNotes(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.JSON(405, gin.H{"message": "Method not allowed"})
		return
	}

	userId := c.Query("user_id")
	if userId == "" {
		c.JSON(400, gin.H{"message": "Bad request"})
		return
	}

	id := 0
	fmt.Sscanf(userId, "%d", &id)

	notes, err := controller.NoteService.GetNotesForUser(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Notes not found"})
		return
	}

	c.JSON(200, notes)
}

// POST /notes
func (controller *NoteController) CreateNote(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.JSON(405, gin.H{"message": "Method not allowed"})
		return
	}

	var json struct {
		UserId  int    `json:"user_id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
		return
	}

	err := controller.NoteService.CreateNoteForUser(json.UserId, json.Title, json.Content)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error while creating note"})
		return
	}

	c.JSON(200, gin.H{"message": "Note created successfully", "user_id": json.UserId})
}

// PUT /notes
func (controller *NoteController) UpdateNote(c *gin.Context) {
	if c.Request.Method != "PUT" {
		c.JSON(405, gin.H{"message": "Method not allowed"})
		return
	}

	var json struct {
		UserId  int    `json:"user_id"`
		NoteId  int    `json:"note_id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
		return
	}

	err := controller.NoteService.UpdateNoteForUser(json.UserId, json.NoteId, json.Title, json.Content)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error while updating note"})
		return
	}

	c.JSON(200, gin.H{"message": "Note updated successfully", "user_id": json.UserId})
}

// DELETE /notes
func (controller *NoteController) DeleteNote(c *gin.Context) {
	if c.Request.Method != "DELETE" {
		c.JSON(405, gin.H{"message": "Method not allowed"})
		return
	}

	var json struct {
		UserId int `json:"user_id"`
		NoteId int `json:"note_id"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
		return
	}

	err := controller.NoteService.DeleteNoteForUser(json.UserId, json.NoteId)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error while deleting note"})
		return
	}

	c.JSON(200, gin.H{"message": "Note deleted successfully", "user_id": json.UserId})
}
