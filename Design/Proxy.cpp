// The Proxy pattern is a structural design pattern that provides a surrogate or placeholder for another object to control access to it. The proxy acts as an intermediate layer between the client and the real object, adding features like lazy initialization, access control, logging, caching, or security checks.

// Subject Interface
#include <iostream>
#include <memory>

using namespace std;

class Image {
  public:
    virtual void display() = 0;
    virtual ~Image() {}
};

// Real Object
class RealImage : public Image {
  private:
    string filename;

    void loadFromDisk() {
      cout << "Loading Image:" << filename << "from disk..." << endl;
    }

  public:
    RealImage(const string& file) : filename(file) {
      loadFromDisk();
    }

    void display() override {
      cout << "Displaying image: " << filename << endl;
    }
};

class ProxyImage : public Image {
  private:
    string filename;
    unique_ptr<RealImage> realImage;
  public:
    ProxyImage(const string& file) : filename(file) {}

    void display() override {
      if(!realImage) {
        realImage = make_unique<RealImage>(filename);
      }

      realImage->display();
    }
};

int main() {
  cout << "Creating ProxyImage object....\n";
  ProxyImage proxyimg("test.jpg");

  cout << "\n First display call:\n";
  proxyimg.display();

  cout << "\n Second display call:\n";
  proxyimg.display();

  return 0;
}
