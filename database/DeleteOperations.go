package database

import (
	"database/sql"
	"fmt"
)

func DeleteDept(dbPointer *sql.DB, id string) error {
	query := "DELETE FROM department WHERE department_name = ?"
	stmt, err := dbPointer.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no department found with id %d", id)
	}
	fmt.Printf("Deleted %d department(s)\n", rowsAffected)
	return nil
}
