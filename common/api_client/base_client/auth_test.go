package base_client

import (
	"reflect"
	"testing"
)

func Test_GetAuthChallenge(t *testing.T) {
	type args struct {
		header string
	}
	tests := []struct {
		name string
		args args
		want *AuthChallenge
	}{
		{
			name: "dockerhub",
			args: args{header: "Bearer realm=\"https://auth.docker.io/token\",service=\"registry.docker.io\",scope=\"repository:samalba/my-app:pull,push\""},
			want: &AuthChallenge{
				Endpoint: "https://auth.docker.io/token",
				Service:  "registry.docker.io",
				Scope:    "repository:samalba/my-app:pull,push",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAuthChallenge(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthChallenge() = %v, want %v", got, tt.want)
			}
		})
	}
}
