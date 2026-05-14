# André Karrlein — Portfolio (v2 Modernized)

> **Clean • Modern • Focused**

Single-site repository for the professional portfolio at [karrlein.com](https://www.karrlein.com).

## What's New in v2 (May 2026)

- **Fully cleaned up**: Removed `motorsport/` and `politik/` sub-sites (future dedicated repos if needed)
- **Restructured to root**: Simpler deployment (`index.html` + `impressum.html` + `images/` directly at root)
- **Modernized DevOps**: Updated GitHub Actions to v4 with best practices
- **Enhanced**: Full SEO meta tags, Open Graph, Twitter Cards, JSON-LD structured data, accessibility improvements, performance optimizations (lazy loading, font preconnect, image dimensions)
- **Preserved 100%**: Exact same beautiful dark glassmorphism design, typography (Space Grotesk + Inter), colors (#A855F7), layout, content, and interactions

## Repository Structure

```
.
├── .github/workflows/deploy.yml   # Modern S3 deployment
├── index.html                     # Main portfolio (single-page, enhanced)
├── impressum.html                 # Legal / Imprint
├── images/                        # Portrait + referee photos
├── README.md                      # This file
└── (no other folders or legacy code)
```

## Quick Start (Local Preview)

1. Clone the `v2-modern-portfolio` branch
2. Open `index.html` in any modern browser
3. No build step, no dependencies — Tailwind loads from CDN

## Deployment

Automatic via GitHub Actions on push to `v2-modern-portfolio` (or master after merge):
- Syncs root contents → `s3://karrlein.com`
- Uses latest AWS actions + recommended security practices

## Design System (Preserved)
- **Primary accent**: `#A855F7` (purple)
- **Typography**: Space Grotesk (headings), Inter (body)
- **Visual style**: Dark theme, glassmorphism cards, elegant hover states, Material Symbols icons
- **Responsive**: Mobile-first with Tailwind

## Content Highlights
- Solution Architect at Red Bull (Media House, DDD • Scale • Cloud since 2020)
- Previous: Senior Technical Lead at FLYERALARM (e-commerce & automation)
- Referee (authority, integrity, focus)
- Sim Racer in Le Mans Ultimate (GT3 & Hypercar endurance)
- Political Enthusiast (technology + governance)

## Editing

Edit `index.html` directly. Keep Tailwind utility classes consistent. Test on mobile + desktop.

## Contact

X / Twitter: [@rbak44](https://x.com/rbak44)  
LinkedIn: [andre-karrlein](https://www.linkedin.com/in/andre-karrlein/)

---

**Orbital Command — 2026**  
*This v2 branch is the clean, modern foundation for karrlein.com.*