/**
 * TCP Echo Server
 * Demonstrates basic socket programming in C.
 * Server echoes back whatever the client sends.
 *
 * Uses fork() to test server and client in the same process.
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <sys/wait.h>

#define PORT 9876
#define BUFFER_SIZE 256

void run_server(void) {
    int server_fd = socket(AF_INET, SOCK_STREAM, 0);
    assert(server_fd >= 0);

    int opt = 1;
    setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));

    struct sockaddr_in addr = {
        .sin_family = AF_INET,
        .sin_addr.s_addr = INADDR_ANY,
        .sin_port = htons(PORT)
    };

    assert(bind(server_fd, (struct sockaddr*)&addr, sizeof(addr)) == 0);
    assert(listen(server_fd, 1) == 0);

    // Accept one connection
    int client_fd = accept(server_fd, NULL, NULL);
    assert(client_fd >= 0);

    char buffer[BUFFER_SIZE];
    ssize_t bytes = read(client_fd, buffer, BUFFER_SIZE);
    if (bytes > 0) {
        write(client_fd, buffer, bytes);  // echo back
    }

    close(client_fd);
    close(server_fd);
}

void run_client(const char* message) {
    // Small delay to let server start
    usleep(100000);  // 100ms

    int sock = socket(AF_INET, SOCK_STREAM, 0);
    assert(sock >= 0);

    struct sockaddr_in addr = {
        .sin_family = AF_INET,
        .sin_port = htons(PORT)
    };
    inet_pton(AF_INET, "127.0.0.1", &addr.sin_addr);

    assert(connect(sock, (struct sockaddr*)&addr, sizeof(addr)) == 0);

    write(sock, message, strlen(message));

    char buffer[BUFFER_SIZE] = {0};
    ssize_t bytes = read(sock, buffer, BUFFER_SIZE);
    assert(bytes > 0);
    assert(strcmp(buffer, message) == 0);

    close(sock);
}

int main(void) {
    pid_t pid = fork();

    if (pid == 0) {
        // Child: client
        run_client("Hello, Echo Server!");
        exit(0);
    } else {
        // Parent: server
        run_server();
        int status;
        waitpid(pid, &status, 0);
        assert(WIFEXITED(status) && WEXITSTATUS(status) == 0);
    }

    printf("All tests passed!\n");
    return 0;
}
