let development = process.env.NODE_ENV !== 'production'

export default {
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,
  loading: false,

  // Target: https://go.nuxtjs.dev/config-target
  target: 'static',

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    titleTemplate: '%s - OUStory',
    title: 'OUStory',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: 'https://cdn.jsdelivr.net/npm/@mdi/font@5.x/css/materialdesignicons.min.css' },
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    "@assets/roboto/roboto.css",
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    "~/plugins/date.js",
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/vuetify
    '@nuxtjs/vuetify',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    ['@nuxtjs/proxy', { ws: false }],
    '@nuxtjs/axios',
    ['@nuxtjs/i18n', {
      strategy: 'prefix_except_default',
      locales: [
        {
          code: 'mn',
          file: 'mn-MN.js',
          name: 'Монгол',
          flag: 'mn',
        },
        {
          code: 'en',
          file: 'en-US.js',
          name: 'English',
          flag: 'gb',
        },
      ],
      lazy: true,
      langDir: 'lang/',
      defaultLocale: 'mn',
    }],
  ],

  // Vuetify module configuration: https://go.nuxtjs.dev/config-vuetify
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    defaultAssets: false,
    theme: {
      themes: {
        light: {
          primary: '#ff7e00'
        }
      }
    },
    icons: {
      iconfont: 'mdi', // default - only for display purposes
    },
  },

  ...development && {
    proxy: {
      '/api': {
        target: 'http://localhost:8000',
        ws: true,
        changeOrigin: true,
        secure: false
      },
      '/pub': {
        target: 'http://localhost:8000',
        ws: true,
        changeOrigin: true,
        secure: false
      }
    }
  },

  axios: {
    proxy: true,
    baseURL: development ? 'http://localhost:8000' : 'https://domain/api'
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  }
}
