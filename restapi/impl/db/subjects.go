package db

import (
	"database/sql"
	"permissions/models"
)

func AddSubject(
	tx *sql.Tx,
	subjectId models.ExternalSubjectID,
	subjectType models.SubjectType,
) (*models.SubjectOut, error) {

	// Update the database.
	query := `INSERT INTO subjects (subject_id, subject_type) VALUES ($1, $2)
            RETURNING id, subject_id, subject_type`
	row := tx.QueryRow(query, string(subjectId), string(subjectType))

	// Return the subject information.
	var subjectDto SubjectDto
	if err := row.Scan(&subjectDto.ID, &subjectDto.SubjectID, &subjectDto.SubjectType); err != nil {
		return nil, err
	}
	return subjectDto.ToSubjectOut(), nil
}

func UpdateSubject(
	tx *sql.Tx,
	id models.InternalSubjectID,
	subjectId models.ExternalSubjectID,
	subjectType models.SubjectType,
) (*models.SubjectOut, error) {

	// Update the database.
	query := `UPDATE subjects SET subject_id = $1, subject_type = $2
            WHERE id = $3
            RETURNING id, subject_id, subject_type`
	row := tx.QueryRow(query, string(subjectId), string(subjectType), string(id))

	// Return the subject information.
	var subjectDto SubjectDto
	if err := row.Scan(&subjectDto.ID, &subjectDto.SubjectID, &subjectDto.SubjectType); err != nil {
		return nil, err
	}
	return subjectDto.ToSubjectOut(), nil
}

func SubjectIdExists(tx *sql.Tx, subjectId models.ExternalSubjectID) (bool, error) {

	// Query the database.
	query := "SELECT count(*) FROM subjects WHERE subject_id = $1"
	row := tx.QueryRow(query, string(subjectId))

	// Get the result.
	var count uint32
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func SubjectExists(tx *sql.Tx, id models.InternalSubjectID) (bool, error) {

	// Query the database.
	query := "SELECT count(*) FROM subjects WHERE id = $1"
	row := tx.QueryRow(query, string(id))

	// Get the result.
	var count uint32
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func DuplicateSubjectExists(
	tx *sql.Tx,
	id models.InternalSubjectID,
	subjectId models.ExternalSubjectID,
) (bool, error) {

	// Query the database.
	query := "SELECT count(*) FROM subjects WHERE id != $1 and subject_id = $2"
	row := tx.QueryRow(query, string(id), string(subjectId))

	// Get the result.
	var count uint32
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func ListSubjects(tx *sql.Tx) ([]*models.SubjectOut, error) {

	// Query the database.
	query := "SELECT id, subject_id, subject_type FROM subjects"
	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Get the list of subjects.
	subjects := make([]*models.SubjectOut, 0)
	for rows.Next() {
		var subjectDto SubjectDto
		if err := rows.Scan(&subjectDto.ID, &subjectDto.SubjectID, &subjectDto.SubjectType); err != nil {
			return nil, err
		}
		subjects = append(subjects, subjectDto.ToSubjectOut())
	}

	return subjects, nil
}
