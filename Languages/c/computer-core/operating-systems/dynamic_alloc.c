#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {

	int rows = 3;
	int col = 20;

	char **star = (char**)malloc(rows * sizeof(char*));

	for(int i = 0; i < rows; i++) {
		star[i] = (char*) malloc(col * sizeof(char));

	}
