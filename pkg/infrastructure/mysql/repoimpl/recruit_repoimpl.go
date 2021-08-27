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

//select all recruit
func (recruitI *recruitRepoImpl) SelectRecruits(ctx context.Context, Param *recruitR.Parameters) ([]*recruitR.Recruits, error) {
	query := "SELECT"
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
func convertToRecruits(rows *sql.Rows) ([]*recruitR.Recruits, error) {
	var recruits []*recruitR.Recruits
	for rows.Next() {
		var recruit *recruitR.Recruits
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
