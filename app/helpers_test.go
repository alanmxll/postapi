package app

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/alanmxll/postapi/app/models"
)

func Test_parse(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		r    *http.Request
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parse(tt.args.w, tt.args.r, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sendResponse(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		in1    *http.Request
		data   interface{}
		status int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendResponse(tt.args.w, tt.args.in1, tt.args.data, tt.args.status)
		})
	}
}

func Test_mapPostToJson(t *testing.T) {
	type args struct {
		p *models.Post
	}
	tests := []struct {
		name string
		args args
		want models.JsonPost
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapPostToJson(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapPostToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
