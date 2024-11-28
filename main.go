package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Define a struct to hold the data for the template
type PageVariables struct {
	Time string
}

func main() {
	// Connect to MySQL database
	dsn := "root:Triveni@1234@tcp(127.0.0.1:3306)/time_api"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Handle requests to "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the current time in Toronto timezone
		location, err := time.LoadLocation("America/Toronto")
		if err != nil {
			http.Error(w, "Error loading Toronto timezone", http.StatusInternalServerError)
			return
		}
		currentTime := time.Now().In(location).Format("2006-01-02 15:04:05")

		// Insert the current time into the database
		_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", currentTime)
		if err != nil {
			http.Error(w, "Error logging time to database", http.StatusInternalServerError)
			return
		}

		// Prepare the data for the HTML template
		pageVariables := PageVariables{
			Time: currentTime,
		}

		// Parse and execute the template
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		// Render the template with the current time
		err = tmpl.Execute(w, pageVariables)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	})

	// Start the web server
	port := ":8080"
	fmt.Println("Starting server on http://localhost" + port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
