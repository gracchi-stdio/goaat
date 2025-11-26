# Goaat Project Brainstorming

## Project Vision
A Go backend for editing **Astro Starlight** content repositories.
- **Target**: Multi-tenant (many repos, many users).
- **Hosting**: Vultr VPS (Self-hosted/PaaS).
- **Core Function**: Authenticate users, clone repos, edit markdown/frontmatter, sync with GitHub.

## Architecture
- **Stack**: Echo v4 + Templ + PostgreSQL (pgx/v5) + SQLC.
- **Runtime**: Podman 5 containers.
- **Auth**: GitHub OAuth (Goth).

## Data Model Draft

```mermaid
erDiagram
    USER ||--o{ REPOSITORY : owns
    USER ||--o{ EDITOR : "is assigned"
    REPOSITORY ||--o{ EDITOR : "has"

    USER {
        int64 id
        string github_id
        string email
        string name
        string avatar_url
        string oauth_token "Encrypted"
    }

    REPOSITORY {
        int64 id
        int64 owner_id FK
        string github_owner
        string github_repo
        string branch
        string content_path "e.g. src/content/docs"
        string clone_path "Local filesystem path"
        timestamp last_synced_at
    }

    EDITOR {
        int64 id
        int64 repo_id FK
        int64 user_id FK
        string role "editor, viewer"
    }
```

## Workflows

### 1. Repo Registration
1. **Owner** logs in.
2. Registers repo (Owner/Name).
3. System validates access & clones repo to server storage.

### 2. Collaboration
1. Owner invites collaborators (via email/GitHub username).
2. Collaborators log in and see assigned repos.

### 3. Editing & Sync
1. **Edit**: User modifies content in web UI.
2. **Save**: Writes to local file system in container.
3. **Concurrency**: Optimistic locking (check file hash/timestamp before save).
4. **Sync**:
   - **Pull**: Before editing session.
   - **Push**: After save/commit.

## Git Authentication Strategy (Push/Pull Access)

To push changes back to GitHub on behalf of different owners, we need credentials.

### Option A: OAuth Access Tokens (Simplest Start)
- Request `repo` scope during login.
- Store the OAuth Access Token (encrypted) in the DB.
- Use token for HTTPS git operations:
  `git push https://oauth2:<TOKEN>@github.com/owner/repo.git`
- **Pros**: Easy to implement with Goth.
- **Cons**: Broad permissions (access to all user's repos). Token expiry/refresh handling needed.

### Option B: GitHub App (Robust)
- Create a "Goaat" GitHub App.
- User installs App on specific repositories.
- App gets an **Installation Token**.
- **Pros**: Granular access (only installed repos), higher rate limits, no user password storage.
- **Cons**: More complex setup than standard OAuth.

### Option C: SSH Deploy Keys
- Generate SSH key per repo on server.
- User manually adds public key to GitHub Repo Settings.
- **Pros**: Standard Git auth.
- **Cons**: High friction for users (manual setup).

## Decision
**Selected for MVP: Option A (OAuth Access Tokens)**

1.  **Update Auth**: Configure Goth to request `repo` scope.
2.  **Database**: Add `access_token` (encrypted) to `users` table.
3.  **Git Operations**: Use HTTPS authentication with the stored token.
    ```bash
    git clone https://oauth2:TOKEN@github.com/user/repo.git
    ```
4.  **Future Migration**: Move to Option B (GitHub App) for production/SaaS to improve security and granularity.
