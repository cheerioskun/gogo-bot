package main

import (
	"fmt"

	"github.com/profclems/go-dotenv"
)

func prepareCommandsAndUsages() {
	prefixChar = dotenv.GetString("PREFIX")
	scheduleCommand = prefixChar + "schedule"
	schCommand = prefixChar + "sch"
	scheduleCommandUsage = fmt.Sprintf("Command Name: %s\nUsage: %s <Day> or %s <Day>\nHere day can be any of 5 weekdays and is optional with default value of present day.\n", scheduleCommand, scheduleCommand, schCommand)
}
