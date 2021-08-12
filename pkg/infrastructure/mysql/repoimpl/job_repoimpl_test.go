package repoimpl_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"

	jobM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
	"github.com/satorunooshie/swipe-shukatu/pkg/infrastructure/mysql/repoimpl"
)

func Test_jobRepoImpl_Select(t *testing.T) {
	t.Helper()
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		query   func(s sqlmock.Sqlmock)
		args    args
		want    []*jobM.Job
		wantErr bool
	}{
		{
			name: "valid case: success to job type",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, name FROM m_job_type").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"))
			},
			args: args{
				ctx: context.Background(),
			},
			want: []*jobM.Job{
				{
					ID:   1,
					Name: "test",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid case: failed to job type",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, name FROM m_job_type").
					WillReturnError(errors.New("failed"))
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid case: failed to job type by sql.ErrNoRows",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, name FROM m_job_type").
					WillReturnError(sql.ErrNoRows)
			},
			args: args{
				ctx: context.Background(),
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
			r := repoimpl.NewJobRepoImpl(db)
			got, err := r.Select(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				// if err is sql.ErrNoRows, we don't care about the error message.
				if err == nil {
					return
				}
				t.Errorf("userRepository.FetchUser() error = %v, wantErr %v", err, tt.wantErr)
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

func Test_jobRepoImpl_Insert(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *jobM.Job
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
		})
	}
}

func Test_jobRepoImpl_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *jobM.Job
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
		})
	}
}

func Test_jobRepoImpl_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *jobM.Job
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
		})
	}
}

func Test_convertToJobs(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name    string
		args    args
		want    []*jobM.Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
