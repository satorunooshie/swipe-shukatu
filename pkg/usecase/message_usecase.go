//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"
	"time"

	"github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	"github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
	messageR "github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
)

type MessageUseCase interface {
	Select(ctx context.Context, rID int32) (*messageR.Ltd, []*messageR.Message, error)
	InsertMessage(ctx context.Context, entity *model.Message) error
	InsertRemind(ctx context.Context, entity *model.Message, ExecuteAt time.Time) error
	Update(ctx context.Context, entity *model.Message) error
	Delete(ctx context.Context, entity *model.Message) error
}

type messageUseCase struct {
	jobRepository     repository.JobRepository
	ltdRepository     repository.LtdRepository
	messageRepository repository.MessageRepository
	recruitRepository repository.RecruitRepository
}

// NewMessageUsecase
func NewMessageUsecase(
	jobR repository.JobRepository,
	ltdR repository.LtdRepository,
	messageR repository.MessageRepository,
	recR repository.RecruitRepository,
) MessageUseCase {
	return &messageUseCase{
		jobRepository:     jobR,
		ltdRepository:     ltdR,
		messageRepository: messageR,
		recruitRepository: recR,
	}
}

// Select
func (messageU *messageUseCase) Select(ctx context.Context, rID int32) (*messageR.Ltd, []*messageR.Message, error) {
	ltd, messages, err := messageU.messageRepository.Select(ctx, rID)
	if err != nil {
		return nil, nil, err
	}
	return ltd, messages, nil
}

// Insert
func (messageU *messageUseCase) InsertMessage(ctx context.Context, entity *model.Message) error {
	err := messageU.messageRepository.InsertMessage(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

func (messageU *messageUseCase) InsertRemind(ctx context.Context, entity *model.Message, ExecuteAt time.Time) error {
	err := messageU.messageRepository.InsertRemind(ctx, entity, ExecuteAt)
	if err != nil {
		return err
	}
	return nil
}

// Update
func (messageU *messageUseCase) Update(ctx context.Context, entity *model.Message) error {
	err := messageU.messageRepository.Update(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (messageU *messageUseCase) Delete(ctx context.Context, entity *model.Message) error {
	err := messageU.messageRepository.Delete(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
