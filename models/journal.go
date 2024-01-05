package models

import (
    "gorm.io/gorm"
)

// Journal represents the journal structure.
type Journal struct {
    UserID  uint `gorm:"primary_key"`
    EntryID uint `gorm:"primary_key"`
}

// BeforeCreate is a GORM hook that is called before creating a journal record.
// This can be used to implement any pre-insert logic.
func (journal *Journal) BeforeCreate(tx *gorm.DB) (err error) {
    // Custom logic before creating a journal entry
    return
}
