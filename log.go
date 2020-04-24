// BSD 2-Clause License
//
// Copyright (c) 2020 Don Owens <don@regexguy.com>.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

// A logger that logs to a file or any io.Writer. If the io.Writer looks like
// syslog, then certain parts of the log line will not be generated, as syslog
// is expected to cover those.
//
// This logger looks like syslog, in that it allows for specifying severities
// when logging. Unlike syslog, however, instead of setting a default severity
// for logging, a severity threshold is specified at logger creation time that
// limits logging to that severity and anything more important. The intention is
// to enable it to be a drop-in replacement, in terms of the API, for the
// standard log/syslog or log.
//
// Installation:
//   go get github.com/cuberat/go-log
package log

// TODO: SeverityFromString

import (
    "fmt"
    "io"
    "os"
    "path"
    // "runtime"
    "strings"
    "time"
)

type Severity int

// Severities
const (
    LOG_EMERG Severity = iota
    LOG_ALERT
    LOG_CRIT
    LOG_ERR
    LOG_WARNING
    LOG_NOTICE
    LOG_INFO
    LOG_DEBUG
)

var (
    sev_string_to_sev map[string]Severity
    default_logger *Logger
)

type SyslogLike interface {
    Alert(m string) error
    Close() error
    Crit(m string) error
    Debug(m string) error
    Emerg(m string) error
    Err(m string) error
    Info(m string) error
    Notice(m string) error
    Warning(m string) error
    Write(b []byte) (int, error)
}

type TimestampFunc func() (string)

func init() {
    sev_string_to_sev = map[string]Severity{
        "log_emerg": LOG_EMERG,
        "log_alert": LOG_ALERT,
        "log_crit": LOG_CRIT,
        "log_err": LOG_ERR,
        "log_error": LOG_ERR,
        "log_warning": LOG_WARNING,
        "log_warn": LOG_WARNING,
        "log_notice": LOG_NOTICE,
        "log_not": LOG_NOTICE,
        "log_info": LOG_INFO,
        "log_debug": LOG_DEBUG,
    }

    default_logger = New(os.Stderr, LOG_DEBUG, "")
}

func SeverityFromString(sev_string string) (Severity, error) {
    check_sev := strings.ToLower(sev_string)
    if sev, ok := sev_string_to_sev[check_sev]; ok {
        return sev, nil
    }
    check_sev = "log_" + check_sev
    if sev, ok := sev_string_to_sev[check_sev]; ok {
        return sev, nil
    }

    return Severity(0), fmt.Errorf("Unknown severity %q", sev_string)
}

func SetOutput(w io.Writer) {
    default_logger.SetOutput(w)
}

func SetSeverityThreshold(sev_thresh Severity) {
    default_logger.SetSeverityThreshold(sev_thresh)
}

func SetPrefix(prefix string) {
    default_logger.SetPrefix(prefix)
}

// Creates a logger from an io.Writer, with the given severity threshold and
// prefix string.
func New(w io.Writer, sev_thresh Severity, prefix string) (*Logger) {
    if prefix == "" {
        prefix = fmt.Sprintf("%s [%d] ", path.Base(os.Args[0]), os.Getpid())
    }

    l := new(Logger)
    l.ts_func = default_ts_func
    l.set_output(w)
    l.SetSeverityThreshold(sev_thresh)
    l.SetPrefix(prefix)

    l.lock_chan = make(chan bool, 1)

    return l
}

// Creates a logger that will write to the specified file path, with the given
// severity threshold and prefix string
func NewFromFile(file_path string, sev_thresh Severity, prefix string) (*Logger,
    error) {
    fh, err := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }

    return New(fh, sev_thresh, prefix), nil
}

func default_ts_func() string {
    t := time.Now().UTC()
    return t.Format(time.RFC3339)
}

func Alert(m string) error {
    return default_logger.log_sev(1, LOG_ALERT, m)
}

func Alertf(format string, v ...interface{}) error {
    return default_logger.log_sevf(1, LOG_ALERT, format, v...)
}

func Crit(m string) error {
    return default_logger.log_sev(1, LOG_CRIT, m)
}

func Critf(format string, v ...interface{}) error {
    return default_logger.log_sevf(1, LOG_CRIT, format, v...)
}

func Debug(m string) error {
    return default_logger.log_sev(1, LOG_DEBUG, m)
}

func Debugf(format string, v ...interface{}) error {
    return default_logger.log_sevf(1, LOG_DEBUG, format, v...)
}

func Emerg(m string) error {
    return default_logger.log_sev(1, LOG_EMERG, m)
}

func Emergf(format string, v ...interface{}) error {
    return default_logger.log_sevf(1, LOG_EMERG, format, v...)
}

func Err(m string) error {
    return default_logger.log_sev(1, LOG_ERR, m)
}

func Errf(format string, v ...interface{}) error {
    return default_logger.log_sevf(1, LOG_ERR, format, v...)
}

func Info(m string) error {
    return default_logger.log_sev(1, LOG_INFO, m)
}

func Infof(format string, v ...interface{}) error {
    return default_logger.log_sevf(1, LOG_INFO, format, v...)
}

func Notice(m string) error {
    return default_logger.log_sev(1, LOG_NOTICE, m)
}

func Noticef(format string, v ...interface{}) error {
    return default_logger.log_sevf(1, LOG_NOTICE, format, v...)
}

func Warning(m string) error {
    return default_logger.log_sev(1, LOG_WARNING, m)
}

func Warningf(format string, v ...interface{}) error {
    return default_logger.log_sevf(1, LOG_WARNING, format, v...)
}

// FIXME: fix up call_depth
// Equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
    default_logger.outputv(1, v...)
    os.Exit(1)
}

// Equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
    default_logger.outputf(1, format, v...)
    os.Exit(1)
}

// Equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
    default_logger.outputlnv(1, v...)
    os.Exit(1)
}

func Panic(v ...interface{}) {
    default_logger.outputv(1, v...)
    panic(fmt.Sprint(v...))
}

func Panicf(format string, v ...interface{}) {
    default_logger.outputf(1, format, v...)
    panic(fmt.Sprintf(format, v...))
}

func Panicln(v ...interface{}) {
    default_logger.outputlnv(1, v...)
    panic(fmt.Sprintln(v...))
}

// Prints to the logger. Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
    default_logger.outputv(1, v...)
}

// Prints to the logger. Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
    default_logger.outputf(1, format, v...)
}

// Prints to the logger. Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
    default_logger.outputlnv(1, v...)
}
