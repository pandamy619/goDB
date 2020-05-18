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


##### New creates a new Logger
```go
NewLog(path string, formatLogFile string, prefix string) (*log.Logger, *os.File)
```

##### Close log file
```go
CloseLog(file *os.File)
```