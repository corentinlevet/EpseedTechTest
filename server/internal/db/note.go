package db

import "time"

type Note struct {
	ID        uint      `gorm:"primaryKey"`
	UserId    uint      `gorm:"not null"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func CreateNote(UserId uint, Title string, Content string) error {
	note := Note{UserId: UserId, Title: Title, Content: Content}
	result := DbInstance.Create(&note)
	return result.Error
}

func UpdateNoteForUser(UserId uint, NoteId uint, Title string, Content string) error {
	var note Note
	result := DbInstance.Where("user_id = ? AND id = ?", UserId, NoteId).First(&note)
	if result.Error != nil {
		return result.Error
	}
	note.Title = Title
	note.Content = Content
	result = DbInstance.Save(&note)
	return result.Error
}

func DeleteNoteForUser(UserId uint, NoteId uint) error {
	var note Note
	result := DbInstance.Where("user_id = ? AND id = ?", UserId, NoteId).First(&note)
	if result.Error != nil {
		return result.Error
	}
	result = DbInstance.Delete(&note)
	return result.Error
}

func GetNotesForUser(UserId uint) ([]Note, error) {
	var notes []Note
	result := DbInstance.Where("user_id = ?", UserId).Find(&notes)
	return notes, result.Error
}

func GetLastNoteId() (uint, error) {
	var note Note
	result := DbInstance.Last(&note)
	return note.ID, result.Error
}
