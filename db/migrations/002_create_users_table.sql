-- Migration: Create users table
-- Created: 2025-11-22
-- Description: Create users table for OAuth authentication

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    github_id TEXT UNIQUE NOT NULL,
    email TEXT NOT NULL,
    name TEXT NOT NULL,
    avatar_url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_github_id ON users(github_id);
CREATE INDEX idx_users_email ON users(email);
