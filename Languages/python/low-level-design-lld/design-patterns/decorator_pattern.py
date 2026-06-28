"""
Decorator Pattern

Attaches additional responsibilities to an object dynamically.
DataSource decorators add encryption and compression transparently.
"""

from abc import ABC, abstractmethod
import base64
import zlib


class DataSource(ABC):
    @abstractmethod
    def write(self, data: str) -> bytes:
        pass

    @abstractmethod
    def read(self, data: bytes) -> str:
        pass


class RawDataSource(DataSource):
    def write(self, data: str) -> bytes:
        return data.encode()

    def read(self, data: bytes) -> str:
        return data.decode()


class DataSourceDecorator(DataSource):
    def __init__(self, wrapped: DataSource):
        self._wrapped = wrapped


class EncryptionDecorator(DataSourceDecorator):
    """Simple base64 'encryption' for demonstration."""

    def write(self, data: str) -> bytes:
        inner = self._wrapped.write(data)
        return base64.b64encode(inner)

    def read(self, data: bytes) -> str:
        decoded = base64.b64decode(data)
        return self._wrapped.read(decoded)


class CompressionDecorator(DataSourceDecorator):
    def write(self, data: str) -> bytes:
        inner = self._wrapped.write(data)
        return zlib.compress(inner)

    def read(self, data: bytes) -> str:
        decompressed = zlib.decompress(data)
        return self._wrapped.read(decompressed)


if __name__ == "__main__":
    original = "Hello, Design Patterns!"

    # Raw
    raw = RawDataSource()
    written = raw.write(original)
    assert raw.read(written) == original

    # Encrypted
    encrypted = EncryptionDecorator(RawDataSource())
    written = encrypted.write(original)
    assert written != original.encode()  # data is transformed
    assert encrypted.read(written) == original

    # Compressed + Encrypted (stacked decorators)
    stacked = CompressionDecorator(EncryptionDecorator(RawDataSource()))
    written = stacked.write(original)
    assert stacked.read(written) == original
    assert written != original.encode()

    # Order matters: Encrypt then Compress
    alt_stack = EncryptionDecorator(CompressionDecorator(RawDataSource()))
    written = alt_stack.write(original)
    assert alt_stack.read(written) == original

    print("All tests passed!")
