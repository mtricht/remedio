package database

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Medication struct {
	gorm.Model
	Name             string
	Time             time.Time
	Supply           int
	LastNotification sql.NullTime
	LastTaken        sql.NullTime
}
