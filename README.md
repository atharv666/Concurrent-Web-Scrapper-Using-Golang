# 🌐 Concurrent Web Scraper using Go

A **lightweight and concurrent web scraper** built in **Golang** that fetches multiple website titles simultaneously using **goroutines** and **WaitGroups**.  
This project demonstrates how concurrency in Go can drastically improve performance for I/O-bound tasks like web scraping.

---

## 🚀 Features

- ⚡ **Concurrent Execution** using Goroutines  
- 🔄 **Synchronization** handled via `sync.WaitGroup`  
- 💬 **Channel Communication** for safe data exchange between goroutines  
- 🧠 **HTML Parsing** with `golang.org/x/net/html`  
- 🧱 **Error Handling** for network and parsing issues  
- 📋 **Clean Output** showing each URL with its corresponding title

---

## 🧰 Tech Stack

| Component | Description |
|------------|-------------|
| **Language** | Go (Golang) |
| **Concurrency Model** | Goroutines + WaitGroup |
| **HTML Parsing** | `golang.org/x/net/html` |
| **HTTP Requests** | Go's `net/http` package |

---

## 🧩 Project Structure


---

## ⚙️ How It Works

1. A list of URLs is defined in the program.  
2. Each URL is passed to a **goroutine** that:
   - Fetches the HTML content using `http.Get()`
   - Parses the HTML document to extract the `<title>` tag
   - Sends the result back through a **channel**
3. The main thread waits for all goroutines to finish using a **WaitGroup**.  
4. Once all results are received, they are printed neatly in the console.

---

## 🧪 Example Output

```bash
PS D:\mygolang-lco\concurrent_webscrapper> go run main.go

https://www.geeksforgeeks.org -> GeeksforGeeks | Your All-in-One Learning Portal
https://github.com -> GitHub: Let’s build from here · GitHub
https://www.firefox.com -> Get Firefox for desktop and mobile — Firefox.com
https://www.google.com -> Google
https://golang.org -> The Go Programming Language
https://www.microsoft.com -> Microsoft – Cloud, Computers, Apps & Gaming
