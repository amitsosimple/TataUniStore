/**** Amit Chatter (amitsosimple@gmail.com) ****/

package log

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
)

// GeneralLogger exported
var GeneralLogger *log.Logger

// ErrorLogger exported
var ErrorLogger *log.Logger

func init() {
	usr, _ := user.Current()
	serviceDataFolder := filepath.Join(usr.HomeDir, "ProductionCatalog")
	if _, err := os.Stat(serviceDataFolder); os.IsNotExist(err) {
		_ = os.Mkdir(serviceDataFolder, os.ModePerm)
	}

	logFile := filepath.Join(serviceDataFolder, "general-log.log")

	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		os.Create(logFile)
	}

	generalLog, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	GeneralLogger = log.New(generalLog, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(generalLog, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}

// Regex to extract just the function name (and not the module path)
var RE_stripFnPreamble = regexp.MustCompile(`^.*\.(.*)$`)

// Trace Functions
func Enter() string {
	fnName := "<unknown>"
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		fnName = RE_stripFnPreamble.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
	}

	GeneralLogger.Printf("Entering to func %s\n", fnName)
	return fnName
}

func Exit(s string) {
	GeneralLogger.Printf("Exiting from func %s\n", s)
}

