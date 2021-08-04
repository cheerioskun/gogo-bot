package main

import "fmt"

func makeRemindString(class *Class) string {
	allRelevantRoles := ""
	for _, role := range class.Sections {
		allRelevantRoles += fmt.Sprintf("<@&%s> ", roleNameToRoleId[role])
	}
	var auxInfo string
	if class.MeetLink == "Unknown" {
		auxInfo = "Goenka links add kar"
	}
	return fmt.Sprintf("%s\n%s\n%s for %s is starting soon!\nHere's the link to join: %s\n%s",
		REMINDER_INTRO,
		allRelevantRoles,
		class.ClassType,
		class.SubjectName, class.MeetLink, auxInfo)
}
