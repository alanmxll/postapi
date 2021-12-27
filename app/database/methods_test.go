package database

import (
	"reflect"
	"testing"

	"github.com/alanmxll/postapi/app/models"
)

func TestDB_CreatePost(t *testing.T) {
	type args struct {
		p *models.Post
	}
	tests := []struct {
		name    string
		d       *DB
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.CreatePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("DB.CreatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_GetPosts(t *testing.T) {
	tests := []struct {
		name    string
		d       *DB
		want    []*models.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.GetPosts()
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.GetPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.GetPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_UpdatePost(t *testing.T) {
	type args struct {
		p *models.Post
	}
	tests := []struct {
		name    string
		d       *DB
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.UpdatePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("DB.UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_DeletePost(t *testing.T) {
	type args struct {
		ID *int64
	}
	tests := []struct {
		name    string
		d       *DB
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.DeletePost(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("DB.DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
