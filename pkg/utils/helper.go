package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func Debug(obj any) {
	raw, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(raw))
}

func LocalTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")

	return LocalTime().In(loc)
}

func ConvertStringTimeToTime(t string) time.Time {
	layout := time.RFC3339
	result, err := time.Parse(layout, t)
	if err != nil {
		log.Printf("Error: Parse time failed: %s", err.Error())
	}

	loc, _ := time.LoadLocation("Asia/Bangkok")
	result = result.In(loc)

	return result
}
