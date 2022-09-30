package service

import (
	"testing"

	"github.com/voltgizerz/go-redis/config"
)

func TestBook_ReadBook(t *testing.T) {
	type fields struct {
		Redis config.Redis
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				Redis: tt.fields.Redis,
			}
			b.ReadBook()
		})
	}
}
