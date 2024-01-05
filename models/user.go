package models

type User struct {
    ID          uint   `gorm:"primary_key" json:"id"`
    PhoneNumber string `gorm:"unique;not null" json:"phoneNumber"`
    Name        string `gorm:"not null" json:"name"`
    Birthday    string `gorm:"not null" json:"birthday"`
}
