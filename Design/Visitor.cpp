#include <iostream>
#include <memory>
#include <vector>

class FileVisitor;

class FileElement {
  public:
    virtual ~FileElement() = default;
    virtual void accept(FileVisitor& visitor) = 0;i
};

class TextFile : public FileElement {
  public:
    void accept(FileVisitor& visitor) override;
    void displayText() { std::cout << "Displaying text file contents.\n"; }
};

class ImageFile : public FileElement {
  public:
    void accept(FileVisitor& visitor) override;
    void renderImage() { std::cout << "Rendering image file.\n"; }
};


class FileVisitor {
  public:
    virtual ~FileVisitor() = default;
    virtual void visit(TextFile& file) = 0;
    virtual void visit(ImageFile& file) = 0;
};

class PrintVistor : public FileVisitor {
  public:
    void visit(TextFile& file) override {
      std::cout << "Printing text file...\n";
      file.displayText();
    }

    void visit(ImageFile& file) override {
      std::cout << "Printing image file...\n";
      file.renderImage();
    }
};

class SaveVisitor : public FileVisitor {
  public:
    void visit(TextFile& file) override {
      std::cout << "saving text file...\n";
    }

    void visit(ImageFile& file) override {
      std::cout << "Saving image file...\n";
    }

};

void TextFile::accept(FileVisitor& visitor) {
    visitor.visit(*this);
}

void ImageFile::accept(FileVisitor& visitor) {
    visitor.visit(*this);
}

int main() {
    std::vector<std::unique_ptr<FileElement>> files;
    files.push_back(std::make_unique<TextFile>());
    files.push_back(std::make_unique<ImageFile>());

    PrintVisitor printVisitor;
    SaveVisitor saveVisitor;

    std::cout << "== Printing Files ==\n";
    for (auto& file : files) {
        file->accept(printVisitor);
    }

    std::cout << "\n== Saving Files ==\n";
    for (auto& file : files) {
        file->accept(saveVisitor);
    }

    return 0;
}

