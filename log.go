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
    // stdlog "log"
    // "log/syslog"
    "os"
    "path"
    "runtime"
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

// type SyslogLike interface {
//     Alert(m string) error
//     Close() error
//     Crit(m string) error
//     Debug(m string) error
//     Emerg(m string) error
//     Err(m string) error
//     Info(m string) error
//     Notice(m string) error
//     Warning(m string) error
//     Write(b []byte) (int, error)
// }

type TimestampFunc func() (string)

type Logger struct {
    severity_thresh Severity
    writer io.Writer
    ts_func TimestampFunc
    prefix string
    lock_chan chan bool
    // syslog_writer *syslog.Writer
    // logger *log.Logger
}

// func NewFromLogger(logger *log.Logger, sev Severity) (*Logger) {
//     l := new(Logger)
//     l.logger = logger
//     l.severity_thresh = sev
//
//     return l
// }

// func NewFromSyslog(syslog_writer *syslog.Writer, sev Severity) (*Logger) {
//     l := new(Logger)
//     l.syslog_writer = syslog_writer
//     l.severity_thresh = sev
//
//     return l
// }

func New(w io.Writer, sev Severity, prefix string) (*Logger) {
    if prefix == "" {
        prefix = fmt.Sprintf("%s [%d] ", path.Base(os.Args[0]), os.Getpid())
    }

    l := new(Logger)
    l.ts_func = default_ts_func
    l.writer = w
    l.severity_thresh = sev
    l.prefix = prefix
    l.lock_chan = make(chan bool, 1)

    return l

    // logger := log.New(w, prefix, log.LstdFlags | log.Lshortfile | log.LUTC)
    // return NewFromLogger(logger, sev)
}

func (l *Logger) Alert(m string) error {
    if l.severity_thresh < LOG_ALERT {
        return nil
    }
    return l.Output(1, m)
}

func (l *Logger) Alertf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_ALERT {
        return nil
    }

    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Crit(m string) error {
    if l.severity_thresh < LOG_CRIT {
        return nil
    }
    return l.Output(1, m)
}

func (l *Logger) Critf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_CRIT {
        return nil
    }

    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(m string) error {
    if l.severity_thresh < LOG_DEBUG {
        return nil
    }
    return l.Output(1, m)
}

func (l *Logger) Debugf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_DEBUG {
        return nil
    }

    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Emerg(m string) error {
    if l.severity_thresh < LOG_EMERG {
        return nil
    }
    return l.Output(1, m)
}

func (l *Logger) Emergf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_EMERG {
        return nil
    }

    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Err(m string) error {
    if l.severity_thresh < LOG_ERR {
        return nil
    }
    return l.Output(1, m)
}

func (l *Logger) Errf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_ERR {
        return nil
    }

    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(m string) error {
    if l.severity_thresh < LOG_INFO {
        return nil
    }
    return l.Output(1, m)
}

func (l *Logger) Infof(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_INFO {
        return nil
    }

    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Notice(m string) error {
    if l.severity_thresh < LOG_NOTICE {
        return nil
    }
    return l.Output(1, m)
}

func (l *Logger) Noticef(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_NOTICE {
        return nil
    }

    return l.Output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Warning(m string) error {
    if l.severity_thresh < LOG_WARNING {
        return nil
    }
    return l.Output(1, m)
}

func (l *Logger) Warningf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_WARNING {
        return nil
    }

    return l.Output(1, fmt.Sprintf(format, v...))
}

// func (l *Logger) Fatal(v ...interface{}) {
//
// }
//
// func (l *Logger) Fatalf(format string, v ...interface{}) {
// }

func (l *Logger) get_lock() {
    l.lock_chan <- true
}

func (l *Logger) release_lock() {
    <-l.lock_chan
}

func (l *Logger) Output(calldepth int, s string) error {
    ts := l.ts_func()

    l.get_lock()
    defer l.release_lock()

    _, file_name, line, _ := runtime.Caller(calldepth + 1)
    source := fmt.Sprintf("%s:%d", path.Base(file_name), line)

    out_str := fmt.Sprintf("%s %s%s", ts, l.prefix, source)

    _, err := fmt.Fprintf(l.writer, "%s %s\n", out_str, s)
    return err

    // if l.logger != nil {
    //     return l.logger.Output(calldepth + 1, s)
    // }

    // if l.syslog_writer != nil {
    //
    // }

    // return fmt.Errorf("no logger defined")
}

// func (l *Logger) Panic(v ...interface{}) {
//
// }
//
// func (l *Logger) Panicf(format string, v ...interface{}) {
//
// }
//
// func (l *Logger) Panicln(v ...interface{}) {
//
// }
//
// func (l *Logger) Prefix() string {
//
// }
//
// func (l *Logger) Print(v ...interface{}) {
//
// }
//
// func (l *Logger) Printf(format string, v ...interface{}) {
//
// }

func default_ts_func() string {
    t := time.Now().UTC()
    return t.Format(time.RFC3339)
}
