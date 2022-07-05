package authorization

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

type MockReadCloser struct {
	mock.Mock
}

func (m *MockReadCloser) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func (m *MockReadCloser) Close() error {
	return nil
}

func Test_Authenticate(t *testing.T) {

	cases := []struct {
		testName string
		client   MockClient
		err      string
	}{
		{"Success",
			MockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
				resp := oauthResponse{AccessToken: "123", ExpiresIn: 3600}
				body, _ := json.Marshal(resp)
				r := ioutil.NopCloser(bytes.NewReader(body))
				return &http.Response{StatusCode: http.StatusOK, Body: r}, nil
			}},
			"",
		},
		{"Client error",
			MockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{}, errors.New("default error")
			}},
			"default error",
		},
		{"ReadAll error",
			MockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{Body: &MockReadCloser{}}, nil
			}},
			"read error",
		},
		{"Invalid json error",
			MockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
				json := ``
				r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
				return &http.Response{Body: r}, nil
			}},
			"unexpected end of JSON input",
		},
		{"Status code error",
			MockClient{DoFunc: func(req *http.Request) (*http.Response, error) {
				json := `{"error":"invalid token"}`
				r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
				return &http.Response{StatusCode: http.StatusBadRequest, Body: r}, nil
			}},
			"invalid token",
		},
	}

	for _, tc := range cases {
		t.Run(tc.testName, func(t *testing.T) {
			caradhras := New("", "", Homologation)
			client = &tc.client
			token, err := caradhras.GetAccessToken()
			if tc.err != "" {
				assert.Equal(t, err.Error(), tc.err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, token, "123")
				assert.Equal(t, caradhras.IsExpired(), false)

				caradhras.ExpireAccessToken()
				assert.Equal(t, caradhras.IsExpired(), true)
			}
		})
	}
}
