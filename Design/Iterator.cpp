// The Iterator pattern is a behavioral design pattern that provides a way to access elements of a collection sequentially without exposing its internal structure. It promotes decoupling between the collection and its traversal logic.
//

#include <iostream>
#include <vector>
#include <memory>

using namespace std;

class Iterator {
  public:
    virtual bool hasNext() = 0;
    virtual int next() = 0;
    virtual ~Iterator() {}
};

class NumberIterator : public Iterator {
  private:
    vector<int> numbers;
    size_t index;

  public:
    NumberIterator(const vector<int>& nums) : numbers(nums), index(0) {}

    bool hasNext() override {
      return index < numbers.size();
    }

    int next() override {
      return hasNext() ? numbers[index++] : -1;
    }
};

class IterableCollection {
  public:
    virtual shared_ptr<Iterator> createIterator() = 0;
    virtual ~IterableCollection() {}
};

class NumberCollection : public IterableCollection {
  private:
    vector<int> numbers;

  public:
    void addNumber(int num) {
      numbers.push_back(num);
    }

    shared_ptr<Iterator> createIterator() override {
      return make_shared<NumberIterator>(numbers);
    }
};

int main() {
  NumberCollection collection;
  collection.addNumber(10);
  collection.addNumber(20);
  collection.addNumber(30);

  shared_ptr<Iterator> it = collection.createIterator();

  while(it->hasNext()) {
    cout << it->next() << " ";
  }

  return 0;
}
