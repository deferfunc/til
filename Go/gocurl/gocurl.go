package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const logDir = "./tmp"

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		url := args[0]
		handleRequest(url)
	} else {
		handleInteractiveMode()
	}
}

func handleRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(body))

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Println("Error creating log directory:", err)
		return
	}

	timestamp := time.Now().Format("2006-01-02_150405")
	logFile := filepath.Join(logDir, timestamp+".txt")

	file, err := os.Create(logFile)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, url)
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, string(body))
	writer.Flush()
}

func handleInteractiveMode() {
	files, err := os.ReadDir(logDir)
	if err != nil {
		fmt.Println("Error reading log directory:", err)
		return
	}

	sort.Slice(files, func(i, j int) bool {
		infoI, _ := files[i].Info()
		infoJ, _ := files[j].Info()
		return infoI.ModTime().After(infoJ.ModTime())
	})

	fmt.Println("Recent logs:")
	for i, file := range files {
		if i >= 10 {
			break
		}
		info, _ := file.Info()
		content, err := os.ReadFile(filepath.Join(logDir, file.Name()))
		if err != nil {
			fmt.Println("Error reading log file:", err)
			continue
		}
		lines := strings.Split(string(content), "\n")
		if len(lines) > 0 {
			url := lines[0]
			fmt.Printf("%d. %s %s\n", i+1, info.ModTime().Format("2006-01-02 15:04:05"), url)
		}
	}

	fmt.Print("Select a log number to re-request: ")
	var choice int
	fmt.Scan(&choice)

	if choice < 1 || choice > len(files) || choice > 10 {
		fmt.Println("Invalid choice")
		return
	}

	selectedFile := files[choice-1]
	content, err := os.ReadFile(filepath.Join(logDir, selectedFile.Name()))
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) > 0 {
		url := lines[0]
		handleRequest(url)
	}
}
