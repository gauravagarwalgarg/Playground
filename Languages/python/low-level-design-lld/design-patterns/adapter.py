"""
Adapter Pattern

Converts the interface of a class into another interface clients expect.
Adapts an XML data source to the JSON interface the client code requires.
"""

import json
from dataclasses import dataclass


@dataclass
class XMLData:
    """Legacy system returns data as XML strings."""

    def fetch(self) -> str:
        return "<user><name>Alice</name><age>30</age><city>NYC</city></user>"

    def parse_xml(self, xml: str) -> dict[str, str]:
        """Minimal XML parser for demo (handles flat key-value elements)."""
        import re
        result = {}
        # Match leaf elements (those whose content has no child tags)
        for match in re.finditer(r"<(\w+)>([^<]+)</\1>", xml):
            result[match.group(1)] = match.group(2)
        return result


class JSONClient:
    """Client expects data as JSON-compatible dict."""

    def process(self, data: dict) -> str:
        return json.dumps(data, sort_keys=True)


class XMLToJSONAdapter:
    """Adapts XMLData interface to dict output expected by JSONClient."""

    def __init__(self, xml_source: XMLData):
        self._source = xml_source

    def fetch_as_dict(self) -> dict[str, str]:
        xml_string = self._source.fetch()
        return self._source.parse_xml(xml_string)


if __name__ == "__main__":
    # Legacy XML source
    xml_source = XMLData()
    raw = xml_source.fetch()
    assert "<user>" in raw

    # Adapter converts to dict
    adapter = XMLToJSONAdapter(xml_source)
    data = adapter.fetch_as_dict()
    assert data == {"name": "Alice", "age": "30", "city": "NYC"}

    # Client consumes adapted data
    client = JSONClient()
    result = client.process(data)
    parsed = json.loads(result)
    assert parsed["name"] == "Alice"
    assert parsed["age"] == "30"

    print("All tests passed!")
