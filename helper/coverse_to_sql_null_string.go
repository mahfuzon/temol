package helper

import "database/sql"

func ConverseToSqlNullString(value string) sql.NullString {
	isValid := true
	if value == "" {
		isValid = false
	}
	return sql.NullString{
		String: value,
		Valid:  isValid,
	}
}
