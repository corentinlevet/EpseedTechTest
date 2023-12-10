package services

import (
	"epseed/internal/models"
)

type NoteService struct{}

func (s *NoteService) GetNotesForUser(userId int) ([]*models.Note, error) {
	notes, err := models.GetNotesByUserId(userId)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *NoteService) CreateNoteForUser(userId int, title string, content string) error {
	err := models.CreateNoteForUser(userId, title, content)
	if err != nil {
		return err
	}

	return nil
}

func (s *NoteService) UpdateNoteForUser(userId int, noteId int, title string, content string) error {
	err := models.UpdateNoteForUser(userId, noteId, title, content)
	if err != nil {
		return err
	}

	return nil
}

func (s *NoteService) DeleteNoteForUser(userId int, noteId int) error {
	err := models.DeleteNoteForUser(userId, noteId)
	if err != nil {
		return err
	}

	return nil
}
