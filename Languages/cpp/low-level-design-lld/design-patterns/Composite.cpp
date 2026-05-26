#include <iostream>
#include <vector>
#include <memory>

using namespace std;

class FileSystemComponent {
  public:
    virtual void showDetails(int indent = 0) const = 0;
    virtual ~FileSystemComponent() {}
};

class File : public FileSystemComponent {
  private:
    string name;

  public:
    File(const string& name) : name(name) {}

    void showDetails(int indent = 0) const override {
      cout << string(indent, ' ') << "- File: " << name << endl;
    }
};

class Folder : public FileSystemComponent {
  private:
    string name;
    vector<unique_ptr<FileSystemComponent>> children;

  public:
    Folder(const string& name) : name(name) {}

    void add(unique_ptr<FileSystemComponent> component) {
      children.push_back(move(component));
    }

    void showDetails(int indent = 0) const override {
      cout << string(indent, ' ') << "+ Folder: " << name << endl;

      for(const auto& child : children) {
        child->showDetails(indent + 2);
      }
    }
};

int main() {
  unique_ptr<FileSystemComponent> file1 = make_unique<File>("document.txt");
  unique_ptr<FileSystemComponent> file2 = make_unique<File>("image.png");
  unique_ptr<FileSystemComponent> file3 = make_unique<File>("video.mp4");

  unique_ptr<Folder> root = make_unique<Folder>("Root");
  unique_ptr<Folder> subFolder1 = make_unique<Folder>("Pictures");
  unique_ptr<Folder> subFolder2 = make_unique<Folder>("Videos");

  subFolder1->add(move(file2));
  subFolder1->add(move(file3));

  root->add(move(file1));
  root->add(move(subFolder1));   // Adding Pictures folder to Root
  root->add(move(subFolder2));   // Adding Videos folder to Root

  root->showDetails();

  return 0;
}

