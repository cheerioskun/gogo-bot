package main

import (
	"strings"
)

func getScheduleForADay(roles []string, weekDay string) []*Class {
	var classesForRequiredDay []*Class
	for _, class := range classes {
		if class.Weekday == strings.Title(weekDay) && anyCommon(class.Sections, roles) {
			classesForRequiredDay = append(classesForRequiredDay, class)
		}
	}
	return classesForRequiredDay
}
