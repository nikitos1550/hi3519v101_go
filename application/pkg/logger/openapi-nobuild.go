
package logger

import (
	"io"
	"io/ioutil"
)

var (
        apiWriter  io.Writer
)


func initApiLog() {
	apiWriter = ioutil.Discard
}
