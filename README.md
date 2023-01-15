# mlog

Minimal log which only adds `Debug` family functions to the standard `log`.

## Usage

```go
package main

import (
	log "github.com/zachcheung/mlog"
)

func main() {
	log.Debug("It's not showed by default")
	log.EnableDebug()
	log.Debug("hello world")
	log.Debugf("hello %s", "world")
	log.Debugln("hello", "world")
}
```

Output:

```
2023/01/14 19:42:02 [DEBUG] hello world
2023/01/14 19:42:02 [DEBUG] hello world
2023/01/14 19:42:02 [DEBUG] hello world
```

## Installation

```shell
go get github.com/zachcheung/mlog
```

## License

MIT

## Author

Zach Cheung
