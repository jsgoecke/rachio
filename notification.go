package rachio

// import "encoding/json"

// Represents the notification event - https://rachio.readme.io/docs/publicnotificationwebhook_event_type
type Event struct {
	ID             string `json:"id"`
	Summary        string `json:"summary"`
	Hidden         bool   `json:"hidden"`
	ExternalID     string `json:"externalId"`
	Category       string `json:"category"`
	Type           string `json:"type"`
	SubType        string `json:"subType"`
	DeviceID       string `json:"deviceId"`
	EventDate      int64  `json:"eventDate"`
	CreateDate     int64  `json:"createDate"`
	LastUpdateDate int64  `json:"lastUpdateDate"`
}
