package repository

import (
	"database/sql"
	"errors"

	"queueflow/internal/models"

	"github.com/google/uuid"
)

type JobRepository struct {
	db *sql.DB
}

func NewJobRepository(db *sql.DB) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

// Create new job
func (r *JobRepository) Create(job models.Job) error {

	query := `
		INSERT INTO jobs
		(
			id,
			type,
			payload,
			status,
			created_at,
			updated_at
		)
		VALUES
		($1,$2,$3,$4,NOW(),NOW())
	`

	_, err := r.db.Exec(
		query,
		job.ID,
		job.Type,
		job.Payload,
		job.Status,
	)

	return err
}

// Get job by ID
func (r *JobRepository) GetByID(id uuid.UUID) (*models.Job, error) {

	query := `
		SELECT
			id,
			type,
			payload,
			status,
			created_at,
			updated_at
		FROM jobs
		WHERE id=$1
	`

	row := r.db.QueryRow(query, id)

	var job models.Job

	err := row.Scan(
		&job.ID,
		&job.Type,
		&job.Payload,
		&job.Status,
		&job.CreatedAt,
		&job.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("job not found")
	}

	if err != nil {
		return nil, err
	}

	return &job, nil
}

// Update job status
func (r *JobRepository) UpdateStatus(
	id uuid.UUID,
	status string,
) error {

	query := `
		UPDATE jobs
		SET
			status=$1,
			updated_at=NOW()
		WHERE id=$2
	`

	_, err := r.db.Exec(
		query,
		status,
		id,
	)

	return err
}

// Get jobs list
func (r *JobRepository) List(skip *int, take *int) ([]models.Job, error) {

	query := `
		SELECT
			id,
			type,
			payload,
			status,
			created_at,
			updated_at
		FROM jobs
	`

	// if skip != nil && take != nil {
	// 	query += " LIMIT $1 OFFSET $2"
	// }

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		err := rows.Scan(
			&job.ID,
			&job.Type,
			&job.Payload,
			&job.Status,
			&job.CreatedAt,
			&job.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}
