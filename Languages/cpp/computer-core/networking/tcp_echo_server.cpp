/**
 * Networking: TCP Echo Server (POSIX sockets)
 *
 * Demonstrates: socket(), bind(), listen(), accept(), read(), write()
 * Compile: g++ -std=c++20 -o tcp_echo tcp_echo_server.cpp -pthread
 * Run server: ./tcp_echo
 * Test: echo "hello" | nc localhost 8080
 */
#include <iostream>
#include <cstring>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <thread>
using namespace std;

constexpr int PORT = 8080;
constexpr int BUFFER_SIZE = 1024;

void handleClient(int clientFd) {
    char buffer[BUFFER_SIZE];
    ssize_t bytesRead;
    while ((bytesRead = read(clientFd, buffer, BUFFER_SIZE - 1)) > 0) {
        buffer[bytesRead] = '\0';
        cout << "Received: " << buffer;
        write(clientFd, buffer, bytesRead); // echo back
    }
    close(clientFd);
    cout << "Client disconnected." << endl;
}

int main() {
    int serverFd = socket(AF_INET, SOCK_STREAM, 0);
    if (serverFd < 0) { perror("socket"); return 1; }

    int opt = 1;
    setsockopt(serverFd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));

    sockaddr_in addr{};
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = INADDR_ANY;
    addr.sin_port = htons(PORT);

    if (bind(serverFd, (sockaddr*)&addr, sizeof(addr)) < 0) {
        perror("bind"); return 1;
    }

    if (listen(serverFd, 5) < 0) {
        perror("listen"); return 1;
    }

    cout << "TCP Echo Server listening on port " << PORT << endl;
    cout << "Test with: echo 'hello' | nc localhost " << PORT << endl;
    cout << "Press Ctrl+C to stop." << endl;

    while (true) {
        sockaddr_in clientAddr{};
        socklen_t clientLen = sizeof(clientAddr);
        int clientFd = accept(serverFd, (sockaddr*)&clientAddr, &clientLen);
        if (clientFd < 0) { perror("accept"); continue; }

        cout << "Client connected." << endl;
        thread(handleClient, clientFd).detach();
    }

    close(serverFd);
    return 0;
}
