package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Logging Framework LLD
// Demonstrates: Singleton (logger instance), Strategy (output targets),
//               Chain of Responsibility (log levels), Observer (log listeners).

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

func (l LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[l]
}

type LogEntry struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
	Logger    string
}

func (e LogEntry) Format() string {
	return fmt.Sprintf("[%s] %s [%s] %s",
		e.Timestamp.Format("2006-01-02 15:04:05.000"),
		e.Level, e.Logger, e.Message)
}

// LogSink - where logs are written to (Strategy pattern)
type LogSink interface {
	Write(entry LogEntry)
	Name() string
}

// ConsoleSink writes to stdout
type ConsoleSink struct{}

func (s *ConsoleSink) Name() string { return "console" }
func (s *ConsoleSink) Write(entry LogEntry) {
	fmt.Println(entry.Format())
}

// MemorySink stores logs in memory (for testing/auditing)
type MemorySink struct {
	mu      sync.Mutex
	entries []LogEntry
}

func NewMemorySink() *MemorySink {
	return &MemorySink{}
}

func (s *MemorySink) Name() string { return "memory" }
func (s *MemorySink) Write(entry LogEntry) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.entries = append(s.entries, entry)
}
func (s *MemorySink) Entries() []LogEntry {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]LogEntry{}, s.entries...)
}

// FilterSink only passes logs at or above a certain level
type FilterSink struct {
	minLevel LogLevel
	target   LogSink
}

func NewFilterSink(minLevel LogLevel, target LogSink) *FilterSink {
	return &FilterSink{minLevel: minLevel, target: target}
}

func (s *FilterSink) Name() string { return fmt.Sprintf("filter(%s→%s)", s.minLevel, s.target.Name()) }
func (s *FilterSink) Write(entry LogEntry) {
	if entry.Level >= s.minLevel {
		s.target.Write(entry)
	}
}

// Logger - the main logging interface
type Logger struct {
	mu       sync.RWMutex
	name     string
	level    LogLevel
	sinks    []LogSink
}

func NewLogger(name string, level LogLevel) *Logger {
	return &Logger{
		name:  name,
		level: level,
	}
}

func (l *Logger) AddSink(sink LogSink) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.sinks = append(l.sinks, sink)
}

func (l *Logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *Logger) log(level LogLevel, msg string, args ...interface{}) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if level < l.level {
		return
	}

	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   fmt.Sprintf(msg, args...),
		Logger:    l.name,
	}

	for _, sink := range l.sinks {
		sink.Write(entry)
	}
}

func (l *Logger) Debug(msg string, args ...interface{}) { l.log(DEBUG, msg, args...) }
func (l *Logger) Info(msg string, args ...interface{})  { l.log(INFO, msg, args...) }
func (l *Logger) Warn(msg string, args ...interface{})  { l.log(WARN, msg, args...) }
func (l *Logger) Error(msg string, args ...interface{}) { l.log(ERROR, msg, args...) }
func (l *Logger) Fatal(msg string, args ...interface{}) { l.log(FATAL, msg, args...) }

// LoggerFactory - Singleton that manages named loggers
type LoggerFactory struct {
	mu      sync.Mutex
	loggers map[string]*Logger
}

var (
	factory     *LoggerFactory
	factoryOnce sync.Once
)

func GetFactory() *LoggerFactory {
	factoryOnce.Do(func() {
		factory = &LoggerFactory{loggers: make(map[string]*Logger)}
	})
	return factory
}

func (f *LoggerFactory) GetLogger(name string) *Logger {
	f.mu.Lock()
	defer f.mu.Unlock()
	if logger, exists := f.loggers[name]; exists {
		return logger
	}
	logger := NewLogger(name, DEBUG)
	f.loggers[name] = logger
	return logger
}

func main() {
	// Setup: create a logger with multiple sinks
	logger := NewLogger("app.server", INFO)
	consoleSink := &ConsoleSink{}
	memorySink := NewMemorySink()
	errorSink := NewMemorySink()

	logger.AddSink(consoleSink)
	logger.AddSink(memorySink)
	logger.AddSink(NewFilterSink(ERROR, errorSink))

	// Log messages at various levels
	fmt.Println("--- Logging ---")
	logger.Debug("this should not appear (below INFO level)")
	logger.Info("server started on port %d", 8080)
	logger.Warn("connection pool at %d%% capacity", 80)
	logger.Error("failed to connect to database: %s", "timeout")
	logger.Fatal("unrecoverable error: %s", "out of memory")

	// Verify memory sink captured everything at INFO+
	entries := memorySink.Entries()
	if len(entries) != 4 { // INFO, WARN, ERROR, FATAL (DEBUG filtered by logger level)
		panic(fmt.Sprintf("FAIL: expected 4 log entries, got %d", len(entries)))
	}
	fmt.Printf("\nPASS: memory sink has %d entries (DEBUG filtered)\n", len(entries))

	// Verify error sink only has ERROR+
	errorEntries := errorSink.Entries()
	if len(errorEntries) != 2 { // ERROR, FATAL
		panic(fmt.Sprintf("FAIL: error sink expected 2, got %d", len(errorEntries)))
	}
	for _, e := range errorEntries {
		if e.Level < ERROR {
			panic("FAIL: error sink has non-error entry")
		}
	}
	fmt.Printf("PASS: error sink has %d entries (only ERROR+)\n", len(errorEntries))

	// Test log level change at runtime
	logger.SetLevel(ERROR)
	logger.Info("this should not appear after level change")
	if len(memorySink.Entries()) != 4 { // shouldn't increase
		panic("FAIL: INFO should be filtered after level change")
	}
	fmt.Println("PASS: runtime level change works")

	// Test singleton factory
	f := GetFactory()
	l1 := f.GetLogger("auth")
	l2 := f.GetLogger("auth")
	if l1 != l2 {
		panic("FAIL: factory should return same instance")
	}
	fmt.Println("PASS: singleton factory returns same logger")

	// Concurrent logging
	var wg sync.WaitGroup
	concurrentLogger := NewLogger("concurrent", DEBUG)
	concurrentMemory := NewMemorySink()
	concurrentLogger.AddSink(concurrentMemory)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			concurrentLogger.Info("request %d processed", id)
		}(i)
	}
	wg.Wait()

	concEntries := concurrentMemory.Entries()
	if len(concEntries) != 100 {
		panic(fmt.Sprintf("FAIL: expected 100 concurrent entries, got %d", len(concEntries)))
	}
	fmt.Printf("PASS: %d concurrent log entries captured\n", len(concEntries))

	// Verify entry content
	found := false
	for _, e := range concEntries {
		if strings.Contains(e.Message, "request 0") || strings.Contains(e.Message, "request 99") {
			found = true
			break
		}
	}
	if found {
		fmt.Println("PASS: log messages contain correct content")
	}
}
