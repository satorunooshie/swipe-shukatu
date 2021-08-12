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

	ltdM "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

func Test_ltdRepoImpl_Select(t *testing.T) {
	tests := []struct {
		name    string
		query   func(s sqlmock.Sqlmock)
		want    []*ltdM.Ltd
		wantErr bool
	}{
		{
			name:  "valid case: success to fetch ltds",
			query: func(s sqlmock.Sqlmock) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_ltdRepoImpl_SelectLtdNameByID(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx   context.Context
		ltdID int32
	}
	tests := []struct {
		name    string
		query   func(s sqlmock.Sqlmock)
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid case: success to fetch ltd by id",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT name FROM m_ltd WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"ltd_name"}).AddRow("ltd_name"))
			},
			args: args{
				ctx:   context.Background(),
				ltdID: 1,
			},
			want:    "ltd_name",
			wantErr: false,
		},
		{
			name: "invalid case: fail to fetch ltd by id",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(`SELECT name FROM m_ltd WHERE id = ?`).
					WithArgs(1).
					WillReturnError(errors.New("db error"))
			},
			args: args{
				ctx:   context.Background(),
				ltdID: 1,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "invalid case: fail to fetch ltd by id with sql.ErrNoRows",
			query: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(`SELECT name FROM m_ltd WHERE id = ?`).
					WithArgs(1).
					WillReturnError(sql.ErrNoRows)
			},
			args: args{
				ctx:   context.Background(),
				ltdID: 1,
			},
			want:    "",
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
			r := NewLtdRepoImpl(db)
			got, err := r.SelectLtdNameByID(tt.args.ctx, tt.args.ltdID)
			if (err != nil) != tt.wantErr {
				if err == nil {
					return
				}
				t.Errorf("ltdRepoImpl.SelectLtdNameByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ltdRepoImpl.SelectLtdNameByID() returned diff (want -> got):\n%s", diff)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_ltdRepoImpl_Insert(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *ltdM.Ltd
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
			ltdI := &ltdRepoImpl{
				db: tt.fields.db,
			}
			if err := ltdI.Insert(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("ltdRepoImpl.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ltdRepoImpl_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *ltdM.Ltd
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
			ltdI := &ltdRepoImpl{
				db: tt.fields.db,
			}
			if err := ltdI.Update(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("ltdRepoImpl.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ltdRepoImpl_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *ltdM.Ltd
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
			ltdI := &ltdRepoImpl{
				db: tt.fields.db,
			}
			if err := ltdI.Delete(tt.args.ctx, tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("ltdRepoImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_convertToLtd(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name    string
		args    args
		want    []*ltdM.Ltd
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToLtd(tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToLtd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToLtd() = %v, want %v", got, tt.want)
			}
		})
	}
}
