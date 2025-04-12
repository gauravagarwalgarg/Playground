#include <iostream>
#include <vector>
#include <map>
#include <string>
#include <algorithm>
#include <set>
#include <cassert>
using namespace std;

struct Node{
   Node* next;
   Node* prev;
   int value;
   int key;
   Node(Node* p, Node* n, int k, int val):prev(p),next(n),key(k),value(val){};
   Node(int k, int val):prev(NULL),next(NULL),key(k),value(val){};
};

class Cache{
   
   protected: 
   map<int,Node*> mp; //map the key to the node in the linked list
   int cp;  //capacity
   Node* tail; // double linked list tail pointer
   Node* head; // double linked list head pointer
   virtual void set(int, int) = 0; //set function
   virtual int get(int) = 0; //get function

};
class LRUCache : public Cache {
    public:
        LRUCache(int capacity) {
            cp = capacity;
            tail = nullptr;
            head = nullptr;
        }
        
        void moveToFront(Node* node) {
            if(node == head) {
                return;
            }
            
            if(node->prev) node->prev->next = node->next;
            if(node->next) node->next->prev = node->prev;
            
            if(node == tail) tail = node->prev;
            
            
            node->next = head;
            node->prev = nullptr;
            
            if(head) 
                head->prev = node;
            
            head = node;
            
            if(!tail) 
                tail = head;
        }
        
        int get(int key) {
            if(mp.find(key) == mp.end())
                return -1;
            
            Node* node = mp[key];
            moveToFront(node);
            return node->value;
        }
        
        void set(int key, int value) {
            if(mp.find(key) != mp.end()) {
                Node* node = mp[key];
                node->value = value;
                moveToFront(node);
            } else {
                if(mp.size() == cp) {
                    mp.erase(tail->key);
                    if(tail->prev) {
                        tail = tail->prev;
                        delete tail->next;
                        tail->next = nullptr;
                    } else {
                        delete tail;
                        tail = nullptr;
                        head = nullptr;
                    }
                }
                
                Node* newNode = new Node(key, value);
                newNode->next = head;
                if(head) head->prev = newNode;
                head = newNode;
                
                if(!tail) tail = head;
                
                mp[key] = newNode;
            }
        }
};
int main() {
   int n, capacity,i;
   cin >> n >> capacity;
   LRUCache l(capacity);
   for(i=0;i<n;i++) {
      string command;
      cin >> command;
      if(command == "get") {
         int key;
         cin >> key;
         cout << l.get(key) << endl;
      } 
      else if(command == "set") {
         int key, value;
         cin >> key >> value;
         l.set(key,value);
      }
   }
   return 0;
}

