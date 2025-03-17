#include <stdio.h>

int main() {
  int l, w, area;
  printf("Enter length: ");
  scanf("%d", &l);
  printf("Enter width: ");
  scanf("%d", &w);
  area = l * w;
  printf("The area is %d\n", area);
  return 0;
}

