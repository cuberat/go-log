package main

import (
    "fmt"
    log "github.com/cuberat/go-log"
    syslog "log/syslog"
)

func main() {
    sys_logger, err := syslog.New(syslog.LOG_INFO | syslog.LOG_LOCAL6, "foo-> ")
    if err != nil {
        panic(fmt.Sprintf("Couldn't connect to syslog: %s", err))
    }
    logger := log.New(sys_logger, log.LOG_INFO, "")

    logger.Alert("my alert msg")
}
