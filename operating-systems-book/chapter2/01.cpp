#include <chrono>
#include <csignal>
#include <cstdlib>
#include <cstring>
#include <iostream>
#include <thread>

void print_usage();
void handler(int signal);
void spin(int ms);

static bool RUNNING = true;

int main(int argc, char **argv)
{
	// set signal handling
	struct sigaction sig_handler;
	sig_handler.sa_handler = handler;
	sig_handler.sa_flags = 0;
	sigemptyset(&sig_handler.sa_mask);

	sigaction(SIGINT, &sig_handler, NULL);

	if (argc < 2) {
		print_usage();
		exit(1);
	}

	while (RUNNING) {
		spin(1000);
		std::cout << argv[1] << "\n";
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

void handler(int signal)
{
	RUNNING = false;
}
