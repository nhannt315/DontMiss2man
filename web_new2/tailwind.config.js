module.exports = {
  purge: [
    "./src/pages/**/*.tsx",
    "./src/components/**/*.tsx",
    "./src/containers/**/*.tsx",
  ],
  darkMode: false, // or 'media' or 'class'
  theme: {
    fontFamily: {
      futura: "Futura",
      body: ["Helvetica", "Lato", "sans-serif"],
    },
    extend: {},
    screens: {
      xs: { max: "639px" },
      sm: { min: "640px", max: "767px" },
      md: { min: "768px", max: "1023px" },
      lg: { min: "1024px", max: "1279px" },
      xl: { min: "1280px", max: "1535px" },
      "2xl": { min: "1536px" },
    },
  },
  variants: {
    extend: {
      // ref: https://tailwindcss.com/docs/hover-focus-and-other-states#first-child
      borderWidth: ["first", "last"],
      // ref: https://tailwindcss.com/docs/hover-focus-and-other-states#disabled
      cursor: ["disabled"],
      opacity: ["disabled", "active"],
      // ref: https://tailwindcss.com/docs/hover-focus-and-other-states#focus-visible
      ringWidth: ["focus-visible"],
      ringColor: ["focus-visible"],
      padding: ["last"],
      margin: ["first"],
    },
  },
  plugins: [],
};
