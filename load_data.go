package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var timeslotMap map[string]*Timeslot
var classes []*Class

func unmarshalTimeSlots() {
	var timeslotPath = "data/timeslots.json"
	jsonFile, err := os.Open(timeslotPath)

	if err != nil {
		log.Fatalf("Error in opening timeslot json file, %v.", err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &timeslotMap)

	// Testing, will remove before going into production
	for _, v := range timeslotMap {
		// Before Parsing for Testing => fmt.Println(v)
		v.parseTime()
		// After Parsing for Testing => fmt.Println(v)
	}
}

func unmarshalClasses() {
	sections := [2]string{"it-a", "it-b"}
	days := [7]string{"monday.json", "tuesday.json", "wednesday.json", "thursday.json", "friday.json"}
	basePath := "data/classes"
	slash := "/"
	for _, section := range sections {
		for _, day := range days {
			path := basePath + slash + section + slash + day
			jsonFile, err := os.Open(path)
			if err != nil {
				log.Fatalf("Error in opening timeslot json file, %v.", err)
			}
			var classesForOneDayForOneSection []*Class
			byteValue, _ := ioutil.ReadAll(jsonFile)
			json.Unmarshal(byteValue, &classesForOneDayForOneSection)
			for _, class := range classesForOneDayForOneSection {
				// fmt.Println(class)
				for _, timeslotNo := range class.TimeslotNos {
					class.Timeslots = append(class.Timeslots, timeslotMap[timeslotNo])
				}
				// fmt.Println(class, class.Timeslots)
			}
			classes = append(classes, classesForOneDayForOneSection...)
		}
	}
}
