package models

import "example.com/rest-api/database"

func (e *Event) Register(userID int64) error {
	query := `
		INSERT INTO registrations (event_id, user_id)
		VALUES (?, ?)
	`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userID)

	if err != nil {
		return err
	}

	return err
}

func (e *Event) CancelRegistration(userID int64) error {
	query := `
		DELETE FROM registrations WHERE event_id = ? AND user_id = ?
	`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userID)

	if err != nil {
		return err
	}

	return err
}
