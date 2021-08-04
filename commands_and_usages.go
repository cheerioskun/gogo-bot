package main

import "fmt"

func prepareCommandsAndUsages() {
	scheduleCommand = prefixChar + "schedule"
	scheduleCommandUsage = fmt.Sprintf("Command Name: %s\nUsage: %sschedule <Day>\nHere day can be any of 5 weekdays and is optional with default value of present day\n.", scheduleCommand, prefixChar)
}
