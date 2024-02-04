/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors')

export default {
    content: [
        "./index.html",
        "./src/**/*.{vue,css}"
    ],
    theme: {
        colors: {
            ...colors,
            'primary': '#16161e',
            'secondary': '#1a1b26',
            'tertiary': '#24283b',
            'theme': '#00ff00',
        },
        fontFamily: {
            'sans': ['"Courier New"']
        },
        extend: {},
    },
    plugins: [],
}

