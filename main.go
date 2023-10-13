package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
    "net/http"
    "math/rand"
    "time"
    "log"
)

var db *sql.DB

func main() {
    r := gin.Default()

    // Initialize the SQLite database
    initDB()

    // Define routes
    r.POST("/shorten", shortenURL)
    r.GET("/:shortURL", redirectURL)

    r.Run(":8080")
}

func initDB() {
    var err error
    db, err = sql.Open("sqlite3", "./url-shortener.db")
    if err != nil {
        log.Fatal("Failed to open the database:", err)
    }
    
    // Create the table if it doesn't exist
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS urls (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            short TEXT NOT NULL,
            long TEXT NOT NULL
        )
    `)
    if err != nil {
        log.Fatal("Failed to create the table:", err)
    }
}

func shortenURL(c *gin.Context) {
    longURL := c.PostForm("longURL")
    if !isValidURL(longURL) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
        return
    }

    shortURL := generateShortURL()
    
    // Store the URL mapping in the database
    _, err := db.Exec("INSERT INTO urls (short, long) VALUES (?, ?)", shortURL, longURL)
    if err != nil {
        log.Println("Failed to insert into the database:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"shortURL": shortURL})
}

func redirectURL(c *gin.Context) {
    shortURL := c.Param("shortURL")

    // Retrieve the original URL from the database
    var longURL string
    err := db.QueryRow("SELECT long FROM urls WHERE short = ?", shortURL).Scan(&longURL)
    if err != nil {
        log.Println("Failed to retrieve the URL from the database:", err)
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }

    // Redirect the user to the original URL
    c.Redirect(http.StatusFound, longURL)
}

func generateShortURL() string {
    // Generate a random short URL
    characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    rand.Seed(time.Now().UnixNano())
    result := make([]byte, 6)
    for i := range result {
        result[i] = characters[rand.Intn(len(characters))]
    }
    return string(result)
}

func isValidURL(url string) bool {
		// Implement URL validation logic here
		// You can use a URL parsing library or regex for validation
		// Return true if the URL is valid, false otherwise
		// For example, you can use a regex pattern to validate URLs:
		// pattern := regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+`)
		// return pattern.MatchString(url)
		return true
}
