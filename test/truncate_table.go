package test

import "gorm.io/gorm"

func TruncateTableUsers(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE USERS")
}
