//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"
	"log"

	recruitM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	recruitR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type recruitRepoImpl struct {
	db *sql.DB
}

func NewRecruitRepoImpl(db *sql.DB) recruitR.RecruitRepository {
	return &recruitRepoImpl{
		db,
	}
}

// Select all recruit.
func (recruitI *recruitRepoImpl) SelectRecruits(ctx context.Context, Param *recruitR.Parameters) ([]*recruitR.Recruit, error) {
	query := "SELECT recruit.ltd_id AS ltd_id, recruit.id AS recruit_id, m_ltd.name AS name, m_ltd.name AS name, m_ltd.employee_number AS employee_number, m_ltd.average_age AS average_age, m_job_type.name AS job_type, m_educational_background.name AS educational_background, recruit.average_salary AS average_salary, recruit.starting_salary AS starting_salary, ltd_image.image_path AS image FROM `recruit`,`ltd_image`,`m_ltd`,`m_job_type`,`m_educational_background` WHERE recruit.job_type_id = m_job_type.id AND recruit.ltd_id = m_ltd.id AND recruit.educational_background_id = m_educational_background.id AND recruit.ltd_id = ltd_image.ltd_id limit 100"
	rows, err := recruitI.db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[INFO] get all recruit: ", err)
			return nil, nil
		}
		return nil, err
	}
	return convertToRecruits(rows)
}

// SelectDetail fetch the recruit.
func (recruitI *recruitRepoImpl) SelectRecruitForMessage(ctx context.Context, rID int32) (*recruitM.Recruit, error) {
	q := "SELECT id, ltd_id, job_type_id FROM recruit WHERE id = ? AND deleted_at IS NULL"
	var r recruitM.Recruit
	err := recruitI.db.QueryRowContext(ctx, q, rID).Scan(
		&r.ID,
		&r.LtdID,
		&r.JobTypeID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[INFO] recruit: ", err)
			return nil, nil
		}
		return nil, err
	}
	return &r, nil
}

// Insert
func (recruitI *recruitRepoImpl) Insert(ctx context.Context, entity *recruitM.Recruit) error {
	stmt, err := recruitI.db.Prepare("INSERT INTO recruit () VALUES ()")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Update
func (recruitI *recruitRepoImpl) Update(ctx context.Context, entity *recruitM.Recruit) error {
	stmt, err := recruitI.db.Prepare("UPDATE recruit SET WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Delete
func (recruitI *recruitRepoImpl) Delete(ctx context.Context, entity *recruitM.Recruit) error {
	stmt, err := recruitI.db.Prepare("DELETE FROM recruit WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// convertToRecruits
func convertToRecruits(rows *sql.Rows) ([]*recruitR.Recruit, error) {
	var recruits []*recruitR.Recruit
	for rows.Next() {
		var recruit *recruitR.Recruit
		err := rows.Scan(
		// Need to scan field
		)
		if err != nil {
			return nil, err
		}
		recruits = append(recruits, recruit)
	}
	return recruits, nil
}
