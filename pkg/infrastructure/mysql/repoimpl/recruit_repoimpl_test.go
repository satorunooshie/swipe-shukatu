//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package repoimpl

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	recruitM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

func Test_recruitRepoImpl_Select(t *testing.T) {
	t.Helper()
	t.Parallel()
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*recruitM.Recruit
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recruitI := &recruitRepoImpl{
				db: tt.fields.db,
			}
			got, err := recruitI.Select(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("recruitRepoImpl.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recruitRepoImpl.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recruitRepoImpl_SelectRecruitForMessage(t *testing.T) {
	t.Helper()
	t.Parallel()
	type args struct {
		ctx context.Context
		rID int32
	}
	tests := []struct {
		name    string
		query   func(s sqlmock.Sqlmock)
		args    args
		want    *recruitM.Recruit
		wantErr bool
	}{
		{
			name: "valid case: success to recult for message",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, ltd_id, job_type_id FROM recruit WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "ltd_id", "job_type_id"}).AddRow(1, 1, 1))
			},
			args: args{
				ctx: context.Background(),
				rID: 1,
			},
			want: &recruitM.Recruit{
				ID:        1,
				LtdID:     1,
				JobTypeID: 1,
			},
			wantErr: false,
		},
		{
			name: "invalid case: failed to fetch data",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, ltd_id, job_type_id FROM recruit WHERE id = ?").
					WithArgs(1).
					WillReturnError(errors.New("failed to fetch data"))
			},
			args: args{
				ctx: context.Background(),
				rID: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid case: failed to fetch data with sql.ErrNoRows",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, ltd_id, job_type_id FROM recruit WHERE id = ?").
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
			r := NewRecruitRepoImpl(db)
			got, err := r.SelectRecruitForMessage(tt.args.ctx, tt.args.rID)
			if (err != nil) != tt.wantErr {
				if err == nil {
					return
				}
				t.Errorf("RecruitRepoImpl.SelectRecruitForMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("userRepository.FetchUser() returned diff (want -> got):\n%s", diff)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_recruitRepoImpl_Insert(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *recruitM.Recruit
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
			recruitI := &recruitRepoImpl{
				db: tt.fields.db,
			}
			if err := recruitI.Insert(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("recruitRepoImpl.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_recruitRepoImpl_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *recruitM.Recruit
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
			recruitI := &recruitRepoImpl{
				db: tt.fields.db,
			}
			if err := recruitI.Update(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("recruitRepoImpl.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_recruitRepoImpl_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *recruitM.Recruit
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
			recruitI := &recruitRepoImpl{
				db: tt.fields.db,
			}
			if err := recruitI.Delete(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("recruitRepoImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_convertToRecruits(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name    string
		args    args
		want    []*recruitM.Recruit
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToRecruits(tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToRecruits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToRecruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
