package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Template Method Pattern: Define the skeleton of an algorithm in a method,
// deferring some steps to subclasses (in Go: via interfaces + embedding).
// Example: Data export pipeline (CSV, JSON, XML).

// Step interface - the varying parts of the algorithm
type ExportSteps interface {
	Header() string
	FormatRow(data map[string]string) string
	Footer() string
	Extension() string
}

// Template - the fixed algorithm skeleton
type DataExporter struct {
	steps ExportSteps
}

func NewDataExporter(steps ExportSteps) *DataExporter {
	return &DataExporter{steps: steps}
}

// Export is the template method - defines the algorithm structure
func (e *DataExporter) Export(dataset []map[string]string) string {
	var builder strings.Builder

	// Step 1: Write header
	builder.WriteString(e.steps.Header())
	builder.WriteString("\n")

	// Step 2: Write each row (hook for subclass formatting)
	for _, row := range dataset {
		builder.WriteString(e.steps.FormatRow(row))
		builder.WriteString("\n")
	}

	// Step 3: Write footer
	footer := e.steps.Footer()
	if footer != "" {
		builder.WriteString(footer)
		builder.WriteString("\n")
	}

	return builder.String()
}

// Checksum is a shared utility step
func (e *DataExporter) Checksum(content string) string {
	h := sha256.Sum256([]byte(content))
	return hex.EncodeToString(h[:8])
}

// Concrete: CSV Export
type CSVExport struct {
	columns []string
}

func (c *CSVExport) Extension() string { return ".csv" }
func (c *CSVExport) Header() string    { return strings.Join(c.columns, ",") }
func (c *CSVExport) Footer() string    { return "" }
func (c *CSVExport) FormatRow(data map[string]string) string {
	vals := make([]string, len(c.columns))
	for i, col := range c.columns {
		vals[i] = data[col]
	}
	return strings.Join(vals, ",")
}

// Concrete: JSON Export
type JSONExport struct {
	columns []string
}

func (j *JSONExport) Extension() string { return ".json" }
func (j *JSONExport) Header() string    { return "[" }
func (j *JSONExport) Footer() string    { return "]" }
func (j *JSONExport) FormatRow(data map[string]string) string {
	pairs := make([]string, 0, len(j.columns))
	for _, col := range j.columns {
		pairs = append(pairs, fmt.Sprintf(`"%s":"%s"`, col, data[col]))
	}
	return "  {" + strings.Join(pairs, ", ") + "},"
}

func main() {
	dataset := []map[string]string{
		{"name": "Alice", "email": "alice@example.com", "role": "admin"},
		{"name": "Bob", "email": "bob@example.com", "role": "user"},
		{"name": "Charlie", "email": "charlie@example.com", "role": "user"},
	}
	columns := []string{"name", "email", "role"}

	// CSV export using the template
	csvExporter := NewDataExporter(&CSVExport{columns: columns})
	csvOutput := csvExporter.Export(dataset)
	fmt.Println("--- CSV Export ---")
	fmt.Println(csvOutput)

	if strings.Contains(csvOutput, "name,email,role") && strings.Contains(csvOutput, "Alice,alice@example.com,admin") {
		fmt.Println("PASS: CSV export")
	} else {
		panic("FAIL: CSV export")
	}

	// JSON export using the same template
	jsonExporter := NewDataExporter(&JSONExport{columns: columns})
	jsonOutput := jsonExporter.Export(dataset)
	fmt.Println("--- JSON Export ---")
	fmt.Println(jsonOutput)

	if strings.Contains(jsonOutput, "[") && strings.Contains(jsonOutput, `"name":"Bob"`) {
		fmt.Println("PASS: JSON export")
	} else {
		panic("FAIL: JSON export")
	}

	// Same algorithm, different formats - that's template method
	fmt.Printf("CSV checksum: %s\n", csvExporter.Checksum(csvOutput))
	fmt.Printf("JSON checksum: %s\n", jsonExporter.Checksum(jsonOutput))
}
