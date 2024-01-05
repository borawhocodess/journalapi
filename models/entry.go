package models

// Entry represents the entry structure.
type Entry struct {
    ID   uint   `gorm:"primary_key"`
    Text string `gorm:"type:text;not null"`
}
