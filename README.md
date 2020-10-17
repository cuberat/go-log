

# log
`import "github.com/cuberat/go-log"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package log implements a logger that logs to a file or any io.Writer. If the
io.Writer looks like syslog, then certain parts of the log line will not be
generated, as syslog is expected to cover those.

This logger looks like syslog, in that it allows for specifying severities
when logging. Unlike syslog, however, instead of setting a default severity
for logging, a severity threshold is specified at logger creation time that
limits logging to that severity and anything more important. This provides a
convenient way to control the verbosity of logging.

The intention is to enable this to be a drop-in replacement, interface-wise,
for the standard log/syslog or log. In addition, it provides most of the
methods offered by the standard log interface.

You may create a new logger using New() or NewFromFile(), or, if you want to
use the default logger, just call methods directly on the package, which will
write to os.Stderr.

A Logger can be used simultaneously from multiple goroutines; it guarantees
to serialize access to the Writer.

The interface for this logging package is in an alpha state. That is, it may
have some non backward compatible changes in upcoming minor releases.

Installation:


	go get github.com/cuberat/go-log




## <a name="pkg-index">Index</a>
* [func Alert(m string) error](#Alert)
* [func Alertf(format string, v ...interface{}) error](#Alertf)
* [func Crit(m string) error](#Crit)
* [func Critf(format string, v ...interface{}) error](#Critf)
* [func Debug(m string) error](#Debug)
* [func Debugf(format string, v ...interface{}) error](#Debugf)
* [func Emerg(m string) error](#Emerg)
* [func Emergf(format string, v ...interface{}) error](#Emergf)
* [func Err(m string) error](#Err)
* [func Errf(format string, v ...interface{}) error](#Errf)
* [func Errorf(format string, v ...interface{}) error](#Errorf)
* [func ErrorfDepth(call_depth int, format string, v ...interface{}) error](#ErrorfDepth)
* [func Fatal(v ...interface{})](#Fatal)
* [func Fatalf(format string, v ...interface{})](#Fatalf)
* [func Fatalln(v ...interface{})](#Fatalln)
* [func Info(m string) error](#Info)
* [func Infof(format string, v ...interface{}) error](#Infof)
* [func Notice(m string) error](#Notice)
* [func Noticef(format string, v ...interface{}) error](#Noticef)
* [func Panic(v ...interface{})](#Panic)
* [func Panicf(format string, v ...interface{})](#Panicf)
* [func Panicln(v ...interface{})](#Panicln)
* [func Print(v ...interface{}) error](#Print)
* [func Printf(format string, v ...interface{}) error](#Printf)
* [func Println(v ...interface{}) error](#Println)
* [func SetOutput(w io.Writer)](#SetOutput)
* [func SetPrefix(prefix string)](#SetPrefix)
* [func SetSeverityThreshold(sev_thresh Severity)](#SetSeverityThreshold)
* [func SetTimestampFunc(f TimestampFunc)](#SetTimestampFunc)
* [func Warning(m string) error](#Warning)
* [func Warningf(format string, v ...interface{}) error](#Warningf)
* [type Logger](#Logger)
  * [func New(w io.Writer, sev_thresh Severity, prefix string) *Logger](#New)
  * [func NewFromFile(file_path string, sev_thresh Severity, prefix string) (*Logger, error)](#NewFromFile)
  * [func (l *Logger) Alert(m string) error](#Logger.Alert)
  * [func (l *Logger) Alertf(format string, v ...interface{}) error](#Logger.Alertf)
  * [func (l *Logger) Crit(m string) error](#Logger.Crit)
  * [func (l *Logger) Critf(format string, v ...interface{}) error](#Logger.Critf)
  * [func (l *Logger) Debug(m string) error](#Logger.Debug)
  * [func (l *Logger) Debugf(format string, v ...interface{}) error](#Logger.Debugf)
  * [func (l *Logger) Emerg(m string) error](#Logger.Emerg)
  * [func (l *Logger) Emergf(format string, v ...interface{}) error](#Logger.Emergf)
  * [func (l *Logger) Err(m string) error](#Logger.Err)
  * [func (l *Logger) Errf(format string, v ...interface{}) error](#Logger.Errf)
  * [func (l *Logger) Errorf(format string, v ...interface{}) error](#Logger.Errorf)
  * [func (l *Logger) ErrorfDepth(call_depth int, format string, v ...interface{}) error](#Logger.ErrorfDepth)
  * [func (l *Logger) Fatal(v ...interface{})](#Logger.Fatal)
  * [func (l *Logger) Fatalf(format string, v ...interface{})](#Logger.Fatalf)
  * [func (l *Logger) Fatalln(v ...interface{})](#Logger.Fatalln)
  * [func (l *Logger) Info(m string) error](#Logger.Info)
  * [func (l *Logger) Infof(format string, v ...interface{}) error](#Logger.Infof)
  * [func (l *Logger) Notice(m string) error](#Logger.Notice)
  * [func (l *Logger) Noticef(format string, v ...interface{}) error](#Logger.Noticef)
  * [func (l *Logger) Panic(v ...interface{})](#Logger.Panic)
  * [func (l *Logger) Panicf(format string, v ...interface{})](#Logger.Panicf)
  * [func (l *Logger) Panicln(v ...interface{})](#Logger.Panicln)
  * [func (l *Logger) Print(v ...interface{}) error](#Logger.Print)
  * [func (l *Logger) Printf(format string, v ...interface{}) error](#Logger.Printf)
  * [func (l *Logger) Println(v ...interface{}) error](#Logger.Println)
  * [func (l *Logger) SetOutput(w io.Writer)](#Logger.SetOutput)
  * [func (l *Logger) SetPrefix(prefix string)](#Logger.SetPrefix)
  * [func (l *Logger) SetSeverityThreshold(sev_thresh Severity)](#Logger.SetSeverityThreshold)
  * [func (l *Logger) SetTimestampFunc(f TimestampFunc)](#Logger.SetTimestampFunc)
  * [func (l *Logger) Warning(m string) error](#Logger.Warning)
  * [func (l *Logger) Warningf(format string, v ...interface{}) error](#Logger.Warningf)
  * [func (l *Logger) Write(b []byte) (int, error)](#Logger.Write)
* [type Severity](#Severity)
  * [func SeverityFromString(sev_string string) (Severity, error)](#SeverityFromString)
* [type SyslogLike](#SyslogLike)
* [type TimestampFunc](#TimestampFunc)


#### <a name="pkg-files">Package files</a>
[log.go](/src/github.com/cuberat/go-log/log.go) [logger.go](/src/github.com/cuberat/go-log/logger.go) 





## <a name="Alert">func</a> [Alert](/src/target/log.go?s=6571:6597#L187)
``` go
func Alert(m string) error
```
Logs a message with severity LOG_ALERT.



## <a name="Alertf">func</a> [Alertf](/src/target/log.go?s=6751:6801#L193)
``` go
func Alertf(format string, v ...interface{}) error
```
Logs a message with severity LOG_ALERT. Arguments are handled in the manner
of fmt.Printf.



## <a name="Crit">func</a> [Crit](/src/target/log.go?s=6912:6937#L198)
``` go
func Crit(m string) error
```
Logs a message with severity LOG_CRIT.



## <a name="Critf">func</a> [Critf](/src/target/log.go?s=7089:7138#L204)
``` go
func Critf(format string, v ...interface{}) error
```
Logs a message with severity LOG_CRIT. Arguments are handled in the manner of
fmt.Printf.



## <a name="Debug">func</a> [Debug](/src/target/log.go?s=7249:7275#L209)
``` go
func Debug(m string) error
```
Logs a message with severity LOG_DEBUG.



## <a name="Debugf">func</a> [Debugf](/src/target/log.go?s=7429:7479#L215)
``` go
func Debugf(format string, v ...interface{}) error
```
Logs a message with severity LOG_DEBUG. Arguments are handled in the manner
of fmt.Printf.



## <a name="Emerg">func</a> [Emerg](/src/target/log.go?s=7591:7617#L220)
``` go
func Emerg(m string) error
```
Logs a message with severity LOG_EMERG.



## <a name="Emergf">func</a> [Emergf](/src/target/log.go?s=7771:7821#L226)
``` go
func Emergf(format string, v ...interface{}) error
```
Logs a message with severity LOG_EMERG. Arguments are handled in the manner
of fmt.Printf.



## <a name="Err">func</a> [Err](/src/target/log.go?s=7931:7955#L231)
``` go
func Err(m string) error
```
Logs a message with severity LOG_ERR.



## <a name="Errf">func</a> [Errf](/src/target/log.go?s=8105:8153#L237)
``` go
func Errf(format string, v ...interface{}) error
```
Logs a message with severity LOG_ERR. Arguments are handled in the manner of
fmt.Printf.



## <a name="Errorf">func</a> [Errorf](/src/target/log.go?s=10829:10879#L327)
``` go
func Errorf(format string, v ...interface{}) error
```
Returns an error like `fmt.Errorf`, but prepended with the source file name
and line number.



## <a name="ErrorfDepth">func</a> [ErrorfDepth](/src/target/log.go?s=11277:11363#L336)
``` go
func ErrorfDepth(
    call_depth int,
    format string,
    v ...interface{},
) error
```
Returns an error like `Errorf`, but allows you to specify a call depth. For
instance, passing a value of 1 as the call_depth will cause the source file
and line number to correspond to where the enclosing function is called,
instead of where `ErrorDepth` is called. `ErrorDepth(0, ...)` is equivalent
to calling `Errorf`.



## <a name="Fatal">func</a> [Fatal](/src/target/log.go?s=9317:9345#L275)
``` go
func Fatal(v ...interface{})
```
Equivalent to Print() followed by a call to os.Exit(1).



## <a name="Fatalf">func</a> [Fatalf](/src/target/log.go?s=9462:9506#L281)
``` go
func Fatalf(format string, v ...interface{})
```
Equivalent to Printf() followed by a call to os.Exit(1).



## <a name="Fatalln">func</a> [Fatalln](/src/target/log.go?s=9632:9662#L287)
``` go
func Fatalln(v ...interface{})
```
Equivalent to Println() followed by a call to os.Exit(1).



## <a name="Info">func</a> [Info](/src/target/log.go?s=8262:8287#L242)
``` go
func Info(m string) error
```
Logs a message with severity LOG_INFO.



## <a name="Infof">func</a> [Infof](/src/target/log.go?s=8439:8488#L248)
``` go
func Infof(format string, v ...interface{}) error
```
Logs a message with severity LOG_INFO. Arguments are handled in the manner of
fmt.Printf.



## <a name="Notice">func</a> [Notice](/src/target/log.go?s=8600:8627#L253)
``` go
func Notice(m string) error
```
Logs a message with severity LOG_NOTICE.



## <a name="Noticef">func</a> [Noticef](/src/target/log.go?s=8783:8834#L259)
``` go
func Noticef(format string, v ...interface{}) error
```
Logs a message with severity LOG_NOTICE. Arguments are handled in the manner
of fmt.Printf.



## <a name="Panic">func</a> [Panic](/src/target/log.go?s=9777:9805#L293)
``` go
func Panic(v ...interface{})
```
Equivalent to Print() followed by a call to panic().



## <a name="Panicf">func</a> [Panicf](/src/target/log.go?s=9932:9976#L299)
``` go
func Panicf(format string, v ...interface{})
```
Equivalent to Printf() followed by a call to panic().



## <a name="Panicln">func</a> [Panicln](/src/target/log.go?s=10121:10151#L305)
``` go
func Panicln(v ...interface{})
```
Equivalent to Println() followed by a call to panic().



## <a name="Print">func</a> [Print](/src/target/log.go?s=10300:10334#L311)
``` go
func Print(v ...interface{}) error
```
Prints to the logger. Arguments are handled in the manner of fmt.Print.



## <a name="Printf">func</a> [Printf](/src/target/log.go?s=10459:10509#L316)
``` go
func Printf(format string, v ...interface{}) error
```
Prints to the logger. Arguments are handled in the manner of fmt.Printf.



## <a name="Println">func</a> [Println](/src/target/log.go?s=10643:10679#L321)
``` go
func Println(v ...interface{}) error
```
Prints to the logger. Arguments are handled in the manner of fmt.Println.



## <a name="SetOutput">func</a> [SetOutput](/src/target/log.go?s=4833:4860#L127)
``` go
func SetOutput(w io.Writer)
```
Sets the writer where logging output should go for the default logger.



## <a name="SetPrefix">func</a> [SetPrefix](/src/target/log.go?s=5251:5280#L139)
``` go
func SetPrefix(prefix string)
```
Sets the prefix to add to the beginning of each log line (after the
timestamp) for the default logger.



## <a name="SetSeverityThreshold">func</a> [SetSeverityThreshold](/src/target/log.go?s=5038:5084#L133)
``` go
func SetSeverityThreshold(sev_thresh Severity)
```
Sets the severity threshold for the default logger. Anything less important
(further down the list of severities) will not be logged.



## <a name="SetTimestampFunc">func</a> [SetTimestampFunc](/src/target/log.go?s=5456:5494#L145)
``` go
func SetTimestampFunc(f TimestampFunc)
```
Sets the timestamp generator function for the default logger. This will be
called to generate the timestamp for each log line.



## <a name="Warning">func</a> [Warning](/src/target/log.go?s=8949:8977#L264)
``` go
func Warning(m string) error
```
Logs a message with severity LOG_WARNING.



## <a name="Warningf">func</a> [Warningf](/src/target/log.go?s=9135:9187#L270)
``` go
func Warningf(format string, v ...interface{}) error
```
Logs a message with severity LOG_WARNING. Arguments are handled in the manner
of fmt.Printf.




## <a name="Logger">type</a> [Logger](/src/target/logger.go?s=1875:2044#L33)
``` go
type Logger struct {
    // contains filtered or unexported fields
}
```
A Logger represents an active logging object that generates lines of output
to an io.Writer. Each logging operation makes a single call to the Writer's
Write method or to the Writer's severity-related method, if it implements the
SyslogLike interface. A Logger can be used simultaneously from multiple
goroutines; it guarantees to serialize access to the Writer.







### <a name="New">func</a> [New](/src/target/log.go?s=5678:5745#L153)
``` go
func New(w io.Writer, sev_thresh Severity, prefix string) *Logger
```
Creates a logger from an io.Writer, with the given severity threshold and
prefix string.


### <a name="NewFromFile">func</a> [NewFromFile](/src/target/log.go?s=6172:6263#L171)
``` go
func NewFromFile(file_path string, sev_thresh Severity, prefix string) (*Logger,
    error)
