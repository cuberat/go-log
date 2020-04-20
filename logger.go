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
)

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

// Logs a message with severity LOG_ALERT.
func (l *Logger) Alert(m string) error {
    if l.severity_thresh < LOG_ALERT {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslog(l.syslog_writer.Alert, 1, m)
    }

    return l.output(1, m)
}

// Logs a message with severity LOG_ALERT. Arguments are handled in the manner
// of fmt.Printf.
func (l *Logger) Alertf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_ALERT {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslogf(l.syslog_writer.Alert, 1, format, v...)
    }

    return l.output(1, fmt.Sprintf(format, v...))
}

// Logs a message with severity LOG_CRIT.
func (l *Logger) Crit(m string) error {
    if l.severity_thresh < LOG_CRIT {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslog(l.syslog_writer.Crit, 1, m)
    }
    return l.output(1, m)
}

// Logs a message with severity LOG_CRIT. Arguments are handled in the manner of
// fmt.Printf.
func (l *Logger) Critf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_CRIT {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslogf(l.syslog_writer.Crit, 1, format, v...)
    }

    return l.output(1, fmt.Sprintf(format, v...))
}

type syslog_func func(m string) error

func (l *Logger) out_syslog(log_func syslog_func, call_depth int,
    m string) error {

    output := l.get_output(call_depth + 1, m, flag_is_syslog)

    l.get_lock()
    defer l.release_lock()

    return log_func(output)
}

func (l *Logger) out_syslogf(log_func syslog_func, call_depth int,
    format string, v ...interface{}) error {

    m := fmt.Sprintf(format, v...)
    return l.out_syslog(log_func, call_depth + 1, m)
}

// Logs a message with severity LOG_DEBUG.
func (l *Logger) Debug(m string) error {
    if l.severity_thresh < LOG_DEBUG {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslog(l.syslog_writer.Debug, 1, m)
    }
    return l.output(1, m)
}

// Logs a message with severity LOG_DEBUG. Arguments are handled in the manner
// of fmt.Printf.
func (l *Logger) Debugf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_DEBUG {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslogf(l.syslog_writer.Debug, 1, format, v...)
    }
    return l.output(1, fmt.Sprintf(format, v...))
}

// Logs a message with severity LOG_EMERG.
func (l *Logger) Emerg(m string) error {
    if l.severity_thresh < LOG_EMERG {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslog(l.syslog_writer.Emerg, 1, m)
    }
    return l.output(1, m)
}

// Logs a message with severity LOG_EMERG. Arguments are handled in the manner
// of fmt.Printf.
func (l *Logger) Emergf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_EMERG {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslogf(l.syslog_writer.Emerg, 1, format, v...)
    }

    return l.output(1, fmt.Sprintf(format, v...))
}

// Logs a message with severity LOG_ERR.
func (l *Logger) Err(m string) error {
    if l.severity_thresh < LOG_ERR {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslog(l.syslog_writer.Err, 1, m)
    }

    return l.output(1, m)
}

// Logs a message with severity LOG_ERR. Arguments are handled in the manner of
// fmt.Printf.
func (l *Logger) Errf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_ERR {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslogf(l.syslog_writer.Err, 1, format, v...)
    }

    return l.output(1, fmt.Sprintf(format, v...))
}

// Logs a message with severity LOG_INFO.
func (l *Logger) Info(m string) error {
    if l.severity_thresh < LOG_INFO {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslog(l.syslog_writer.Info, 1, m)
    }

    return l.output(1, m)
}

// Logs a message with severity LOG_INFO. Arguments are handled in the manner of
// fmt.Printf.
func (l *Logger) Infof(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_INFO {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslogf(l.syslog_writer.Info, 1, format, v...)
    }

    return l.output(1, fmt.Sprintf(format, v...))
}

// Logs a message with severity LOG_NOTICE.
func (l *Logger) Notice(m string) error {
    if l.severity_thresh < LOG_NOTICE {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslog(l.syslog_writer.Notice, 1, m)
    }

    return l.output(1, m)
}

