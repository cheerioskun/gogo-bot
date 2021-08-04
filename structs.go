package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type PairString struct {
	roleName, roleEnvName string
}
type Timeslot struct {
	StartTimeString string `json:"start_time"`
	EndTimeString   string `json:"end_time"`
	startHour       int
	startMinute     int
	endHour         int
	endMinute       int
}

type Class struct {
	ClassType        string   `json:"type"`
	SubjectName      string   `json:"subject"`
	SubjectShorthand string   `json:"shortname"`
	Weekday          string   `json:"day"`
	TimeslotNos      []string `json:"timeslots"`
	Timeslots        []*Timeslot
	Sections         []string `json:"sections"`
	FacultyName      string   `json:"faculty_name"`
	MeetLink         string   `json:"meet_link"`
}

func (ts *Timeslot) parseTime() error {
	if ts.StartTimeString == "" || ts.EndTimeString == "" {
		log.Printf("timeslots in db not valid")
		return fmt.Errorf("timeslot object is invalid")
	}
	startTimeArray := strings.Split(ts.StartTimeString, ":")
	endTimeArray := strings.Split(ts.EndTimeString, ":")
	if len(startTimeArray) != 2 || len(endTimeArray) != 2 {
		log.Printf("timeslots in db not valid")
		return fmt.Errorf("timeslot object is invalid")
	}
	var err error
	ts.startHour, err = strconv.Atoi(startTimeArray[0])
	if err != nil {
		log.Printf("could not parse time string")
		return fmt.Errorf("error parsing time string: %q", err)
	}
	ts.startMinute, err = strconv.Atoi(startTimeArray[1])
	if err != nil {
		log.Printf("could not parse time string")
		return fmt.Errorf("error parsing time string: %q", err)
	}
	ts.endHour, err = strconv.Atoi(endTimeArray[0])
	if err != nil {
		log.Printf("could not parse time string")
		return fmt.Errorf("error parsing time string: %q", err)
	}
	ts.endMinute, err = strconv.Atoi(endTimeArray[1])
	if err != nil {
		log.Printf("could not parse time string")
		return fmt.Errorf("error parsing time string: %q", err)
	}
	return nil
}

func (ts *Timeslot) isComingUp() bool {
	// Check for 5 minutes from now
	now := time.Now().Add(time.Minute * 5)
	nowHour, nowMinute := now.Local().Hour(), now.Local().Minute()
	// Check if it falls within start + 5 minutes. This will only happen once if we're checking every 5 minutes
	return (ts.startHour <= nowHour && ts.startMinute <= nowMinute) &&
		(ts.startMinute+5 > nowMinute) &&
		(ts.endHour > nowHour)
}
