//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"
	"log"

	jobM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	jobR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type jobRepoImpl struct {
	db *sql.DB
}

func NewJobRepoImpl(db *sql.DB) jobR.JobRepository {
	return &jobRepoImpl{
		db,
	}
}

// Select return job's id, name list.
func (jobI *jobRepoImpl) Select(ctx context.Context) ([]*jobM.Job, error) {
	rows, err := jobI.db.QueryContext(ctx, "SELECT id, name FROM m_job_type")
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[INFO] job: ", err)
			return nil, nil
		}
		return nil, err
	}
	return convertToJobs(rows)
}

// Insert
func (jobI *jobRepoImpl) Insert(ctx context.Context, entity *jobM.Job) error {
	stmt, err := jobI.db.Prepare("INSERT INTO job () VALUES ()")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Update
func (jobI *jobRepoImpl) Update(ctx context.Context, entity *jobM.Job) error {
	stmt, err := jobI.db.Prepare("UPDATE job SET WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Delete
func (jobI *jobRepoImpl) Delete(ctx context.Context, entity *jobM.Job) error {
	stmt, err := jobI.db.Prepare("DELETE FROM job WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// convertToJobs
func convertToJobs(rows *sql.Rows) ([]*jobM.Job, error) {
	var jobs []*jobM.Job
	for rows.Next() {
		var job jobM.Job
		err := rows.Scan(
			&job.ID,
			&job.Name,
		)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
	}
	return jobs, nil
}
