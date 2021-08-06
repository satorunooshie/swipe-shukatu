//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"
	"log"

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
func (messageI *messageRepoImpl) Insert(ctx context.Context, entity *messageM.Message) error {
	stmt, err := messageI.db.Prepare("INSERT INTO message () VALUES ()")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(); err != nil {
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
