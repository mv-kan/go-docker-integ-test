package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/assert"
)

var (
	RSYSLOG_DAEMON_PORT = "10123"
)
var (
	LOG_HOST                             = "localhost"
	LOG_PORT                      string = RSYSLOG_DAEMON_PORT
	RSYSLOG_DAEMON_CONTAINER_NAME        = "test-remote-logger"
	THIS_APP_CONTAINER_NAME              = "test-sendthing"
)

func runThisApp(pool *dockertest.Pool) *dockertest.Resource {
	resource, err := pool.BuildAndRunWithOptions(
		"./Dockerfile.debug",
		&dockertest.RunOptions{
			Privileged: true,
			Name:       THIS_APP_CONTAINER_NAME,
		}, func(hc *docker.HostConfig) {
			hc.NetworkMode = "host"
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	return resource
}

func runRemoteSyslogDaemon(pool *dockertest.Pool) *dockertest.Resource {
	resource, err := pool.BuildAndRunWithOptions(
		"./Dockerfile.rsyslog_daemon",
		&dockertest.RunOptions{
			Privileged: true,
			Name:       RSYSLOG_DAEMON_CONTAINER_NAME,
		}, func(hc *docker.HostConfig) {
			hc.NetworkMode = "host"
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	err = pool.Retry(func() error {
		// ping to ensure that the server is up and running
		_, err := net.Dial("tcp", net.JoinHostPort("localhost", LOG_PORT))
		return err
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	return resource
}

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	allContainers := []*dockertest.Resource{}

	rsyslogDaemon := runRemoteSyslogDaemon(pool)
	allContainers = append(allContainers, rsyslogDaemon)

	code := m.Run()

	// remove all containers
	for _, container := range allContainers {
		if err := pool.Purge(container); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
		// hard kill the container in 3 minutes (180 Seconds)
		container.Expire(180)
	}
	os.Exit(code)
}

func TestSendThingGeneral(t *testing.T) {
	fmt.Println("Send logs to ", LOG_HOST+":"+LOG_PORT)
	err := SendThing(LOG_HOST + ":" + LOG_PORT)
	assert.Nil(t, err)

}
