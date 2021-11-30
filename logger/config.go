package logger

import (
	"log"
	"os"
	"time"
)

type Config struct {
	LogLevel LogLevel
	Color    string
	SlowThreshold 	time.Duration
}

var logLevel2Color = map[LogLevel]string {
	Info:  Green,
	Warn:  Yellow,
	Error: Red,
}

var logLevel2Str = map[LogLevel]string {
	Info:  logLevel2Color[Info] + "[info]" + logLevel2Color[Info] + "%s\n" + Reset,
	Warn:  logLevel2Color[Warn] + "[warm]" + logLevel2Color[Warn] + "%s\n" + Reset,
	Error: logLevel2Color[Error] + "[error]" + logLevel2Color[Error] + "%s\n" + Reset,
}

var logLevel2Logger = map[LogLevel]*log.Logger {
	Info:  log.New(os.Stdout, logLevel2Color[Info], log.LstdFlags|log.Lshortfile),
	Warn:  log.New(os.Stdout, logLevel2Color[Warn], log.LstdFlags|log.Lshortfile),
	Error: log.New(os.Stdout, logLevel2Color[Error], log.LstdFlags|log.Lshortfile),
}