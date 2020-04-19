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

package log

import (
    "fmt"
    "io"
    "os"
    "path"
    "runtime"
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

type Logger struct {
    severity_thresh Severity
    writer io.Writer
    ts_func TimestampFunc
    prefix string
    lock_chan chan bool
    syslog_writer SyslogLike
}

const (
    flag_is_syslog uint32 = 1 << iota
)

func (l *Logger) set_output(w io.Writer) {
    l.writer = w
    if sysl, ok := w.(SyslogLike); ok {
        l.syslog_writer = sysl
    } else {
        l.syslog_writer = nil
    }
}

func New(w io.Writer, sev Severity, prefix string) (*Logger) {
    if prefix == "" {
        prefix = fmt.Sprintf("%s [%d] ", path.Base(os.Args[0]), os.Getpid())
    }

    l := new(Logger)
    l.ts_func = default_ts_func
    l.set_output(w)
    l.severity_thresh = sev
    l.prefix = prefix
    l.lock_chan = make(chan bool, 1)

    return l
}

func (l *Logger) Alert(m string) error {
    if l.severity_thresh < LOG_ALERT {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Alert(l.get_output(1, m, flag_is_syslog))
    }
    return l.Output(1, m)
}

func (l *Logger) Alertf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_ALERT {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Alert(l.get_output(1, fmt.Sprintf(format, v...),
            flag_is_syslog))
    }
    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Crit(m string) error {
    if l.severity_thresh < LOG_CRIT {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Crit(l.get_output(1, m, flag_is_syslog))
    }
    return l.Output(1, m)
}

func (l *Logger) Critf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_CRIT {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Crit(l.get_output(1, fmt.Sprintf(format, v...),
            flag_is_syslog))
    }
    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(m string) error {
    if l.severity_thresh < LOG_DEBUG {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Debug(l.get_output(1, m, flag_is_syslog))
    }
    return l.Output(1, m)
}

func (l *Logger) Debugf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_DEBUG {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Debug(l.get_output(1, fmt.Sprintf(format, v...),
            flag_is_syslog))
    }
    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Emerg(m string) error {
    if l.severity_thresh < LOG_EMERG {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Emerg(l.get_output(1, m, flag_is_syslog))
    }
    return l.Output(1, m)
}

func (l *Logger) Emergf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_EMERG {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Emerg(l.get_output(1, fmt.Sprintf(format, v...),
            flag_is_syslog))
    }
    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Err(m string) error {
    if l.severity_thresh < LOG_ERR {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Err(l.get_output(1, m, flag_is_syslog))
    }
    return l.Output(1, m)
}

func (l *Logger) Errf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_ERR {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Err(l.get_output(1, fmt.Sprintf(format, v...),
            flag_is_syslog))
    }
    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(m string) error {
    if l.severity_thresh < LOG_INFO {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Info(l.get_output(1, m, flag_is_syslog))
    }
    return l.Output(1, m)
}

func (l *Logger) Infof(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_INFO {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Alert(l.get_output(1, fmt.Sprintf(format, v...),
            flag_is_syslog))
    }
    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Notice(m string) error {
    if l.severity_thresh < LOG_NOTICE {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Notice(l.get_output(1, m, flag_is_syslog))
    }
    return l.Output(1, m)
}

func (l *Logger) Noticef(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_NOTICE {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Notice(l.get_output(1, fmt.Sprintf(format, v...),
            flag_is_syslog))
    }
    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Warning(m string) error {
    if l.severity_thresh < LOG_WARNING {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Warning(l.get_output(1, m, flag_is_syslog))
    }
    return l.Output(1, m)
}

func (l *Logger) Warningf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_WARNING {
        return nil
    }
    if l.syslog_writer != nil {
        return l.syslog_writer.Warning(l.get_output(1,
            fmt.Sprintf(format, v...), flag_is_syslog))
    }
    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Write(b []byte) (int, error) {
    if l.syslog_writer != nil {
        str := l.get_output(1, string(b), flag_is_syslog)
        _, err := l.syslog_writer.Write([]byte(str))
        return len(b), err
    }

    err := l.Output(1, string(b))

    return len(b), err
}

func (l *Logger) Fatal(v ...interface{}) {
    l.Print(v...)
    os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
    l.Printf(format, v...)
    os.Exit(1)
}

func (l *Logger) Fatalln(v ...interface{}) {
    l.Println(v...)
    os.Exit(1)
}

func (l *Logger) get_lock() {
    l.lock_chan <- true
}

func (l *Logger) release_lock() {
    <-l.lock_chan
}

func (l *Logger) get_output(call_depth int, s string, flags uint32) string {
    parts := make([]string, 0, 4)

    // Leave out the timestamp and prefix if the writer looks like syslog, since
    // syslog will add these itself.
    if (flags & flag_is_syslog) == 0 {
        ts := l.ts_func()
        parts = append(parts, ts + " ")
        parts = append(parts, l.prefix)
    }

    _, file_name, line, _ := runtime.Caller(call_depth + 1)
    source := fmt.Sprintf("%s:%d", path.Base(file_name), line)
    parts = append(parts, source + ": ")

    if !strings.HasSuffix(s, "\n") {
        s += "\n"
    }

    parts = append(parts, s)

    return strings.Join(parts, "")
}

func (l *Logger) Output(call_depth int, s string) error {
    out_str := l.get_output(call_depth + 1, s, 0)

    l.get_lock()
    defer l.release_lock()

    _, err := fmt.Fprint(l.writer, out_str)
    return err
}

func (l *Logger) Panic(v ...interface{}) {
    if l.syslog_writer != nil {
        str := l.get_output(1, fmt.Sprint(v...), flag_is_syslog)
        l.syslog_writer.Write([]byte(str))
    } else {
        l.Output(1, fmt.Sprint(v...))
    }

    panic(fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
    if l.syslog_writer != nil {
        str := l.get_output(1,
            fmt.Sprintf(format, v...), flag_is_syslog)
        l.syslog_writer.Write([]byte(str))
    } else {
        l.Output(1, fmt.Sprintf(format, v...))
    }

    panic(fmt.Sprintf(format, v...))
}


func (l *Logger) Panicln(v ...interface{}) {
    if l.syslog_writer != nil {
        str := l.get_output(1, fmt.Sprintln(v...), flag_is_syslog)
        l.syslog_writer.Write([]byte(str))
    } else {
        l.Output(1, fmt.Sprint(v...))
    }

    panic(fmt.Sprintln(v...))
}

func (l *Logger) Print(v ...interface{}) {
    m := fmt.Sprint(v...)
    l.Write([]byte(m))
}

func (l *Logger) Printf(format string, v ...interface{}) {
    m := fmt.Sprintf(format, v...)
    l.Write([]byte(m))
}

func (l *Logger) Println(v ...interface{}) {
    m := fmt.Sprintln(v...)
    l.Write([]byte(m))
}

func default_ts_func() string {
    t := time.Now().UTC()
    return t.Format(time.RFC3339)
}
