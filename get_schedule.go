package main

import "strings"

func getScheduleForADay(roles []string, weekDay string) []*Class {
	var classesForReqiredDay []*Class
	for _, class := range classes {
		if class.Weekday == strings.ToTitle(weekDay) && contains(class.Sections, roles) {
			classesForReqiredDay = append(classesForReqiredDay, class)
		}
	}
	return classesForReqiredDay
}
