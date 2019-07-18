#include <stdio.h>

int sum(int x, int y) {
	return x + y;
}

void main(void){
	printf("Hello world!\n");
	int x = 1, y = 3;
	printf("sum %d + %d = %d\n", x, y, sum(x, y));
}
