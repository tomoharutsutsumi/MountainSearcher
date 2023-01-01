package database

import (
  domain "mountain-searcher/domain"
  "database/sql"
  // "html/template"
  // "log"

  _ "github.com/go-sql-driver/mysql"
)

// func GetMountains() []domain.Mountain {
//   var mountains []domain.Mountain
//   db, err := sql.Open("mysql", "example:example@tcp(db:3306)/mountain_searcher")
//   defer db.Close()
//   if err != nil {
//     panic(err.Error())
//   }
//   rows, _ := db.Query("SELECT * FROM mountains")
//   for rows.Next() {
//     mountain := domain.Mountain{}
//     err = rows.Scan(&mountain.Id, &mountain.Name, &mountain.Prefecture, &mountain.HasTentSite)
//     if err != nil {
//       panic(err.Error())
//     }
//     mountains = append(mountains, mountain)
//   }
//   // if err := tmpl.ExecuteTemplate(w, "index.html", mountains); err != nil {
// 	// 	log.Fatal(err)
// 	// }
//   return mountains
// }

//　引数多いからrefactoringしたい。おそらくqueryのモデル作ってやる感じでいいかもしらん
func SearchMountains(Prefecture string, HasTentSite bool) []domain.Mountain {
  var mountains []domain.Mountain
  db, err := sql.Open("mysql", "example:example@tcp(db:3306)/mountain_searcher")
  defer db.Close()
  if err != nil {
    panic(err.Error())
  }
  rows, _ := db.Query("SELECT * FROM mountains WHERE prefecture = ? AND hasTentSite = ?", Prefecture, HasTentSite)
  for rows.Next() {
    mountain := domain.Mountain{}
    err = rows.Scan(&mountain.Id, &mountain.Name, &mountain.Prefecture, &mountain.HasTentSite, &mountain.Latitude, &mountain.Longitude)
    if err != nil {
      panic(err.Error())
    }
    mountains = append(mountains, mountain)
  }

  return mountains
}