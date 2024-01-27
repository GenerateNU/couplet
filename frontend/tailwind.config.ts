const defaultTheme = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./App.{js,jsx,ts,tsx}", "./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    // TODO: Get colors from designers
    colors: {
      black: "#000000",
      gold: "#a4804a",
      red: "#c8102e",
      white: "#ffffff",
    },
    // TODO: Get fonts from designers
    fontFamily: {
      sans: [...defaultTheme.fontFamily.sans],
      display: [...defaultTheme.fontFamily.sans],
    },
  },
  plugins: [],
};