# Docker Commands for Cube Orchestrator

This document contains useful Docker commands and examples for working with the cube-orchestrator project.

## Running PostgreSQL Container

To run a PostgreSQL container for development and testing:

```bash
docker run -it -p 5432:5432 --name cube-orchestrator-db -e POSTGRES_USER=cube -e POSTGRES_PASSWORD=secret postgres
```

This command:

- Runs PostgreSQL in interactive mode (`-it`)
- Maps port 5432 from container to host (`-p 5432:5432`)
- Names the container `cube-orchestrator-db` (`--name cube-orchestrator-db`)
- Sets database user to `cube` (`-e POSTGRES_USER=cube`)
- Sets database password to `secret` (`-e POSTGRES_PASSWORD=secret`)

## Connecting to PostgreSQL

To connect to the PostgreSQL database:

```bash
psql -h localhost -p 5432 -U cube
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
