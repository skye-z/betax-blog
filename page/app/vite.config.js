import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'
import { VitePWA } from 'vite-plugin-pwa'

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    outDir: '../app_dist',
    rollupOptions: {
      output: {
        manualChunks: {
          icon: ['@vicons/antd', '@vicons/fa']
        }
      }
    }
  },
  plugins: [
    vue(), AutoImport({
      imports: [
        'vue',
        {
          'naive-ui': [
            'useDialog',
            'useMessage',
            'useNotification',
            'useLoadingBar'
          ]
        }
      ]
    }),
    Components({
      resolvers: [NaiveUiResolver()]
    }),
    VitePWA({ 
      registerType: 'autoUpdate',
      injectRegister: 'auto',
      workbox: {
        globPatterns: ['**/*.{js,css,html,png,svg,jpg,ico,json}'],
        runtimeCaching: [
          {
            urlPattern: /^.*\/app\/api.*/g,
            handler: 'CacheFirst',
            options: {
              cacheName: 'api-data-cache',
              expiration: {
                maxEntries: 50,
                maxAgeSeconds: 30 * 24 * 60 * 60
              },
              cacheableResponse: {
                statuses: [200]
              }
            }
          }
        ]
      },
      manifest: false
    })
  ],
  server: {
    proxy: {
      '/api': {
        target: "http://192.168.1.160:9800/api",
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    },
  }
})
