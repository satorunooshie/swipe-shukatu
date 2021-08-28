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
func (messageI *messageRepoImpl) Select(ctx context.Context, rID int32) ([]*messageM.Message, error) {
	rows, err := messageI.db.QueryContext(ctx, "SELECT content, image_path, created_at FROM message WHERE recruit_id = ?", rID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[INFO] message: ", err)
			return nil, nil
		}
		return nil, err
	}
	return convertToMessages(rows)
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
func convertToMessages(rows *sql.Rows) ([]*messageM.Message, error) {
	var messages []*messageM.Message
	for rows.Next() {
		var message messageM.Message
		err := rows.Scan(
			&message.Content,
			&message.ImagePath,
			&message.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}
	return messages, nil
}
