package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Timeslot struct {
	StartTimeString string `json:"start_time"`
	EndTimeString   string `json:"end_time"`
	startHour       int
	startMinute     int
	endHour         int
	endMinute       int
}

type Class struct {
	ClassType        string `json:"type"`
	SubjectName      string `json:"subject"`
	SubjectShorthand string `json:"shortname"`
	Weekday          string `json:"day"`
	TimeslotNos      []int  `json:"timeslots"`
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
	return nil
}

func (ts *Timeslot) isActive() bool {
	now := time.Now()
	nowHour, nowMinute := now.Local().Hour(), now.Local().Minute()
	return (ts.startHour <= nowHour && ts.startMinute <= nowMinute) &&
		(ts.endHour >= nowHour && ts.endMinute > nowMinute)
}
