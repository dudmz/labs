#include <stdio.h>
#include <stdlib.h>
#include <strings.h>

typedef unsigned char *byte_pointer;

void show_bytes(byte_pointer start, size_t len)
{
	int i = 0;

	for (; i < len; i++)
		printf("%.2x", start[i]);

	printf("\n");
}

void show_int(int x)
{
	show_bytes((byte_pointer) &x, sizeof(int));
}

void show_float(float x)
{
	show_bytes((byte_pointer) &x, sizeof(float));
}

void show_pointer(void *x)
{
	show_bytes((byte_pointer) &x, sizeof(void *));
}

int main(int argc, char **argv)
{
	char *str = malloc(sizeof(char)*100);
	if (str == NULL)
		return 1;

	strcpy(str, "Hello, world!\n");

	show_int(32);
	show_float(2.3);
	show_pointer(str);

	free(str);

	return 0;
}

