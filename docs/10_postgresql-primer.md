# PostgreSQL Primer

This document provides a comprehensive guide to PostgreSQL commands and concepts for the cube-orchestrator project.

## What is PostgreSQL?

PostgreSQL is a powerful, open-source object-relational database system with a strong reputation for reliability, feature robustness, and performance. It's often used in container orchestration projects for storing cluster state, task metadata, and configuration data.

## Basic Connection

```bash
# Connect to PostgreSQL database
psql -h localhost -p 5432 -U cube

# Connect with password prompt
psql -h localhost -p 5432 -U cube -W

# Connect to specific database
psql -h localhost -p 5432 -U cube -d cube_orchestrator
```

## PostgreSQL Meta-Commands (\d commands)

Meta-commands in PostgreSQL start with a backslash (`\`) and provide quick ways to explore and manage your database structure.

### Basic Describe Commands

```sql
-- List all relations (tables, views, sequences)
\d

-- Describe specific table structure
\d tablename

-- Example output:
cube=# \d
      List of relations
Schema | Name | Type  | Owner
--------+------+-------+-------
public | book | table | cube
(1 row)

cube=# \d book
                Table "public.book"
  Column |          Type          | Modifiers
---------+------------------------+-----------
 isbn    | character(13)          | not null
 title   | character varying(240) | not null
 author  | character varying(140) |
Indexes:
    "book_pkey" PRIMARY KEY, btree (isbn)
```

### Specialized Describe Commands

```sql
-- List only tables
\dt

-- List only views
\dv

-- List only sequences
\ds

-- List only indexes
\di

-- List functions
\df

-- List users/roles
\du

-- List databases
\l

-- Show current database and user
\conninfo
```

## Database and Schema Management

```sql
-- Create database
CREATE DATABASE cube_orchestrator;

-- Connect to database
\c cube_orchestrator

-- Create schema
CREATE SCHEMA orchestrator;

-- Set search path
SET search_path TO orchestrator, public;

-- Show current schema
SELECT current_schema();
```

## Table Operations

### Creating Tables

```sql
-- Basic table creation
CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    state VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table with constraints
CREATE TABLE workers (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    ip_address INET NOT NULL,
    cores INTEGER CHECK (cores > 0),
    memory_mb INTEGER CHECK (memory_mb > 0),
    disk_gb INTEGER CHECK (disk_gb > 0),
    status VARCHAR(20) DEFAULT 'idle',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table with foreign keys
CREATE TABLE task_assignments (
    id UUID PRIMARY KEY,
    task_id UUID REFERENCES tasks(id) ON DELETE CASCADE,
    worker_id UUID REFERENCES workers(id) ON DELETE CASCADE,
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(task_id, worker_id)
);
```

### Modifying Tables

```sql
-- Add column
ALTER TABLE tasks ADD COLUMN priority INTEGER DEFAULT 1;

-- Drop column
ALTER TABLE tasks DROP COLUMN priority;

-- Rename column
ALTER TABLE tasks RENAME COLUMN name TO task_name;

-- Change column type
ALTER TABLE tasks ALTER COLUMN state TYPE VARCHAR(100);

-- Add constraint
ALTER TABLE tasks ADD CONSTRAINT check_state 
    CHECK (state IN ('pending', 'running', 'completed', 'failed'));
```

## Data Operations

### Inserting Data

```sql
-- Insert single record
INSERT INTO workers (id, name, ip_address, cores, memory_mb, disk_gb)
VALUES (gen_random_uuid(), 'worker-1', '192.168.1.100', 4, 8192, 100);

-- Insert multiple records
INSERT INTO tasks (id, name, state) VALUES
(gen_random_uuid(), 'task-1', 'pending'),
(gen_random_uuid(), 'task-2', 'running'),
(gen_random_uuid(), 'task-3', 'completed');
```

### Querying Data

```sql
-- Basic select
SELECT * FROM workers;

-- Select with conditions
SELECT name, cores, memory_mb FROM workers WHERE cores >= 4;

-- Join tables
SELECT t.name as task_name, w.name as worker_name
FROM tasks t
JOIN task_assignments ta ON t.id = ta.task_id
JOIN workers w ON ta.worker_id = w.id;

-- Aggregate functions
SELECT state, COUNT(*) as task_count
FROM tasks
GROUP BY state
ORDER BY task_count DESC;
```

### Updating Data

```sql
-- Update single record
UPDATE tasks SET state = 'running' WHERE name = 'task-1';

-- Update with conditions
UPDATE workers SET status = 'busy'
WHERE id IN (SELECT worker_id FROM task_assignments);
```

### Deleting Data

```sql
-- Delete specific records
DELETE FROM tasks WHERE state = 'completed';

-- Delete with joins
DELETE FROM task_assignments
WHERE task_id IN (SELECT id FROM tasks WHERE state = 'failed');
```

## Indexes and Performance

```sql
-- Create index
CREATE INDEX idx_tasks_state ON tasks(state);

-- Create composite index
CREATE INDEX idx_workers_cores_memory ON workers(cores, memory_mb);

-- Create unique index
CREATE UNIQUE INDEX idx_workers_name ON workers(name);

-- Show indexes for a table
\d+ workers

-- Analyze query performance
EXPLAIN ANALYZE SELECT * FROM tasks WHERE state = 'pending';
```

## Useful PostgreSQL Functions for Orchestrators

```sql
-- Generate UUIDs
SELECT gen_random_uuid();

-- Current timestamp
SELECT now(), current_timestamp;

-- JSON operations (useful for storing task configurations)
CREATE TABLE task_configs (
    id UUID PRIMARY KEY,
    task_id UUID REFERENCES tasks(id),
    config JSONB
);

-- Query JSON data
SELECT config->>'image' as docker_image
FROM task_configs
WHERE config ? 'image';

-- Array operations (useful for storing worker capabilities)
CREATE TABLE worker_capabilities (
    worker_id UUID REFERENCES workers(id),
    capabilities TEXT[]
);

-- Query arrays
SELECT * FROM worker_capabilities
WHERE 'docker' = ANY(capabilities);
```

## Database Administration

```sql
-- Show database size
SELECT pg_size_pretty(pg_database_size('cube_orchestrator'));

-- Show table sizes
SELECT 
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) as size
FROM pg_tables
WHERE schemaname = 'public'
ORDER BY pg_total_relation_size(schemaname||'.'||tablename) DESC;

-- Show active connections
SELECT pid, usename, datname, client_addr, state
FROM pg_stat_activity
WHERE datname = 'cube_orchestrator';

-- Backup database
-- (Run from command line)
pg_dump -h localhost -U cube cube_orchestrator > backup.sql

-- Restore database
-- (Run from command line)
psql -h localhost -U cube cube_orchestrator < backup.sql
```

## Common PostgreSQL Data Types for Orchestrators

```sql
-- Text types
VARCHAR(n)      -- Variable length string with limit
TEXT            -- Variable length string without limit
CHAR(n)         -- Fixed length string

-- Numeric types
INTEGER         -- 4-byte integer
BIGINT          -- 8-byte integer
DECIMAL(p,s)    -- Exact decimal
REAL            -- 4-byte floating point

-- Date/Time types
TIMESTAMP       -- Date and time
TIMESTAMPTZ     -- Date and time with timezone
INTERVAL        -- Time interval

-- Other useful types
UUID            -- Universally unique identifier
JSONB           -- Binary JSON (efficient storage and indexing)
INET            -- IP address
BOOLEAN         -- True/false
ARRAY           -- Array of any data type
```

## Best Practices for Orchestrator Databases

1. **Use UUIDs for distributed systems** - Avoids ID conflicts across nodes
2. **Index frequently queried columns** - Especially state, timestamps, and foreign keys
3. **Use JSONB for flexible configurations** - Store task definitions, worker metadata
4. **Implement proper constraints** - Ensure data integrity
5. **Use transactions for atomic operations** - Critical for state management
6. **Regular backups** - Essential for production orchestrators
7. **Monitor performance** - Use EXPLAIN ANALYZE for slow queries
8. **Connection pooling** - Use tools like PgBouncer for high-concurrency workloads

## Exit Commands

```sql
-- Exit psql
\q

-- Clear screen
\! clear
```

This primer covers the essential PostgreSQL knowledge needed for building and managing data in your cube orchestrator project.
