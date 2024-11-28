# Go API with MySQL and Time Logging Tutorial

This tutorial will guide you through the steps to build a Go-based API that returns the current time in Toronto, stores the time in a MySQL database, and serves an `index.html` page displaying the time.

## Objectives

- **Create a Go API Endpoint**: Develop an endpoint that returns the current time in Toronto in JSON format.
- **MySQL Database Integration**: Connect to a MySQL database and store the time data for each API request.
- **Time Zone Handling**: Ensure that the time returned is accurately adjusted to Toronto's timezone.
  
---

## Tasks

### 1. Set Up MySQL Database

First, you need to set up MySQL to store the logged times.

#### 1.1 Install MySQL

- Install MySQL on your system if you haven't already. You can find installation instructions on the [MySQL official website](https://dev.mysql.com/downloads/installer/).

#### 1.2 Create a New Database

- Log into MySQL using the following command:

```
mysql -u root -p
```

- Create a new database for the project:
```
CREATE DATABASE time_api;
```
### 1.3 Create a Table for Storing Timestamps

- Use the new database:

```
USE time_api;
```
- Create a table called time_log to store the logged times:
```
CREATE TABLE time_log (
  id INT AUTO_INCREMENT PRIMARY KEY,
  timestamp DATETIME NOT NULL
);
```
### 2. API Development

Now, let's create the Go application that will serve the API and interact with the MySQL database.

#### 2.1 Install Go and Set Up Project

Make sure you have Go installed on your system. You can download it from the [Go website](https://golang.org/dl/).

- Create a new Go project:

```
mkdir go-time-api
cd go-time-api
go mod init go-time-api
```

#### 2.2 Write Go Application Code

Create a new file called `main.go` and write the following code:

```
 Refer to the main.go file go throught the code/
```
#### This code does the following:
- Connects to a MySQL database called time_api.
- Defines a /current-time endpoint that returns the current time in Toronto in JSON format.
- Logs each request to the time_log table in the database.


### 3. Time Zone Conversion

The ```time``` package in Go will handle time zone conversion for us. In the code above, we use ```time.LoadLocation("America/Toronto")``` to load Toronto's time zone and adjust the current time accordingly.

### 4. Database Connection

In this section, we'll establish a connection to the MySQL database from our Go application using the `github.com/go-sql-driver/mysql` package.

#### 4.1 Install MySQL Driver for Go

First, install the MySQL driver package:

```
go get -u github.com/go-sql-driver/mysql
```
#### 4.2 Connect to the Database

In your `main.go` file, use the following code to establish a connection to the MySQL database:

```go
// Connect to MySQL database
db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/time_api")
if err != nil {
	log.Fatal(err)
}
defer db.Close()
```

#### Explanation:

- ```root:password``` are the MySQL user credentials (replace ```root``` with your ``` MySQL username```, and ```passwor```d with your ```MySQL password```).
- ```localhost:3306``` is the MySQL server address and port (adjust if you are using a different host).
- ```time_api``` is the name of the database you created earlier.

#### 4.3 Handling Database Errors

To ensure that the connection to the database is successful, it's important to handle any potential errors during the connection process. Use `db.Ping()` to verify that the database is accessible.

```go
// Check the connection to the database
if err := db.Ping(); err != nil {
	log.Fatal("Error connecting to database: ", err)
}
```

#### Explanation:

- ```db.Ping()``` sends a simple query to the database to verify that the connection is working.
- If the connection fails, the program will terminate with an error message using ```log.Fatal()```.

### 5. Serve HTML Page with Current Time

In this section, we will modify the Go application to serve an `index.html` file that displays information about the use case and the current time.

#### 5.1 Create the `index.html` File

First, create a new file called `index.html` inside a `templates` folder. The `index.html` file will display a message about the use case and show the current time.

Create the `templates` folder and the `index.html` file:

```bash
mkdir templates
touch templates/index.html
```
- Refer to the ```templates/index.html``` file go throught the code.

#### 5.2 Serve the HTML Page in Go

In this step, we will modify the Go application (`main.go`) to serve the `index.html` file from the `templates` folder.

Add the following code in your `main.go` to serve static files (such as `index.html`) from the `templates` folder:

```go
// Serve the HTML page (index.html)
http.Handle("/", http.FileServer(http.Dir("./templates")))

fmt.Println("Server running on http://localhost:8080")
log.Fatal(http.ListenAndServe(":8080", nil))
```
#### Explanation:

- ```http.Handle("/", http.FileServer(http.Dir("./templates"))):``` This line tells the Go server to serve static files from the templates folder when a request is made to the root URL``` /```.
- The index.html file will be automatically served when you navigate to ```http://localhost:8080/```.

#### 5.3 Update the current-time API to Log the Time

Now, we need to ensure that the `/current-time` API endpoint logs the current time into the MySQL database every time it is called. The current time will also be returned in the response as JSON.

Update the `/current-time` API endpoint in your `main.go` file with the following code:

```go
http.HandleFunc("/current-time", func(w http.ResponseWriter, r *http.Request) {
    // Get current time in Toronto
    location, err := time.LoadLocation("America/Toronto")
    if err != nil {
        http.Error(w, "Error loading timezone", http.StatusInternalServerError)
        return
    }
    currentTime := time.Now().In(location)

    // Insert current time into the database
    _, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", currentTime)
    if err != nil {
        http.Error(w, "Error inserting time into database", http.StatusInternalServerError)
        return
    }

    // Respond with the current time in JSON
    response := TimeResponse{Time: currentTime.Format(time.RFC3339)}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
})
```
- This ensures that each time the /current-time API is accessed, the current time is inserted into the MySQL database.

### 6. Error Handling
- Proper error handling is implemented in the code for both the MySQL operations and time zone conversions. If any error occurs, the server will respond with an appropriate HTTP status code (500 for internal server errors).

### 7. Running the Application
- After writing the code, you can run the Go application using the following command:
```bash
Copy code
go run main.go
```
This will start the server on http://localhost:8080.

### 8. Verify the Logs in MySQL
- To verify that the time is being logged in the MySQL database, you can execute the following commands inside the MySQL shell:
```bash
USE time_api;
SELECT * FROM time_log;
```
- You should see a list of timestamps corresponding to each API request.


### Conclusion

#### In this tutorial, you've learned how to:

- Build a Go API to provide the current time in Toronto.
- Store each request's timestamp in a MySQL database.
- Serve a simple HTML page displaying the current time.
- Handle time zone conversions using Go's time package.
