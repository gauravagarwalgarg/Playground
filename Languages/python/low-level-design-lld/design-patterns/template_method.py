"""
Template Method Pattern

Defines the skeleton of an algorithm in a base class, letting subclasses
override specific steps. DataParser with CSV/JSON/XML implementations.
"""

from abc import ABC, abstractmethod
import json


class DataParser(ABC):
    """Template: parse flow is fixed, steps are customizable."""

    def parse(self, raw: str) -> list[dict]:
        data = self.read_data(raw)
        data = self.process_data(data)
        data = self.hook_validate(data)
        return data

    @abstractmethod
    def read_data(self, raw: str) -> list[dict]:
        pass

    def process_data(self, data: list[dict]) -> list[dict]:
        """Default: strip whitespace from string values."""
        return [
            {k: v.strip() if isinstance(v, str) else v for k, v in row.items()}
            for row in data
        ]

    def hook_validate(self, data: list[dict]) -> list[dict]:
        """Hook: subclasses can override for custom validation."""
        return data


class CSVParser(DataParser):
    def read_data(self, raw: str) -> list[dict]:
        lines = raw.strip().split("\n")
        headers = [h.strip() for h in lines[0].split(",")]
        return [
            dict(zip(headers, [v.strip() for v in line.split(",")]))
            for line in lines[1:]
        ]


class JSONParser(DataParser):
    def read_data(self, raw: str) -> list[dict]:
        return json.loads(raw)

    def hook_validate(self, data: list[dict]) -> list[dict]:
        """JSON parser validates all records have same keys."""
        if data:
            keys = set(data[0].keys())
            return [row for row in data if set(row.keys()) == keys]
        return data


class XMLParser(DataParser):
    def read_data(self, raw: str) -> list[dict]:
        import re
        records = []
        for item in re.finditer(r"<item>(.*?)</item>", raw, re.DOTALL):
            record = {}
            for field in re.finditer(r"<(\w+)>(.*?)</\1>", item.group(1)):
                record[field.group(1)] = field.group(2)
            records.append(record)
        return records


if __name__ == "__main__":
    # CSV
    csv_data = "name, age\nAlice, 30\nBob, 25"
    csv_result = CSVParser().parse(csv_data)
    assert csv_result == [{"name": "Alice", "age": "30"}, {"name": "Bob", "age": "25"}]

    # JSON
    json_data = '[{"name": " Alice ", "age": 30}, {"name": "Bob", "age": 25}]'
    json_result = JSONParser().parse(json_data)
    assert json_result[0]["name"] == "Alice"

    # XML
    xml_data = "<data><item><name>Alice</name><age>30</age></item></data>"
    xml_result = XMLParser().parse(xml_data)
    assert xml_result == [{"name": "Alice", "age": "30"}]

    print("All tests passed!")
