# URL Shortener

URL Shortener is a simple and efficient web service for transforming long and unwieldy URLs into concise, user-friendly links. It provides a convenient way to create, manage, and access shortened URLs.

## Features:

- Shorten long URLs with a single API call.
- Redirect users to the original URL using the shortened link.
- Securely store URL mappings in a SQLite database.
- Generate unique, six-character short aliases for each URL.
- Error handling and validation to ensure a smooth user experience.

## Getting Started

Prerequisites:

- Go programming language
- SQLite database

Installation:

1. Clone this repository:

   git clone https://github.com/your-username/url-shortener.git
   cd url-shortener

2. Initialize the SQLite database:

   sqlite3 url-shortener.db

3. Build and run the project:

   go build
   ./url-shortener

The URL Shortener service will start on http://localhost:8080.

## Usage

Shorten a URL:

To shorten a URL, make a POST request to the /shorten endpoint with the longURL parameter:

   curl -X POST -d "longURL=https://www.example.com" http://localhost:8080/shorten

You will receive a JSON response with the shortened URL:

   {"shortURL":"Ff5ypq"}

Access a Shortened URL:

To access a shortened URL, simply enter it in your browser's address bar, e.g., http://localhost:8080/Ff5ypq. You will be redirected to the original URL.

Database Schema

The project uses an SQLite database to store URL mappings. The schema consists of a single table named urls with the following structure:

- id: Integer, primary key
- short: Text, the short URL
- long: Text, the original URL

## Contributing

Contributions are welcome! If you have suggestions, feature requests, or found a bug, please open an issue or create a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
