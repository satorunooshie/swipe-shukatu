//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"
	"log"

	likeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	likeR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type likeRepoImpl struct {
	db *sql.DB
}

func NewLikeRepoImpl(db *sql.DB) likeR.LikeRepository {
	return &likeRepoImpl{
		db,
	}
}

// Select

func (likeI *likeRepoImpl) Select(ctx context.Context, UID string) ([]*likeM.Like, error) {
	rows, err := likeI.db.QueryContext(ctx, "SELECT user_id, recruit_id, created_at, updated_at FROM `like` WHERE user_id = ?", UID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[INFO] like: ", err)
			return nil, nil
		}
		return nil, err
	}
	return convertToLikes(rows)
}

// Insert
func (likeI *likeRepoImpl) Insert(ctx context.Context, entity *likeM.Like, UID string) error {
	query := "INSERT INTO `like`(user_id, recruit_id) VALUES (?,?)"
	stmt, err := likeI.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(UID, entity.RecruitID); err != nil {
		return err
	}
	return nil
}

// Update
func (likeI *likeRepoImpl) Update(ctx context.Context, entity *likeM.Like) error {
	stmt, err := likeI.db.PrepareContext(ctx, "UPDATE like SET WHERE ")
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
func (likeI *likeRepoImpl) Delete(ctx context.Context, entity *likeM.Like) error {
	stmt, err := likeI.db.PrepareContext(ctx, "DELETE FROM like WHERE ")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// convertToLike
func convertToLikes(rows *sql.Rows) ([]*likeM.Like, error) {
	var likes []*likeM.Like
	for rows.Next() {
		var like likeM.Like
		err := rows.Scan(
			&like.UID,
			&like.RecruitID,
			&like.CreatedAt,
			&like.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		likes = append(likes, &like)
	}
	return likes, nil
}
