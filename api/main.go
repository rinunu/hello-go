package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"time"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rinunu/hello-go/api/rss"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<a href=\"/rss/sources/1/articles\">/rss/sources/1/articles</a>")
}

func initDb() *sqlx.DB {
	var db, err = sqlx.Open("mysql", "root:root@tcp(mysql:3306)/hello_go")

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db;
}

func main() {
	db := initDb()

	r := mux.NewRouter()

	rssService := rss.NewService(db)

	setupTestData(rssService)

	r.HandleFunc("/", IndexHandler)

	rss.MakeHandler(r.PathPrefix("/rss").Subrouter(), rssService)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func setupTestData(rssService *rss.RssService) {
	rssService.Subscribe("https://blog.jetbrains.com/kotlin/feed/")
	rssService.Subscribe("https://www.suzukikenichi.com/blog/feed/")
	rssService.Subscribe("https://www.digitaltrends.com/feed/")
	rssService.Subscribe("https://gizmodo.com/rss")
}
