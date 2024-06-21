/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
 
    // Or if using `src` directory:
    "./src/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        inter: ['Inter', 'sans-serif'],
      },
    },
    colors:{
      blue: {
        100: '#E1E2F8',
        200: '#C3C5F1',
        300: '#B4B7ED',
        400: '#A5A9E9',
        500: '#959CE5',
        600: '#858EE1',
        700: '#7381DD',
        800: '#6174D9',
        900: '#4C68D5'
      },
      black:{
        100: '#ABB0B9',
        200: '#979DA9',
        300: '#838B98',
        400: '#707988',
        500: '#5D6778',
        600: '#4B5669',
        700: '#39465A',
        800: '#27364B',
        900: '#0C1024'
      },
      white:{
        100: '#FAFBFF',
        200: '#F1F4F9',
        300: '#ECF0F5',
        400: '#E2E8F0',
        900: '#FFFFFF'
      }
    }
  },
  plugins: [],
}