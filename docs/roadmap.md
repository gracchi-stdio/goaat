# Goaat Development Roadmap

## Overview
Build a multi-tenant backend for editing Astro Starlight content repositories.

---

## Phase 1: Foundation 🔄
*Establish core infrastructure, authentication, and a clean minimal UI.*

### 1A: Core Infrastructure ✅

| Task | Status | Notes |
|------|--------|-------|
| Project structure (Echo + Templ + SQLC) | ✅ Done | `cmd/server/main.go` entry point |
| PostgreSQL + pgx/v5 connection | ✅ Done | Graceful degradation if unavailable |
| GitHub OAuth login (Goth) | ✅ Done | Basic auth, no `repo` scope yet |
| Typed session management | ✅ Done | `internal/auth/session.go` |
| Graceful shutdown | ✅ Done | Signal handling, pool cleanup |

### 1B: UI Foundation 🔄

| Task | Status | Owner | Notes |
|------|--------|-------|-------|
| 1B.1 Clean base layout | ⬜ Todo | You | Header, nav, footer structure |
| 1B.2 Navigation component | ⬜ Todo | You | Links: Home, Profile, Login/Logout |
| 1B.3 Home page (public) | ⬜ Todo | You | Simple landing for unauthenticated |
| 1B.4 Dashboard page (auth) | ⬜ Todo | Together | User's main view after login |
| 1B.5 User profile page | 🚧 In Progress | You | [Implementation Guide](tasks/1B.5-profile-page-implementation.md) |
| 1B.6 Flash messages | ⬜ Todo | Copilot | Success/error notifications |
| 1B.7 404 & error pages | ⬜ Todo | You | Custom error templates |

### Current Focus: Task 1B.1

---

## Learning Goals (Phase 1B)

Each task teaches a Go/Templ concept:

| Task | You'll Learn |
|------|--------------|
| 1B.1 Layout | Templ components, HTMX boost, View Transitions |
| 1B.2 Navigation | Conditional rendering (`if`), auth context |
| 1B.3 Home page | Creating a new page, route registration |
| 1B.4 Dashboard | Authenticated routes, middleware |
| 1B.5 Profile | Reading from DB, passing data to templates |
| 1B.6 Flash messages | Session flash values, HTMX swaps |
| 1B.7 Error pages | Echo error handler, custom templates |

### UI Stack
- **Templ**: Type-safe HTML components
- **HTMX**: AJAX navigation (`hx-boost`, `hx-swap`)
- **Shoelace**: Web components (buttons, cards, icons)
- **View Transitions API**: Smooth page animations

---

## Phase 2: Repository Management 🔄
*Allow users to register and manage repos.*

| Task | Status | Notes |
|------|--------|-------|
| 2.1 Store OAuth access token | ⬜ Todo | Add `access_token` to users table (encrypted) |
| 2.2 Request `repo` scope | ⬜ Todo | Update Goth config |
| 2.3 Create `repositories` table | ⬜ Todo | Migration + SQLC queries |
| 2.4 Create `editors` table | ⬜ Todo | Role-based access (owner/editor) |
| 2.5 Repo registration endpoint | ⬜ Todo | `POST /repos` - validate & clone |
| 2.6 Clone repo to filesystem | ⬜ Todo | Use `go-git` library |
| 2.7 List user's repos UI | ⬜ Todo | Dashboard page |

### Implementation Order
```
2.1 → 2.2 → 2.3 → 2.4 → 2.6 → 2.5 → 2.7
```

---

## Phase 3: Content Browsing
*Navigate and view markdown files.*

| Task | Status | Notes |
|------|--------|-------|
| 3.1 File tree endpoint | ⬜ Todo | List files in `content_path` |
| 3.2 File tree UI component | ⬜ Todo | Sidebar navigation |
| 3.3 Read markdown file | ⬜ Todo | Parse frontmatter + content |
| 3.4 Display markdown | ⬜ Todo | Viewer with syntax highlighting |

---

## Phase 4: Content Editing
*Edit markdown files with frontmatter.*

| Task | Status | Notes |
|------|--------|-------|
| 4.1 Markdown editor component | ⬜ Todo | Consider Monaco, CodeMirror, or simple textarea |
| 4.2 Frontmatter form | ⬜ Todo | Title, description, sidebar config |
| 4.3 Save file endpoint | ⬜ Todo | Write to local clone |
| 4.4 Optimistic locking | ⬜ Todo | Check file hash before save |
| 4.5 Validation | ⬜ Todo | Starlight frontmatter schema |

---

## Phase 5: Git Sync
*Push changes back to GitHub.*

| Task | Status | Notes |
|------|--------|-------|
| 5.1 Commit changes | ⬜ Todo | `go-git` commit with user info |
| 5.2 Push to GitHub | ⬜ Todo | HTTPS auth with stored token |
| 5.3 Pull/sync endpoint | ⬜ Todo | Refresh from remote |
| 5.4 Conflict detection | ⬜ Todo | Show diff if remote changed |

---

## Phase 6: Collaboration
*Invite editors and manage access.*

| Task | Status | Notes |
|------|--------|-------|
| 6.1 Invite editor endpoint | ⬜ Todo | By email or GitHub username |
| 6.2 Pending invitations | ⬜ Todo | Match on login |
| 6.3 Remove editor | ⬜ Todo | Owner action |
| 6.4 Activity log | ⬜ Todo | Who edited what, when |

---

## Phase 7: Polish & Production
*Prepare for deployment.*

| Task | Status | Notes |
|------|--------|-------|
| 7.1 Error handling & logging | ⬜ Todo | Structured logging |
| 7.2 Rate limiting | ⬜ Todo | Prevent abuse |
| 7.3 Token encryption | ⬜ Todo | AES-GCM for access tokens |
| 7.4 HTTPS & security headers | ⬜ Todo | Production config |
| 7.5 Backup strategy | ⬜ Todo | Database + cloned repos |
| 7.6 GitHub App migration | ⬜ Future | Better than OAuth tokens |

---

## Current Focus
**Phase 2: Repository Management**

Next task: `2.1 Store OAuth access token`

---

## Technical Decisions Log

| Date | Decision | Rationale |
|------|----------|-----------|
| 2025-11-25 | OAuth tokens for MVP (use interface for future GitHub App integration) | Simpler than GitHub App, upgrade later  |
| 2025-11-25 | Clone repos to disk | Full git history, offline edits, simpler than API |
| 2025-11-25 | Optimistic locking | Concurrency is rare, avoid complex locking |
| 2025-11-25 | HTTPS git auth | Works with OAuth tokens, no SSH key management |

---

## File Structure (Planned)

```
/data/repos/
  └── {repo_id}/
      ├── .git/
      └── src/content/docs/   ← Starlight content

internal/
  ├── auth/                   ← OAuth, sessions ✅
  ├── repository/             ← Git operations (planned)
  │   ├── service.go          ← Interface
  │   └── git.go              ← go-git implementation
  ├── content/                ← Markdown parsing (planned)
  │   ├── service.go
  │   └── markdown.go
  └── web/
      └── handlers/
          ├── repos.go        ← Repo CRUD (planned)
          └── content.go      ← File browsing/editing (planned)
```
