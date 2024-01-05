package model

import (
	"time"
)

type Meme struct {
	ID          uint      `JSON:"id"`
	DriveId     string    `JSON:"driveId"`
	Name        string    `JSON:"name"`
	Description string    `JSON:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
