// learn testcontainers + testify
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"observer/pkg/env"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testContainer testcontainers.Container
var baseUrl string

func TestMain(m *testing.M) {
	env.GetEnv()

	appName := os.Getenv("APP_NAME")
	appPort := os.Getenv("APP_PORT")

	ctx := context.Background()

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        fmt.Sprintf("%s:latest", appName),
			ExposedPorts: []string{fmt.Sprintf("%s/tcp", appPort)}, // container port only
			WaitingFor:   wait.ForListeningPort(fmt.Sprintf("%s/tcp", appPort)),
			AutoRemove:   true,
		},
		Started: true,
	})
	testContainer = container
	if err != nil {
		panic("Can't run Docker container. Probably Docker core isn't run.")
	}
	mappedPort, _ := container.MappedPort(ctx, appPort)
	baseUrl = fmt.Sprintf("http://localhost:%s", mappedPort.Port())

	code := m.Run()
	container.Terminate(ctx)
	os.Exit(code)
}

func TestBasic(t *testing.T) {
	resp, err := http.Get(baseUrl)
	require.NoError(t, err)

	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, "Hi mate! It's main page afaik.", string(body))
}

func TestPage(t *testing.T) {
	testCases := []struct {
		name        string
		page        int
		status_code int
	}{
		{"invalid page", -1, 404},
		{"valid page", 10, 200},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("%s?page=%d", baseUrl, tc.page)
			resp, err := http.Get(url)
			require.NoError(t, err)

			defer resp.Body.Close()

			assert.Equal(t, tc.status_code, resp.StatusCode)

			if tc.status_code == 200 {
				body, _ := io.ReadAll(resp.Body)
				assert.Contains(t, string(body), fmt.Sprintf("#%d", tc.page))
			}
		})
	}
}
