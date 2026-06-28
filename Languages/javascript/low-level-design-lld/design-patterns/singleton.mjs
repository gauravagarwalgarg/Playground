/**
 * Singleton Design Pattern
 * Uses closure and module pattern for a single instance.
 */

class Database {
  constructor() {
    if (Database._instance) {
      return Database._instance;
    }
    this.connection = "PostgreSQL@localhost:5432";
    Database._instance = this;
  }

  query(sql) {
    return `Executing: ${sql}`;
  }

  static getInstance() {
    if (!Database._instance) {
      new Database();
    }
    return Database._instance;
  }
}

export { Database };

// Tests
import { test } from "node:test";
import assert from "node:assert/strict";

test("Singleton - same instance", () => {
  const db1 = Database.getInstance();
  const db2 = Database.getInstance();
  assert.strictEqual(db1, db2);
});

test("Singleton - correct connection", () => {
  const db = Database.getInstance();
  assert.strictEqual(db.connection, "PostgreSQL@localhost:5432");
});

test("Singleton - query works", () => {
  const db = Database.getInstance();
  assert.strictEqual(db.query("SELECT 1"), "Executing: SELECT 1");
});

test("Singleton - new returns same instance", () => {
  const db1 = new Database();
  const db2 = new Database();
  assert.strictEqual(db1, db2);
});
