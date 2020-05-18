### Usage log package

#### import package
```go
import "logging"
```

#### create new log

required aguments

* `path` - path to save a log file
* `formatLogFile` - format log files

optional arguments
* `prefix` - prefix for the standard logger (default to "prefix")

```go
logger, file := logging.NewLog(PathLog, FormatLogFile, "prefix")

<some code>

logging.CloseLog(file)
```