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

// TODO:
//   - add tests for syslog (Sysloglike) for all severity levels
//      - create an type that implements Sysloglike

package log_test

import (
    "bytes"
    "fmt"
    "io"
    log "github.com/cuberat/go-log"
    "strings"
    "testing"
)

type LogTestFunc func(*log.Logger) ()
type LogPkgTestFunc func() ()

type LogFunc func(string) error

type LogLevelTester struct {
    Key string
    LogIt LogTestFunc
    LogSev log.Severity
}
type LogPkgLevelTester struct {
    Key string
    LogIt LogPkgTestFunc
    LogSev log.Severity
}

type SevTestStr struct {
    Name string
    ExpectedSev log.Severity
}

type SyslogLikeLogger struct {
    Writer io.Writer
}

func TestSeverityFromString(t *testing.T) {
    tests := []*SevTestStr{
        &SevTestStr{"emerg", log.LOG_EMERG},
        &SevTestStr{"log_emerg", log.LOG_EMERG},
        &SevTestStr{"alert", log.LOG_ALERT},
        &SevTestStr{"crit", log.LOG_CRIT},
        &SevTestStr{"err", log.LOG_ERR},
        &SevTestStr{"error", log.LOG_ERR},
        &SevTestStr{"warning", log.LOG_WARNING},
        &SevTestStr{"warn", log.LOG_WARNING},
        &SevTestStr{"log_warn", log.LOG_WARNING},
        &SevTestStr{"notice", log.LOG_NOTICE},
        &SevTestStr{"info", log.LOG_INFO},
        &SevTestStr{"debug", log.LOG_DEBUG},
    }

    for _, tester := range tests {
        test_func := func(t *testing.T) {
            found_sev, err := log.SeverityFromString(tester.Name)
            if err != nil {
                t.Errorf("Got error from SeverityFromString for name %q: %s",
                    tester.Name, err)
            }
            if found_sev != tester.ExpectedSev {
                t.Errorf("Incorrect severity: got %d, expected %d", found_sev,
                    tester.ExpectedSev)
            }
        }

        if !t.Run(tester.Name, test_func) {
            t.Errorf("t.Run() failed when testing %q", tester.Name)
        }
    }
}

func TestSeverityLogLevels(t *testing.T) {
    _, _, _, all, _ := get_level_before_after("emerg")
    for _, tester := range all {
        test_func := get_level_test_func(tester)
        if !t.Run(tester.Key, test_func) {
            t.Errorf("t.Run() failed when testing %q", tester.Key)
        }
    }
}

func get_level_test_func(conf *LogLevelTester) (func(*testing.T)) {
    test_func := func (t *testing.T) {
        buffer := new(bytes.Buffer)
        logger := log.New(buffer, conf.LogSev, "")
        key := conf.Key

        before, current, after, _, log_all := get_level_before_after(key)

        log_all(logger)
        log_result := buffer.String()
        // t.Logf("got log: %s", log_result)

        for _, tester := range before {
            if !strings.Contains(log_result, tester.Key) {
                t.Errorf("log should contain %q", tester.Key)
            }
        }

        if !strings.Contains(log_result, current.Key) {
            t.Errorf("log should contain %q", current.Key)
        }

        for _, tester := range after {
            if strings.Contains(log_result, tester.Key) {
                t.Errorf("log should NOT contain %q", tester.Key)
            }
        }
    }

    return test_func
}

func TestSeverityLogLevelsSyslog(t *testing.T) {
    _, _, _, all, _ := get_level_before_after("emerg")
    for _, tester := range all {
        test_func := get_syslog_like_level_test_func(tester)
        if !t.Run(tester.Key, test_func) {
            t.Errorf("t.Run() failed when testing %q", tester.Key)
        }
    }
}

func get_syslog_like_level_test_func(conf *LogLevelTester) (func(*testing.T)) {
    test_func := func (t *testing.T) {
        buffer := new(bytes.Buffer)
        syslog_logger := new(SyslogLikeLogger)
        syslog_logger.Writer = buffer
        logger := log.New(syslog_logger, conf.LogSev, "")
        key := conf.Key

        before, current, after, _, log_all := get_level_before_after(key)

        log_all(logger)
        log_result := buffer.String()
        t.Logf("got log: %s", log_result)

        for _, tester := range before {
            if !strings.Contains(log_result, tester.Key) {
                t.Errorf("log should contain %q", tester.Key)
            }
        }

        if !strings.Contains(log_result, current.Key) {
            t.Errorf("log should contain %q", current.Key)
        }

        for _, tester := range after {
            if strings.Contains(log_result, tester.Key) {
                t.Errorf("log should NOT contain %q", tester.Key)
            }
        }
    }

    return test_func
}

