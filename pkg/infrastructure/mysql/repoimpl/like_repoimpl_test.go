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

	likeM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

func Test_likeRepoImpl_Select(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
		UID string
	}
	tests := []struct {
		name    string
		query   func(s sqlmock.Sqlmock)
		args    args
		want    []*likeM.Like
		wantErr bool
	}{
		{
			name: "[正常系]like取得",
			query: func(s sqlmock.Sqlmock) {
				testT := time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC)
				columns := []string{"user_id", "recruit_id", "created_at", "updated_at"}
				s.ExpectQuery("SELECT user_id, recruit_id, created_at, updated_at FROM like WHERE user_id = ?").
					WithArgs("testUid").
					WillReturnRows(sqlmock.NewRows(columns).AddRow("testUid", 1, testT, testT))
			},
			args: args{
				ctx: context.Background(),
				UID: "testUid",
			},
			want: []*likeM.Like{
				{
					UID:        "testUid",
					RecruitID:  1,
					CreatedAt:  time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
					UpdateddAt: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{
			name: "[異常系]like取得エラー",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT user_id, recruit_id, created_at, updated_at FROM like WHERE user_id = ?").
					WithArgs("testUid").
					WillReturnError(errors.New("error"))
			},
			args: args{
				ctx: context.Background(),
				UID: "testUid",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "[異常系]sqlエラー",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT user_id, recruit_id, created_at, updated_at FROM like WHERE user_id = ?").
					WithArgs("testUid").
					WillReturnError(sql.ErrNoRows)
			},
			args: args{
				ctx: context.Background(),
				UID: "testUid",
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
			r := NewLikeRepoImpl(db)
			got, err := r.Select(tt.args.ctx, tt.args.UID)
			if (err != nil) != tt.wantErr {
				if err == nil {
					return
				}
				t.Errorf("likeRepoImpl.Select() \n error = %v, \nwantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("likeRepoImpl.Select() result diff %s", diff)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_likeRepoImpl_Insert(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *likeM.Like
		UID    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases. preprecontext(insert)の適切なtestの書き方迷走してるので一旦放置
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			likeI := &likeRepoImpl{
				db: tt.fields.db,
			}
			if err := likeI.Insert(tt.args.ctx, tt.args.entity, tt.args.UID); (err != nil) != tt.wantErr {
				t.Errorf("likeRepoImpl.Insert() \n error = %v, \n wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_likeRepoImpl_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *likeM.Like
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
			likeI := &likeRepoImpl{
				db: tt.fields.db,
			}
			if err := likeI.Update(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("likeRepoImpl.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_likeRepoImpl_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *likeM.Like
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
			likeI := &likeRepoImpl{
				db: tt.fields.db,
			}
			if err := likeI.Delete(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("likeRepoImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_convertToLikes(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name    string
		args    args
		want    []*likeM.Like
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToLikes(tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToLikes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToLikes() = %v, want %v", got, tt.want)
			}
		})
	}
}
