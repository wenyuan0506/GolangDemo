package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var (
	once    sync.Once
	logFile *os.File
)

// Logger 是 HTTP middleware，記錄詳細請求與回應，並將日誌寫入檔案
func Logger(next http.Handler) http.Handler {
	once.Do(initLogger)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 捕捉 request body
		reqBody := readRequestBody(r)

		// 觸發 handler 並攔截 response
		lrw := newLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)

		// 組合完整請求資訊
		fullURL := buildFullURL(r)
		respBody := lrw.bodyBuffer.String()

		// 寫入日誌

		log.Printf("%s %s -> %d %d bytes\nRequest: %s\nResponse: %s",
			r.Method, fullURL,
			lrw.statusCode, lrw.bytesWritten,
			reqBody, respBody,
		)
	})
}

// initLogger 初始化日誌：載入 .env，設定目錄/保留，建立當天檔案，清理舊檔
func initLogger() {
	godotenv.Load()

	dir := getLogDir()
	os.MkdirAll(dir, 0755)

	retention := getRetentionDays()
	openTodayLog(dir)
	cleanupOldLogs(dir, retention)

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// getLogDir 讀取 LOG_DIR，預設 "logs"
func getLogDir() string {
	dir := os.Getenv("LOG_DIR")
	if strings.TrimSpace(dir) == "" {
		dir = "logs"
	}
	return dir
}

// getRetentionDays 讀取 LOG_RETENTION_DAYS，預設 7
func getRetentionDays() int {
	days := 7
	if s := os.Getenv("LOG_RETENTION_DAYS"); s != "" {
		if v, err := strconv.Atoi(s); err == nil && v > 0 {
			days = v
		}
	}
	return days
}

// openTodayLog 開啟或建立今日日誌檔，並設定全域 logFile
func openTodayLog(dir string) {
	today := time.Now().Format("2006-01-02")
	path := filepath.Join(dir, today+".log")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("無法開啟日誌檔：%v", err)
	}
	logFile = f
}

// cleanupOldLogs 刪除超過指定天數的 .log 檔
func cleanupOldLogs(dir string, days int) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("日誌清理失敗：%v", err)
		return
	}
	cutoff := time.Now().AddDate(0, 0, -days)
	for _, f := range files {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".log") {
			continue
		}
		base := strings.TrimSuffix(f.Name(), ".log")
		if t, err := time.Parse("2006-01-02", base); err == nil && t.Before(cutoff) {
			os.Remove(filepath.Join(dir, f.Name()))
		}
	}
}

// loggingResponseWriter 攔截回應
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int
	bodyBuffer   *bytes.Buffer
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK, 0, &bytes.Buffer{}}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	lrw.bodyBuffer.Write(b)
	n, err := lrw.ResponseWriter.Write(b)
	lrw.bytesWritten += n
	return n, err
}

// readRequestBody 讀取並還原 request body
func readRequestBody(r *http.Request) string {
	if r.Body == nil {
		return ""
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("讀取 request body 失敗：%v", err)
		return ""
	}
	// 還原供 handler 使用
	r.Body = io.NopCloser(bytes.NewBuffer(b))
	return string(b)
}

// buildFullURL 組合 Path 和 Query
func buildFullURL(r *http.Request) string {
	url := r.URL.Path
	if q := r.URL.RawQuery; q != "" {
		url += "?" + q
	}
	return url
}
