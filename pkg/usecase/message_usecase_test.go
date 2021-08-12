//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	"github.com/satorunooshie/swipe-shukatu/pkg/domain/repository"
	"github.com/satorunooshie/swipe-shukatu/pkg/domain/repository/mock_repository"
)

func Test_messageUseCase_Select(t *testing.T) {
	t.Parallel()
	var (
		rID int32 = 1
		lID int32 = 1
	)
	type fields struct {
		jobRepository     func(m *mock_repository.MockJobRepository)
		ltdRepository     func(m *mock_repository.MockLtdRepository)
		messageRepository func(m *mock_repository.MockMessageRepository)
		recruitRepository func(m *mock_repository.MockRecruitRepository)
	}
	type args struct {
		ctx context.Context
		rID int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*MessageResponse
		wantErr bool
	}{
		{
			name: "valid case: success to fetch messages",
			fields: fields{
				jobRepository: func(m *mock_repository.MockJobRepository) {
					m.EXPECT().Select(gomock.Any()).Return([]*model.Job{
						{
							ID:   1,
							Name: "job1",
						},
					}, nil)
				},
				ltdRepository: func(m *mock_repository.MockLtdRepository) {
					m.EXPECT().SelectLtdNameByID(gomock.Any(), lID).Return("ltd1", nil)
				},
				messageRepository: func(m *mock_repository.MockMessageRepository) {
					m.EXPECT().Select(gomock.Any(), rID).Return([]*model.Message{
						{
							Content:   "message1",
							ImagePath: "image1",
							CreatedAt: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
						},
					}, nil)
				},
				recruitRepository: func(m *mock_repository.MockRecruitRepository) {
					m.EXPECT().SelectRecruitForMessage(gomock.Any(), rID).Return(&model.Recruit{
						ID:        1,
						LtdID:     1,
						JobTypeID: 1,
					}, nil)
				},
			},
			args: args{
				ctx: context.Background(),
				rID: 1,
			},
			want: []*MessageResponse{
				{
					LtdID:     1,
					RecruitID: 1,
					Name:      "ltd1",
					JobType:   "job1",
					Content:   "message1",
					Image:     "image1",
					CreatedAt: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{
			name: "invalid case: failed to fetch job",
			fields: fields{
				jobRepository: func(m *mock_repository.MockJobRepository) {
					m.EXPECT().Select(gomock.Any()).Return(nil, errors.New("failed to fetch job"))
				},
				ltdRepository:     func(m *mock_repository.MockLtdRepository) {},
				messageRepository: func(m *mock_repository.MockMessageRepository) {},
				recruitRepository: func(m *mock_repository.MockRecruitRepository) {},
			},
			args: args{
				ctx: context.Background(),
				rID: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid case: failed to message",
			fields: fields{
				jobRepository: func(m *mock_repository.MockJobRepository) {
					m.EXPECT().Select(gomock.Any()).Return([]*model.Job{
						{
							ID:   1,
							Name: "job1",
						},
					}, nil)
				},
				ltdRepository: func(m *mock_repository.MockLtdRepository) {},
				messageRepository: func(m *mock_repository.MockMessageRepository) {
					m.EXPECT().Select(gomock.Any(), rID).Return(nil, errors.New("failed to message"))
				},
				recruitRepository: func(m *mock_repository.MockRecruitRepository) {
					m.EXPECT().SelectRecruitForMessage(gomock.Any(), rID).Return(&model.Recruit{
						ID:        1,
						LtdID:     1,
						JobTypeID: 1,
					}, nil).AnyTimes()
				},
			},
			args: args{
				ctx: context.Background(),
				rID: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid case: failed to fetch recruit",
			fields: fields{
				jobRepository: func(m *mock_repository.MockJobRepository) {
					m.EXPECT().Select(gomock.Any()).Return([]*model.Job{
						{
							ID:   1,
							Name: "job1",
						},
					}, nil)
				},
				ltdRepository: func(m *mock_repository.MockLtdRepository) {},
				messageRepository: func(m *mock_repository.MockMessageRepository) {
					m.EXPECT().Select(gomock.Any(), rID).Return([]*model.Message{
						{
							Content:   "message1",
							ImagePath: "image1",
							CreatedAt: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
						},
					}, nil).AnyTimes()
				},
				recruitRepository: func(m *mock_repository.MockRecruitRepository) {
					m.EXPECT().SelectRecruitForMessage(gomock.Any(), rID).Return(nil, errors.New("failed to fetch recruit"))
				},
			},
			args: args{
				ctx: context.Background(),
				rID: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid case: failed to fetch response with ctx timeout",
			fields: fields{
				jobRepository: func(m *mock_repository.MockJobRepository) {
					m.EXPECT().Select(gomock.Any()).Return([]*model.Job{
						{
							ID:   1,
							Name: "job1",
						},
					}, nil).AnyTimes()
				},
				ltdRepository:     func(m *mock_repository.MockLtdRepository) {},
				messageRepository: func(m *mock_repository.MockMessageRepository) {},
				recruitRepository: func(m *mock_repository.MockRecruitRepository) {},
			},
			args: args{
				ctx: func() context.Context {
					ctx, cancel := context.WithTimeout(context.Background(), 0)
					_ = cancel
					return ctx
				}(),
				rID: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid case: failed to fetch ltd name",
			fields: fields{
				jobRepository: func(m *mock_repository.MockJobRepository) {
					m.EXPECT().Select(gomock.Any()).Return([]*model.Job{
						{
							ID:   1,
							Name: "job1",
						},
					}, nil)
				},
				ltdRepository: func(m *mock_repository.MockLtdRepository) {
					m.EXPECT().SelectLtdNameByID(gomock.Any(), rID).Return("", errors.New("failed to fetch ltd name"))
				},
				messageRepository: func(m *mock_repository.MockMessageRepository) {
					m.EXPECT().Select(gomock.Any(), rID).Return([]*model.Message{
						{
							Content:   "message1",
							ImagePath: "image1",
							CreatedAt: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
						},
					}, nil).AnyTimes()
				},
				recruitRepository: func(m *mock_repository.MockRecruitRepository) {
					m.EXPECT().SelectRecruitForMessage(gomock.Any(), rID).Return(&model.Recruit{
						ID:        1,
						LtdID:     1,
						JobTypeID: 1,
					}, nil)
				},
			},
			args: args{
				ctx: context.Background(),
				rID: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockJobRepository := mock_repository.NewMockJobRepository(ctrl)
			mockLtdRepository := mock_repository.NewMockLtdRepository(ctrl)
			mockMessageRepository := mock_repository.NewMockMessageRepository(ctrl)
			mockRecruitRepository := mock_repository.NewMockRecruitRepository(ctrl)

			tt.fields.jobRepository(mockJobRepository)
			tt.fields.ltdRepository(mockLtdRepository)
			tt.fields.messageRepository(mockMessageRepository)
			tt.fields.recruitRepository(mockRecruitRepository)

			messageU := NewMessageUsecase(mockJobRepository, mockLtdRepository, mockMessageRepository, mockRecruitRepository)

			got, err := messageU.Select(tt.args.ctx, tt.args.rID)
			if (err != nil) != tt.wantErr {
				t.Errorf("messageUseCase.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("messageUseCase.Select() got = %v, want %v diff %s", got, tt.want, diff)
			}
		})
	}
}

func Test_messageUseCase_Insert(t *testing.T) {
	type fields struct {
		jobRepository     repository.JobRepository
		ltdRepository     repository.LtdRepository
		messageRepository repository.MessageRepository
		recruitRepository repository.RecruitRepository
	}
	type args struct {
		ctx    context.Context
		entity *model.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messageU := &messageUseCase{
				jobRepository:     tt.fields.jobRepository,
				ltdRepository:     tt.fields.ltdRepository,
				messageRepository: tt.fields.messageRepository,
				recruitRepository: tt.fields.recruitRepository,
			}
			if err := messageU.Insert(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("messageUseCase.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_messageUseCase_Update(t *testing.T) {
	type fields struct {
		jobRepository     repository.JobRepository
		ltdRepository     repository.LtdRepository
		messageRepository repository.MessageRepository
		recruitRepository repository.RecruitRepository
	}
	type args struct {
		ctx    context.Context
		entity *model.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messageU := &messageUseCase{
				jobRepository:     tt.fields.jobRepository,
				ltdRepository:     tt.fields.ltdRepository,
				messageRepository: tt.fields.messageRepository,
				recruitRepository: tt.fields.recruitRepository,
			}
			if err := messageU.Update(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("messageUseCase.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_messageUseCase_Delete(t *testing.T) {
	type fields struct {
		jobRepository     repository.JobRepository
		ltdRepository     repository.LtdRepository
		messageRepository repository.MessageRepository
		recruitRepository repository.RecruitRepository
	}
	type args struct {
		ctx    context.Context
		entity *model.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messageU := &messageUseCase{
				jobRepository:     tt.fields.jobRepository,
				ltdRepository:     tt.fields.ltdRepository,
				messageRepository: tt.fields.messageRepository,
				recruitRepository: tt.fields.recruitRepository,
			}
			if err := messageU.Delete(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("messageUseCase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
