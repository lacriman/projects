# 🧪 Go Practice Projects

This repository contains a collection of small, focused projects written in Go, designed to help solidify core concepts and improve practical skills through repetition and experimentation. Each folder is an isolated CLI-based application that showcases a particular aspect of the language.

## 📁 Project Structure

Each subfolder contains a standalone `main.go` file with zero dependencies beyond the Go standard library. Projects are easy to run, read, and modify.

### Projects Included

| Project                        | Description                                                                        |
| ------------------------------ | ---------------------------------------------------------------------------------- |
| `atm-simulator`                | Simulates an ATM interface with PIN check, withdraw, deposit, and balance handling |
| `calculator`                   | Basic CLI calculator that performs arithmetic operations                           |
| `concurrent-square-calculator` | Demonstrates concurrency using goroutines and channels                             |
| `fibonacci`                    | Generates a Fibonacci sequence up to a given number                                |
| `fizzbuzz`                     | Classic FizzBuzz challenge using loops and conditionals                            |
| `grade-tracker`                | Accepts and validates user input to calculate average grade                        |
| `concurrent-division-manager`  | Divides integers concurrently with error handling using WaitGroup                  |
| `api-mock-pinger`              | Pings mock APIs concurrently using goroutines and random delays                    |
| `file-line-counter`            | Counts lines in multiple files concurrently using bufio.Scanner                    |
| `word-counter`                 | Counts specific word occurrences across files using Mutex and goroutines           |

> More projects coming soon...

🔗 GitHub: [github.com/lacriman/projects](https://github.com/lacriman/projects)

---

## 🚀 Getting Started

### ✅ Requirements

- Go 1.20 or higher  
  [Install Go](https://go.dev/doc/install)

### ▶️ Run a Project

Each project is self-contained. To run any of them:

```bash
go run ./<folder-name>/main.go
