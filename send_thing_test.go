package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"github.com/ory/dockertest"
	"github.com/stretchr/testify/assert"
)

var (
	LOG_HOST = "localhost"
	LOG_PORT string
)

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.BuildAndRunWithOptions(
		"./Dockerfile",
		&dockertest.RunOptions{
			Privileged:   true,
			Name:         "test-remote-logger",
			ExposedPorts: []string{"10123"}})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	tmp := resource.GetPort("10123/tcp")
	fmt.Println(tmp)
	err = pool.Retry(func() error {
		LOG_PORT = resource.GetPort("10123/tcp")
		// ping to ensure that the server is up and running
		_, err := net.Dial("tcp", net.JoinHostPort("localhost", LOG_PORT))
		return err
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	// hard kill the container in 3 minutes (180 Seconds)
	resource.Expire(180)

	os.Exit(code)
}

func TestSendThingGeneral(t *testing.T) {
	fmt.Println("Send logs to ", LOG_HOST+":"+LOG_PORT)
	err := SendThing(LOG_HOST + ":" + LOG_PORT)
	assert.Nil(t, err)
}
