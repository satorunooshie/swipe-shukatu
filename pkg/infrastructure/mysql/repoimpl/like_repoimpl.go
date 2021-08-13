//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

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

func (likeI *likeRepoImpl) Select(ctx context.Context) ([]*likeM.Like, error) {
	UUID := ""
	rows, err := likeI.db.QueryContext(ctx, "SELECT * FROM like WHERE user_id = ?", UUID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[INFO] like: ", err)
			return nil, nil
		}
		return nil, err
	}
	return convertToLike(rows)
}

// Insert
func (likeI *likeRepoImpl) Insert(ctx context.Context, entity *likeM.Like) error {
	t := time.Now()
	str := fmt.Sprintf("(%s,%d,%s,%s)", entity.UID, entity.RecruitID, t, t)
	stmt, err := likeI.db.Prepare("INSERT INTO like VALUES ?")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(str); err != nil {
		return err
	}
	return nil
}

// Update
func (likeI *likeRepoImpl) Update(ctx context.Context, entity *likeM.Like) error {
	stmt, err := likeI.db.Prepare("UPDATE like SET  WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Delete
func (likeI *likeRepoImpl) Delete(ctx context.Context, entity *likeM.Like) error {
	stmt, err := likeI.db.Prepare("DELETE FROM like WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// convertToLike
func convertToLike(rows *sql.Rows) ([]*likeM.Like, error) {
	var likes []*likeM.Like
	for rows.Next() {
		var like likeM.Like
		err := rows.Scan(
			&like.UID,
			&like.RecruitID,
		)
		if err != nil {
			return nil, err
		}
		likes = append(likes, &like)
	}
	return likes, nil
}
