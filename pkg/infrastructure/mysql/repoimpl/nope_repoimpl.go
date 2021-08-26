//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"
	"log"

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

func (nopeI *nopeRepoImpl) Select(ctx context.Context, UID string) ([]*nopeM.Nope, error) {
	rows, err := nopeI.db.QueryContext(ctx, "SELECT user_id, recruit_id, created_at, updated_at FROM `nope` WHERE user_id = ?", UID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[INFO] nope: %v", err)
			return nil, nil
		}
		return nil, err
	}
	return convertToNopes(rows)
}

// Insert
func (nopeI *nopeRepoImpl) Insert(ctx context.Context, entity *nopeM.Nope) error {
	query := "INSERT INTO `nope`(user_id, recruit_id) VALUES (?,?)"
	stmt, err := nopeI.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(entity.UID, entity.RecruitID); err != nil {
		return err
	}
	return nil
}

// Update
func (nopeI *nopeRepoImpl) Update(ctx context.Context, entity *nopeM.Nope) error {
	stmt, err := nopeI.db.PrepareContext(ctx, "UPDATE nope SET WHERE ")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Delete
func (nopeI *nopeRepoImpl) Delete(ctx context.Context, entity *nopeM.Nope) error {
	stmt, err := nopeI.db.PrepareContext(ctx, "DELETE FROM nope WHERE ")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// convertToNope
func convertToNopes(rows *sql.Rows) ([]*nopeM.Nope, error) {
	var nopes []*nopeM.Nope
	for rows.Next() {
		var nope nopeM.Nope
		err := rows.Scan(
			&nope.UID,
			&nope.RecruitID,
			&nope.CreatedAt,
			&nope.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		nopes = append(nopes, &nope)
	}
	return nopes, nil
}
