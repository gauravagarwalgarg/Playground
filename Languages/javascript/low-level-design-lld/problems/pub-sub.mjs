/**
 * In-Memory Pub/Sub System
 *
 * A publish-subscribe messaging system where:
 * - Publishers send messages to topics (not directly to subscribers)
 * - Subscribers express interest in topics and receive relevant messages
 * - The broker decouples publishers from subscribers
 *
 * Components:
 * - MessageBroker: Central coordinator managing topics and routing
 * - Topic: Manages a set of subscribers for a specific channel
 * - Subscriber: Receives messages from subscribed topics
 *
 * Run: node --test Languages/javascript/low-level-design-lld/problems/pub-sub.mjs
 */

import { describe, it, beforeEach } from "node:test";
import assert from "node:assert/strict";

// ──────────────────────────────────────────────────────────────
// Subscriber
// ──────────────────────────────────────────────────────────────

class Subscriber {
  #id;
  #messages;

  /**
   * @param {string} id Unique identifier for this subscriber
   */
  constructor(id) {
    this.#id = id;
    this.#messages = [];
  }

  get id() {
    return this.#id;
  }

  /**
   * Called by the topic when a new message is published.
   * @param {string} topicName - The topic the message came from
   * @param {*} message - The message payload
   */
  receive(topicName, message) {
    this.#messages.push({ topic: topicName, message, timestamp: Date.now() });
  }

  /**
   * Returns all received messages.
   */
  getMessages() {
    return [...this.#messages];
  }

  /**
   * Returns messages filtered by topic.
   */
  getMessagesByTopic(topicName) {
    return this.#messages
      .filter((m) => m.topic === topicName)
      .map((m) => m.message);
  }

  /**
   * Clears all received messages.
   */
  clear() {
    this.#messages = [];
  }
}

// ──────────────────────────────────────────────────────────────
// Topic
// ──────────────────────────────────────────────────────────────

class Topic {
  #name;
  #subscribers;

  /**
   * @param {string} name Unique topic name
   */
  constructor(name) {
    this.#name = name;
    this.#subscribers = new Map(); // id -> subscriber
  }

  get name() {
    return this.#name;
  }

  get subscriberCount() {
    return this.#subscribers.size;
  }

  /**
   * Adds a subscriber to this topic.
   * @param {Subscriber} subscriber
   */
  addSubscriber(subscriber) {
    this.#subscribers.set(subscriber.id, subscriber);
  }

  /**
   * Removes a subscriber from this topic.
   * @param {string} subscriberId
   * @returns {boolean} true if the subscriber was found and removed
   */
  removeSubscriber(subscriberId) {
    return this.#subscribers.delete(subscriberId);
  }

  /**
   * @param {string} subscriberId
   * @returns {boolean}
   */
  hasSubscriber(subscriberId) {
    return this.#subscribers.has(subscriberId);
  }

  /**
   * Publishes a message to all subscribers of this topic.
   * @param {*} message
   */
  publish(message) {
    for (const subscriber of this.#subscribers.values()) {
      subscriber.receive(this.#name, message);
    }
  }
}

// ──────────────────────────────────────────────────────────────
// MessageBroker
// ──────────────────────────────────────────────────────────────

class MessageBroker {
  #topics;

  constructor() {
    this.#topics = new Map(); // topicName -> Topic
  }

  /**
   * Creates a new topic. No-op if the topic already exists.
   * @param {string} topicName
   * @returns {Topic}
   */
  createTopic(topicName) {
    if (!this.#topics.has(topicName)) {
      this.#topics.set(topicName, new Topic(topicName));
    }
    return this.#topics.get(topicName);
  }

  /**
   * Subscribes a subscriber to a topic.
   * Creates the topic if it doesn't exist.
   * @param {string} topicName
   * @param {Subscriber} subscriber
   */
  subscribe(topicName, subscriber) {
    const topic = this.createTopic(topicName);
    topic.addSubscriber(subscriber);
  }

  /**
   * Unsubscribes a subscriber from a topic.
   * @param {string} topicName
   * @param {string} subscriberId
   * @returns {boolean} true if successfully unsubscribed
   */
  unsubscribe(topicName, subscriberId) {
    const topic = this.#topics.get(topicName);
    if (!topic) return false;
    return topic.removeSubscriber(subscriberId);
  }

  /**
   * Publishes a message to all subscribers of a topic.
   * @param {string} topicName
   * @param {*} message
   * @throws {Error} if the topic doesn't exist
   */
  publish(topicName, message) {
    const topic = this.#topics.get(topicName);
    if (!topic) {
      throw new Error(`Topic "${topicName}" does not exist`);
    }
    topic.publish(message);
  }

  /**
   * Returns the number of subscribers for a topic.
   */
  getSubscriberCount(topicName) {
    const topic = this.#topics.get(topicName);
    return topic ? topic.subscriberCount : 0;
  }

  /**
   * Checks if a topic exists.
   */
  hasTopic(topicName) {
    return this.#topics.has(topicName);
  }