```
Creates a logger that will write to the specified file path, with the given
severity threshold and prefix string





### <a name="Logger.Alert">func</a> (\*Logger) [Alert](/src/target/logger.go?s=3097:3135#L82)
``` go
func (l *Logger) Alert(m string) error
```
Logs a message with severity LOG_ALERT.




### <a name="Logger.Alertf">func</a> (\*Logger) [Alertf](/src/target/logger.go?s=3424:3486#L95)
``` go
func (l *Logger) Alertf(format string, v ...interface{}) error
```
Logs a message with severity LOG_ALERT. Arguments are handled in the manner
of fmt.Printf.




### <a name="Logger.Crit">func</a> (\*Logger) [Crit](/src/target/logger.go?s=3756:3793#L107)
``` go
func (l *Logger) Crit(m string) error
```
Logs a message with severity LOG_CRIT.




### <a name="Logger.Critf">func</a> (\*Logger) [Critf](/src/target/logger.go?s=4078:4139#L119)
``` go
func (l *Logger) Critf(format string, v ...interface{}) error
```
Logs a message with severity LOG_CRIT. Arguments are handled in the manner of
fmt.Printf.




### <a name="Logger.Debug">func</a> (\*Logger) [Debug](/src/target/logger.go?s=4408:4446#L131)
``` go
func (l *Logger) Debug(m string) error
```
Logs a message with severity LOG_DEBUG.




### <a name="Logger.Debugf">func</a> (\*Logger) [Debugf](/src/target/logger.go?s=4734:4796#L143)
``` go
func (l *Logger) Debugf(format string, v ...interface{}) error
```
Logs a message with severity LOG_DEBUG. Arguments are handled in the manner
of fmt.Printf.




### <a name="Logger.Emerg">func</a> (\*Logger) [Emerg](/src/target/logger.go?s=5066:5104#L154)
``` go
func (l *Logger) Emerg(m string) error
```
Logs a message with severity LOG_EMERG.




### <a name="Logger.Emergf">func</a> (\*Logger) [Emergf](/src/target/logger.go?s=5392:5454#L166)
``` go
func (l *Logger) Emergf(format string, v ...interface{}) error
```
Logs a message with severity LOG_EMERG. Arguments are handled in the manner
of fmt.Printf.




### <a name="Logger.Err">func</a> (\*Logger) [Err](/src/target/logger.go?s=5723:5759#L178)
``` go
func (l *Logger) Err(m string) error
```
Logs a message with severity LOG_ERR.




### <a name="Logger.Errf">func</a> (\*Logger) [Errf](/src/target/logger.go?s=6042:6102#L191)
``` go
func (l *Logger) Errf(format string, v ...interface{}) error
```
Logs a message with severity LOG_ERR. Arguments are handled in the manner
of fmt.Printf.




### <a name="Logger.Errorf">func</a> (\*Logger) [Errorf](/src/target/logger.go?s=10415:10477#L351)
``` go
func (l *Logger) Errorf(format string, v ...interface{}) error
```
Returns an error like `fmt.Errorf`, but prepended with the source file name
and line number.




### <a name="Logger.ErrorfDepth">func</a> (\*Logger) [ErrorfDepth](/src/target/logger.go?s=10862:10960#L360)
``` go
func (l *Logger) ErrorfDepth(
    call_depth int,
    format string,
    v ...interface{},
) error
```
Returns an error like `Errorf`, but allows you to specify a call depth. For
instance, passing a value of 1 as the call_depth will cause the source file
and line number to correspond to where the enclosing function is called,
instead of where `ErrorDepth` is called. `ErrorDepth(0, ...)` is equivalent
to calling `Errorf`.




### <a name="Logger.Fatal">func</a> (\*Logger) [Fatal](/src/target/logger.go?s=8899:8939#L299)
``` go
func (l *Logger) Fatal(v ...interface{})
```
Equivalent to Print() followed by a call to os.Exit(1).




### <a name="Logger.Fatalf">func</a> (\*Logger) [Fatalf](/src/target/logger.go?s=9056:9112#L305)
``` go
func (l *Logger) Fatalf(format string, v ...interface{})
```
Equivalent to Printf() followed by a call to os.Exit(1).




### <a name="Logger.Fatalln">func</a> (\*Logger) [Fatalln](/src/target/logger.go?s=9225:9267#L311)
``` go
func (l *Logger) Fatalln(v ...interface{})
```
Equivalent to Println() followed by a call to os.Exit(1).




### <a name="Logger.Info">func</a> (\*Logger) [Info](/src/target/logger.go?s=6368:6405#L203)
``` go
func (l *Logger) Info(m string) error
```
Logs a message with severity LOG_INFO.




### <a name="Logger.Infof">func</a> (\*Logger) [Infof](/src/target/logger.go?s=6691:6752#L216)
``` go
func (l *Logger) Infof(format string, v ...interface{}) error
```
Logs a message with severity LOG_INFO. Arguments are handled in the manner
of fmt.Printf.




### <a name="Logger.Notice">func</a> (\*Logger) [Notice](/src/target/logger.go?s=7022:7061#L228)
``` go
func (l *Logger) Notice(m string) error
```
Logs a message with severity LOG_NOTICE.




### <a name="Logger.Noticef">func</a> (\*Logger) [Noticef](/src/target/logger.go?s=7353:7416#L241)
``` go
func (l *Logger) Noticef(format string, v ...interface{}) error
```
Logs a message with severity LOG_NOTICE. Arguments are handled in the
manner of fmt.Printf.




### <a name="Logger.Panic">func</a> (\*Logger) [Panic](/src/target/logger.go?s=9369:9409#L317)
``` go
func (l *Logger) Panic(v ...interface{})
```
Equivalent to Print() followed by a call to panic().




### <a name="Logger.Panicf">func</a> (\*Logger) [Panicf](/src/target/logger.go?s=9523:9579#L323)
``` go
func (l *Logger) Panicf(format string, v ...interface{})
```
Equivalent to Printf() followed by a call to panic().




### <a name="Logger.Panicln">func</a> (\*Logger) [Panicln](/src/target/logger.go?s=9711:9753#L329)
``` go
func (l *Logger) Panicln(v ...interface{})
```
Equivalent to Println() followed by a call to panic().




### <a name="Logger.Print">func</a> (\*Logger) [Print](/src/target/logger.go?s=9889:9935#L335)
``` go
func (l *Logger) Print(v ...interface{}) error
```
Prints to the logger. Arguments are handled in the manner of fmt.Print.




### <a name="Logger.Printf">func</a> (\*Logger) [Printf](/src/target/logger.go?s=10047:10109#L340)
``` go
func (l *Logger) Printf(format string, v ...interface{}) error
```
Prints to the logger. Arguments are handled in the manner of fmt.Printf.




### <a name="Logger.Println">func</a> (\*Logger) [Println](/src/target/logger.go?s=10230:10278#L345)
``` go
func (l *Logger) Println(v ...interface{}) error
```
Prints to the logger. Arguments are handled in the manner of fmt.Println.




### <a name="Logger.SetOutput">func</a> (\*Logger) [SetOutput](/src/target/logger.go?s=2329:2368#L56)
``` go
func (l *Logger) SetOutput(w io.Writer)
```
Sets the writer where logging output should go.




### <a name="Logger.SetPrefix">func</a> (\*Logger) [SetPrefix](/src/target/logger.go?s=2696:2737#L68)
``` go
func (l *Logger) SetPrefix(prefix string)
```
Sets the prefix to add to the beginning of each log line (after the
timestamp).




### <a name="Logger.SetSeverityThreshold">func</a> (\*Logger) [SetSeverityThreshold](/src/target/logger.go?s=2511:2569#L62)
``` go
func (l *Logger) SetSeverityThreshold(sev_thresh Severity)
```
Sets the severity threshold. Anything less important (further down the list
of severities) will not be logged.




### <a name="Logger.SetTimestampFunc">func</a> (\*Logger) [SetTimestampFunc](/src/target/logger.go?s=2980:3030#L77)
``` go
func (l *Logger) SetTimestampFunc(f TimestampFunc)
```
Sets the timestamp generator function. This will be called to generate the
timestamp for each log line.




### <a name="Logger.Warning">func</a> (\*Logger) [Warning](/src/target/logger.go?s=7691:7731#L253)
``` go
func (l *Logger) Warning(m string) error
```
Logs a message with severity LOG_WARNING.




### <a name="Logger.Warningf">func</a> (\*Logger) [Warningf](/src/target/logger.go?s=8026:8090#L266)
``` go
func (l *Logger) Warningf(format string, v ...interface{}) error
```
Logs a message with severity LOG_WARNING. Arguments are handled in the
manner of fmt.Printf.




### <a name="Logger.Write">func</a> (\*Logger) [Write](/src/target/logger.go?s=8347:8392#L278)
``` go
func (l *Logger) Write(b []byte) (int, error)
```
Writes a log message.




## <a name="Severity">type</a> [Severity](/src/target/log.go?s=2831:2848#L55)
``` go
type Severity int
```
The Severity type.


``` go
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
```
Severities to be passed to New() or SetSeverityThreshold().







### <a name="SeverityFromString">func</a> [SeverityFromString](/src/target/log.go?s=4376:4436#L113)
``` go
func SeverityFromString(sev_string string) (Severity, error)
```
Converts a severity name to a Severity that can be passed to New() and
SetSeverityThreshold(). This is useful for allowing specification of the
severity on the command line an creating a logger that uses that threshold.





## <a name="SyslogLike">type</a> [SyslogLike](/src/target/log.go?s=3285:3554#L76)
``` go
type SyslogLike interface {
    Alert(m string) error
    Crit(m string) error
    Debug(m string) error
    Emerg(m string) error
    Err(m string) error
    Info(m string) error
    Notice(m string) error
    Warning(m string) error
    Write(b []byte) (int, error)
}
```
If the io.Writer passed to New() or SetOutput() implements the SyslogLike
interface, it is treated as if it actually is a standard log/syslog logger.










## <a name="TimestampFunc">type</a> [TimestampFunc](/src/target/log.go?s=3594:3628#L89)
``` go
type TimestampFunc func() string
```
Timestamp generator function type.














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
