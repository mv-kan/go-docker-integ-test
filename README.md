# go-docker-integ-test
Golang with Docker container for integration testing. Container is running inner systemd systemctl (has to be run in container privileged mode) and rsyslog port is 10123.

## enter the test container
```
docker exec -it test-remote-logger /bin/bash
```
test-remote-logger - is hardcoded at send_thing_test.go

## How test container is runned? 

- in privileged mode
- in host mode (all ports of container are mapped to host's ports)
