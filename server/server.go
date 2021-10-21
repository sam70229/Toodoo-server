package server

import (
	"log"
	"net/http"

	_ "Toodoo/database"
	"Toodoo/logger"
	"Toodoo/routes"
)

func Run() {
    // f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    // checkErr(err)

    // defer f.Close()

    // log.SetOutput(f)

    logger.Info.Println("Starting server...")

    // router 讓網址可以對應處理函式
    // http.HandleFunc("/hello", hello)
    // http.HandleFunc("/servers", testJson)
    // 監聽 8080 port    
    // controller.Connect()
    // database.Connect()
    router := routes.NewRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}