#include <chrono>
#include <cstring>
#include <iostream>
#include <thread>

void print_usage();
void spin(int ms);

int main(int argc, char **argv)
{
	if (argc < 2) {
		print_usage();
		exit(1);
	}

	char *str = (char*) std::malloc(20 * sizeof(char));
	std::strcpy(str, argv[1]);

	while (1) {
		spin(1000);
		std::cout << str << "\n";
	}

	return 0;
}

void print_usage()
{
	std::cerr << "usage: ./cpu <string>\n";
}

void spin(int ms)
{
	std::this_thread::sleep_for(std::chrono::milliseconds(ms));
}
