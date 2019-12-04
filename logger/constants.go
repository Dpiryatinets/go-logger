package logger

const fatal = "fatal"
const errorLevel = "error"
const warn = "warn"
const info = "info"
const verbose = "verbose"
const debug = "debug"

var levels = map[string]int{
	fatal:      5,
	errorLevel: 4,
	warn:       3,
	info:       2,
	verbose:    1,
	debug:      0,
}
