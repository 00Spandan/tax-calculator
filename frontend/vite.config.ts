import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

const bffProxyTarget = process.env.VITE_BFF_BASE_URL ?? 'http://localhost:8081'

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api': {
        target: bffProxyTarget,
        changeOrigin: true,
      },
    },
  },
})
