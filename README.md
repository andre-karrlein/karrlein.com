# André Karrlein — Portfolio

> **Clean • Modern • Focused**

Single-site repository for the professional portfolio at [karrlein.com](https://www.karrlein.com).

## Overview

This repository contains the complete source for the modern portfolio website. It has been cleaned up to a single, maintainable site with:

- Professional dark glassmorphism design
- Full SEO, accessibility, and performance optimizations
- Modern GitHub Actions deployment with CloudFront invalidation
- Simple root-level structure for easy editing and deployment

## Repository Structure

```
.
├── .github/workflows/deploy.yml   # Modern S3 + CloudFront deployment
├── index.html                     # Main portfolio (single-page)
├── impressum.html                 # Legal / Imprint
├── images/                        # Portrait + referee photos
├── README.md                      # This file
└── (clean single-site structure)
```

## Quick Start (Local Preview)

1. Clone the repository
2. Open `index.html` in any modern browser
3. No build step required — Tailwind CSS loads from CDN

## Deployment

Automatic via GitHub Actions on push to `master`:
- Syncs root contents → `s3://karrlein.com`
- Invalidates CloudFront cache automatically
- Uses latest AWS actions (v4) with best practices and concurrency control

**Required GitHub Secrets** (set in repo Settings → Secrets and variables → Actions):
- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `CLOUDFRONT_DISTRIBUTION_ID` (your CloudFront distribution ID)

## Design System
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

Edit `index.html` directly. Keep Tailwind utility classes consistent. Test responsiveness on mobile + desktop.

## Contact

X / Twitter: [@rbak44](https://x.com/rbak44)  
LinkedIn: [andre-karrlein](https://www.linkedin.com/in/andre-karrlein/)

---

**Orbital Command — 2026**