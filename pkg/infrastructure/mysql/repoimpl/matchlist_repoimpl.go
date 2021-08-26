//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"
	"log"

	matchlistM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
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
func (matchlistI *matchlistRepoImpl) Select(ctx context.Context, UID string) ([]*matchlistM.Matchlist, error) {
	queryforlike := "SELECT recruit.ltd_id, recruit_id, m_ltd.name as name, m_job_type.name as job_type, ltd_image.image_path as image FROM `recruit`,`like`,`m_ltd`,`m_job_type`,`ltd_image` where like.user_id=? AND like.recruit_id=recruit.id AND recruit.ltd_id=m_ltd.id AND recruit.job_type_id=m_job_type.id AND recruit.ltd_id=ltd_image.ltd_id AND recruit_id NOT IN (select recruit_id FROM `message` where user_id = ?)"
	queryforsuperlike := "SELECT recruit.ltd_id, recruit_id, m_ltd.name as name, m_job_type.name as job_type, ltd_image.image_path as image FROM `recruit`,`like`,`m_ltd`,`m_job_type`,`ltd_image` where like.user_id=? AND like.recruit_id=recruit.id AND recruit.ltd_id=m_ltd.id AND recruit.job_type_id=m_job_type.id AND recruit.ltd_id=ltd_image.ltd_id AND recruit_id NOT IN (select recruit_id FROM `message` where user_id = ?)"
	jointableRowForMatchListFromLike, err := matchlistI.db.QueryContext(ctx, queryforlike, UID, UID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[INFO] sql err at matchlistRepoImpl.select: failed to get matchlist from like: ", err)
			return nil, nil
		}
		return nil, err
	}
	jointableRowForMatchListFromSuperlike, err := matchlistI.db.QueryContext(ctx, queryforsuperlike, UID, UID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[INFO] sql err at matchlistRepoImpl.select: failed to get matchlist from superlike: ", err)
			return nil, nil
		}
		return nil, err
	}
	return convertToMatchlist(jointableRowForMatchListFromLike, jointableRowForMatchListFromSuperlike)
}

// Insert
func (matchlistI *matchlistRepoImpl) Insert(ctx context.Context, entity *matchlistM.Matchlist) error {
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
func (matchlistI *matchlistRepoImpl) Update(ctx context.Context, entity *matchlistM.Matchlist) error {
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
func (matchlistI *matchlistRepoImpl) Delete(ctx context.Context, entity *matchlistM.Matchlist) error {
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
func convertToMatchlist(likerows, superlikerows *sql.Rows) ([]*matchlistM.Matchlist, error) {
	var matchlists []*matchlistM.Matchlist
	for likerows.Next() {
		var matchlist matchlistM.Matchlist
		err := likerows.Scan(
			&matchlist.LtdID,
			&matchlist.RecruitID,
			&matchlist.Name,
			&matchlist.JobType,
			&matchlist.Image,
		)
		if err != nil {
			return nil, err
		}
		matchlist.Reactiontype = 0
		matchlists = append(matchlists, &matchlist)
	}
	for superlikerows.Next() {
		var matchlist matchlistM.Matchlist
		err := superlikerows.Scan(
			&matchlist.LtdID,
			&matchlist.RecruitID,
			&matchlist.Name,
			&matchlist.JobType,
			&matchlist.Image,
		)
		if err != nil {
			return nil, err
		}
		matchlist.Reactiontype = 1
		matchlists = append(matchlists, &matchlist)
	}
	return matchlists, nil
}
