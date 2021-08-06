//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"

	ltdM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	ltdR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type ltdRepoImpl struct {
	db *sql.DB
}

func NewLtdRepoImpl(db *sql.DB) ltdR.LtdRepository {
	return &ltdRepoImpl{
		db,
	}
}

// Select
func (ltdI *ltdRepoImpl) Select(ctx context.Context) ([]*ltdM.Ltd, error) {
	rows, err := ltdI.db.Query("SELECT * FROM ltd")
	if err != nil {
		return nil, err
	}
	return convertToLtd(rows)
}

// SelectLtdNameByID return the name of Ltd by ID
func (ltdI *ltdRepoImpl) SelectLtdNameByID(ctx context.Context, ltdID int32) (string, error) {
	row := ltdI.db.QueryRowContext(ctx, "SELECT name FROM ltd WHERE id = ?", ltdID)
	var name string
	err := row.Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return name, nil
}

// Insert
func (ltdI *ltdRepoImpl) Insert(ctx context.Context, entity *ltdM.Ltd) error {
	stmt, err := ltdI.db.Prepare("INSERT INTO ltd () VALUES ()")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Update
func (ltdI *ltdRepoImpl) Update(ctx context.Context, entity *ltdM.Ltd) error {
	stmt, err := ltdI.db.Prepare("UPDATE ltd SET WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Delete
func (ltdI *ltdRepoImpl) Delete(ctx context.Context, entity *ltdM.Ltd) error {
	stmt, err := ltdI.db.Prepare("DELETE FROM ltd WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// convertToLtd
func convertToLtd(rows *sql.Rows) ([]*ltdM.Ltd, error) {
	var ltds []*ltdM.Ltd
	for rows.Next() {
		var ltd *ltdM.Ltd
		err := rows.Scan(
		// Need to scan field
		)
		if err != nil {
			return nil, err
		}
		ltds = append(ltds, ltd)
	}
	return ltds, nil
}
