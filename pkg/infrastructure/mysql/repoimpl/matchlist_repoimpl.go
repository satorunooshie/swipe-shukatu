//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"
	"log"

	matchlistR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type matchlistRepoImpl struct {
	db *sql.DB
}

func NewMatchlistRepoImpl(db *sql.DB) matchlistR.MatchlistRepository {
	return &matchlistRepoImpl{
		db,
	}
}

// Select
func (matchlistI *matchlistRepoImpl) Select(ctx context.Context, UID string) ([]*matchlistR.Match, error) {
	queryforlike := "SELECT /*出力するカラム*/ recruit.ltd_id, recruit_id, m_ltd.name as name, ltd_image.image_path as image FROM /*使用するテーブル*/ `recruit`,`like`,`m_ltd`,`ltd_image` where /*条件*/like.user_id=? AND like.recruit_id=recruit.id AND recruit.ltd_id=m_ltd.id AND recruit.ltd_id=ltd_image.ltd_id /*messageテーブルのuseridが存在するrecruitIDをselectしてきてNOT INで該当recruitIDを弾く*/ AND recruit_id NOT IN (select recruit_id FROM `message` where user_id = ?) ORDER BY like.updated_at DESC"
	queryforsuperlike := "SELECT /*出力するカラム*/ recruit.ltd_id, recruit_id, m_ltd.name as name, ltd_image.image_path as image FROM /*使用するテーブル*/` `recruit`,`like`,`m_ltd`,`ltd_image` where /*条件*/ like.user_id=? AND like.recruit_id=recruit.id AND recruit.ltd_id=m_ltd.id AND recruit.ltd_id=ltd_image.ltd_id /*messageテーブルのuseridが存在するrecruitIDをselectしてきてNOT INで該当recruitIDを弾く*/ AND recruit_id NOT IN (select recruit_id FROM `message` where user_id = ?) ORDER BY like.updated_at DESC"
	jointableRowForMatchListFromLike, err := matchlistI.db.QueryContext(ctx, queryforlike, UID, UID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[INFO] sql err at matchlistRepoImpl.select: failed to get matchlist from like: %v", err)
			return nil, nil
		}
		return nil, err
	}
	jointableRowForMatchListFromSuperlike, err := matchlistI.db.QueryContext(ctx, queryforsuperlike, UID, UID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[INFO] sql err at matchlistRepoImpl.select: failed to get matchlist from superlike: %v", err)
			return nil, nil
		}
		return nil, err
	}
	return convertToMatchlist(jointableRowForMatchListFromLike, jointableRowForMatchListFromSuperlike)
}

// Insert
func (matchlistI *matchlistRepoImpl) Insert(ctx context.Context, entity *matchlistR.Match) error {
	stmt, err := matchlistI.db.Prepare("INSERT INTO matchlist () VALUES ()")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Update
func (matchlistI *matchlistRepoImpl) Update(ctx context.Context, entity *matchlistR.Match) error {
	stmt, err := matchlistI.db.Prepare("UPDATE matchlist SET WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Delete
func (matchlistI *matchlistRepoImpl) Delete(ctx context.Context, entity *matchlistR.Match) error {
	stmt, err := matchlistI.db.Prepare("DELETE FROM matchlist WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// convertToMatchlist
func convertToMatchlist(likerows, superlikerows *sql.Rows) ([]*matchlistR.Match, error) {
	var matchlists []*matchlistR.Match
	for likerows.Next() {
		var matchlist matchlistR.Match
		err := likerows.Scan(
			&matchlist.LtdID,
			&matchlist.RecruitID,
			&matchlist.Name,
			&matchlist.Image,
		)
		if err != nil {
			return nil, err
		}
		if matchlist.RecruitID < 0 {
			log.Println("[ERROR] sql error: Invalid recruitID detected.")
		}
		matchlist.Reactiontype = 1
		matchlists = append(matchlists, &matchlist)
	}
	for superlikerows.Next() {
		var matchlist matchlistR.Match
		err := superlikerows.Scan(
			&matchlist.LtdID,
			&matchlist.RecruitID,
			&matchlist.Name,
			&matchlist.Image,
		)
		if err != nil {
			return nil, err
		}
		if matchlist.RecruitID < 0 {
			log.Println("[ERROR] sql error: Invalid recruitID detected.")
		}
		matchlist.Reactiontype = 2
		matchlists = append(matchlists, &matchlist)
	}
	return matchlists, nil
}
