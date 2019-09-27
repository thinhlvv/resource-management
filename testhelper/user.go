package testhelper

import "database/sql"

// RemoveUser is function to help remove testing user.
func RemoveUser(db *sql.DB, userID int, email string) error {
	stmt, err := db.Prepare(`
		DELETE FROM user WHERE id = ? OR email = ?
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID, email)
	if err != nil {
		return err
	}

	return nil
}
