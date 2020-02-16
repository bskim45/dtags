package api_client

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func createSearchMock(url string) ApiClient {
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
			"testdata/quay/search_nginx.json",
			"https://quay.io/api/v1/find/repositories?query=nginx",
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

func TestApiClient_SearchRepo(t *testing.T) {
	type args struct {
		query string
	}
	type want struct {
		count int
		first string
	}

	tests := []struct {
		name     string
		endpoint string
		args     args
		want     want
		wantErr  bool
	}{
		{
			name:     "quay.io/bitnami/nginx",
			endpoint: "https://quay.io",
			args: args{
				query: "nginx",
			},
			want: want{
				count: 10,
				first: "kubernetes-ingress-controller/nginx-ingress-controller",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := createSearchMock(tt.endpoint)
			got, err := c.SearchRepo(tt.args.query, 100)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchRepo() error = %v, wantErr %v", err, tt.wantErr)
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
