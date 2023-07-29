const colors = require('tailwindcss/colors')
const plugin = require('tailwindcss/plugin')

module.exports = {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  corePlugins: {
    preflight: false
  },
  theme: {
    extend: {},
    colors: {
      transparent: 'transparent',
      current: 'currentColor',
      black: colors.black,
      white: colors.white,
      indigo: colors.indigo,
      blue: {
        link: '#0081ff',
        DEFAULT: '#22cce2'
      },
      green: {
        DEFAULT: '#09b66d'
      },
      red: {
        DEFAULT: '#ff3d57'
      },
      orange: {
        DEFAULT: '#ff8a48'
      },
      yellow: {
        DEFAULT: '#fdbf5e'
      },
      gray: {
        lightest: '#8291a9',
        light: '#425365',
        DEFAULT: '#dbe7ee',
        dark: '#1f2935',
        darker: '#161d26'
      }
    }
  },
  plugins: [
    plugin(function ({ addBase, addComponents, theme }) {
      addBase({
        '.font-barlow': { fontFamily: 'Barlow' },
        '.h-15': { height: '3.75rem !important' }, // 高度
        '.h-30': { height: '7.5rem' },
        '.h-120': { height: '30rem' },
        '.w-70': { width: '17.5rem' },
        '.pt-15': { paddingTop: '3.75rem' },
        '.pb-full': { paddingBottom: '100%' },
        '.y-center': {
          top: '50%',
          transform: 'translateY(-50%)'
        },
        '.flex-center-xy': {
          'display': 'flex',
          'justifyContent': 'center',
          'alignItems': 'center'
        },
        '.border-color-base': {
          borderColor: 'rgba(130, 145, 169, 0.25)'
        },
        '.svg-icon': {
          width: '1em',
          height: '1em',
          verticalAlign: '-0.1em', /* 因icon大小被设置为和字体大小一致，而span等标签的下边缘会和字体的基线对齐，故需设置一个往下的偏移比例，来纠正视觉上的未对齐效果 */
          fill: 'currentColor', /* 定义元素的颜色，currentColor是一个变量，这个变量的值就表示当前元素的color值，如果当前元素未设置color值，则从父元素继承 */
          overflow: 'hidden'
        }
      })
      // components
      const textLink = {
        '.text-link': {
          cursor: theme('cursor.pointer'),
          color: theme('colors.blue.link'),
          transition: 'color 0.3s',
          '&:hover': {
            color: '#299bff'
          },
          '&:active': {
            color: '#0065d9'
          }
        }
      }
      addComponents(textLink)
    })
  ]
}