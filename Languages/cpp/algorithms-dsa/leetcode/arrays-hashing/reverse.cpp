void reverse(Node* head)
{
	Node* next, prev = null;

	Node* curr = head;

	while(curr != null)
	{
		next = curr->next;
		curr->next = prev;
		prev = curr;
		curr = next;
	}:q

