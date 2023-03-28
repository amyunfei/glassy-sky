import path from 'path'
import { defineConfig } from 'vite'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'
import react from '@vitejs/plugin-react-swc'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    createSvgIconsPlugin({
      iconDirs: [path.resolve(__dirname, './src/icons/svg')],
      symbolId: 'icon-[dir]-[name]'
    })
  ],
  resolve: {
    alias: [
      { find: /^~/, replacement: '' },
      { find: '@', replacement: path.resolve(__dirname, './src') }
    ]
  },
  css: {
    preprocessorOptions: {
      less: {
        javascriptEnabled: true,
      }
    }  
  },
  server: {
    proxy: {
      '/adminServer': {
        target: 'http://localhost:9999',
        rewrite: path => path.replace(/^\/adminServer/, '')
      }
    }
  }
})
