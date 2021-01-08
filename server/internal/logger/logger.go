package logger

import (
	"log"
	"os"
	"sync"
)

var once sync.Once
var instance *log.Logger

//Get the db instance
func Get() *log.Logger {

	once.Do(func() { // <-- atomic, does not allow repeating
		logger := log.New(os.Stdout, "Vincivia API ", log.LstdFlags|log.Lshortfile)
		instance = logger
		instance.Println("Logger singleton instantiated")
	})

	return instance
}
