package database

import (
	"database/sql"
	"fmt"
	"reflect"
)

func GenericUpdate(dbPointer *sql.DB, tableName string, data interface{}) error {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return fmt.Errorf("data must be a non-nil pointer to a struct")
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("data must be a pointer to a struct")
	}

	var setClauses []string
	var values []interface{}
	var idValue interface{}
	var idColumn string

	// Iterate over fields to build SET clauses and get ID
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		columnName := field.Tag.Get("db") // Use struct tags for column names
		if columnName == "" {
			columnName = field.Name // Fallback to field name if no tag is present
		}

		fieldValue := val.Field(i)

		if i == 0 { // Assuming first field is the ID
			idColumn = columnName
			idValue = fieldValue.Interface()
			continue // Skip adding ID to SET clauses
		}

		setClauses = append(setClauses, fmt.Sprintf("%s = ?", columnName))
		values = append(values, fieldValue.Interface())
	}

	values = append(values, idValue) // Add ID value for WHERE clause

	query := fmt.Sprintf("UPDATE `%s` SET %s WHERE %s = ?",
		tableName,
		sqlJoin(setClauses),
		idColumn,
	)

	stmt, err := dbPointer.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		fmt.Printf("Error came: %v\n", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("Updated %d record(s) in table %s\n", rowsAffected, tableName)
	return nil
}
