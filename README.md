# André Karrlein — Portfolio (v3)

> **Clean • Modern • Animated • Focused**

Single-site repository for the professional portfolio at [karrlein.com](https://www.karrlein.com).

## Overview

This repository contains the complete source for the modern portfolio website.

**v3 Highlights (feat/v3-enhanced-animations branch):**
- Cool, meaningful animations (scroll reveals, magnetic buttons, 3D hero portrait tilt, progress bar, back-to-top)
- Fully split structure: `index.html` + `css/styles.css` + `js/script.js`
- Production-ready: Tailwind CDN removed to eliminate warnings
- All original design, content, and glassmorphism preserved
- Fully accessible (respects `prefers-reduced-motion`)

## Repository Structure

```
.
├── .github/workflows/             # Deployment workflows
├── index.html                     # Main portfolio (single-page)
├── css/
│   └── styles.css                 # Custom styles + all animations
├── js/
│   └── script.js                  # Interactive animations (magnetic, tilt, scroll reveals, etc.)
├── impressum.html                 # Legal / Imprint
├── images/                        # Portrait + referee photos
├── README.md                      # This file
└── (clean, maintainable structure)
```

## Quick Start (Local Preview)

1. Clone the repository
2. Checkout the `feat/v3-enhanced-animations` branch (recommended for latest animations)
3. Open `index.html` in any modern browser
4. No build step required — everything is static and self-contained

**For the full Tailwind-powered preview** (with all utilities), open the self-contained preview file in `/artifacts/karrlein-v3-preview.html`.

## Main Branch vs Feature Branch

- `master` — Stable production version
- `feat/v3-enhanced-animations` — Latest version with cool animations and split file structure (recommended for development)

## Deployment

Automatic via GitHub Actions on push to `master`:
- Syncs root contents → S3
- Invalidates CloudFront cache
- Uses latest AWS actions with best practices

**Required GitHub Secrets** (set in repo Settings → Secrets and variables → Actions):
- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `CLOUDFRONT_DISTRIBUTION_ID`

## Design System
- **Primary accent**: `#A855F7` (purple)
- **Typography**: Space Grotesk (headings), Inter (body)
- **Visual style**: Dark theme, glassmorphism cards, elegant hover states, Material Symbols icons
- **Animations**: Subtle, performant, meaningful (scroll-triggered, magnetic, 3D tilt)
- **Responsive**: Mobile-first

## Content Highlights
- Solution Architect at Red Bull (Media House, DDD • Scale • Cloud since 2020)
- Previous: Senior Technical Lead at FLYERALARM (e-commerce & automation)
- Referee (authority, integrity, focus)
- Sim Racer in Le Mans Ultimate (GT3 & Hypercar endurance)
- Political Enthusiast (technology + governance)

## Editing

- Edit `index.html` for content/structure
- Edit `css/styles.css` for styling and animations
- Edit `js/script.js` for interactive behavior
- Keep classes consistent with the design system

## Contact

X / Twitter: [@rbak44](https://x.com/rbak44)  
LinkedIn: [andre-karrlein](https://www.linkedin.com/in/andre-karrlein/)

---

**Orbital Command — 2026**