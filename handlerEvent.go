package main

import (
	"encoding/json"
	"fmt"
)

func addEventHandler(guest1, guest2, subject, date string) string {
	url := "http://profiler:8081/addEvent/" + guest1 + "/" + guest2 + "/" + date + "/" + subject
	res := postDataProfiler(url)

	return res
}

func delEventHandler(guest1, guest2, date string) string {
	url := "http://profiler:8081/delEvent/" + guest1 + "/" + guest2 + "/" + date
	resp := postDataProfiler(url)
	return resp
}

func listEventHandler(userID string) []Event {
	resp := getDataProfiler(userID, "http://profiler:8081/listEvent/"+userID)
	var events []Event
	err := json.Unmarshal([]byte(resp[12:len(resp)-1]), &events)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return events
}