func TestPkgSeverityLogLevels(t *testing.T) {
    _, _, _, all, _ := get_pkg_level_before_after("emerg")
    for _, tester := range all {
        test_func := get_pkg_level_test_func(tester)
        if !t.Run(tester.Key, test_func) {
            t.Errorf("t.Run() failed when testing %q", tester.Key)
        }
    }
}

func get_pkg_level_test_func(conf *LogPkgLevelTester) (func(*testing.T)) {
    test_func := func (t *testing.T) {
        buffer := new(bytes.Buffer)
        key := conf.Key

        log.SetOutput(buffer)
        log.SetSeverityThreshold(conf.LogSev)

        before, current, after, _, log_all := get_pkg_level_before_after(key)

        log_all()
        log_result := buffer.String()
        // t.Logf("got log: %s", log_result)

        for _, tester := range before {
            if !strings.Contains(log_result, tester.Key) {
                t.Errorf("log should contain %q", tester.Key)
            }
        }

        if !strings.Contains(log_result, current.Key) {
            t.Errorf("log should contain %q", current.Key)
        }

        for _, tester := range after {
            if strings.Contains(log_result, tester.Key) {
                t.Errorf("log should NOT contain %q", tester.Key)
            }
        }
    }

    return test_func
}

func TestLevelEmerg(t *testing.T) {
    buffer := new(bytes.Buffer)
    logger := log.New(buffer, log.LOG_EMERG, "")
    key := "emerg"

    before, current, after, _, log_all := get_level_before_after(key)

    log_all(logger)
    log_result := buffer.String()
    // t.Logf("got log: %s", log_result)

    for _, tester := range before {
        if !strings.Contains(log_result, tester.Key) {
            t.Errorf("log should contain %q", tester.Key)
        }
    }

    if !strings.Contains(log_result, current.Key) {
        t.Errorf("log should contain %q", current.Key)
    }

    for _, tester := range after {
        if strings.Contains(log_result, tester.Key) {
            t.Errorf("log should NOT contain %q", tester.Key)
        }
    }
}

func get_level_before_after(level string) ([]*LogLevelTester,
    *LogLevelTester, []*LogLevelTester, []*LogLevelTester,
    func(*log.Logger)()) {

    var (
        before []*LogLevelTester
        current *LogLevelTester
        after []*LogLevelTester
    )

    config := []*LogLevelTester{
        &LogLevelTester{"emerg", log_emerg, log.LOG_EMERG},
        &LogLevelTester{"alert", log_alert, log.LOG_ALERT},
        &LogLevelTester{"crit", log_crit, log.LOG_CRIT},
        &LogLevelTester{"err", log_err, log.LOG_ERR},
        &LogLevelTester{"warning", log_warning, log.LOG_WARNING},
        &LogLevelTester{"notice", log_notice, log.LOG_NOTICE},
        &LogLevelTester{"info", log_info, log.LOG_INFO},
        &LogLevelTester{"debug", log_debug, log.LOG_DEBUG},
    }

    found := false
    for _, tester := range config {
        if tester.Key == level {
            current = tester
            found = true
            continue
        }
        if found {
            after = append(after, tester)
            continue
        }
        before = append(before, tester)
    }

    log_all := func(logger *log.Logger) {
        for _, tester := range config {
            tester.LogIt(logger)
        }
    }

    return before, current, after, config, log_all
}

