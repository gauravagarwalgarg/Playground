"""Simple TCP echo server + client using the socket module.

Demonstrates: bind, listen, accept, send, recv.
Run with: python tcp_server_client.py [server|client]
"""

import socket
import sys
import threading

HOST = "127.0.0.1"
PORT = 9999

def run_server() -> None:
    """TCP echo server: accepts connections and echoes messages back."""
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as server:
        server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        server.bind((HOST, PORT))
        server.listen(5)
        print(f"[Server] Listening on {HOST}:{PORT}")

        while True:
            conn, addr = server.accept()
            print(f"[Server] Connection from {addr}")
            threading.Thread(target=handle_client, args=(conn, addr), daemon=True).start()

def handle_client(conn: socket.socket, addr: tuple) -> None:
    """Handle a single client connection."""
    with conn:
        while True:
            data = conn.recv(1024)
            if not data:
                print(f"[Server] {addr} disconnected")
                break
            message = data.decode()
            print(f"[Server] Received from {addr}: {message}")
            conn.sendall(f"Echo: {message}".encode())

def run_client() -> None:
    """TCP client: sends messages and prints server responses."""
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client:
        client.connect((HOST, PORT))
        print(f"[Client] Connected to {HOST}:{PORT}")

        messages = ["Hello, Server!", "How are you?", "Goodbye!"]
        for msg in messages:
            client.sendall(msg.encode())
            response = client.recv(1024).decode()
            print(f"[Client] Sent: {msg} -> Got: {response}")

    print("[Client] Disconnected")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python tcp_server_client.py [server|client]")
        print("\nRunning demo (server + client in threads)...")
        # Demo mode: run both in threads
        server_thread = threading.Thread(target=run_server, daemon=True)
        server_thread.start()
        import time; time.sleep(0.5)
        run_client()
    elif sys.argv[1] == "server":
        run_server()
    elif sys.argv[1] == "client":
        run_client()
    else:
        print("Usage: python tcp_server_client.py [server|client]")
