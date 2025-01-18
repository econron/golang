package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_packRequestToMessage(t *testing.T) {
	test1Body := map[string]int{
		"server": 1,
		"order":  1,
	}
	test1, _ := json.Marshal(test1Body)
	type args struct {
		r []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Message
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				r: test1,
			},
			want: &Message{
				ServerID: 1,
				Order:    1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "http://localhost:8000", bytes.NewReader(tt.args.r))
			got, err := packRequestToMessage(request)
			if (err != nil) != tt.wantErr {
				t.Errorf("packRequestToMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("packRequestToMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}