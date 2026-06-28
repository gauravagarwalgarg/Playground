"""
LeetCode #721 - Accounts Merge
Topic: Union-Find
Difficulty: Medium

Given a list of accounts where each account has a name and emails, merge
accounts that share at least one common email. Return merged accounts with
emails sorted.

Approach: Union-Find on email ownership. Map each email to an index, union
emails that belong to the same account, then group by root.

Time Complexity: O(n * k * α(n)) where k = avg emails per account
Space Complexity: O(n * k)
"""

from collections import defaultdict


class UnionFind:
    def __init__(self, n: int):
        self.parent = list(range(n))
        self.rank = [0] * n

    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x: int, y: int) -> None:
        px, py = self.find(x), self.find(y)
        if px == py:
            return
        if self.rank[px] < self.rank[py]:
            px, py = py, px
        self.parent[py] = px
        if self.rank[px] == self.rank[py]:
            self.rank[px] += 1


def accounts_merge(accounts: list[list[str]]) -> list[list[str]]:
    email_to_id: dict[str, int] = {}
    email_to_name: dict[str, str] = {}
    idx = 0

    # Assign unique id to each email
    for account in accounts:
        name = account[0]
        for email in account[1:]:
            if email not in email_to_id:
                email_to_id[email] = idx
                idx += 1
            email_to_name[email] = name

    uf = UnionFind(idx)

    # Union all emails within same account
    for account in accounts:
        first_id = email_to_id[account[1]]
        for email in account[2:]:
            uf.union(first_id, email_to_id[email])

    # Group emails by root
    groups: dict[int, list[str]] = defaultdict(list)
    for email, eid in email_to_id.items():
        groups[uf.find(eid)].append(email)

    # Build result
    result = []
    for root, emails in groups.items():
        name = email_to_name[emails[0]]
        result.append([name] + sorted(emails))
    return result


if __name__ == "__main__":
    accounts = [
        ["John", "johnsmith@mail.com", "john_newyork@mail.com"],
        ["John", "johnsmith@mail.com", "john00@mail.com"],
        ["Mary", "mary@mail.com"],
        ["John", "johnnybravo@mail.com"],
    ]
    merged = accounts_merge(accounts)
    # Sort for comparison
    merged_sorted = sorted([sorted(a) for a in merged])
    expected = sorted([
        sorted(["John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"]),
        sorted(["John", "johnnybravo@mail.com"]),
        sorted(["Mary", "mary@mail.com"]),
    ])
    assert merged_sorted == expected

    # Single account
    assert accounts_merge([["Bob", "bob@mail.com"]]) == [["Bob", "bob@mail.com"]]

    print("All tests passed!")
