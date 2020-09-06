package main
package main

import (
  log "github.com/sirupsen/logrus"
  "docker.io/go-docker/api/types"
)

func main() {

  client := &Client{
    client: newMockClient(errorMock(http.StatusInternalServerError, "Server error")),
  }

  err := client.ContainerStart(context.Background(), "nothing", types.ContainerStartOptions{})
  if err == nil || err.Error() != "Error reponse from daemon: Server error" {
    t.Fatalf("expected a Server Error, got %v", err)
  }

  expectedURL := "/containers/container_id/start"
  client := &Client{
    client: newMockClient(func(req *http.Request) (*http.Response, error) {
      if !strings.HasPrefix(req.URL.Path, expectedURL) {
        return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
      }
    }),
  }

}
