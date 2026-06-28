package main

import "fmt"

// Composite Pattern: Compose objects into tree structures to represent
// part-whole hierarchies. Lets clients treat individual objects and
// compositions uniformly.
// Example: File system (files and directories).

// Component interface
type FileSystemNode interface {
	Name() string
	Size() int64
	Display(indent string)
}

// Leaf: File
type File struct {
	name string
	size int64
}

func NewFile(name string, size int64) *File {
	return &File{name: name, size: size}
}

func (f *File) Name() string { return f.name }
func (f *File) Size() int64  { return f.size }
func (f *File) Display(indent string) {
	fmt.Printf("%s📄 %s (%d bytes)\n", indent, f.name, f.size)
}

// Composite: Directory
type Directory struct {
	name     string
	children []FileSystemNode
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) Name() string { return d.name }

func (d *Directory) Size() int64 {
	var total int64
	for _, child := range d.children {
		total += child.Size()
	}
	return total
}

func (d *Directory) Display(indent string) {
	fmt.Printf("%s📁 %s/ (%d bytes)\n", indent, d.name, d.Size())
	for _, child := range d.children {
		child.Display(indent + "  ")
	}
}

func (d *Directory) Add(node FileSystemNode) {
	d.children = append(d.children, node)
}

func (d *Directory) Remove(name string) {
	for i, child := range d.children {
		if child.Name() == name {
			d.children = append(d.children[:i], d.children[i+1:]...)
			return
		}
	}
}

func main() {
	// Build a file tree
	root := NewDirectory("project")

	src := NewDirectory("src")
	src.Add(NewFile("main.go", 1500))
	src.Add(NewFile("handler.go", 2200))
	src.Add(NewFile("model.go", 800))

	tests := NewDirectory("tests")
	tests.Add(NewFile("main_test.go", 900))
	tests.Add(NewFile("handler_test.go", 1100))

	root.Add(src)
	root.Add(tests)
	root.Add(NewFile("go.mod", 120))
	root.Add(NewFile("README.md", 450))

	// Display the tree
	root.Display("")

	// Verify sizes roll up correctly
	expectedSize := int64(1500 + 2200 + 800 + 900 + 1100 + 120 + 450)
	if root.Size() == expectedSize {
		fmt.Printf("\nPASS: total size = %d bytes\n", root.Size())
	} else {
		panic(fmt.Sprintf("FAIL: expected %d, got %d", expectedSize, root.Size()))
	}

	// Leaf and composite share the same interface
	var node FileSystemNode = src
	fmt.Printf("PASS: src dir size via interface = %d\n", node.Size())
}
