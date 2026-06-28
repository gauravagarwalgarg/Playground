#include <iostream>
#include <memory>

using namespace std;

class Shape {
  public:
    virtual unique_ptr<Shape> clone() const = 0;
    virtual void draw() const = 0;
    virtual ~Shape() {}
};

class Circle : public Shape {
  private:
    int radius;

  public:
    Circle(int r) : radius(r) {}

    unique_ptr<Shape> clone() const override {
      return make_unique<Circle>(*this);
    }

    void draw() const override {
      cout << "Drawing Circle with radius " << radius << endl;
    }
};

class Rectangle : public Shape {
  private:
    int height, width;

  public:
    Rectangle(int w, int h) : width(w), height(h) {}

    unique_ptr<Shape> clone() const override {
      return make_unique<Rectangle>(*this);
    }

    void draw() const override {
      cout << "Drawing Rectangle with width " << width << " and height " << height << endl;
    }
};

int main() {
  unique_ptr<Shape> circle = make_unique<Circle>(10);
  unique_ptr<Shape> rectangle = make_unique<Rectangle>(5, 8);

  unique_ptr<Shape> clonedCircle = circle->clone();
  unique_ptr<Shape> cloneRectangle = rectangle->clone();

  circle->draw();
  clonedCircle->draw();

  rectangle->draw();
  cloneRectangle->draw();

  return 0;
}
