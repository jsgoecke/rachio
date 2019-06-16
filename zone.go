package rachio

// import "encoding/json"

// Represents the zone - https://rachio.readme.io/docs/publiczoneid
type Zone struct {
	ZoneNumber           int     `json:"zoneNumber"`
	Runtime              int     `json:"runtime"`
	ID                   string  `json:"id"`
	RootZoneDepth        int     `json:"rootZoneDepth"`
	Efficiency           float64 `json:"efficiency"`
	FixedRuntime         int     `json:"fixedRuntime"`
	ImageURL             string  `json:"imageUrl"`
	ScheduleDataModified bool    `json:"scheduleDataModified"`
	CustomShade          struct {
		Name string `json:"name"`
	} `json:"customShade"`
	DepthOfWater               float64 `json:"depthOfWater"`
	ManagementAllowedDepletion float64 `json:"managementAllowedDepletion"`
	CustomCrop                 struct {
		Name        string  `json:"name"`
		Coefficient float64 `json:"coefficient"`
	} `json:"customCrop"`
	SaturatedDepthOfWater float64 `json:"saturatedDepthOfWater"`
	CustomSoil            struct {
		Name string `json:"name"`
	} `json:"customSoil"`
	CustomNozzle struct {
		Name          string  `json:"name"`
		InchesPerHour float64 `json:"inchesPerHour"`
	} `json:"customNozzle"`
	AvailableWater             float64 `json:"availableWater"`
	YardAreaSquareFeet         int     `json:"yardAreaSquareFeet"`
	RuntimeNoMultiplier        int     `json:"runtimeNoMultiplier"`
	MaxRuntime                 int     `json:"maxRuntime"`
	Name                       string  `json:"name"`
	Enabled                    bool    `json:"enabled"`
	LastWateredDate            int64   `json:"lastWateredDate"`
	WateringAdjustmentRuntimes struct {
		Num1 int `json:"1"`
		Num2 int `json:"2"`
		Num3 int `json:"3"`
		Num4 int `json:"4"`
		Num5 int `json:"5"`
	} `json:"wateringAdjustmentRuntimes"`
}
