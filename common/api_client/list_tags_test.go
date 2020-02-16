package api_client

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func createMock(url string) ApiClient {
	c, err := New(url)
	if err != nil {
		return nil
	}

	httpmock.ActivateNonDefault(c.GetBaseClient().Resty.GetClient())
	httpmock.Reset()

	mocks := []struct {
		dataPath string
		fakeUrl  string
	}{
		{
			"testdata/list_tags_hub_ubuntu.json",

			"https://registry.hub.docker.com/v2/repositories/library/ubuntu/tags/",
		},
		{
			"testdata/list_tags_elastic_elasticsearch.json",
			"https://docker.elastic.co/v2/elasticsearch/elasticsearch/tags/list",
		},
		{
			"testdata/quay/list_tags_nginx.json",
			"https://quay.io/api/v1/repository/bitnami/nginx/tag/",
		},
	}

	for _, mock := range mocks {
		fixture, _ := ioutil.ReadFile(mock.dataPath)
		response := httpmock.NewStringResponse(200, string(fixture))
		response.Header.Set("Content-Type", "application/json")
		responder := httpmock.ResponderFromResponse(response)
		httpmock.RegisterResponder("GET", mock.fakeUrl, responder)
	}

	return c
}

func TestApiClient_ListTags(t *testing.T) {
	type args struct {
		name string
	}
	type want struct {
		count int
		first string
	}

	tests := []struct {
		name    string
		baseUrl string
		args    args
		want    want
		wantErr bool
	}{
		{
			name:    "ubuntu",
			baseUrl: "https://registry.hub.docker.com",
			args: args{
				name: "library/ubuntu",
			},
			want: want{
				count: 100,
				first: "xenial-20200114",
			},
		},
		{
			name:    "elasticsearch",
			baseUrl: "https://docker.elastic.co",
			args: args{
				name: "elasticsearch/elasticsearch",
			},
			want: want{
				count: 191,
				first: "master-SNAPSHOT",
			},
		},
		{
			name:    "quay.io/bitnami/nginx",
			baseUrl: "https://quay.io",
			args: args{
				name: "bitnami/nginx",
			},
			want: want{
				count: 50,
				first: "latest",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := createMock(tt.baseUrl)
			got, err := c.ListTags(tt.args.name, 100)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			//noinspection GoNilness
			if assert.NotNil(t, got) {
				assert.Equal(t, tt.want.count, len(*got))

				if assert.NotNil(t, *got) {
					assert.Equal(t, tt.want.first, (*got)[0])
				}
			}
		})
	}
	httpmock.DeactivateAndReset()
}
