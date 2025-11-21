# Database Structure

## Overview
This directory contains all database-related SQL files organized by purpose.

## Directory Structure

### `migrations/`
Sequential migration files that track schema changes over time.
- **Naming**: `001_description.sql`, `002_description.sql`, etc.
- **Purpose**: Version control for database schema
- **Usage**: Apply in order for fresh database setup

### `schema/`
Current database schema organized by domain/table.
- **Files**: One file per table or related group
- **Purpose**: Source of truth for current schema
- **Usage**: Referenced by SQLC for code generation

### `queries/`
SQLC query definitions organized by domain.
- **Files**: Queries grouped by table/feature (e.g., `authors.sql`)
- **Purpose**: Type-safe Go code generation
- **Syntax**: Use SQLC annotations (`-- name: QueryName :one`)

## SQLC Code Generation

Generate type-safe Go database code:
```bash
cd db
sqlc generate
```

This creates Go code in `internal/db/` with:
- Type-safe query methods
- Database models (structs)
- Query interface

## Adding New Tables

1. Create migration: `db/migrations/00X_add_tablename.sql`
2. Update schema: `db/schema/tablename.sql`
3. Add queries: `db/queries/tablename.sql`
4. Run: `cd db && sqlc generate`
