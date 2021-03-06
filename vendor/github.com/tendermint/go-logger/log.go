package logger

import (
	"os"

	. "github.com/tendermint/go-common"
	"github.com/tendermint/log15"
)

//var rootHandler log15.Handler
var mainHandler log15.Handler
var bypassHandler log15.Handler

func init() {
	Reset()
}

// You might want to call this after resetting tendermint/go-config.
func Reset() {

	var logLevel string = "debug"
	if config != nil {
		logLevel = config.GetString("log_level")
	}

	// main handler
	//handlers := []log15.Handler{}
	mainHandler = log15.LvlFilterHandler(
		getLevel(logLevel),
		log15.StreamHandler(os.Stdout, log15.TerminalFormat()),
	)
	//handlers = append(handlers, mainHandler)

	// bypass handler for not filtering on global logLevel.
	bypassHandler = log15.StreamHandler(os.Stdout, log15.TerminalFormat())
	//handlers = append(handlers, bypassHandler)

	// Set rootHandler.
	//rootHandler = log15.MultiHandler(handlers...)

	// By setting handlers on the root, we handle events from all loggers.
	log15.Root().SetHandler(mainHandler)
}

// See go-wire/log for an example of usage.
func MainHandler() log15.Handler {
	return mainHandler
}

func BypassHandler() log15.Handler {
	return bypassHandler
}

func New(ctx ...interface{}) log15.Logger {
	return NewMain(ctx...)
}

func NewMain(ctx ...interface{}) log15.Logger {
	return log15.Root().New(ctx...)
}

func NewBypass(ctx ...interface{}) log15.Logger {
	bypass := log15.New(ctx...)
	bypass.SetHandler(bypassHandler)
	return bypass
}

func getLevel(lvlString string) log15.Lvl {
	lvl, err := log15.LvlFromString(lvlString)
	if err != nil {
		Exit(Fmt("Invalid log level %v: %v", lvlString, err))
	}
	return lvl
}

//----------------------------------------
// Exported from log15

var LvlFilterHandler = log15.LvlFilterHandler
var LvlDebug = log15.LvlDebug
var LvlInfo = log15.LvlInfo
var LvlWarn = log15.LvlWarn
var LvlError = log15.LvlError
