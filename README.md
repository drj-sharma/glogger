Glogger: A logging library for Golang, simple and easy to use compatible with the standard library logger.


## Installation & testing

Use below as package to install glogger in your Go project

RUN
```
go get github.com/drj-sharma/glogger@latest
```

and use below package as logger, see example for more details
```
github.com/drj-sharma/glogger
```
&nbsp;

example

```bash
import (
	log "github.com/drj-sharma/glogger"
)
var (
	logging = log.Logger{}
	logger  = logging.GetLogger()
)

func main() {
    logger.Info("This is info logger")
    logger.Debug("This is debug logger")
    logger.Fatal("This is fatal logger")
    // example log
    "[time] - LOGGER_LEVEL - LOG_MSG"
}
```

---
&nbsp;

Easiest way to test this library by using Docker

```bash
docker build -t app .

docker run app
```

OR

Use Makefile (https://askubuntu.com/questions/161104/how-do-i-install-make)

``` bash
make docker-build

make docker-run
    
