package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    // Render the home html page from static folder
    http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
    // Render the course html page
    http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
    // Render the about html page
    http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
    // Render the contact html page
    http.ServeFile(w, r, "static/contact.html")
}

func main() {
    // Root (/) should serve the homepage instead of just a text message
    http.HandleFunc("/", homePage)

    http.HandleFunc("/home", homePage)
    http.HandleFunc("/courses", coursePage)
    http.HandleFunc("/about", aboutPage)
    http.HandleFunc("/contact", contactPage)

    fmt.Println("🚀 Server running on port 8081...")
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        log.Fatal(err)
    }
}

