package main

import "fmt"

// Flyweight Pattern: Minimizes memory usage by sharing as much data as possible
// with similar objects. Separates intrinsic (shared) from extrinsic (unique) state.
// Example: Text rendering - share font/style objects across characters.

// Flyweight: shared intrinsic state
type CharacterStyle struct {
	FontFamily string
	FontSize   int
	Bold       bool
	Italic     bool
	Color      string
}

func (cs *CharacterStyle) Render(char rune, x, y int) {
	fmt.Printf("    '%c' at (%d,%d) [%s %dpx bold=%v italic=%v color=%s]\n",
		char, x, y, cs.FontFamily, cs.FontSize, cs.Bold, cs.Italic, cs.Color)
}

// Flyweight Factory: caches and reuses style objects
type StyleFactory struct {
	styles map[string]*CharacterStyle
}

func NewStyleFactory() *StyleFactory {
	return &StyleFactory{styles: make(map[string]*CharacterStyle)}
}

func (f *StyleFactory) GetStyle(font string, size int, bold, italic bool, color string) *CharacterStyle {
	key := fmt.Sprintf("%s-%d-%v-%v-%s", font, size, bold, italic, color)
	if style, exists := f.styles[key]; exists {
		return style // Reuse existing
	}
	style := &CharacterStyle{
		FontFamily: font,
		FontSize:   size,
		Bold:       bold,
		Italic:     italic,
		Color:      color,
	}
	f.styles[key] = style
	return style
}

func (f *StyleFactory) UniqueStyles() int {
	return len(f.styles)
}

// Character: combines flyweight (style) with extrinsic state (position)
type Character struct {
	Char  rune
	Style *CharacterStyle // shared flyweight
	X, Y  int             // extrinsic state (position)
}

// Document: collection of characters
type Document struct {
	chars   []Character
	factory *StyleFactory
}

func NewDocument() *Document {
	return &Document{factory: NewStyleFactory()}
}

func (d *Document) AddChar(char rune, x, y int, font string, size int, bold, italic bool, color string) {
	style := d.factory.GetStyle(font, size, bold, italic, color)
	d.chars = append(d.chars, Character{Char: char, Style: style, X: x, Y: y})
}

func (d *Document) Render() {
	for _, ch := range d.chars {
		ch.Style.Render(ch.Char, ch.X, ch.Y)
	}
}

func (d *Document) CharCount() int { return len(d.chars) }
func (d *Document) UniqueStyles() int { return d.factory.UniqueStyles() }

func main() {
	doc := NewDocument()

	// Add characters - many share the same style
	// "Hello" in Arial 12px black
	text := "Hello, World!"
	for i, ch := range text {
		doc.AddChar(ch, i*10, 0, "Arial", 12, false, false, "black")
	}

	// "IMPORTANT" in Arial 16px bold red
	important := "IMPORTANT"
	for i, ch := range important {
		doc.AddChar(ch, i*12, 20, "Arial", 16, true, false, "red")
	}

	// "note" in Courier 10px italic gray
	note := "This is a longer piece of text that should share styles efficiently across all its characters"
	for i, ch := range note {
		doc.AddChar(ch, i*8, 40, "Courier", 10, false, true, "gray")
	}

	fmt.Printf("Total characters: %d\n", doc.CharCount())
	fmt.Printf("Unique style objects: %d\n", doc.UniqueStyles())

	// Memory savings: instead of 100+ style objects, we have only 3
	if doc.UniqueStyles() == 3 {
		fmt.Println("PASS: only 3 unique style objects for all characters")
	} else {
		panic(fmt.Sprintf("FAIL: expected 3 styles, got %d", doc.UniqueStyles()))
	}

	// Render a sample
	fmt.Println("\n--- Render first 5 chars ---")
	for i := 0; i < 5 && i < len(doc.chars); i++ {
		ch := doc.chars[i]
		ch.Style.Render(ch.Char, ch.X, ch.Y)
	}

	// Verify sharing: same style object reference for same parameters
	style1 := doc.factory.GetStyle("Arial", 12, false, false, "black")
	style2 := doc.factory.GetStyle("Arial", 12, false, false, "black")
	if style1 == style2 {
		fmt.Println("PASS: same parameters return same object (pointer equality)")
	} else {
		panic("FAIL: should return cached style")
	}

	// Memory comparison
	withoutFlyweight := doc.CharCount() // Each char would have its own style
	withFlyweight := doc.UniqueStyles()
	savings := float64(withoutFlyweight-withFlyweight) / float64(withoutFlyweight) * 100
	fmt.Printf("Memory savings: %.1f%% (%d objects → %d shared styles)\n",
		savings, withoutFlyweight, withFlyweight)
	fmt.Println("PASS: flyweight pattern complete")
}
