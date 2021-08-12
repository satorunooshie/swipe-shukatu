//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/satorunooshie/swipe-shukatu/pkg/usecase"
	"github.com/satorunooshie/swipe-shukatu/pkg/usecase/mock_usecase"
)

func Test_messageHandler_HandleSelect(t *testing.T) {
	t.Parallel()
	var rID int32 = 1
	type fields struct {
		messageUseCase func(m *mock_usecase.MockMessageUseCase)
	}
	tests := []struct {
		name     string
		fields   fields
		want     *MessageResponses
		wantErr  bool
		wantCode int
	}{
		{
			name: "valid case: 200 status",
			fields: fields{
				messageUseCase: func(m *mock_usecase.MockMessageUseCase) {
					m.EXPECT().Select(gomock.Any(), rID).Return([]*usecase.MessageResponse{
						{
							LtdID:     1,
							RecruitID: 1,
							Name:      "name",
							JobType:   "job",
							Content:   "content",
							Image:     "image",
							CreatedAt: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
						},
					}, nil)
				},
			},
			want: &MessageResponses{
				Messages: []*MessageResponse{
					{
						LtdID:     1,
						RecruitID: 1,
						Name:      "name",
						JobType:   "job",
						Content:   "content",
						Image:     "image",
						CreatedAt: time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC),
					},
				},
			},
			wantCode: http.StatusOK,
		},
		{
			name: "invalid case: 500 status",
			fields: fields{
				messageUseCase: func(m *mock_usecase.MockMessageUseCase) {
					m.EXPECT().Select(gomock.Any(), rID).Return(nil, errors.New("error"))
				},
			},
			want:     &MessageResponses{},
			wantErr:  true,
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_usecase.NewMockMessageUseCase(ctrl)
			tt.fields.messageUseCase(m)
			handler := NewMessageHandler(m)

			server := httptest.NewServer(handler.HandleSelect())
			defer server.Close()

			res, err := http.Get(server.URL + "/message/" + strconv.Itoa(int(rID)))
			if err != nil {
				t.Errorf("http.Get(%q) returned error %v", server.URL+"/"+strconv.Itoa(int(rID)), err)
				return
			}
			if res.StatusCode != tt.wantCode {
				t.Errorf("http.Get(%q) returned status %v", server.URL+"/"+strconv.Itoa(int(rID)), res.Status)
				return
			}
			got := &MessageResponses{}
			err = json.NewDecoder(res.Body).Decode(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageHandler.HandleSelect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("handler.HandleSelect() returned diff (want -> got):\n%s", diff)
				return
			}
		})
	}
}

func Test_messageHandler_HandleInsert(t *testing.T) {
}

func Test_messageHandler_HandleUpdate(t *testing.T) {
}

func Test_messageHandler_HandleDelete(t *testing.T) {
}
