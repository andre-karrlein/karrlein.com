# André Karrlein — Portfolio (v3)

> **Clean • Modern • Animated • Focused**

Single-site repository for the professional portfolio at [karrlein.com](https://www.karrlein.com).

## Overview

This repository contains the complete source for the modern portfolio website.

**Current Status (Working):** The site is fully functional using only `css/styles.css`. All animations, glassmorphism, and design are working perfectly. The Tailwind build is **optional** for advanced users who want full utility classes.

**v3 Highlights:**
- Cool animations (scroll reveals, magnetic buttons, 3D tilt, progress bar, back-to-top)
- Clean split: `index.html` + `css/styles.css` + `js/script.js`
- No Tailwind CDN warning
- Fully accessible

## Quick Start

Just open `index.html` in your browser. Everything works out of the box.

## Optional: Full Tailwind Build (Advanced)

If you want 100% Tailwind utilities, run:

```bash
npm install -D tailwindcss @tailwindcss/cli
npx @tailwindcss/cli -i ./input.css -o ./css/tailwind.css --minify
```

Then **manually** add this line in `index.html`:
```html
<link rel="stylesheet" href="css/tailwind.css">
```

**Note:** Only do this if you need extra Tailwind classes. The current `css/styles.css` already covers the full design.

## Repository Structure

```
.
├── tailwind.config.js
├── input.css
├── index.html
├── css/styles.css          ← Currently active (working design)
├── js/script.js
└── ...
```

## Contact

X: [@rbak44](https://x.com/rbak44)  
LinkedIn: [andre-karrlein](https://www.linkedin.com/in/andre-karrlein/)

---

**Orbital Command — 2026**