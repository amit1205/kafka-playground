# Overview
Apache Kafka is a distributed event streaming platform designed to handle high-throughput, low-latency data feeds in real time. It acts as a central nervous system for data, allowing different parts of a software system to communicate asynchronously and reliably by publishing and subscribing to streams of "events" or "messages."

It is more than just a traditional message queue; it combines three core capabilities:

1. Messaging: Publish and subscribe to streams of records.

2. Storage: Durably store streams of records in the order they were generated.

3. Stream Processing: Process streams of records in real time as they occur.

# Core Kafka Concepts

Kafka's architecture is built around a few key components:

- Events (Records): The basic unit of data. An event records the fact that "something happened," such as a user clicking a button, a sensor reading a temperature, or a service placing an order.


- Producers: Client applications that publish (write) events to Kafka topics. They are responsible for sending data streams to the Kafka cluster.


- Consumers: Client applications that subscribe to (read and process) events from Kafka topics. They read the data at their own pace.


- Topics: The way events are categorized and organized. A topic is a named, immutable, ordered sequence of records, similar to a labeled folder or log file.

- Partitions: Each topic is divided into one or more partitions. This is the main mechanism that allows Kafka to scale horizontally and ensures fault tolerance. Events within a single partition are guaranteed to be in order, but the order is not guaranteed across different partitions of the same topic.

- Brokers (Kafka Servers): The servers that run Kafka. A Kafka Cluster is typically made up of multiple brokers. Brokers receive events from producers, store them on disk, replicate them across the cluster for durability, and serve them to consumers.

# Key Architectural Advantages

Kafka's design provides significant benefits over traditional message queues:

|Feature            |   Description    | Benefit |
|-------------------|------------------|---------|
|Durability | Messages are persisted to disk for a configurable period (e.g., 7 days or more) and are not deleted immediately after being consumed.|Allows multiple independent applications to read the same data stream and provides data recovery capabilities.|
|Scalability | Topics are partitioned across multiple brokers, allowing both producers and consumers to read and write data in parallel across the entire cluster.| Handles millions of messages per second with high throughput.|
|Decoupling | Producers and consumers are fully decoupled. They do not know about each other, only about the topic in Kafka. Producers don't wait for consumers.| Systems can evolve independently and are resilient to the failure of other systems.|
|Ordering | Ordering is guaranteed within a single partition. By assigning related events (e.g., all events for a single user ID) to the same partition, you guarantee sequential processing.| Essential for applications like financial transactions or event sourcing.|

---

# Primary Use Cases

Kafka is used by thousands of companies (including Netflix, Uber, and LinkedIn) to handle mission-critical data flows:

- Real-Time Data Pipelines: Moving data between different systems (databases, analytics platforms, microservices) instantly and reliably.

- Microservices Communication: Serving as a central, asynchronous messaging bus to facilitate communication between decoupled services.

- Activity Tracking: Recording user activity (clicks, searches, views) on websites or apps for real-time monitoring and analytics.

- Log Aggregation: Collecting log files and operational metrics from all systems into a centralized feed for analysis and alerting.

- Event Sourcing: Storing a chronological, immutable log of all state changes to a system.