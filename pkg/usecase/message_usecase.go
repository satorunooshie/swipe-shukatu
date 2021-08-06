//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"
	"log"
	"time"

	"github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	"github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
	"golang.org/x/sync/errgroup"
)

type MessageUseCase interface {
	Select(ctx context.Context, rID int32) ([]*MessageResponse, error)
	Insert(ctx context.Context, entity *model.Message) error
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
func (messageU *messageUseCase) Select(ctx context.Context, rID int32) ([]*MessageResponse, error) {
	var (
		messages []*model.Message
		rec      *model.Recruit
		err      error
	)

	js, err := messageU.jobRepository.Select(ctx)
	if err != nil {
		return nil, err
	}

	jm := make(map[int32]string)

	for _, j := range js {
		jm[j.ID] = j.Name
	}

	eg, ctx := errgroup.WithContext(ctx)
	// TODO::認証middlewareが入った段階でここでuser idでの絞り込みも追加する
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("[WARN]fetch message's goroutine is canceld")
			return ctx.Err()
		default:
			messages, err = messageU.messageRepository.Select(ctx, rID) // content, image_path, created_at
			if err != nil {
				return err
			}
			return nil
		}
	})

	eg.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("[WARN] fetch recruit's goroutine is canceld")
			return ctx.Err()
		default:
			rec, err = messageU.recruitRepository.SelectRecruitForMessage(ctx, rID) // id, ltd_id, job_type_id
			if err != nil {
				return err
			}
			return nil
		}
	})

	if err = eg.Wait(); err != nil {
		return nil, err
	}

	ltdName, err := messageU.ltdRepository.SelectLtdNameByID(ctx, rec.LtdID)
	if err != nil {
		return nil, err
	}

	ms := make([]*MessageResponse, 0, len(messages))
	for i, m := range messages {
		var mr MessageResponse
		job := jm[rec.JobTypeID]
		mr.LtdID = rec.LtdID
		mr.RecruitID = rID
		mr.Name = ltdName
		mr.JobType = job
		mr.Content = m.Content
		mr.Image = m.ImagePath
		mr.CreatedAt = m.CreatedAt
		ms[i] = &mr
	}
	return ms, nil
}

// Insert
func (messageU *messageUseCase) Insert(ctx context.Context, entity *model.Message) error {
	err := messageU.messageRepository.Insert(ctx, entity)
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

type MessageResponse struct {
	LtdID     int32
	RecruitID int32
	Name      string
	JobType   string
	Content   string
	Image     string
	CreatedAt time.Time
}
