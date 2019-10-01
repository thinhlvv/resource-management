package testhelper

import "database/sql"

// RemoveResource is function to help remove testing user.
func RemoveResource(db *sql.DB, id int, name string) error {
	stmt, err := db.Prepare(`
		DELETE FROM resource WHERE id = ? OR name = ?
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, name)
	if err != nil {
		return err
	}

	return nil
}
