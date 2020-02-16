package common

import (
	"net/url"
	"reflect"
	"testing"
)

func TestCheckDockerHubLibrary(t *testing.T) {
	type args struct {
		endpoint string
		q        *ImageName
	}
	tests := []struct {
		name string
		args args
		want *ImageName
	}{
		{
			name: "library/python",
			args: args{
				endpoint: "index.docker.io",
				q: &ImageName{
					Endpoint: "",
					Repo:     "",
					Name:     "python",
				},
			},
			want: &ImageName{
				Repo: "library",
				Name: "python",
			},
		},
		{
			name: "quay.io/bitnami/nginx",
			args: args{
				endpoint: "quay.io",
				q: &ImageName{
					Endpoint: "quay.io",
					Repo:     "bitnami",
					Name:     "nginx",
				},
			},
			want: &ImageName{
				Endpoint: "quay.io",
				Repo:     "bitnami",
				Name:     "nginx",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckDockerHubLibrary(tt.args.endpoint, tt.args.q)
			if !reflect.DeepEqual(tt.args.q, tt.want) {
				t.Errorf("GetImageName() = %v, want %v", tt.args.q, tt.want)
			}
		})
	}
}

func TestImageName_GetImageName(t *testing.T) {
	type fields struct {
		Endpoint string
		Repo     string
		Name     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ubuntu",
			fields: fields{
				Repo: "library",
				Name: "ubuntu",
			},
			want: "library/ubuntu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := ImageName{
				Endpoint: tt.fields.Endpoint,
				Repo:     tt.fields.Repo,
				Name:     tt.fields.Name,
			}
			if got := i.GetImageName(); got != tt.want {
				t.Errorf("GetImageName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDockerHub(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "index.docker.io", args: args{endpoint: "https://index.docker.io"}, want: true},
		{name: "quay.io", args: args{endpoint: "https://quay.io"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDockerHub(tt.args.endpoint); got != tt.want {
				t.Errorf("IsDockerHub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsQuay(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "index.docker.io", args: args{endpoint: "https://index.docker.io"}, want: false},
		{name: "quay.io", args: args{endpoint: "https://quay.io"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsQuay(tt.args.endpoint); got != tt.want {
				t.Errorf("IsQuay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *ImageName
	}{
		{
			name: "ubuntu",
			args: args{name: "ubuntu"},
			want: &ImageName{
				Name: "ubuntu",
			},
		},
		{
			name: "bitnami/nginx",
			args: args{name: "bitnami/nginx"},
			want: &ImageName{
				Repo: "bitnami",
				Name: "nginx",
			},
		},
		{
			name: "quay.io/bitnami/nginx",
			args: args{name: "quay.io/bitnami/nginx"},
			want: &ImageName{
				Endpoint: "quay.io",
				Repo:     "bitnami",
				Name:     "nginx",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildUrl(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			name: "https://registry.hub.docker.com/",
			args: args{endpoint: "https://registry.hub.docker.com/"},
			want: &url.URL{
				Scheme: "https",
				Host:   "registry.hub.docker.com",
				Path:   "/",
			},
			wantErr: false,
		},
		{
			name: "quay.io",
			args: args{endpoint: "quay.io"},
			want: &url.URL{
				Scheme: "https",
				Host:   "quay.io",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildUrl(tt.args.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
