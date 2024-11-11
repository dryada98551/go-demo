package log

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// 定義日誌級別常數
const (
	LogLevelInfo    = "INFO"
	LogLevelWarning = "WARNING"
	LogLevelError   = "ERROR"
)

// LogEntry 定義日誌條目的結構
type LogEntry struct {
	Message  string `json:"message"`
	Severity string `json:"severity"`
	User     string `json:"user,omitempty"`
	Time     string `json:"time"`
	// 可以根據需要添加更多字段
}

// INFO 寫入日誌到標準輸出
func INFO(message string) {

	hostname, err := os.Hostname()

	entry := LogEntry{
		Message:  message,
		Severity: "INFO",
		User:     hostname,
		Time:     time.Now().Format(time.RFC3339),
	}
	
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Fatalf("Error marshaling log entry: %v", err)
	}

	log.Println(string(jsonEntry))
}

// WARNING 寫入日誌到標準輸出
func WARNING(message string) {

	hostname, err := os.Hostname()

	entry := LogEntry{
		Message:  message,
		Severity: "WARNING",
		User:     hostname,
		Time:     time.Now().Format(time.RFC3339),
	}

	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Fatalf("Error marshaling log entry: %v", err)
	}

	log.Println(string(jsonEntry))
}

// ERROR 寫入日誌到標準輸出
func ERROR(message string) {

	hostname, err := os.Hostname()

	entry := LogEntry{
		Message:  message,
		Severity: "ERROR",
		User:     hostname,
		Time:     time.Now().Format(time.RFC3339),
	}

	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Fatalf("Error marshaling log entry: %v", err)
	}

	log.Println(string(jsonEntry))
}
