package main

import (
	"fmt"
	"log/syslog"
)

func SendThing(raddr string) error {
	sysLog, err := syslog.Dial("tcp", raddr,
		syslog.LOG_DEBUG, "sendthing")
	if err != nil {
		return err
	}
	fmt.Fprintf(sysLog, "This is a daemon warning with sendthing.")
	err = sysLog.Emerg("HAHAHAHAHAAHHAHAHAHAHA")
	if err != nil {
		return err
	}
	sysLog.Close()
	return nil
}
