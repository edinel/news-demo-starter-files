package tides

import (
	"net/http"
)

type Client struct {
	http     *http.Client
	key      string
	PageSize int
}

func NewClient(httpClient *http.Client, key string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, key, pageSize}
}

/*
type tideResult struct {
	Count    int `json:"count"`
	Units    any `json:"units"`
	Stations []struct {
		Tidal      bool   `json:"tidal"`
		Greatlakes bool   `json:"greatlakes"`
		Shefcode   string `json:"shefcode"`
		Details    struct {
			ID           string  `json:"id"`
			Established  string  `json:"established"`
			Removed      string  `json:"removed"`
			Noaachart    string  `json:"noaachart"`
			Timemeridian int     `json:"timemeridian"`
			Timezone     float64 `json:"timezone"`
			Origyear     string  `json:"origyear"`
			Self         string  `json:"self"`
		} `json:"details"`
		Sensors struct {
			Units   string `json:"units"`
			Sensors []struct {
				Status    int     `json:"status"`
				Refdatum  string  `json:"refdatum"`
				SensorID  string  `json:"sensorID"`
				Name      string  `json:"name"`
				Elevation float64 `json:"elevation"`
				Message   string  `json:"message"`
				Dcp       int     `json:"dcp"`
			} `json:"sensors"`
			Self string `json:"self"`
		} `json:"sensors"`
		Floodlevels struct {
			Self string `json:"self"`
		} `json:"floodlevels"`
		Datums struct {
			Self string `json:"self"`
		} `json:"datums"`
		Supersededdatums struct {
			Self string `json:"self"`
		} `json:"supersededdatums"`
		HarmonicConstituents struct {
			Self string `json:"self"`
		} `json:"harmonicConstituents"`
		Benchmarks struct {
			Self string `json:"self"`
		} `json:"benchmarks"`
		TidePredOffsets struct {
			Self string `json:"self"`
		} `json:"tidePredOffsets"`
		OfsMapOffsets struct {
			Self string `json:"self"`
		} `json:"ofsMapOffsets"`
		State        string `json:"state"`
		Timezone     string `json:"timezone"`
		Timezonecorr int    `json:"timezonecorr"`
		Observedst   bool   `json:"observedst"`
		Stormsurge   bool   `json:"stormsurge"`
		Nearby       struct {
			Self string `json:"self"`
		} `json:"nearby"`
		Forecast        bool    `json:"forecast"`
		Outlook         bool    `json:"outlook"`
		HTFhistorical   bool    `json:"HTFhistorical"`
		NonNavigational bool    `json:"nonNavigational"`
		ID              string  `json:"id"`
		Name            string  `json:"name"`
		Lat             float64 `json:"lat"`
		Lng             float64 `json:"lng"`
		Affiliations    string  `json:"affiliations"`
		Portscode       string  `json:"portscode"`
		Products        struct {
			Self string `json:"self"`
		} `json:"products"`
		Disclaimers struct {
			Self string `json:"self"`
		} `json:"disclaimers"`
		Notices struct {
			Self string `json:"self"`
		} `json:"notices"`
		Self     string `json:"self"`
		Expand   string `json:"expand"`
		TideType string `json:"tideType"`
	} `json:"stations"`
	Self any `json:"self"`
}
*/
