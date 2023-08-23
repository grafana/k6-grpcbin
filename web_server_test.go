package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebServer(t *testing.T) {
	t.Parallel()

	client := webTestClient{addr: ":8080"}
	code := client.GetIndex(t)
	assert.Equal(t, code, 200)
}

type webTestClient struct {
	addr string
}

func (client *webTestClient) GetIndex(t *testing.T) int {
	resp, err := http.Get("http://localhost" + client.addr)
	require.NoError(t, err)
	return resp.StatusCode
}
