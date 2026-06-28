#include <stdio.h>

void ModifyString(char* buff) {
  int i = 0, j = 0;

  for(i = 0; buff[i] != '\0'; i++) {
    if((i + 1) % 3 !=0) {
      buff[j++] = buff[i];
    }
  }
    buff[j] = '\0';
}


int main() {

  char str[] = "abcdefghi";

  ModifyString(str);

  printf("%s", str);

  return 0;
}
