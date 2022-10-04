package model

type Person struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null;type:varchar(50)" json:"first_name"`
	LastName  string `gorm:"not null;type:varchar(50)" json:"last_name"`
}
