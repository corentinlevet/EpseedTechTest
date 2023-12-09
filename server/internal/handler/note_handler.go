package handler

import (
	"encoding/json"
	"epseed/internal/db"
	"fmt"
	"net/http"
	"strconv"
)

type CreateNoteRequest struct {
	UserId  uint   `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Message string `json:"message"`
	NoteId  uint   `json:"note_id"`
}

// POST /notes/create
func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var createNoteRequest CreateNoteRequest
	err := json.NewDecoder(r.Body).Decode(&createNoteRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.CreateNote(createNoteRequest.UserId, createNoteRequest.Title, createNoteRequest.Content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	noteId, err := db.GetLastNoteId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(Response{
		Message: "Note créée avec succès",
		NoteId:  noteId,
	})

	w.Write(returnJson)
}

type DeleteNoteRequest struct {
	UserId uint `json:"user_id"`
	NoteId uint `json:"note_id"`
}

// DELETE /notes/delete
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var deleteNoteRequest DeleteNoteRequest
	err := json.NewDecoder(r.Body).Decode(&deleteNoteRequest)

	if err != nil {
		fmt.Println("Erreur lors du décodage du body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.DeleteNoteForUser(deleteNoteRequest.UserId, deleteNoteRequest.NoteId)
	if err != nil {
		fmt.Println("Erreur lors de la suppression de la note:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(Response{
		Message: "Note supprimée avec succès",
	})

	w.Write(returnJson)
}

// GET /notes/get
func GetNotesForUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseUint(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	notes, err := db.GetNotesForUser(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(notes)

	w.Write(returnJson)
}

type UpdateNoteRequest struct {
	UserId  uint   `json:"user_id"`
	NoteId  uint   `json:"note_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// PUT /notes/update
func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var updateNoteRequest UpdateNoteRequest
	err := json.NewDecoder(r.Body).Decode(&updateNoteRequest)

	if err != nil {
		fmt.Println("Erreur lors du décodage du body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.UpdateNoteForUser(updateNoteRequest.UserId, updateNoteRequest.NoteId, updateNoteRequest.Title, updateNoteRequest.Content)
	if err != nil {
		fmt.Println("Erreur lors de la mise à jour de la note:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(Response{
		Message: "Note mise à jour avec succès",
		NoteId:  updateNoteRequest.NoteId,
	})

	w.Write(returnJson)
}
