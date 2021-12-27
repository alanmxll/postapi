package database

import "testing"

func TestDB_Open(t *testing.T) {
	tests := []struct {
		name    string
		d       *DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Open(); (err != nil) != tt.wantErr {
				t.Errorf("DB.Open() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Close(t *testing.T) {
	tests := []struct {
		name    string
		d       *DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Close(); (err != nil) != tt.wantErr {
				t.Errorf("DB.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
