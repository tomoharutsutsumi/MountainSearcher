package domain

import (
  domain "mountain-searcher/domain/mountain"
  "encoding/json"
  "fmt"
  "io"
  "net/http"
  "time"

  // "context"

  // "github.com/hectormalot/omgo"
)

type Mountain struct {
  Id           int
  Name         string
  Prefecture   string
  HasTentSite  bool
  Latitude     float64
  Longitude    float64
  WeatherName  string
}

func getWeather(m Mountain, day string) string {
  url := "https://api.open-meteo.com/v1/forecast"
  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
	  panic(err.Error())
	}
  q := req.URL.Query()
  q.Add("latitude", fmt.Sprintf("%g", m.Latitude))
  q.Add("longitude", fmt.Sprintf("%g", m.Longitude)) 
  q.Add("start_date", day)
  q.Add("end_date", day)
  q.Add("daily", "weathercode")
  q.Add("timezone", "Asia/Tokyo")
  req.URL.RawQuery = q.Encode()
  client := http.DefaultClient
  resp, err := client.Do(req)
  body, err := io.ReadAll(resp.Body)
  var w domain.Weather
  fmt.Println(string(body))
  //b := []byte(`{"latitude":45.2,"longitude":141.25,"generationtime_ms":0.28502941131591797,"utc_offset_seconds":32400,"timezone":"Asia/Tokyo","timezone_abbreviation":"JST","elevation":1681.0,"daily_units":{"time":"iso8601","weathercode":"wmo code"},"daily":{"time":["2022-12-01"],"weathercode":[73]}}`)
  b := []byte(`{"daily":{"time":["2022-12-01"],"weathercode":[73]}}`)
  if err := json.Unmarshal(b, &w); err != nil {
		panic(err.Error())
	}
  fmt.Println(w.Daily.Weathercode[0])
  defer resp.Body.Close()
  return domain.GetWeatherName(w.Daily.Weathercode[0])
}

//{"error":true,"reason":"Parameter 'start_date' is out of allowed range from 2022-06-08 to 2023-01-10"}
//=> 2022/12/27時点　とりあえず一週間先のdateはvalidateする=> 値オブジェクトにしてビジネスルール追加

func GetMountainsWithWeather(mountains []Mountain, day string) ([]Mountain, []string) {
  var mountainsWithWeather []Mountain
  isValid, messages := dateForClimbIsValid(day)
  fmt.Println(isValid)
  if isValid {
    for _, m := range mountains {
      m.WeatherName = getWeather(m, day)
      mountainsWithWeather = append(mountainsWithWeather, m)
    }
    return mountainsWithWeather, messages
  } else {
    return mountainsWithWeather, messages
  }
}

func dateForClimbIsValid(day string) (bool, []string) {
  messages := []string{}
  errorCount := 0
  
  if day == "" {
    errorCount++
    messages = append(messages, "登山日が入力されていません")
    return false, messages
  }
  
	layout := "2006-01-02"
	currentTime := time.Now()
	today, _ := time.Parse(layout, currentTime.Format(layout))
	paramDay, _ := time.Parse(layout, day)

  if paramDay.Before(today) {
    errorCount++
    messages = append(messages, "登山日が過去の日付となっています")
    return false, messages
  }

  twoWeeksLater := 13

  if int(paramDay.Sub(today).Hours() / 24) > twoWeeksLater {
    errorCount++
    messages = append(messages, "入力できる登山日は今日から2週間以内になります")
    return false, messages
  }

  return true, messages
}

// const (
//   Shizuoka int = iota + 22
// )

//https://paiza.hatenablog.com/entry/2021/11/04/130000
// https://www.gsi.go.jp/kihonjohochousa/kihonjohochousa41140.html

// https://www.rescuework.jp/ordering/
// https://www.momonayama.net/hundred_mt_list_data/list.html
// https://montrek55.com/tentsite/