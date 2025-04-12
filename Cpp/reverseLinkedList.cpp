#include <iostream>

using namespace std;

struct Node {
	int data;
	Node* next;
	Node(int val) : data(val), next(nullptr) {}
}

Node* reverse(Node* head) {
	Node* prev = nullptr;
	Node* next = nullptr;
	Node* curr = head;

	while(curr != nullptr)
	{
		next = curr->next;
		curr->next = prev;
		prev = curr;
		curr = next;
	}

	return prev;
}


void Reverse(Node* head)
{
	Node* prev = nullptr;
	Node* next = nullptr;
	Node* curr = head;

	while( curr != nullptr)
	{
		next = curr->next;
		curr->next = prev;
		prev = curr;
		curr = next;
	}

	return prev;
}

void printList(Node* head)
{
	while(head != nullptr) {
		cout << head->data << " ";
		head = head->next;
	}

	cout << endl;
}

int main() {
    Node* head = new Node(1);
    head->next = new Node(2);
    head->next->next = new Node(3);
    head->next->next->next = new Node(4);
    
    cout << "Original List: ";
    printList(head);

    head = reverse(head);

    cout << "Reversed List: ";
    printList(head);

    return 0;
}

