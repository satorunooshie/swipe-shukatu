//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"

	messageM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

func Test_messageRepoImpl_Select(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
		rID int32
	}
	tests := []struct {
		name    string
		query   func(s sqlmock.Sqlmock)
		args    args
		want    []*messageM.Message
		wantErr bool
	}{
		{
			name: "valid case: success to fetch all messages",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT content, image_path, created_at FROM message WHERE recruit_id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"content", "image_path", "created_at"}).AddRow("test", "test", time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC)))
			},
			args: args{
				ctx: context.Background(),
				rID: 1,
			},
			want: []*messageM.Message{
				{
					Content:   "test",
					ImagePath: "test",
					CreatedAt: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{
			name: "invalid case: failed to fetch all messages",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT content, image_path, created_at FROM message WHERE recruit_id = ?").
					WithArgs(1).
					WillReturnError(errors.New("error"))
			},
			args: args{
				ctx: context.Background(),
				rID: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid case: failed to fetch all messages by sql.ErrNoRows",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT content, image_path, created_at FROM message WHERE recruit_id = ?").
					WithArgs(1).
					WillReturnError(sql.ErrNoRows)
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
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.query(mock)
			r := NewMessageRepoImpl(db)
			got, err := r.Select(tt.args.ctx, tt.args.rID)
			if (err != nil) != tt.wantErr {
				if err == nil {
					return
				}
				t.Errorf("messageRepoImpl.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("messageRepoImpl.Select() result diff %s", diff)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_messageRepoImpl_Insert(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *messageM.Message
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
			messageI := &messageRepoImpl{
				db: tt.fields.db,
			}
			if err := messageI.Insert(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("messageRepoImpl.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_messageRepoImpl_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *messageM.Message
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
			messageI := &messageRepoImpl{
				db: tt.fields.db,
			}
			if err := messageI.Update(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("messageRepoImpl.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_messageRepoImpl_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *messageM.Message
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
			messageI := &messageRepoImpl{
				db: tt.fields.db,
			}
			if err := messageI.Delete(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("messageRepoImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_convertToMessages(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name    string
		args    args
		want    []*messageM.Message
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToMessages(tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}
