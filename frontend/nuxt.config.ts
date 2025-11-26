// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-05-01',
  devtools: { enabled: true },
  ssr: false,
  modules: ['@nuxtjs/tailwindcss'],

  // Nuxt 4 uses app/ directory by default, but we keep srcDir for simplicity
  future: {
    compatibilityVersion: 4
  },

  nitro: {
    preset: 'static'
  },
  
  app: {
    baseURL: '/',
    head: {
      title: 'เกรียนเวฟ App',
      htmlAttrs: {
        lang: 'th'
      },
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no' },
        { name: 'theme-color', content: '#FF69B4' },
        { name: 'description', content: 'แอปเกรียนเวฟ ธีมอีสานสีชมพู' }
      ],
      link: [
        { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
        { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' },
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Prompt:wght@400;600;700&display=swap' }
      ]
    }
  },

  runtimeConfig: {
    public: {
      apiBase: process.env.API_BASE || 'http://localhost:8080'
    }
  }
})
