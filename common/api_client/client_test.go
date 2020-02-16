package api_client

import (
	"reflect"
	"testing"

	"github.com/bskim45/dtags/common/api_client/general"
	"github.com/bskim45/dtags/common/api_client/hub"
)

func TestNew(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    reflect.Type
		wantErr bool
	}{
		{
			name:    "hub",
			args:    args{"https://index.docker.io"},
			want:    reflect.TypeOf(&hub.HubApiClient{}),
			wantErr: false,
		},
		{
			name:    "hub",
			args:    args{"https://docker.elastic.co"},
			want:    reflect.TypeOf(&general.GeneralApiClient{}),
			wantErr: false,
		},
		{
			name:    "hub",
			args:    args{"https:/localhost"},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(reflect.TypeOf(got), tt.want) {
				t.Errorf("New() got = %v, want %v", reflect.TypeOf(got), tt.want)
			}
		})
	}
}
