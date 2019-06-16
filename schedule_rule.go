package rachio

// import "encoding/json"

// Represents the schedule rule - https://rachio.readme.io/docs/publicscheduleruleid
type ScheduleRule struct {
	ID    string `json:"id"`
	Zones []struct {
		ZoneID    string `json:"zoneId"`
		Duration  int    `json:"duration"`
		SortOrder int    `json:"sortOrder"`
	} `json:"zones"`
	ScheduleJobTypes []string `json:"scheduleJobTypes"`
	Operator         string   `json:"operator"`
	CycleSoakStatus  string   `json:"cycleSoakStatus"`
	StartDate        int64    `json:"startDate"`
	Name             string   `json:"name"`
	Enabled          bool     `json:"enabled"`
	StartDay         int      `json:"startDay"`
	StartMonth       int      `json:"startMonth"`
	StartYear        int      `json:"startYear"`
	TotalDuration    int      `json:"totalDuration"`
	EndDate          int64    `json:"endDate"`
	EtSkip           bool     `json:"etSkip"`
	ExternalName     string   `json:"externalName"`
	CycleSoak        bool     `json:"cycleSoak"`
}
