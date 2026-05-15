# André Karrlein — Portfolio (v3)

> **Clean • Modern • Animated • Focused**

Single-site repository for the professional portfolio at [karrlein.com](https://www.karrlein.com).

## Overview

This repository contains the complete source for the modern portfolio website.

**v3 Highlights (feat/v3-enhanced-animations branch):**
- Cool, meaningful animations (scroll reveals, magnetic buttons, 3D hero portrait tilt, progress bar, back-to-top)
- Fully split structure: `index.html` + `css/styles.css` + `js/script.js`
- Production-ready: Tailwind CDN removed (no more warnings)
- All original design, content, and glassmorphism preserved
- Fully accessible (respects `prefers-reduced-motion`)

## Repository Structure

```
.
├── tailwind.config.js             # Production Tailwind config
├── input.css                      # Tailwind entry point
├── package.json                   # Build scripts (optional but recommended)
├── index.html                     # Main portfolio
├── css/
│   └── styles.css                 # Custom styles + animations
├── js/
│   └── script.js                  # Interactive animations
├── impressum.html                 # Legal / Imprint
├── images/                        # Portrait + referee photos
├── README.md                      # This file
└── (clean, maintainable structure)
```

## Quick Start (Local Preview)

1. Clone the repository
2. Checkout the `feat/v3-enhanced-animations` branch
3. Open `index.html` in any modern browser
4. No build step required — fully static

**For the full Tailwind-powered experience** (with all utilities), open the self-contained preview:
`/artifacts/karrlein-v3-preview.html`

## Production Tailwind Build (Fix for "could not determine executable to run")

If you get the error `npx tailwindcss ... could not determine executable to run`, use this **correct command for Tailwind v4**:

```bash
# 1. Install (first time only)
npm install -D tailwindcss @tailwindcss/cli

# 2. Build production CSS (correct command)
npx @tailwindcss/cli -i ./input.css -o ./css/tailwind.css --minify

# 3. (Optional) Add this line in index.html <head> for full utilities
<link rel="stylesheet" href="css/tailwind.css">
```

**Even easier** — add this to `package.json`:

```json
"scripts": {
  "build:css": "@tailwindcss/cli -i input.css -o css/tailwind.css --minify"
}
```

Then just run:
```bash
npm run build:css
```

This gives you:
- Zero console warnings
- Full Tailwind utility classes
- Smaller, optimized payload
- One-time setup

## Main Branch vs Feature Branch

- `master` — Stable production version
- `feat/v3-enhanced-animations` — Latest version with animations + split structure (recommended)

## Deployment

Automatic via GitHub Actions on push to `master` (S3 + CloudFront).

## Design System & Animations
- **Primary accent**: `#A855F7` (purple)
- **Typography**: Space Grotesk (headings), Inter (body)
- **Visual style**: Dark theme, glassmorphism, elegant hovers
- **Animations**: Subtle, performant, meaningful (scroll-triggered, magnetic, 3D tilt)
- **Responsive**: Mobile-first

## Content Highlights
- Solution Architect at Red Bull (Media House, DDD • Scale • Cloud since 2020)
- Previous: Senior Technical Lead at FLYERALARM
- Referee, Sim Racer (Le Mans Ultimate), Political Enthusiast

## Editing
- `index.html` → content & structure
- `css/styles.css` → styling & animations
- `js/script.js` → interactive behavior
- `tailwind.config.js` + `input.css` → Tailwind build

## Contact

X: [@rbak44](https://x.com/rbak44)  
LinkedIn: [andre-karrlein](https://www.linkedin.com/in/andre-karrlein/)

---

**Orbital Command — 2026**