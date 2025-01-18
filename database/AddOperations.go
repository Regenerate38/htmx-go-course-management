package database

import (
	"database/sql"
	"fmt"
	"htmx-go-course-management/model"
	"reflect"
	"strings"
)

func GenericAdd(dbPointer *sql.DB, tableName string, data interface{}) error {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return fmt.Errorf("data must be a non-nil pointer to a struct")
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("data must be a pointer to a struct")
	}

	var columns []string
	var placeholders []string
	var values []interface{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		columnName := field.Tag.Get("db") // Use struct tags for column names
		if columnName == "" {
			columnName = field.Name // Fallback to field name if no tag is present
		}

		columns = append(columns, columnName)
		placeholders = append(placeholders, "?")
		values = append(values, val.Field(i).Interface())
	}

	query := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)",
		tableName,
		sqlJoin(columns),
		sqlJoin(placeholders),
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
	fmt.Printf("Inserted %d record(s) into table %s\n", rowsAffected, tableName)
	return nil
}

func sqlJoin(items []string) string {
	return strings.Join(items, ", ")
}

func InsertDepartment(dbPointer *sql.DB, department *model.Department) error {
	return GenericAdd(dbPointer, "department", department)
}
