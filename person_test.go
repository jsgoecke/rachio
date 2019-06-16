package rachio

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	PersonInfoJSON = `{"id":"3c59a593-04b8-42df-91db-758f4fe4a97f"}`
	PersonJSON     = `{
	"createDate": 1558904620000,
	"id": "3c59a593-04b8-42df-91db-758f4fe4a97f",
	"username": "jason@goecke.net",
	"fullName": "Jason Goecke",
	"email": "jason@goecke.net",
	"devices": [{
		"createDate": 1560698099000,
		"id": "1234",
		"status": "ONLINE",
		"zones": [{
			"id": "1234",
			"zoneNumber": 5,
			"name": "Zone 5",
			"enabled": false,
			"customNozzle": {
				"name": "FIXED_SPRAY_HEAD",
				"inchesPerHour": 1.5
			},
			"customSoil": {
				"name": "LOAM"
			},
			"customSlope": {
				"name": "ZERO_THREE",
				"sortOrder": 0
			},
			"customCrop": {
				"name": "Cool Season Grass",
				"coefficient": 0.8
			},
			"customShade": {
				"name": "LOTS_OF_SUN"
			},
			"availableWater": 0.17,
			"rootZoneDepth": 6.0,
			"managementAllowedDepletion": 0.5,
			"efficiency": 0.8,
			"yardAreaSquareFeet": 500,
			"imageUrl": "https://prod-media-photo.rach.io/1234",
			"lastWateredDuration": 4,
			"lastWateredDate": 1560702507000,
			"scheduleDataModified": false,
			"fixedRuntime": 0,
			"runtimeNoMultiplier": 1391,
			"wateringAdjustmentRuntimes": {
				"1": 2086,
				"2": 1739,
				"3": 1391,
				"4": 1043,
				"5": 695
			},
			"saturatedDepthOfWater": 0.56,
			"depthOfWater": 0.51,
			"maxRuntime": 10800,
			"runtime": 1391
		}],
		"timeZone": "America/Los_Angeles",
		"latitude": 37.493024,
		"longitude": -122.456849,
		"name": "23Alameda-Irrigation-Controller",
		"scheduleRules": [],
		"serialNumber": "AC1567796",
		"macAddress": "009D6BBEDD10",
		"on": true,
		"flexScheduleRules": [{
			"id": "1235",
			"zones": [{
				"zoneId": "1235",
				"sortOrder": 1,
				"duration": 3478
			}],
			"scheduleJobTypes": ["ANY"],
			"operator": "AFTER",
			"summary": "As needed end by sunrise",
			"cycleSoak": false,
			"startDate": 1560668400000,
			"name": "Dynamic Schedule",
			"enabled": true,
			"startDay": 16,
			"startMonth": 6,
			"startYear": 2019,
			"totalDuration": 35201,
			"externalName": "Dynamic Schedule",
			"type": "FLEX"
		}],
		"model": "GENERATION3_16ZONE",
		"scheduleModeType": "MANUAL",
		"deleted": false,
		"homeKitCompatible": false,
		"utcOffset": -25200000
	}],
	"deleted": false
}`
)

func TestStatesSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL

	client, _ := NewClient("mykey")

	Convey("Should get a person", t, func() {
		person, err := client.GetPerson()
		So(err, ShouldBeNil)
		So(person, ShouldNotBeNil)
		So(person.ID, ShouldEqual, "3c59a593-04b8-42df-91db-758f4fe4a97f")
	})

	BaseURL = previousURL
}
