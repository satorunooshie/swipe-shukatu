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

// Select
func (recruitI *recruitRepoImpl) Select(ctx context.Context) ([]*recruitM.Recruit, error) {
	rows, err := recruitI.db.Query("SELECT * FROM recruit")
	if err != nil {
		return nil, err
	}
	return convertToRecruit(rows)
}

// SelectDetail fetch the recruit.
func (recruitI *recruitRepoImpl) SelectDetail(ctx context.Context, rID int32) (*recruitM.Recruit, error) {
	var r recruitM.Recruit
	if err := recruitI.db.QueryRowContext(ctx, "SELECT id, ltd_id, job_type_id FROM recruit WHERE id = ?", rID).Scan(
		r.ID,
		r.LtdID,
		r.JobTypeID,
	); err != nil {
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

// convertToRecruit
func convertToRecruit(rows *sql.Rows) ([]*recruitM.Recruit, error) {
	var recruits []*recruitM.Recruit
	for rows.Next() {
		var recruit *recruitM.Recruit
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
