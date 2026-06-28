"""Minimal HTTP server: stdlib http.server version and raw socket version.

Run: python http_server.py [stdlib|raw]
"""

import sys
import socket
from http.server import HTTPServer, BaseHTTPRequestHandler

HOST = "127.0.0.1"
PORT = 8080

# --- Standard Library Version ---

class SimpleHandler(BaseHTTPRequestHandler):
    """Minimal HTTP handler using http.server."""

    def do_GET(self) -> None:
        self.send_response(200)
        self.send_header("Content-Type", "text/plain")
        self.end_headers()
        self.wfile.write(f"Hello from stdlib server! Path: {self.path}\n".encode())

    def do_POST(self) -> None:
        content_length = int(self.headers.get("Content-Length", 0))
        body = self.rfile.read(content_length).decode()
        self.send_response(200)
        self.send_header("Content-Type", "text/plain")
        self.end_headers()
        self.wfile.write(f"Received POST body: {body}\n".encode())

def run_stdlib_server() -> None:
    """Run HTTP server using http.server module."""
    server = HTTPServer((HOST, PORT), SimpleHandler)
    print(f"[stdlib] HTTP server on http://{HOST}:{PORT}")
    print("Press Ctrl+C to stop")
    server.serve_forever()

# --- Raw Socket Version ---

def run_raw_server() -> None:
    """Run HTTP server using raw sockets to understand the protocol."""
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as server:
        server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        server.bind((HOST, PORT))
        server.listen(5)
        print(f"[raw] HTTP server on http://{HOST}:{PORT}")
        print("Press Ctrl+C to stop")

        while True:
            conn, addr = server.accept()
            with conn:
                request = conn.recv(4096).decode()
                request_line = request.split("\r\n")[0]
                print(f"[raw] {addr}: {request_line}")

                body = f"Hello from raw socket server!\nRequest: {request_line}\n"
                response = (
                    "HTTP/1.1 200 OK\r\n"
                    "Content-Type: text/plain\r\n"
                    f"Content-Length: {len(body)}\r\n"
                    "Connection: close\r\n"
                    "\r\n"
                    f"{body}"
                )
                conn.sendall(response.encode())

if __name__ == "__main__":
    mode = sys.argv[1] if len(sys.argv) > 1 else "stdlib"

    if mode == "stdlib":
        run_stdlib_server()
    elif mode == "raw":
        run_raw_server()
    else:
        print("Usage: python http_server.py [stdlib|raw]")
