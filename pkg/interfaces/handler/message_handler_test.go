//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
package handler

import (
	"net/http"
	"reflect"
	"testing"

	messageU "github.com/satorunooshie/swipe-shukatu/pkg/usecase"
)

func TestNewMessageHandler(t *testing.T) {
	type args struct {
		messageU messageU.MessageUseCase
	}
	tests := []struct {
		name string
		args args
		want MessageHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessageHandler(tt.args.messageU); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessageHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messageHandler_HandleSelect(t *testing.T) {
	type fields struct {
		messageUseCase messageU.MessageUseCase
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messageH := &messageHandler{
				messageUseCase: tt.fields.messageUseCase,
			}
			if got := messageH.HandleSelect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageHandler.HandleSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messageHandler_HandleInsert(t *testing.T) {
	type fields struct {
		messageUseCase messageU.MessageUseCase
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messageH := &messageHandler{
				messageUseCase: tt.fields.messageUseCase,
			}
			if got := messageH.HandleInsert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageHandler.HandleInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messageHandler_HandleUpdate(t *testing.T) {
	type fields struct {
		messageUseCase messageU.MessageUseCase
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messageH := &messageHandler{
				messageUseCase: tt.fields.messageUseCase,
			}
			if got := messageH.HandleUpdate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageHandler.HandleUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messageHandler_HandleDelete(t *testing.T) {
	type fields struct {
		messageUseCase messageU.MessageUseCase
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messageH := &messageHandler{
				messageUseCase: tt.fields.messageUseCase,
			}
			if got := messageH.HandleDelete(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageHandler.HandleDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
