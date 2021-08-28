//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"
	"log"
	"time"

	messageM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	messageR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type messageRepoImpl struct {
	db *sql.DB
}

func NewMessageRepoImpl(db *sql.DB) messageR.MessageRepository {
	return &messageRepoImpl{
		db,
	}
}

// Select return a slice of Message by recruit ID
func (messageI *messageRepoImpl) Select(ctx context.Context, rID int32) (*messageR.Ltd, []*messageR.Message, error) {
	ltdQuery := "SELECT m_ltd.id AS id, m_ltd.name AS name, m_ltd.profile AS profile, m_ltd.employee_number AS employee_number, m_ltd.average_age AS average_age, m_industry.name AS industry , m_ltd.created_at AS created_at, m_ltd.updated_at AS updated_at, m_ltd.deleted_at AS deleted_at FROM `m_ltd`, `recruit`, `m_industry` WHERE  m_ltd.id = recruit.ltd_id AND recruit.id = ? AND m_ltd.industry_id = m_industry.id "
	rows, err := messageI.db.QueryContext(ctx, ltdQuery, rID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[INFO] ltd: ", err)
			return nil, nil, nil
		}
		return nil, nil, err
	}
	messageQuery := "SELECT recruit_id, content, image_path as photo, created_at FROM `message` WHERE recruit_id = ? ORDER BY created_at ASC"
	rows2, err := messageI.db.QueryContext(ctx, messageQuery, rID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[INFO] ltd: ", err)
			return nil, nil, nil
		}
		return nil, nil, err
	}
	return convertToLtdAndMessages(rows, rows2)
}

// Insert
func (messageI *messageRepoImpl) InsertMessage(ctx context.Context, entity *messageM.Message) error {
	switch entity.Type {
	case 1:
		stmt, err := messageI.db.PrepareContext(ctx, "INSERT INTO `message` (user_id,recruit_id,type,content) VALUES (?,?,?,?)")
		if err != nil {
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(entity.UserID, entity.RecruitID, entity.Type, entity.Content); err != nil {
			return err
		}
	case 2:
		stmt, err := messageI.db.PrepareContext(ctx, "INSERT INTO `message` (user_id,recruit_id,type,image_path) VALUES (?,?,?,?)")
		if err != nil {
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(entity.UserID, entity.RecruitID, entity.Type, entity.ImagePath); err != nil {
			return err
		}
	}
	return nil
}

func (messageI *messageRepoImpl) InsertRemind(ctx context.Context, entity *messageM.Message, ExecuteAt time.Time) error {
	tx, err := messageI.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO `message` (user_id,recruit_id,type,content) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	ret, err := stmt.Exec(entity.UserID, entity.RecruitID, entity.Type, entity.Content)
	if err != nil {
		log.Printf("[ERROR] at message_repoimpl.InsertRemind : failed to Exec: %v", err)
		if err = tx.Rollback(); err != nil {
			log.Printf("[ERROR] failed to insertmessage: %v", err)
		}
		return err
	}
	messageID, err := ret.LastInsertId()
	if err != nil {
		return err
	}
	stmt2, err := tx.PrepareContext(ctx, "INSERT INTO `remind_message` (message_id,execute_at) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stmt2.Close()
	if _, err = stmt2.Exec(messageID, ExecuteAt); err != nil {
		if err = tx.Rollback(); err != nil {
			log.Printf("[ERROR] failed to insertmessage: %v", err)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Printf("[ERROR] failed to commit sql: %v", err)
		return err
	}
	return nil
}

// Update
func (messageI *messageRepoImpl) Update(ctx context.Context, entity *messageM.Message) error {
	stmt, err := messageI.db.Prepare("UPDATE message SET WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// Delete
func (messageI *messageRepoImpl) Delete(ctx context.Context, entity *messageM.Message) error {
	stmt, err := messageI.db.Prepare("DELETE FROM message WHERE ")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
		return err
	}
	return nil
}

// convertToMessages
func convertToLtdAndMessages(rows, rows2 *sql.Rows) (*messageR.Ltd, []*messageR.Message, error) {
	var ltd messageR.Ltd
	defer rows.Close()
	err := rows.Scan(
		&ltd.ID,
		&ltd.Name,
		&ltd.Profile,
		&ltd.EmployeeNumber,
		&ltd.AverageNumber,
		&ltd.Industry,
		&ltd.CreatedAt,
		&ltd.UpdatedAt,
		&ltd.DeletedAt,
	)
	if err != nil {
		return nil, nil, err
	}
	var messages []*messageR.Message
	defer rows2.Close()
	for rows2.Next() {
		var message messageR.Message
		err := rows.Scan(
			&message.RecruitID,
			&message.Content,
			&message.ImagePath,
			&message.CreatedAt,
		)
		if err != nil {
			return nil, nil, err
		}
		messages = append(messages, &message)
	}
	return &ltd, messages, nil
}
