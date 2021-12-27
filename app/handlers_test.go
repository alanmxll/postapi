package app

import (
	"net/http"
	"reflect"
	"testing"
)

func TestApp_IndexHandler(t *testing.T) {
	tests := []struct {
		name string
		a    *App
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.IndexHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.IndexHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_CreatePostHandler(t *testing.T) {
	tests := []struct {
		name string
		a    *App
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.CreatePostHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.CreatePostHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_GetPostsHandler(t *testing.T) {
	tests := []struct {
		name string
		a    *App
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetPostsHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.GetPostsHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_UpdatePostHandler(t *testing.T) {
	tests := []struct {
		name string
		a    *App
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.UpdatePostHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.UpdatePostHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_DeletePostHandler(t *testing.T) {
	tests := []struct {
		name string
		a    *App
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.DeletePostHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.DeletePostHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
