import { defineConfig } from 'vite'

export default defineConfig({
  build: {
    outDir: 'public',
    emptyOutDir: false, // Don't clear public/ (Go might have other files)
    rollupOptions: {
      input: 'assets/js/main.js',
      output: {
        entryFileNames: 'js/main.js',
        chunkFileNames: 'js/[name]-[hash].js',
        assetFileNames: (assetInfo) => {
          if (assetInfo.name.endsWith('.css')) {
            return 'css/main.css'
          }
          return 'assets/[name][extname]'
        }
      }
    }
  },

  server: {
    host: '0.0.0.0',
    port: 5173,
    strictPort: true,
    hmr: {
      host: 'localhost',
    },
    watch: {
      usePolling: true,
    },
    proxy: {
      '^/(hello|authors|health|api)': {
        target: 'http://localhost:8080',
        changeOrigin: true
      },
      '^/$': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    }
  },

  css: {
    postcss: './postcss.config.js'
  }
})