func get_pkg_level_before_after(level string) ([]*LogPkgLevelTester,
    *LogPkgLevelTester, []*LogPkgLevelTester, []*LogPkgLevelTester,
    func()()) {

    var (
        before []*LogPkgLevelTester
        current *LogPkgLevelTester
        after []*LogPkgLevelTester
    )

    config := []*LogPkgLevelTester{
        &LogPkgLevelTester{"emerg", pkg_log_emerg, log.LOG_EMERG},
        &LogPkgLevelTester{"alert", pkg_log_alert, log.LOG_ALERT},
        &LogPkgLevelTester{"crit", pkg_log_crit, log.LOG_CRIT},
        &LogPkgLevelTester{"err", pkg_log_err, log.LOG_ERR},
        &LogPkgLevelTester{"warning", pkg_log_warning, log.LOG_WARNING},
        &LogPkgLevelTester{"notice", pkg_log_notice, log.LOG_NOTICE},
        &LogPkgLevelTester{"info", pkg_log_info, log.LOG_INFO},
        &LogPkgLevelTester{"debug", pkg_log_debug, log.LOG_DEBUG},
    }

    found := false
    for _, tester := range config {
        if tester.Key == level {
            current = tester
            found = true
            continue
        }
        if found {
            after = append(after, tester)
            continue
        }
        before = append(before, tester)
    }

    log_all := func() {
        for _, tester := range config {
            tester.LogIt()
        }
    }

    return before, current, after, config, log_all
}

// Test funcs for logger objects
func log_emerg(logger *log.Logger) {
    logger.Emerg("emerg")
}

func log_alert(logger *log.Logger) {
    logger.Alert("alert")
}

func log_crit(logger *log.Logger) {
    logger.Crit("crit")
}

func log_err(logger *log.Logger) {
    logger.Err("err")
}

func log_warning(logger *log.Logger) {
    logger.Warning("warning")
}

func log_notice(logger *log.Logger) {
    logger.Notice("notice")
}

func log_info(logger *log.Logger) {
    logger.Info("info")
}

func log_debug(logger *log.Logger) {
    logger.Debug("debug")
}

// Test funcs for log pkg.
func pkg_log_emerg() {
    log.Emerg("emerg")
}

func pkg_log_alert() {
    log.Alert("alert")
}

func pkg_log_crit() {
    log.Crit("crit")
}

func pkg_log_err() {
    log.Err("err")
}

func pkg_log_warning() {
    log.Warning("warning")
}

func pkg_log_notice() {
    log.Notice("notice")
}

func pkg_log_info() {
    log.Info("info")
}

func pkg_log_debug() {
    log.Debug("debug")
}

func TestOneAlert(t *testing.T) {
    buffer := new(bytes.Buffer)
    logger := log.New(buffer, log.LOG_ALERT, "")
    logger.Alert("foo")
    logger.Alertf("stuff=%s", "bar")
    logger.Crit("boo")

    // alert := logger.Alertf
    // alert("detached func test: type is %T", alert)

    log_str := buffer.String()
    t.Logf("log: %s", log_str)

    if !strings.Contains(log_str, "stuff=bar") {
        t.Error("Output should contain \"stuff=bar\"")
    }

    if strings.Contains(log_str, "boo") {
        t.Error("Output should NOT contain \"boo\"")
    }

    t.Logf("log.LOG_EMERG=%d", log.LOG_EMERG)
    t.Logf("log.LOG_DEBUG=%d", log.LOG_DEBUG)
}

// Syslog-like writer for testing syslog and look-alikes
func (l *SyslogLikeLogger) Alert(m string) error {
    _, err := fmt.Fprintln(l.Writer, m)
    return err
}

func (l *SyslogLikeLogger) Close() error {
    return nil
}

func (l *SyslogLikeLogger) Crit(m string) error {
    _, err := fmt.Fprintln(l.Writer, m)
    return err
}

func (l *SyslogLikeLogger) Debug(m string) error {
    _, err := fmt.Fprintln(l.Writer, m)
    return err
}

func (l *SyslogLikeLogger) Emerg(m string) error {
    _, err := fmt.Fprintln(l.Writer, m)
    return err
}

func (l *SyslogLikeLogger) Err(m string) error {
    _, err := fmt.Fprintln(l.Writer, m)
    return err
}

func (l *SyslogLikeLogger) Info(m string) error {
    _, err := fmt.Fprintln(l.Writer, m)
    return err
}

func (l *SyslogLikeLogger) Notice(m string) error {
    _, err := fmt.Fprintln(l.Writer, m)
    return err
}

func (l *SyslogLikeLogger) Warning(m string) error {
    _, err := fmt.Fprintln(l.Writer, m)
    return err
}

func (l *SyslogLikeLogger) Write(b []byte) (int, error) {
    return fmt.Fprintf(l.Writer, "%s\n", b)
}
