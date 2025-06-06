package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Email     string    `json:"email" gorm:"uniqueIndex" validate:"required,email"`
	Password  string    `json:"-" validate:"required,min=6"` // hashed password, omit json output
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
