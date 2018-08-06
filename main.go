package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"os"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func test(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("templates/index.html", "templates/footer.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("templates/cover.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening well on PORT%s...\n", addr)

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/test", test)
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	http.ListenAndServe(":5000", router)

}
