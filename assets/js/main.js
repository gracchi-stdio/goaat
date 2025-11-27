// Datastar - Hypermedia Framework
import 'https://cdn.jsdelivr.net/gh/starfederation/datastar@1.0.0-RC.6/bundles/datastar.js'

// Import Shoelace components (tree-shaken - only import what you use)
import '@shoelace-style/shoelace/dist/components/button/button.js';
import '@shoelace-style/shoelace/dist/components/icon/icon.js';
import '@shoelace-style/shoelace/dist/components/alert/alert.js';
import '@shoelace-style/shoelace/dist/components/card/card.js';
import '@shoelace-style/shoelace/dist/components/input/input.js';
import '@shoelace-style/shoelace/dist/components/dialog/dialog.js';
import '@shoelace-style/shoelace/dist/components/avatar/avatar.js';
import '@shoelace-style/shoelace/dist/components/badge/badge.js';
import '@shoelace-style/shoelace/dist/components/breadcrumb/breadcrumb.js';
import '@shoelace-style/shoelace/dist/components/breadcrumb-item/breadcrumb-item.js';
import '@shoelace-style/shoelace/dist/components/dropdown/dropdown.js';
import '@shoelace-style/shoelace/dist/components/menu/menu.js';
import '@shoelace-style/shoelace/dist/components/menu-item/menu-item.js';
import '@shoelace-style/shoelace/dist/components/divider/divider.js';
import '@shoelace-style/shoelace/dist/components/select/select.js';
import '@shoelace-style/shoelace/dist/components/option/option.js';
import '@shoelace-style/shoelace/dist/components/textarea/textarea.js';
import '@shoelace-style/shoelace/dist/components/checkbox/checkbox.js';

// Set the base path for Shoelace assets (icons, etc.)
import { setBasePath } from '@shoelace-style/shoelace/dist/utilities/base-path.js';
setBasePath('/node_modules/@shoelace-style/shoelace/dist');

console.log('🚀 Goaat app initialized');

// Set Shoelace theme based on system preference
const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
console.log('🎨 System prefers dark mode:', prefersDark);
document.documentElement.classList.toggle('sl-theme-dark', prefersDark);
document.documentElement.classList.toggle('sl-theme-light', !prefersDark);

// Listen for theme changes
window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
  document.documentElement.classList.toggle('sl-theme-dark', e.matches);
  document.documentElement.classList.toggle('sl-theme-light', !e.matches);
});

// Global alert system
window.showAlert = function(message, variant = 'primary', duration = 5000, icon = 'info-circle') {
  const container = document.getElementById('alert-container');
  if (!container) return;

  const alert = document.createElement('sl-alert');
  alert.variant = variant;
  alert.closable = true;
  alert.duration = duration;
  
  const iconEl = document.createElement('sl-icon');
  iconEl.setAttribute('slot', 'icon');
  iconEl.setAttribute('name', icon);
  
  alert.appendChild(iconEl);
  alert.appendChild(document.createTextNode(message));
  
  container.appendChild(alert);
  alert.toast();
  
  // Remove after it closes
  alert.addEventListener('sl-after-hide', () => {
    alert.remove();
  });
  
  return alert;
};

// Listen for Datastar signals to show alerts
document.addEventListener('datastar-signal', (e) => {
  if (e.detail.eventType === 'showAlert') {
    const { message, variant = 'primary', icon = 'info-circle', duration = 5000 } = e.detail.data;
    window.showAlert(message, variant, duration, icon);
  }
});

// Datastar event handlers for debugging and user feedback
document.addEventListener('datastar-request-start', (e) => {
  console.log('🔄 Datastar request started:', e.detail);
});

document.addEventListener('datastar-request-end', (e) => {
  console.log('✅ Datastar request completed:', e.detail);
});

document.addEventListener('datastar-request-error', (e) => {
  console.error('❌ Datastar request error:', e.detail);
  window.showAlert('Request failed. Please try again.', 'danger', 5000, 'exclamation-triangle');
});

// Handle navigation with view transitions
if (document.startViewTransition) {
  document.addEventListener('datastar-navigation', (e) => {
    console.log('🔀 Navigation event:', e.detail);
  });
}

// Handle active navigation state
function updateActiveNav() {
  const path = window.location.pathname;
  document.querySelectorAll('.sidebar-nav a').forEach(link => {
    const href = link.getAttribute('href');
    if (href && path.startsWith(href)) {
      link.classList.add('active');
    } else {
      link.classList.remove('active');
    }
  });
}

// Update on load and navigation
document.addEventListener('DOMContentLoaded', updateActiveNav);
document.addEventListener('datastar-navigation', updateActiveNav);

// Initialize on DOM ready
document.addEventListener('DOMContentLoaded', () => {
  console.log('✅ DOM loaded');
  
  // Show welcome alert on first load (example)
  const isFirstVisit = !sessionStorage.getItem('welcomed');
  if (isFirstVisit && window.location.pathname.includes('/admin/')) {
    sessionStorage.setItem('welcomed', 'true');
    setTimeout(() => {
      window.showAlert('Welcome to Goaat! 🎉', 'success', 5000, 'check-circle');
    }, 500);
  }
});

