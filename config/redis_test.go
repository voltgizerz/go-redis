package config

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/voltgizerz/go-redis/mocks"
)

func TestRedis_Get(t *testing.T) {
	redisMock := mocks.RedisInterface{}
	redisMock.On("Get").Return("Asd", nil).Once()
}

func TestRedis_Set(t *testing.T) {
	type fields struct {
		client *redis.Client
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Redis{
				client: tt.fields.client,
			}
			r.Set(tt.args.key, tt.args.value)
		})
	}
}