// Logs a message with severity LOG_NOTICE. Arguments are handled in the manner
// of fmt.Printf.
func (l *Logger) Noticef(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_NOTICE {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslogf(l.syslog_writer.Notice, 1, format, v...)
    }

    return l.output(1, fmt.Sprintf(format, v...))
}

// Logs a message with severity LOG_WARNING.
func (l *Logger) Warning(m string) error {
    if l.severity_thresh < LOG_WARNING {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslog(l.syslog_writer.Warning, 1, m)
    }

    return l.output(1, m)
}

// Logs a message with severity LOG_WARNING. Arguments are handled in the manner
// of fmt.Printf.
func (l *Logger) Warningf(format string, v ...interface{}) error {
    if l.severity_thresh < LOG_WARNING {
        return nil
    }
    if l.syslog_writer != nil {
        return l.out_syslogf(l.syslog_writer.Warning, 1, format, v...)
    }

    return l.output(1, fmt.Sprintf(format, v...))
}

func (l *Logger) Write(b []byte) (int, error) {
    if l.syslog_writer != nil {
        str := l.get_output(1, string(b), flag_is_syslog)
        _, err := l.syslog_writer.Write([]byte(str))
        return len(b), err
    }

    err := l.output(1, string(b))

    return len(b), err
}

func (l *Logger) Close() error {
    if closer, ok := l.writer.(io.Closer); ok {
        return closer.Close()
    }

    return fmt.Errorf("don't know how to Close() a %T", l.writer)
}

// Equivalent to Print() followed by a call to os.Exit(1).
func (l *Logger) Fatal(v ...interface{}) {
    l.Print(v...)
    os.Exit(1)
}

// Equivalent to Printf() followed by a call to os.Exit(1).
func (l *Logger) Fatalf(format string, v ...interface{}) {
    l.Printf(format, v...)
    os.Exit(1)
}

// Equivalent to Println() followed by a call to os.Exit(1).
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

func (l *Logger) output(call_depth int, s string) error {
    return l.output_with_flags(call_depth + 1, s, 0)
}

func (l *Logger) output_with_flags(call_depth int, s string,
    flags uint32) error {

    out_str := l.get_output(call_depth + 1, s, flags)

    l.get_lock()
    defer l.release_lock()

    _, err := fmt.Fprint(l.writer, out_str)
    return err
}

// Equivalent to Print() followed by a call to panic().
func (l *Logger) Panic(v ...interface{}) {
    if l.syslog_writer != nil {
        str := l.get_output(1, fmt.Sprint(v...), flag_is_syslog)
        l.syslog_writer.Write([]byte(str))
    } else {
        l.output(1, fmt.Sprint(v...))
    }

    panic(fmt.Sprint(v...))
}

// Equivalent to Printf() followed by a call to panic().
func (l *Logger) Panicf(format string, v ...interface{}) {
    if l.syslog_writer != nil {
        str := l.get_output(1,
            fmt.Sprintf(format, v...), flag_is_syslog)
        l.syslog_writer.Write([]byte(str))
    } else {
        l.output(1, fmt.Sprintf(format, v...))
    }

    panic(fmt.Sprintf(format, v...))
}

// Equivalent to Println() followed by a call to panic().
func (l *Logger) Panicln(v ...interface{}) {
    if l.syslog_writer != nil {
        str := l.get_output(1, fmt.Sprintln(v...), flag_is_syslog)
        l.syslog_writer.Write([]byte(str))
    } else {
        l.output(1, fmt.Sprint(v...))
    }

    panic(fmt.Sprintln(v...))
}

// Prints to the logger. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) {
    m := fmt.Sprint(v...)
    l.Write([]byte(m))
}

// Prints to the logger. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
    m := fmt.Sprintf(format, v...)
    l.Write([]byte(m))
}

// Prints to the logger. Arguments are handled in the manner of fmt.Println.
func (l *Logger) Println(v ...interface{}) {
    m := fmt.Sprintln(v...)
    l.Write([]byte(m))
}
