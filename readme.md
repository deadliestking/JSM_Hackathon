# CDN Simulator

## Overview
A **Content Delivery Network (CDN)** is a system of distributed servers designed to deliver content (like images, videos, and web pages) to users with high availability and performance. This project simulates a CDN using Go, Docker, and load balancing. It demonstrates the caching and load distribution benefits of a CDN.

This project is ideal for showcasing your understanding of distributed systems, caching, and load balancing at a hackathon!

---

## Features
- **Two Edge Servers**: Simulated by `server1` and `server2`.
- **Load Balancer**: Distributes client requests to the least-loaded edge server.
- **Caching Mechanism**:
  - Servers cache content after the first request.
  - Subsequent requests for the same content are served directly from the cache for improved speed.
- **Origin Server Simulation**: If content isn’t in the cache, it’s fetched from the origin server and then cached.
- **Docker Integration**: Simplifies deployment and testing with Docker containers and volumes.

---

## How It Works

### Request Flow
1. A **client** sends a request for content (e.g., `/content1`).
2. The **load balancer** routes the request to the least-loaded edge server.
3. The edge server:
   - Checks if the content is in its **cache**.
   - If not found, fetches it from the **origin server** and caches it.
   - Serves the content to the client.
4. Future requests for the same content are served from the cache, improving speed.

### Example Output
- **First request for `/content2`**:

- **Second request for `/content2`**:


### Serving from cache on server 1: Content of /content2

---

## Benefits Simulated

- **Improved Performance**: Cached content is served faster than fetching from the origin server.
- **Reduced Load**: Distributes traffic between servers, reducing the load on any single server.
- **Scalability**: Demonstrates how real-world CDNs handle massive traffic efficiently.

---

## Getting Started

### Prerequisites
- Docker installed on your system.

### Build and Run

#### 1. Build the Docker Image
```bash
docker build -t cdn-simulator .
```

#### 2. Run the Docker Container
```bash
docker run -d -p 8080:8080 -p 8081:8081 -p 8082:8082 --name my-cdn-simulator cdn-simulator
```

#### 3. Access the Servers
```bash
Load Balancer: http://localhost:8080
Server 1: http://localhost:8081
Server 2: http://localhost:8082
```

### Persistent Caching with Docker Volumes

#### 1. Create a Docker Volume
```bash
docker volume create CDN_SIM
```

#### 2. Run the Container with the Volume
```bash
docker run -d -p 8080:8080 -p 8081:8081 -p 8082:8082 -v CDN_SIM:/root/server1_files --name my-cdn-simulator cdn-simulator
```


### Testing
#### Check Cache Contents
## Access the servers and check if content is cached:
 ```bash
 docker exec -it <container_name> /bin/sh
```

## Navigate to the cache directories:
```bash
cd /root/server1_files
ls
cd /root/server2_files
ls
```

### Simulating Requests
#### end requests via the load balancer:

## First request:
```bash
curl http://localhost:8080/content1
```
## Second request (cached):
```bash
curl http://localhost:8080/content1
```


### Ideas for Future Enhancements
## Geo-Based Load Balancing: Simulate routing requests to the nearest server.
## Metrics and Monitoring: Show real-time server load and cache usage.
## HTTP/2 Support: Improve content delivery efficiency.
## Dashboard: Visualize request patterns and cache hits/misses.