  /**
   * Lists all topic names.
   */
  listTopics() {
    return [...this.#topics.keys()];
  }
}

// ──────────────────────────────────────────────────────────────
// Tests
// ──────────────────────────────────────────────────────────────

describe("Pub/Sub System", () => {
  let broker;
  let sub1, sub2, sub3;

  beforeEach(() => {
    broker = new MessageBroker();
    sub1 = new Subscriber("user-1");
    sub2 = new Subscriber("user-2");
    sub3 = new Subscriber("user-3");
  });

  describe("Topic Management", () => {
    it("should create topics", () => {
      broker.createTopic("news");
      broker.createTopic("sports");
      assert.ok(broker.hasTopic("news"));
      assert.ok(broker.hasTopic("sports"));
      assert.deepEqual(broker.listTopics().sort(), ["news", "sports"]);
    });

    it("should not duplicate topics", () => {
      broker.createTopic("news");
      broker.createTopic("news");
      assert.equal(broker.listTopics().length, 1);
    });

    it("should throw when publishing to non-existent topic", () => {
      assert.throws(
        () => broker.publish("ghost-topic", "hello"),
        /does not exist/
      );
    });
  });

  describe("Message Delivery", () => {
    it("should deliver messages to all subscribers", () => {
      broker.subscribe("orders", sub1);
      broker.subscribe("orders", sub2);

      broker.publish("orders", { id: 1, item: "laptop" });

      assert.deepEqual(sub1.getMessagesByTopic("orders"), [
        { id: 1, item: "laptop" },
      ]);
      assert.deepEqual(sub2.getMessagesByTopic("orders"), [
        { id: 1, item: "laptop" },
      ]);
    });

    it("should deliver multiple messages in order", () => {
      broker.subscribe("events", sub1);

      broker.publish("events", "first");
      broker.publish("events", "second");
      broker.publish("events", "third");

      assert.deepEqual(sub1.getMessagesByTopic("events"), [
        "first",
        "second",
        "third",
      ]);
    });

    it("should not deliver messages to non-subscribers", () => {
      broker.subscribe("vip", sub1);
      broker.createTopic("vip"); // sub2 not subscribed

      broker.publish("vip", "exclusive content");

      assert.equal(sub1.getMessagesByTopic("vip").length, 1);
      assert.equal(sub2.getMessagesByTopic("vip").length, 0);
    });
  });

  describe("Unsubscribe", () => {
    it("should stop message delivery after unsubscribe", () => {
      broker.subscribe("alerts", sub1);
      broker.subscribe("alerts", sub2);

      broker.publish("alerts", "msg-1");
      broker.unsubscribe("alerts", sub1.id);
      broker.publish("alerts", "msg-2");

      // sub1 got only msg-1
      assert.deepEqual(sub1.getMessagesByTopic("alerts"), ["msg-1"]);
      // sub2 got both
      assert.deepEqual(sub2.getMessagesByTopic("alerts"), ["msg-1", "msg-2"]);
    });

    it("should return false when unsubscribing from non-existent topic", () => {
      assert.equal(broker.unsubscribe("nope", "user-1"), false);
    });

    it("should update subscriber count on unsubscribe", () => {
      broker.subscribe("chat", sub1);
      broker.subscribe("chat", sub2);
      assert.equal(broker.getSubscriberCount("chat"), 2);

      broker.unsubscribe("chat", sub1.id);
      assert.equal(broker.getSubscriberCount("chat"), 1);
    });
  });

  describe("Multiple Topics", () => {
    it("should isolate messages between topics", () => {
      broker.subscribe("sports", sub1);
      broker.subscribe("tech", sub1);
      broker.subscribe("sports", sub2);

      broker.publish("sports", "goal!");
      broker.publish("tech", "new release");

      // sub1 subscribed to both
      assert.deepEqual(sub1.getMessagesByTopic("sports"), ["goal!"]);
      assert.deepEqual(sub1.getMessagesByTopic("tech"), ["new release"]);
      // sub2 only subscribed to sports
      assert.deepEqual(sub2.getMessagesByTopic("sports"), ["goal!"]);
      assert.deepEqual(sub2.getMessagesByTopic("tech"), []);
    });

    it("should allow subscribing to multiple topics independently", () => {
      broker.subscribe("A", sub1);
      broker.subscribe("B", sub1);
      broker.subscribe("C", sub1);

      broker.publish("A", 1);
      broker.publish("B", 2);
      broker.publish("C", 3);

      const allMessages = sub1.getMessages();
      assert.equal(allMessages.length, 3);
      assert.deepEqual(
        allMessages.map((m) => m.message),
        [1, 2, 3]
      );
    });
  });

  describe("Subscriber", () => {
    it("should track messages with metadata", () => {
      broker.subscribe("logs", sub1);
      broker.publish("logs", "error occurred");

      const messages = sub1.getMessages();
      assert.equal(messages.length, 1);
      assert.equal(messages[0].topic, "logs");
      assert.equal(messages[0].message, "error occurred");
      assert.ok(messages[0].timestamp > 0);
    });

    it("should clear messages", () => {
      broker.subscribe("data", sub1);
      broker.publish("data", "payload");
      assert.equal(sub1.getMessages().length, 1);

      sub1.clear();
      assert.equal(sub1.getMessages().length, 0);
    });
  });
});
