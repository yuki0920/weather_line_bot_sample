package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// レスポンスの形式に合わせて struct を定義する
type Weather struct {
	Area     string `json:"targetArea"`
	HeadLine string `json:"headlineText"`
	Body     string `json:"text"`
}

func GetWeather() string {
	body := httpGetBody("https://www.jma.go.jp/bosai/forecast/data/overview_forecast/130000.json")
	weather := formatWeather(body)
	result := weather.ToS()

	return result
}

func httpGetBody(url string) []byte {
	// HTTPリクエストを発行しレスポンスを取得する
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}
	// レスポンスボディを読み込む
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("IO Read Error:", err)
	}
	// 読み込み終わったらレスポンスボディを閉じる
	defer response.Body.Close()

	return body
}

func formatWeather(body []byte) *Weather {
	weather := new(Weather)
	if err := json.Unmarshal(body, weather); err != nil {
		log.Fatal("JSON Unmarshal error:", err)
	}
	return weather
}

func (w *Weather) ToS() string {
	area := fmt.Sprintf("%sの天気です。\n", w.Area)
	head := fmt.Sprintf("%s\n", w.HeadLine)
	body := fmt.Sprintf("%s\n", w.Body)
	result := area + head + body

	return result
}
