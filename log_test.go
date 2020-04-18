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

// A syslog-like logger that logs to a file. The intention is to enable it to be
// a drop-in replacement for the standard log/syslog or log.
//
// Installation
//   go get github.com/cuberat/go-log
package log_test

import (
    "bytes"
    log "github.com/cuberat/go-log"
    "strings"
    "testing"
)

func TestAlert(t *testing.T) {
    buffer := new(bytes.Buffer)
    logger := log.New(buffer, log.LOG_ALERT, "")
    logger.Alert("foo")
    logger.Alertf("stuff=%s", "bar")
    logger.Crit("boo")

    log_str := buffer.String()
    // t.Logf("log: %s", log_str)

    if !strings.Contains(log_str, "stuff=bar") {
        t.Error("Output should contain \"stuff=bar\"")
    }

    if strings.Contains(log_str, "boo") {
        t.Error("Output should NOT contain \"boo\"")
    }
}
