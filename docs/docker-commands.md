# Docker Commands for Cube Orchestrator

This document contains useful Docker commands and examples for working with the cube-orchestrator project.

## Running PostgreSQL Container

To run a PostgreSQL container for development and testing:

**Option 1: Run in background (recommended for development)**

```bash
docker run -d -p 5432:5432 --name cube-orchestrator-db -e POSTGRES_USER=cube -e POSTGRES_PASSWORD=secret postgres
```

**Option 2: Run in foreground (see logs directly)**

```bash
docker run -p 5432:5432 --name cube-orchestrator-db -e POSTGRES_USER=cube -e POSTGRES_PASSWORD=secret postgres
```

**Option 3: Access container bash shell**

```bash
# First run the container in background
docker run -d -p 5432:5432 --name cube-orchestrator-db -e POSTGRES_USER=cube -e POSTGRES_PASSWORD=secret postgres

# Then access the bash shell
docker exec -it cube-orchestrator-db bash
```

**Recommended approach explanation:**

- `-d` runs the container in detached mode (background)
- `-p 5432:5432` maps port 5432 from container to host
- `--name cube-orchestrator-db` names the container
- `-e POSTGRES_USER=cube` sets database user to `cube`
- `-e POSTGRES_PASSWORD=secret` sets database password to `secret`

## Connecting to PostgreSQL

**Option 1: Connect from host machine (requires psql client installed)**

```bash
psql -h localhost -p 5432 -U cube
```

**Option 2: Connect from within the container**

```bash
# Access the container shell
docker exec -it cube-orchestrator-db bash

# Then connect to PostgreSQL (inside container)
psql -U cube
```

**Option 3: Direct psql connection via docker exec**

```bash
docker exec -it cube-orchestrator-db psql -U cube
```

## Creating Tables in PostgreSQL

Example of creating a table in PostgreSQL:

```sql
-- Connect to the database
$ psql -h localhost -p 5432 -U cube
Password for user cube: secret

-- Check existing tables
cube=# \d
No relations found.

-- Create a sample table
cube=# CREATE TABLE book (
    isbn char(13) PRIMARY KEY,
    title varchar(240) NOT NULL,
    author varchar(140)
);
CREATE TABLE

-- Verify table creation
cube=# \d
      List of relations
Schema | Name | Type  | Owner
--------+------+-------+-------
public | book | table | cube
(1 row)
```

## Additional Docker Commands

### Container Management

```bash
# List running containers
docker ps

# List all containers (including stopped)
docker ps -a

# Inspect container details (configuration, networks, volumes, etc.)
docker inspect cube-orchestrator-db

# Stop the PostgreSQL container
docker stop cube-orchestrator-db

# Start the PostgreSQL container
docker start cube-orchestrator-db

# Remove the container
docker rm cube-orchestrator-db

# Remove the container and its volumes
docker rm -v cube-orchestrator-db
```

### Image Management

```bash
# List Docker images
docker images

# Pull PostgreSQL image
docker pull postgres

# Remove PostgreSQL image
docker rmi postgres
```

### Docker API Commands

Docker provides a REST API accessible via Unix socket, useful for programmatic container management in orchestrators:

```bash
# Get container information via Docker API (replace container ID)
curl --unix-socket \
   /var/run/docker.sock http://docker/containers/b79998e6bd40/json | jq .

# Get container information using container name
curl --unix-socket \
   /var/run/docker.sock http://docker/containers/cube-orchestrator-db/json | jq .

# List all containers via API
curl --unix-socket \
   /var/run/docker.sock http://docker/containers/json | jq .

# Get container stats (CPU, memory usage)
curl --unix-socket \
   /var/run/docker.sock http://docker/containers/cube-orchestrator-db/stats?stream=false | jq .

# Get system information
curl --unix-socket \
   /var/run/docker.sock http://docker/info | jq .
```

**Note:** These API commands require `jq` for JSON formatting. Install with:

```bash
# Ubuntu/Debian
sudo apt-get install jq

# macOS
brew install jq
```

### Docker Compose (Future Enhancement)

For easier management, consider creating a `docker-compose.yml` file:

```yaml
version: '3.8'
services:
  postgres:
    image: postgres:latest
    container_name: cube-orchestrator-db
    environment:
      POSTGRES_USER: cube
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: cube_orchestrator
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

Then run with:

```bash
docker-compose up -d
```
