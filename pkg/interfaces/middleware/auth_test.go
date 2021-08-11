package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	t.Parallel()
	tests := []struct {
		path    string
		method  string
		handler func() http.HandlerFunc
		want    int
	}{
		{
			path:   "/",
			method: http.MethodGet,
			handler: func() http.HandlerFunc {
				return func(rw http.ResponseWriter, req *http.Request) {}
			},
			want: http.StatusOK,
		},
		{
			path:   "/auth",
			method: http.MethodGet,
			handler: func() http.HandlerFunc {
				return Auth(func(rw http.ResponseWriter, req *http.Request) {})
			},
			want: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			ts := httptest.NewServer(tt.handler())
			defer ts.Close()
			var u bytes.Buffer
			u.WriteString(ts.URL)
			u.WriteString(tt.path)

			req, err := http.NewRequest(tt.method, u.String(), nil)
			if err != nil {
				t.Fatalf("%v", err)
			}
			req.Header.Set("Authorization", "Bearer Test")
			client := &http.Client{}
			res, _ := client.Do(req)
			if res.StatusCode != tt.want {
				t.Errorf("status %v", res.Status)
			}
		})
	}
}
