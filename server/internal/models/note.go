package models

import (
	"epseed/internal/db"
	"time"
)

type Note struct {
	ID        int       `gorm:"primaryKey"`
	UserId    int       `gorm:"not null"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func GetNoteById(id int) (*Note, error) {
	var note Note
	err := db.DbInstance.Where("id = ?", id).First(&note).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func GetNotesByUserId(userId int) ([]*Note, error) {
	var notes []*Note
	err := db.DbInstance.Where("user_id = ?", userId).Find(&notes).Error
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func CreateNoteForUser(userId int, title string, content string) error {
	note := Note{
		UserId:  userId,
		Title:   title,
		Content: content,
	}
	err := db.DbInstance.Create(&note).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateNoteForUser(userId int, noteId int, title string, content string) error {
	note := Note{
		Title:   title,
		Content: content,
	}
	err := db.DbInstance.Model(&note).Where("id = ? AND user_id = ?", noteId, userId).Updates(&note).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteNoteForUser(userId int, noteId int) error {
	err := db.DbInstance.Where("id = ? AND user_id = ?", noteId, userId).Delete(&Note{}).Error
	if err != nil {
		return err
	}
	return nil
}
