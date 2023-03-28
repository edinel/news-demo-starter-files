package tides

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	http     *http.Client
	key      string
	PageSize int
}

func (c *Client) FetchTidePredictions(query, page string) (*Results, error) {
	today := time.Now()
	aweek := today.AddDate(0, 0, 7)
	endpoint := fmt.Sprintf("https://api.tidesandcurrents.noaa.gov/api/prod/datagetter?begin_date=%s&end_date=%s&station=%s&product=predictions&datum=MLLW&time_zone=lst_ldt&interval=hilo&units=english&application=DataAPI_Sample&format=json", today.Format("20060102"), aweek.Format("20060102"), url.QueryEscape(query))

	resp, err := c.http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	res := &Results{}
	return res, json.Unmarshal(body, res)
}

func NewClient(httpClient *http.Client, key string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, key, pageSize}
}

type Results struct {
	Predictions []struct {
		Time  string `json:"t"`
		Value string `json:"v"`
		HiLo  string `json:"type"`
	} `json:"predictions"`
}

func PrintTideStruct(result Results, b *bytes.Buffer) {
	for i, eachPrediction := range result.Predictions {
		if eachPrediction.HiLo == "H" {
			fmt.Printf("At %v, the tide will be %v above normal\n", eachPrediction.Time, eachPrediction.Value)
			b.WriteString(fmt.Sprintf("%v,<br>", eachPrediction))
		} else {
			fmt.Printf("At %v, the tide will be %v below normal\n", eachPrediction.Time, eachPrediction.Value)
			b.WriteString(fmt.Sprintf("At %v, the tide will be %v below normal<br>%v <br>", eachPrediction.Time, eachPrediction.Value, eachPrediction))
		}
		if i < 0 {
			fmt.Errorf("Should Never Happen")
		}
	}
}
