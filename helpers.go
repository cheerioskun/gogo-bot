package main

import "fmt"

func contains(list []string, sublist []string) bool {
	for _, v1 := range sublist {
		for _, v2 := range list {
			if v1 == v2 {
				return true
			}
		}
	}

	return false
}

func make_generic_class_string(class *Class) string {
	startTime := class.Timeslots[0].StartTimeString
	endTime := class.Timeslots[len(class.Timeslots)-1].EndTimeString
	return fmt.Sprintf("Type: %s\nSubject: %s (%s)\nTime: %s-%s\nFaculty Name: %s\nMeet Link: <%s>\n",
		class.ClassType, class.SubjectName,
		class.SubjectShorthand, startTime,
		endTime, class.FacultyName, class.MeetLink)
}
