package main

import (
	"fmt"
	"log"
	"net/http"
)

// create a simple web server that listens on port 8080 and serves a simple HTML page
func main() {
	http.HandleFunc("/", handleHome)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Simple Web Server</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 50px auto;
            padding: 20px;
            background-color: #f5f5f5;
        }
        .container {
            background-color: white;
            padding: 40px;
            border-radius: 10px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
        }
        h1 {
            color: #333;
            text-align: center;
        }
        p {
            color: #666;
            line-height: 1.6;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Welcome to the Simple Web Server!</h1>
        <p>This is a basic HTTP server running on port 8080, built with Go.</p>
        <p>The server is working correctly and serving this HTML page.</p>
    </div>
</body>
</html>
`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
