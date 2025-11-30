/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'media',
  content: [
    "./index.html",
    "./src/**/*.{vue,ts,js,jsx,tsx}"
  ],
  theme: {
    extend: {
      colors: {
        'glass-bg': 'color-mix(in oklab, var(--bg) 80%, transparent)',
        'glass-border': 'color-mix(in oklab, var(--glass-bg) 70%, white)',
      },
      width: {
        'custom-76': '18.75rem',
        'custom-76p8': '26.75rem',
        '48': '12rem',
        '48p8': '20rem',

      },
      margin: {
        '2': '0.6rem',
        '2-custom': '0.55rem',
      },
      boxShadow: {
        'soft-glow': 'inset 1px 1px 0.5px #ffffff15, inset -0.5px -0.5px 0.5px #ffffff10, inset 0 0 2px #ffffff15, 0 1px 2px #ffffff10',
      },
      screens: {
        'Uxs': '30rem',      // 480px
        'Usm': '40rem',      // 640px
        'Usm-44': '44rem',      // 640px
        'Umd': '48rem',      // 768px
        'Umd-60': '60rem',   // 960px (custom midpoint)
        'Ulg-70': '70rem',      // 1024px
        'Uxl': '83rem',      // 1280px
        'U2xl': '96rem',     // 1536px
        '4xl': '130rem',
      },
      backgroundImage: {
        'fade-overlay':
          'linear-gradient(to bottom, rgba(255,255,255,0)_0%, rgba(255,255,255,0.1)_10%, rgba(255,255,255,0.2)_20%, rgba(255,255,255,0.4)_40%, rgba(255,255,255,0.8)_50%, rgba(255,255,255,1)_100%)',
        'fade-overlay-dark':
          'linear-gradient(to bottom, rgba(17,17,17,0)_0%, rgba(17,17,17,0.1)_10%, rgba(17,17,17,0.2)_20%, rgba(17,17,17,0.4)_40%, rgba(17,17,17,0.8)_50%, rgba(17,17,17,1)_100%)',

      },
    },
  },
  plugins: [],
}
