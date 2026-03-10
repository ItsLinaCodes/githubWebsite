# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Running Locally

No build system — serve static files directly:

```bash
python3 -m http.server 8000
# or
npx http-server
```

## Architecture

This is a **zero-dependency, vanilla HTML/CSS/JS** portfolio site. No frameworks, no npm, no build step.

### File Structure

- `index.html` — Main single-page portfolio (~1,064 lines, all-in-one HTML/CSS/JS)
- `projects/pulse/index.html` — Live server monitoring dashboard demo
- `projects/breach/index.html` — Terminal hacking game (CTF/pentest concepts)

### Main Portfolio (`index.html`)

Single-page with anchor-based navigation across 7 sections: Hero, About, Skills, Experience, Projects, Contact, Footer. All CSS and JS are inline in the file.

Key JS systems:
- **Canvas particle background** — animated dot network in hero section
- **Custom cursor** — `#cursor` + `#cursor-ring` track mouse with CSS transitions
- **Scroll reveals** — `IntersectionObserver` adds `.visible` class to `.reveal` elements
- **Typing animation** — loops through role titles in the hero
- **Project spotlight** — `mousemove` on `.project-card` drives a CSS radial gradient via `--mouse-x`/`--mouse-y` custom properties

### Project Demos

**Pulse (`projects/pulse/`)**: Simulated server monitoring dashboard. Uses `setInterval` to generate fake metrics and `<canvas>` for sparkline charts. No real backend.

**BREACH (`projects/breach/`)**: Terminal UI game with a command parser. Simulates a pentest scenario (recon → credential reuse → hash cracking → lateral movement → exfil). All logic is client-side JS with a fake filesystem and fake network state.

## Next Project

**Do not read, reference, or modify any files in `projects/breach/`** when working on the next project.

### Styling Conventions

- CSS custom properties defined on `:root` for colors, fonts, and spacing
- Color palette centers on `--accent: #00ff88` (green) and `--accent-secondary: #0088ff` (blue)
- Glassmorphism cards: `background: rgba(255,255,255,0.03)` + `backdrop-filter: blur`
- Fonts: Space Grotesk (body), JetBrains Mono (code/terminal elements) via Google Fonts
