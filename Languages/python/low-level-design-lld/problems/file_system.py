"""
Low-Level Design: In-Memory File System (LeetCode #588)

Trie-like structure supporting mkdir, ls, addContentToFile, readContentFromFile.
Each node is either a directory (with children) or a file (with content).
"""


class FSNode:
    def __init__(self):
        self.children: dict[str, "FSNode"] = {}
        self.content: str = ""
        self.is_file: bool = False


class FileSystem:
    def __init__(self):
        self.root = FSNode()

    def _traverse(self, path: str) -> FSNode:
        """Navigate to node at path, creating directories as needed."""
        node = self.root
        if path == "/":
            return node
        for part in path.split("/")[1:]:
            if part not in node.children:
                node.children[part] = FSNode()
            node = node.children[part]
        return node

    def ls(self, path: str) -> list[str]:
        node = self._traverse(path)
        if node.is_file:
            return [path.split("/")[-1]]
        return sorted(node.children.keys())

    def mkdir(self, path: str) -> None:
        self._traverse(path)

    def add_content_to_file(self, path: str, content: str) -> None:
        node = self._traverse(path)
        node.is_file = True
        node.content += content

    def read_content_from_file(self, path: str) -> str:
        node = self._traverse(path)
        return node.content


if __name__ == "__main__":
    fs = FileSystem()

    # ls root (empty)
    assert fs.ls("/") == []

    # mkdir
    fs.mkdir("/a/b/c")
    assert fs.ls("/") == ["a"]
    assert fs.ls("/a") == ["b"]
    assert fs.ls("/a/b") == ["c"]

    # Add file
    fs.add_content_to_file("/a/b/c/data.txt", "hello ")
    fs.add_content_to_file("/a/b/c/data.txt", "world")
    assert fs.read_content_from_file("/a/b/c/data.txt") == "hello world"

    # ls on file returns filename
    assert fs.ls("/a/b/c/data.txt") == ["data.txt"]

    # ls directory with multiple items
    fs.mkdir("/a/b/d")
    fs.add_content_to_file("/a/b/readme.md", "# Hi")
    listing = fs.ls("/a/b")
    assert listing == ["c", "d", "readme.md"]

    assert fs.read_content_from_file("/a/b/readme.md") == "# Hi"

    print("All tests passed!")
