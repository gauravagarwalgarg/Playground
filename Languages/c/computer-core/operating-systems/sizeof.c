#include <stdio.h>

#define my_sizeof(type) ((char *)(&type + 1) - (char *)(&type))

int main() {
    int x;
    double y;
    char z;

    printf("Size of int: %lu\n", my_sizeof(x));
    printf("Size of double: %lu\n", my_sizeof(y));
    printf("Size of char: %lu\n", my_sizeof(z));

    return 0;
}

