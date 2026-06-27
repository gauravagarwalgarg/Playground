"""
Builder Pattern

Separates construction of complex objects from their representation.
QueryBuilder uses method chaining for step-by-step SQL query construction.
"""


class QueryBuilder:
    def __init__(self):
        self._table: str = ""
        self._columns: list[str] = ["*"]
        self._conditions: list[str] = []
        self._order_by: str | None = None
        self._limit: int | None = None
        self._joins: list[str] = []

    def table(self, name: str) -> "QueryBuilder":
        self._table = name
        return self

    def select(self, *columns: str) -> "QueryBuilder":
        self._columns = list(columns)
        return self

    def where(self, condition: str) -> "QueryBuilder":
        self._conditions.append(condition)
        return self

    def join(self, table: str, on: str) -> "QueryBuilder":
        self._joins.append(f"JOIN {table} ON {on}")
        return self

    def order_by(self, column: str, desc: bool = False) -> "QueryBuilder":
        direction = "DESC" if desc else "ASC"
        self._order_by = f"{column} {direction}"
        return self

    def limit(self, count: int) -> "QueryBuilder":
        self._limit = count
        return self

    def build(self) -> str:
        if not self._table:
            raise ValueError("Table name is required")
        parts = [f"SELECT {', '.join(self._columns)} FROM {self._table}"]
        parts.extend(self._joins)
        if self._conditions:
            parts.append("WHERE " + " AND ".join(self._conditions))
        if self._order_by:
            parts.append(f"ORDER BY {self._order_by}")
        if self._limit is not None:
            parts.append(f"LIMIT {self._limit}")
        return " ".join(parts)


if __name__ == "__main__":
    # Simple query
    q1 = QueryBuilder().table("users").select("name", "email").build()
    assert q1 == "SELECT name, email FROM users"

    # Complex query with chaining
    q2 = (
        QueryBuilder()
        .table("orders")
        .select("id", "total")
        .join("users", "users.id = orders.user_id")
        .where("total > 100")
        .where("status = 'active'")
        .order_by("total", desc=True)
        .limit(10)
        .build()
    )
    assert "JOIN users ON users.id = orders.user_id" in q2
    assert "WHERE total > 100 AND status = 'active'" in q2
    assert "ORDER BY total DESC" in q2
    assert "LIMIT 10" in q2

    # Default select *
    q3 = QueryBuilder().table("products").build()
    assert q3 == "SELECT * FROM products"

    print("All tests passed!")
