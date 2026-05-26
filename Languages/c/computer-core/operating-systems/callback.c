#include <stdio.h>

typedef int (*operation_callback)(int, int);

int add(int a, int b) {
	return a + b;
}

int multiply(int a, int b) {
	return a * b;
}

void operate(int x, int y, operation_callback callback) {
	printf("Result : %d \n" , callback(x, y);
}

int main() {
	printf("Using add callback:\n");
	operate(5, 3, add)
}

