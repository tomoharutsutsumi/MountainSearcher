package main

import ( 
  d "mountain-searcher/database"
  m "mountain-searcher/domain"
  mailer "mountain-searcher/mailer"
  "net/http"
  "text/template"
  "strconv"
  "fmt"
//   "os"
)

// var tpl *template.Template

// func init() {
//     tpl = template.Must(template.ParseFiles("./static/index.html"))
// }

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// wd, _ := os.Getwd()
	// mountains := d.GetMountains()
	// t, _ := template.ParseFiles("./static/index.html")
	// // t := template.Must(template.ParseFiles(wd + "/static/index.html"))
	// t.Execute(w, mountains)
	// _ = t.ExecuteTemplate(w, wd + "/static/index.html", mountains)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
  Prefecture := r.FormValue("prefecture")
  HasTentSite, _ := strconv.ParseBool(r.FormValue("hasTentSite"))
  DateForClimb := r.FormValue("dateForClimb")

  mountains := d.SearchMountains(Prefecture, HasTentSite)
  mountains = m.GetMountainsWithWeather(mountains, DateForClimb)
  fmt.Println(mountains)
  t, _ := template.ParseFiles("./static/search.html")
  t.Execute(w, mountains)
}

func mailHandler(w http.ResponseWriter, r *http.Request) {
  RequestForService:= r.FormValue("requestForService")
  fmt.Println(RequestForService)
  mailer.Send(RequestForService)
  http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func main() {
  http.HandleFunc("/", mainHandler)
  http.HandleFunc("/search", searchHandler)
  http.HandleFunc("/mail", mailHandler)
  http.ListenAndServe (":3000", nil)
}

// 簡単に感想を投稿できる窓口作りたいかも

// package main

// import (
// 	"net/http"
// 	"text/template"
// 	"fmt"
// )

// //TemperatureDataElem 気温データの一つのデータセット
// type TemperatureDataElem struct {
// 	Label string
// 	Data  []float64
// }

// func mainHandler(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("./static/index.html")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	if err := t.Execute(w, nil); err != nil {
// 		panic(err.Error())
// 	}
// }

// func dataHandler(w http.ResponseWriter, r *http.Request) {
// 	var temperatureData []TemperatureDataElem
// 	temperatureData = append(temperatureData, TemperatureDataElem{
// 		Label: "沖縄県",
// 		Data:  []float64{17.0, 17.1, 18.9, 21.4, 24.0, 26.8, 28.9, 28.7, 27.6, 25.2, 22.1, 18.7},
// 	})
// 	temperatureData = append(temperatureData, TemperatureDataElem{
// 		Label: "東京都",
// 		Data:  []float64{5.2, 5.7, 8.7, 13.9, 18.2, 21.4, 25.0, 26.4, 22.8, 17.5, 12.1, 7.6},
// 	})
// 	// ta, err := template.ParseFiles("./static/index.html")
// 	// fmt.Println(ta)
// 	t, err := template.ParseFiles("./static/index.html")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	if err := t.Execute(w, temperatureData); err != nil {
// 		panic(err.Error())
// 	}
// }

// func main() {
// 	fmt.Println("aaa")
// 	fmt.Println(dataHandler)
// 	fmt.Println("aaaccc")
// 	// fmt.Println(dataHandler())
// 	// http.HandleFunc("/", mainHandler)
// 	http.HandleFunc("/", dataHandler)
// 	http.ListenAndServe(":3000", nil)
// }