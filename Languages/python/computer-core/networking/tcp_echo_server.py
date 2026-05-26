"""
TCP Echo Server & Client
Demonstrates basic socket programming with a simple echo protocol.

The server echoes back whatever the client sends.
"""

import socket
import threading


def start_echo_server(host: str, port: int, ready_event: threading.Event):
    """Start a TCP echo server that handles one client then shuts down."""
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    server_socket.bind((host, port))
    server_socket.listen(1)
    server_socket.settimeout(5.0)
    ready_event.set()

    try:
        conn, addr = server_socket.accept()
        data = conn.recv(1024)
        if data:
            conn.sendall(data)  # echo back
        conn.close()
    finally:
        server_socket.close()


def echo_client(host: str, port: int, message: str) -> str:
    """Send a message to the echo server and return the response."""
    client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client_socket.settimeout(5.0)
    client_socket.connect((host, port))
    client_socket.sendall(message.encode())
    response = client_socket.recv(1024).decode()
    client_socket.close()
    return response


if __name__ == "__main__":
    HOST = "127.0.0.1"
    PORT = 9999

    ready = threading.Event()
    server_thread = threading.Thread(
        target=start_echo_server, args=(HOST, PORT, ready)
    )
    server_thread.start()
    ready.wait()  # wait for server to be ready

    # Test echo
    response = echo_client(HOST, PORT, "Hello, Server!")
    assert response == "Hello, Server!", f"Expected echo, got: {response}"

    server_thread.join(timeout=5)

    print("All tests passed!")
