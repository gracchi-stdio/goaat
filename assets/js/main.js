import htmx from 'htmx.org';
window.htmx = htmx;

// Import Shoelace components (tree-shaken - only import what you use)
import '@shoelace-style/shoelace/dist/components/button/button.js';
import '@shoelace-style/shoelace/dist/components/icon/icon.js';
import '@shoelace-style/shoelace/dist/components/alert/alert.js';
import '@shoelace-style/shoelace/dist/components/card/card.js';
import '@shoelace-style/shoelace/dist/components/input/input.js';
import '@shoelace-style/shoelace/dist/components/dialog/dialog.js';

// Set the base path for Shoelace assets (icons, etc.)
import { setBasePath } from '@shoelace-style/shoelace/dist/utilities/base-path.js';
setBasePath('/node_modules/@shoelace-style/shoelace/dist');

// import "./form-group.js";

// Import your CSS
// import '../css/main.css';

// Your custom JavaScript
console.log('🚀 Goaat app initialized');

// Set Shoelace theme based on system preference
const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
console.log('🎨 System prefers dark mode:', prefersDark);
document.documentElement.classList.toggle('sl-theme-dark', prefersDark);
document.documentElement.classList.toggle('sl-theme-light', !prefersDark);
console.log('🎨 Applied theme class:', document.documentElement.className);

// Listen for theme changes
window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
  document.documentElement.classList.toggle('sl-theme-dark', e.matches);
  document.documentElement.classList.toggle('sl-theme-light', !e.matches);
});

// Example: Add some interactivity
document.addEventListener('DOMContentLoaded', () => {
  console.log('✅ DOM loaded');
  
  // Example: Handle Shoelace events
  document.querySelectorAll('sl-button').forEach(button => {
    button.addEventListener('click', (e) => {
      console.log('Button clicked:', e.target);
    });
  });
});
