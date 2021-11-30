package logger

import (
	"io/ioutil"
	"os"
)

func init() {
	for key, Logger := range logLevel2Logger {
		if key <= Silent {
			Logger.SetOutput(ioutil.Discard)
		} else {
			Logger.SetOutput(os.Stdout)
		}
	}
}
