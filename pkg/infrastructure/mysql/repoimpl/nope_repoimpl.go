//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"

	nopeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	nopeR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type nopeRepoImpl struct {
	db *sql.DB
}

func NewNopeRepoImpl(db *sql.DB) nopeR.NopeRepository {
	return &nopeRepoImpl{
		db,
	}
}

// Select
func (nopeI *nopeRepoImpl) Select(ctx context.Context) ([]*nopeM.Nope, error) {
	rows, err := nopeI.db.Query("SELECT * FROM nope")
	if err != nil {
		return nil, err
	}
	return convertToNope(rows)
}

// Insert
func (nopeI *nopeRepoImpl) Insert(ctx context.Context, entity *nopeM.Nope) error {
	stmt, err := nopeI.db.Prepare("INSERT INTO nope () VALUES ()")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Update
func (nopeI *nopeRepoImpl) Update(ctx context.Context, entity *nopeM.Nope) error {
	stmt, err := nopeI.db.Prepare("UPDATE nope SET WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Delete
func (nopeI *nopeRepoImpl) Delete(ctx context.Context, entity *nopeM.Nope) error {
	stmt, err := nopeI.db.Prepare("DELETE FROM nope WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// convertToNope
func convertToNope(rows *sql.Rows) ([]*nopeM.Nope, error) {
	var nopes []*nopeM.Nope
	for rows.Next() {
		var nope *nopeM.Nope
		err := rows.Scan(
		// Need to scan field
		)
		if err != nil {
			return nil, err
		}
		nopes = append(nopes, nope)
	}
	return nopes, nil
}
