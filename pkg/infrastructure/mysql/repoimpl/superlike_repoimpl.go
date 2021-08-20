//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"

	superlikeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	superlikeR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type superlikeRepoImpl struct {
	db *sql.DB
}

func NewSuperlikeRepoImpl(db *sql.DB) superlikeR.SuperlikeRepository {
	return &superlikeRepoImpl{
		db,
	}
}

// Select
func (superlikeI *superlikeRepoImpl) Select(ctx context.Context) ([]*superlikeM.Superlike, error) {
	rows, err := superlikeI.db.Query("SELECT * FROM superlike")
	if err != nil {
		return nil, err
	}
	return convertToSuperlike(rows)
}

// Insert
func (superlikeI *superlikeRepoImpl) Insert(ctx context.Context, entity *superlikeM.Superlike) error {
	stmt, err := superlikeI.db.Prepare("INSERT INTO superlike () VALUES ()")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Update
func (superlikeI *superlikeRepoImpl) Update(ctx context.Context, entity *superlikeM.Superlike) error {
	stmt, err := superlikeI.db.Prepare("UPDATE superlike SET WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Delete
func (superlikeI *superlikeRepoImpl) Delete(ctx context.Context, entity *superlikeM.Superlike) error {
	stmt, err := superlikeI.db.Prepare("DELETE FROM superlike WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// convertToSuperlike
func convertToSuperlike(rows *sql.Rows) ([]*superlikeM.Superlike, error) {
	var superlikes []*superlikeM.Superlike
	for rows.Next() {
		var superlike *superlikeM.Superlike
		err := rows.Scan(
		// Need to scan field
		)
		if err != nil {
			return nil, err
		}
		superlikes = append(superlikes, superlike)
	}
	return superlikes, nil
}
