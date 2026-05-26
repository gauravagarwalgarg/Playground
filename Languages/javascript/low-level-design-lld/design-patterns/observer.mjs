/**
 * Observer Design Pattern
 * Defines a one-to-many dependency between objects.
 * When one object changes state, all dependents are notified.
 */

class EventManager {
  constructor() {
    this.listeners = new Map();
  }

  subscribe(event, listener) {
    if (!this.listeners.has(event)) {
      this.listeners.set(event, []);
    }
    this.listeners.get(event).push(listener);
  }

  unsubscribe(event, listener) {
    const eventListeners = this.listeners.get(event);
    if (eventListeners) {
      const idx = eventListeners.indexOf(listener);
      if (idx !== -1) eventListeners.splice(idx, 1);
    }
  }

  notify(event, data) {
    const eventListeners = this.listeners.get(event) || [];
    for (const listener of eventListeners) {
      listener.update(event, data);
    }
  }
}

class EmailNotifier {
  constructor() {
    this.messages = [];
  }
  update(event, data) {
    this.messages.push(`Email: [${event}] ${data}`);
  }
}

class LogNotifier {
  constructor() {
    this.logs = [];
  }
  update(event, data) {
    this.logs.push(`Log: [${event}] ${data}`);
  }
}

export { EventManager, EmailNotifier, LogNotifier };

// Tests
import { test } from "node:test";
import assert from "node:assert/strict";

test("Observer - notify subscribers", () => {
  const manager = new EventManager();
  const email = new EmailNotifier();
  const logger = new LogNotifier();

  manager.subscribe("order_placed", email);
  manager.subscribe("order_placed", logger);
  manager.subscribe("order_shipped", email);

  manager.notify("order_placed", "Order #123");
  manager.notify("order_shipped", "Order #123");

  assert.strictEqual(email.messages.length, 2);
  assert.strictEqual(logger.logs.length, 1);
  assert.strictEqual(email.messages[0], "Email: [order_placed] Order #123");
});

test("Observer - unsubscribe", () => {
  const manager = new EventManager();
  const email = new EmailNotifier();

  manager.subscribe("event", email);
  manager.notify("event", "data1");
  assert.strictEqual(email.messages.length, 1);

  manager.unsubscribe("event", email);
  manager.notify("event", "data2");
  assert.strictEqual(email.messages.length, 1); // no new message
});
