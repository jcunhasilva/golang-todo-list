package model

import (
	"time"

	"github.com/go-pg/pg/orm"
)

// BaseModel specifies the common fields accross all database entities
type BaseModel struct {
	CreatedAt time.Time `sql:",notnull"`
	UpdatedAt time.Time `sql:",notnull"`
}

// BeforeInsert default values
func (m *BaseModel) BeforeInsert(db orm.DB) error {
	if m != nil {
		now := time.Now()
		m.CreatedAt = now
		m.UpdatedAt = now
	}

	return nil
}

// BeforeUpdate default values
func (m *BaseModel) BeforeUpdate(db orm.DB) error {
	if m != nil {
		m.UpdatedAt = time.Now()
	}

	return nil
}
