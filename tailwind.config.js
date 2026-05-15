/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./impressum.html"
  ],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        "on-secondary-fixed": "#3e4040",
        "surface-bright": "#2c2c2c",
        "error-container": "#9f0519",
        "primary-fixed": "#A855F7",
        "secondary-fixed-dim": "#d4d4d4",
        "on-surface": "#ffffff",
        "background": "#000000",
        "outline": "#757575",
        "surface-container": "#191919",
        "primary": "#A855F7",
        "primary-container": "#A855F7",
        "on-primary-fixed": "#ffffff",
        "on-tertiary-fixed": "#515050",
        "on-primary": "#ffffff",
        "inverse-surface": "#f9f9f9",
        "secondary-fixed": "#e2e2e2",
        "surface-container-lowest": "#000000",
        "on-secondary": "#505252",
        "on-primary-container": "#ffffff",
        "on-error": "#490006",
        "outline-variant": "#484848",
        "surface-variant": "#262626",
        "secondary-container": "#454747",
        "tertiary-dim": "#edeaea",
        "surface-container-low": "#131313",
        "on-tertiary-container": "#575656",
        "primary-dim": "#9333EA",
        "on-error-container": "#ffa8a3",
        "secondary-dim": "#d4d4d4",
        "on-surface-variant": "#ababab",
        "surface-container-high": "#1f1f1f",
        "tertiary": "#fcf9f8",
        "tertiary-container": "#edeaea",
        "inverse-primary": "#7e22ce",
        "on-background": "#ffffff",
        "surface-tint": "#A855F7",
        "inverse-on-surface": "#555555",
        "tertiary-fixed-dim": "#f3f0ef",
        "primary-fixed-dim": "#9333EA",
        "surface-dim": "#0e0e0e",
        "on-tertiary-fixed-variant": "#6e6d6d",
        "error": "#ff716c",
        "tertiary-fixed": "#ffffff",
        "error-dim": "#d7383b",
        "on-primary-fixed-variant": "#7e22ce",
        "on-secondary-fixed-variant": "#5a5c5c",
        "surface": "#0e0e0e",
        "secondary": "#e2e2e2",
        "surface-container-highest": "#262626",
        "on-tertiary": "#605f5f",
        "on-secondary-container": "#d0d0d0"
      },
      fontFamily: {
        "headline": ["Space Grotesk"],
        "body": ["Inter"],
        "label": ["Inter"]
      },
      borderRadius: {
        "DEFAULT": "0.125rem",
        "lg": "0.25rem",
        "xl": "0.5rem",
        "full": "0.75rem"
      }
    }
  },
  plugins: []
}
