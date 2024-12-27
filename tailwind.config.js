const colors = require("tailwindcss/colors");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["internal/templates/*.html", "static/*.html"],
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};