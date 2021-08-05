package main

import "fmt"

func makeScheduleClassString(roles []string, weekDay string) string {
	classes := getScheduleForADay(roles, weekDay)
	result := "Classes are as follows\n"
	for i, class := range classes {
		result += fmt.Sprintf("Class %d\n%s\n\n", (i + 1), class.make_generic_class_string())
	}
	return result
}
