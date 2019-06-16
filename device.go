package rachio

// import "encoding/json"

// Represents the device - https://rachio.readme.io/docs/publicpersonid
type Device struct {
	ID                string        `json:"id"`
	FlexScheduleRules []interface{} `json:"flexScheduleRules"`
	SerialNumber      string        `json:"serialNumber"`
	HomeKitCompatible bool          `json:"homeKitCompatible"`
	Name              string        `json:"name"`
	Zones             []interface{} `json:"zones"`
	On                bool          `json:"on"`
	Latitude          float64       `json:"latitude"`
	Longitude         float64       `json:"longitude"`
	CreateDate        int64         `json:"createDate"`
	Status            string        `json:"status"`
	Deleted           bool          `json:"deleted"`
	ScheduleModeType  string        `json:"scheduleModeType"`
	Model             string        `json:"model"`
	MacAddress        string        `json:"macAddress"`
	ScheduleRules     []interface{} `json:"scheduleRules"`
}
