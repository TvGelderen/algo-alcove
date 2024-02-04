/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors')

export default {
    content: [
        './view/**/*.templ'
    ],
    theme: {
        colors: {
            ...colors,
            'primary': '#16161e',
            'secondary': '#1a1b26',
            'tertiary': '#24283b',
            'text': '#c0caf5',
            'text-secondary': '#a9b1d6',
            'theme': '#00ff00',
        },
        fontFamily: {
            'sans': ['"Courier New"']
        },
        extend: {},
    },
    plugins: [],
}
