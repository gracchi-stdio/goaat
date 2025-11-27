# Datastar Guide for Goaat

This guide covers how we use [Datastar](https://data-star.dev/) in the Goaat project for reactive UI updates and soft navigation.

## What is Datastar?

Datastar is a lightweight hypermedia framework that enables reactive, real-time web applications using HTML attributes. Instead of writing JavaScript, you declare behavior directly in your HTML using `data-*` attributes.

**Key benefits:**
- No client-side JavaScript to write
- Server-driven UI updates via SSE (Server-Sent Events)
- Works perfectly with Go/Templ server-side rendering
- Soft navigation without full page reloads

---

## Core Concepts

### 1. Signals (`data-signals`)

Signals are reactive variables that Datastar tracks. When a signal changes, any element referencing it automatically updates.

```html
<!-- Initialize signals with JSON object -->
<div data-signals="{count: 0, name: 'World'}">
  <span data-text="$name"></span>
</div>
```

**In Goaat**, we initialize signals on the `app-container` to track navigation state:

```html
<div class="app-container" data-signals="{currentPath: window.location.pathname, currentTitle: 'Dashboard'}">
```

### 2. Event Handling (`data-on:event`)

Attach event listeners that execute expressions when triggered.

```html
<!-- Basic click handler -->
<button data-on:click="$count = $count + 1">
  Increment
</button>

<!-- With modifiers -->
<button data-on:click__prevent="$foo = 'bar'">
  Prevent default behavior
</button>
```

**Common modifiers:**
- `__prevent` - Calls `event.preventDefault()`
- `__stop` - Calls `event.stopPropagation()`
- `__debounce.500ms` - Debounce the handler
- `__throttle.500ms` - Throttle the handler

### 3. Attribute Binding (`data-attr:name`)

Dynamically set HTML attributes based on expressions.

```html
<!-- Set aria-current when path matches -->
<button data-attr:aria-current="$currentPath === '/dashboard' ? 'page' : null">
  Dashboard
</button>

<!-- Set multiple attributes -->
<button data-attr="{disabled: $loading, title: $tooltip}">
  Submit
</button>
```

### 4. Text Binding (`data-text`)

Bind element text content to a signal.

```html
<span data-text="$currentTitle"></span>
<span data-text="`Hello, ${$name}!`"></span>
```

### 5. Conditional Display (`data-show`)

Show or hide elements based on a condition.

```html
<div data-show="$currentTitle !== 'Dashboard'">
  <span data-text="$currentTitle"></span>
</div>
```

### 6. Backend Requests (`@get`, `@post`)

Fetch content from the server and merge it into the DOM.

```html
<!-- GET request -->
<button data-on:click="@get('/api/data')">
  Load Data
</button>

<!-- POST request -->
<button data-on:click="@post('/api/submit')">
  Submit
</button>
```

---

## Goaat Navigation Pattern

Here's how we implement soft navigation with active state tracking:

### Signal Initialization

```html
<!-- In layout.templ, on the app-container -->
<div class="app-container" 
     data-signals={ "{currentPath: window.location.pathname, currentTitle: '" + title + "'}" }>
```

This initializes:
- `currentPath` - The current URL path (from browser)
- `currentTitle` - The current page title (from server)

### Navigation Button

```html
<button type="button"
    data-attr:aria-current="$currentPath === '/admin/dashboard' ? 'page' : null"
    data-on:click__prevent="$currentPath = '/admin/dashboard'; $currentTitle = 'Dashboard'; history.pushState(null, '', '/admin/dashboard'); @get('/admin/dashboard')">
    <sl-icon name="grid"></sl-icon>
    <span>Dashboard</span>
</button>
```

**What happens on click:**
1. `$currentPath = '/admin/dashboard'` - Update the path signal (instant UI feedback)
2. `$currentTitle = 'Dashboard'` - Update the title signal (breadcrumb updates)
3. `history.pushState(...)` - Update browser URL (no reload)
4. `@get('/admin/dashboard')` - Fetch new content from server

### Dynamic Breadcrumb

```html
<sl-breadcrumb>
    <sl-breadcrumb-item>
        <span data-on:click="...">Dashboard</span>
    </sl-breadcrumb-item>
    <!-- Only show if not on Dashboard -->
    <sl-breadcrumb-item data-show="$currentTitle !== 'Dashboard'">
        <span data-text="$currentTitle"></span>
    </sl-breadcrumb-item>
</sl-breadcrumb>
```

### Active State CSS

```css
.sidebar-nav button[aria-current="page"] {
    background: var(--sl-color-primary-50);
    color: var(--sl-color-primary-700);
    font-weight: var(--sl-font-weight-semibold);
}
```

---

## Server Response Format

When Datastar makes a `@get` request, the server returns HTML fragments that get merged into the DOM. In Goaat, we return partial HTML that replaces `#page-content`:

```go
// Handler returns just the page content, not the full layout
func (h *Handler) GetDashboard(c echo.Context) error {
    // For Datastar requests, return partial
    if c.Request().Header.Get("Accept") == "text/html-partial" {
        return Render(c, pages.DashboardContent())
    }
    // For full page loads, return with layout
    return Render(c, pages.Dashboard())
}
```

---

## Common Patterns

### Toast Notifications

```html
<!-- Container that persists across navigations -->
<div id="alert-container">
    <!-- Toasts injected here -->
</div>

<!-- Show toast -->
<button data-on:click="$alerts.push({type: 'success', message: 'Saved!'})">
    Save
</button>
```

### Loading States

```html
<button 
    data-indicator:loading
    data-attr:disabled="$loading"
    data-on:click="@post('/api/save')">
    <span data-show="!$loading">Save</span>
    <span data-show="$loading">Saving...</span>
</button>
```

### Form Binding

```html
<input type="text" data-bind:username />
<span data-text="$username"></span>
```

---

## Debugging Tips

1. **Check signals in DevTools**: Open browser console and type `datastar.signals` to see current signal values

2. **Watch for typos**: Signal names are case-sensitive (`$currentPath` ≠ `$CurrentPath`)

3. **Use browser network tab**: Verify `@get` requests are being made and returning expected HTML

4. **Check attribute syntax**: Use colon syntax `data-on:click`, not `data-on-click`

---

## Resources

- [Datastar Documentation](https://data-star.dev/docs)
- [Datastar Examples](https://data-star.dev/examples)
- [Datastar GitHub](https://github.com/starfederation/datastar)

---

## Goaat-Specific Files

| File | Purpose |
|------|---------|
| `internal/web/templates/layouts/layout.templ` | Main layout with signal initialization |
| `assets/css/pages/authed.css` | Sidebar navigation styles including active state |
| `assets/js/main.js` | Datastar and Shoelace imports |
