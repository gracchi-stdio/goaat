# Shoelace Design Token Strategy

## Overview

Goaat uses [Shoelace](https://shoelace.style/) as its component library. Shoelace provides a comprehensive design token system that allows for consistent theming without writing custom CSS for individual components.

## Why Use Shoelace Tokens Instead of Custom Colors?

### ✅ Benefits

1. **Consistency**: All Shoelace components automatically use the same design tokens
2. **Dark Mode Support**: Tokens automatically adapt when switching themes
3. **Semantic Meaning**: Use `primary`, `success`, `danger`, `warning`, `neutral` instead of arbitrary colors
4. **Accessibility**: Shoelace ensures proper contrast ratios
5. **Maintainability**: Change your entire color scheme by updating 11 variables per color
6. **No Conflicts**: Shoelace tokens use `--sl-` prefix to avoid naming collisions

### ❌ Problems with Custom Colors

- Components won't match your custom colors
- Dark mode requires duplicate work
- Breaks Shoelace's built-in accessibility features
- Creates maintenance burden

## Shoelace Color System

### Color Primitives

Shoelace provides 20+ color palettes, each with 11 shades (50-950):

```
gray, red, orange, amber, yellow, lime, green, emerald,
teal, cyan, sky, blue, indigo, violet, purple, fuchsia, pink, rose
```

Each has values: `50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950`

### Theme Tokens (Semantic Colors)

These map primitives to semantic meanings:

| Token | Default Primitive | Usage |
|-------|------------------|-------|
| `--sl-color-primary-*` | sky | Brand color, primary actions |
| `--sl-color-success-*` | green | Success states, positive feedback |
| `--sl-color-warning-*` | amber | Warnings, caution states |
| `--sl-color-danger-*` | red | Errors, destructive actions |
| `--sl-color-neutral-*` | gray | Text, borders, backgrounds |

### Neutral Special Values

```css
--sl-color-neutral-0: white
--sl-color-neutral-50 through 950: gray shades
--sl-color-neutral-1000: black
```

## How to Customize Colors

### Option 1: Change Brand Color (Recommended)

Edit `assets/css/variables.css`:

```css
:root {
    /* Change primary from sky to purple */
    --sl-color-primary-50: var(--sl-color-purple-50);
    --sl-color-primary-100: var(--sl-color-purple-100);
    --sl-color-primary-200: var(--sl-color-purple-200);
    --sl-color-primary-300: var(--sl-color-purple-300);
    --sl-color-primary-400: var(--sl-color-purple-400);
    --sl-color-primary-500: var(--sl-color-purple-500);
    --sl-color-primary-600: var(--sl-color-purple-600);
    --sl-color-primary-700: var(--sl-color-purple-700);
    --sl-color-primary-800: var(--sl-color-purple-800);
    --sl-color-primary-900: var(--sl-color-purple-900);
    --sl-color-primary-950: var(--sl-color-purple-950);
}
```

### Option 2: Use Primitives Directly (For Specific Cases)

```css
.feature-card {
    background: var(--sl-color-indigo-50);
    border: 1px solid var(--sl-color-indigo-200);
}

.feature-icon {
    color: var(--sl-color-indigo-600);
}
```

### Option 3: Custom Values (Use Sparingly)

Only for brand-specific colors not covered by primitives:

```css
:root {
    /* Custom brand colors */
    --brand-orange: #ff6b35;
    --brand-teal: #004e89;
}
```

## Using Tokens in Your CSS

### ✅ Good Examples

```css
/* Use semantic tokens */
.button-primary {
    background: var(--sl-color-primary-600);
    color: var(--sl-color-neutral-0);
}

/* Use primitives for specific needs */
.badge-new {
    background: var(--sl-color-violet-100);
    color: var(--sl-color-violet-700);
}

/* Use Shoelace spacing */
.card {
    padding: var(--sl-spacing-large);
    gap: var(--sl-spacing-medium);
}

/* Use Shoelace shadows */
.elevated-card {
    box-shadow: var(--sl-shadow-large);
}
```

### ❌ Bad Examples

```css
/* Don't hardcode colors */
.button {
    background: #3b82f6; /* Bad! */
}

/* Don't hardcode spacing */
.card {
    padding: 20px; /* Bad! Use var(--sl-spacing-large) */
}

/* Don't create duplicate token systems */
:root {
    --my-primary-color: blue; /* Bad! Use Shoelace tokens */
}
```

## Available Shoelace Design Tokens

### Colors
- `--sl-color-{theme}-{50-950}` - Theme colors (primary, success, warning, danger, neutral)
- `--sl-color-{primitive}-{50-950}` - Color primitives (20+ palettes)

### Spacing
- `--sl-spacing-3x-small` (2px)
- `--sl-spacing-2x-small` (4px)
- `--sl-spacing-x-small` (8px)
- `--sl-spacing-small` (12px)
- `--sl-spacing-medium` (16px)
- `--sl-spacing-large` (20px)
- `--sl-spacing-x-large` (28px)
- `--sl-spacing-2x-large` (36px)
- `--sl-spacing-3x-large` (48px)
- `--sl-spacing-4x-large` (72px)

### Border Radius
- `--sl-border-radius-small` (3px)
- `--sl-border-radius-medium` (4px)
- `--sl-border-radius-large` (8px)
- `--sl-border-radius-x-large` (16px)
- `--sl-border-radius-circle` (50%)
- `--sl-border-radius-pill` (9999px)

### Shadows
- `--sl-shadow-x-small`
- `--sl-shadow-small`
- `--sl-shadow-medium`
- `--sl-shadow-large`
- `--sl-shadow-x-large`

### Typography
- `--sl-font-sans`, `--sl-font-serif`, `--sl-font-mono`
- `--sl-font-size-2x-small` through `--sl-font-size-4x-large`
- `--sl-font-weight-light`, `-normal`, `-semibold`, `-bold`
- `--sl-line-height-denser`, `-dense`, `-normal`, `-loose`, `-looser`

### Transitions
- `--sl-transition-x-slow` (1000ms)
- `--sl-transition-slow` (500ms)
- `--sl-transition-medium` (250ms)
- `--sl-transition-fast` (150ms)
- `--sl-transition-x-fast` (50ms)

## Dark Theme

Shoelace provides automatic dark mode support. To enable:

```html
<html class="sl-theme-dark">
```

Or toggle programmatically:

```javascript
document.documentElement.classList.toggle('sl-theme-dark');
```

## Custom Application Tokens

Define app-specific tokens in `variables.css` **without** the `--sl-` prefix:

```css
:root {
    /* Layout constraints */
    --max-w-lg: 1024px;
    
    /* App-specific gradients */
    --hero-gradient: linear-gradient(
        135deg,
        var(--sl-color-primary-50),
        var(--sl-color-sky-50)
    );
}
```

## Resources

- [Shoelace Documentation](https://shoelace.style/)
- [Shoelace Design Tokens](https://shoelace.style/tokens/color)
- [Source: light.css](https://github.com/shoelace-style/shoelace/blob/next/src/themes/light.css)
- [Source: dark.css](https://github.com/shoelace-style/shoelace/blob/next/src/themes/dark.css)

## Summary

**Always prefer Shoelace tokens over custom colors.** They provide:
- ✅ Consistency across components
- ✅ Automatic dark mode
- ✅ Accessibility
- ✅ Maintainability
- ✅ Professional design system

Only create custom tokens for application-specific needs that Shoelace doesn't cover.
