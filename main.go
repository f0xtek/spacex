package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type flight struct {
	FlightNum   int    `json:"flight_number"`
	MissionName string `json:"mission_name"`
	LaunchDate  int64  `json:"launch_date_unix"`
	Rocket      struct {
		Name string `json:"rocket_name"`
	} `json:"rocket"`
	LaunchSite struct {
		SiteName string `json:"site_name_long"`
	} `json:"launch_site"`
	Links struct {
		VideoLink string `json:"video_link"`
	} `json:"links"`
	SuccessfulLaunch bool `json:"launch_success"`
}

func main() {
	fd := getFlightDetails()
	t := formatTime(fd.LaunchDate)

	if fd.SuccessfulLaunch {
		fmt.Printf("The latest successful SpaceX mission was %s.\nThe flight number was %d and launched from %s on %s.\nThe rocker used was a %s.\n\nHere is a link to the launch video stream: %s\n\n", fd.MissionName, fd.FlightNum, fd.LaunchSite.SiteName, t, fd.Rocket.Name, fd.Links.VideoLink)
	} else {
		// TODO: handle failed launches
		fmt.Println("No successful launches found!")
	}

}

func getFlightDetails() flight {
	latestLaunch := getLatestLaunch()

	var f flight
	err := json.Unmarshal([]byte(latestLaunch), &f)
	if err != nil {
		log.Fatal("Unmarshal failed", err)
	}

	return f
}

func formatTime(t int64) string {
	// Convert unix time to RFC822
	u := strconv.FormatInt(t, 10)
	i, err := strconv.ParseInt(u, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	tm := time.Unix(i, 0)
	ft := tm.Format(time.RFC822)

	return ft
}

func getLatestLaunch() []byte {
	res, err := http.Get("https://api.spacexdata.com/v3/launches/latest")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}
