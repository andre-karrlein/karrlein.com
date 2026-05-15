// js/script.js - Cool, meaningful animations for André Karrlein Portfolio v3
// - Scroll reveals with stagger
// - Magnetic interactive buttons
// - Mouse-tilt parallax on hero portrait
// - Scroll progress bar
// - Back-to-top with smooth animation
// Fully respects prefers-reduced-motion for accessibility
// No external dependencies - pure vanilla, high performance

document.addEventListener('DOMContentLoaded', () => {
    const prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;

    // ============================================
    // 1. SCROLL PROGRESS INDICATOR (top bar)
    // ============================================
    const progressBar = document.createElement('div');
    progressBar.id = 'scroll-progress';
    progressBar.style.width = '0%';
    document.body.prepend(progressBar);

    let ticking = false;
    window.addEventListener('scroll', () => {
        if (prefersReducedMotion || ticking) return;
        ticking = true;
        
        requestAnimationFrame(() => {
            const scrollTop = window.scrollY;
            const docHeight = document.documentElement.scrollHeight - window.innerHeight;
            const scrollPercent = docHeight > 0 ? Math.min((scrollTop / docHeight) * 100, 100) : 0;
            progressBar.style.width = `${scrollPercent}%`;
            ticking = false;
        });
    }, { passive: true });

    // ============================================
    // 2. SCROLL-REVEAL ANIMATIONS (IntersectionObserver)
    // ============================================
    const revealElements = document.querySelectorAll('.reveal');
    
    if (!prefersReducedMotion && 'IntersectionObserver' in window && revealElements.length > 0) {
        const revealObserver = new IntersectionObserver((entries) => {
            entries.forEach((entry, idx) => {
                if (entry.isIntersecting) {
                    // Staggered reveal for beautiful cascade effect
                    setTimeout(() => {
                        entry.target.classList.add('visible');
                    }, idx * 75);
                    revealObserver.unobserve(entry.target);
                }
            });
        }, {
            threshold: 0.12,
            rootMargin: '-60px 0px -60px 0px'
        });

        revealElements.forEach(el => revealObserver.observe(el));
    } else {
        // Immediate show for reduced motion or no IO support
        revealElements.forEach(el => el.classList.add('visible'));
    }

    // ============================================
    // 3. MAGNETIC BUTTONS (subtle cursor attraction + shine)
    // ============================================
    const magneticBtns = document.querySelectorAll('.magnetic-btn');
    
    magneticBtns.forEach((btn) => {
        let raf = null;
        
        btn.addEventListener('mousemove', (e) => {
            if (prefersReducedMotion) return;
            if (raf) cancelAnimationFrame(raf);
            
            raf = requestAnimationFrame(() => {
                const rect = btn.getBoundingClientRect();
                const relX = ((e.clientX - rect.left) / rect.width - 0.5) * 2;
                const relY = ((e.clientY - rect.top) / rect.height - 0.5) * 2;
                
                // Gentle magnetic pull
                const translateX = relX * 6;
                const translateY = relY * 4;
                btn.style.transform = `translate(${translateX}px, ${translateY}px)`;
                
                // Dynamic shine position
                btn.style.setProperty('--mouse-x', `${(e.clientX - rect.left) / rect.width * 100}%`);
                btn.style.setProperty('--mouse-y', `${(e.clientY - rect.top) / rect.height * 100}%`);
            });
        });
        
        btn.addEventListener('mouseleave', () => {
            if (raf) cancelAnimationFrame(raf);
            btn.style.transition = 'transform 0.5s cubic-bezier(0.23, 1, 0.32, 1)';
            btn.style.transform = 'translate(0, 0)';
            
            setTimeout(() => {
                btn.style.transition = 'transform 0.25s cubic-bezier(0.23, 1, 0.32, 1)';
            }, 500);
        });
    });

    // ============================================
    // 4. HERO PORTRAIT 3D TILT + PARALLAX (mouse reactive)
    // ============================================
    const heroContainer = document.querySelector('.hero-image-container');
    const heroImg = heroContainer?.querySelector('img');
    
    if (heroContainer && heroImg && !prefersReducedMotion) {
        let tiltRaf = null;
        
        heroContainer.addEventListener('mousemove', (e) => {
            if (tiltRaf) cancelAnimationFrame(tiltRaf);
            
            tiltRaf = requestAnimationFrame(() => {
                const rect = heroContainer.getBoundingClientRect();
                const centerX = rect.width / 2;
                const centerY = rect.height / 2;
                const mouseX = e.clientX - rect.left;
                const mouseY = e.clientY - rect.top;
                
                const rotX = ((mouseY - centerY) / centerY) * -7.5;
                const rotY = ((mouseX - centerX) / centerX) * 7.5;
                
                heroContainer.style.transform = `perspective(1200px) rotateX(${rotX}deg) rotateY(${rotY}deg)`;
            });
        });
        
        heroContainer.addEventListener('mouseleave', () => {
            if (tiltRaf) cancelAnimationFrame(tiltRaf);
            heroContainer.style.transition = 'transform 0.9s cubic-bezier(0.23, 1, 0.32, 1)';
            heroContainer.style.transform = 'perspective(1200px) rotateX(0deg) rotateY(0deg)';
            
            setTimeout(() => {
                heroContainer.style.transition = 'transform 0.12s ease-out';
            }, 900);
        });
        
        // Subtle breathing animation on load
        setTimeout(() => {
            if (!heroContainer.matches(':hover')) {
                heroContainer.style.transition = 'transform 4s ease-in-out';
                heroContainer.style.transform = 'perspective(1200px) rotateX(1.5deg) rotateY(-1deg)';
                
                setTimeout(() => {
                    if (!heroContainer.matches(':hover')) {
                        heroContainer.style.transform = 'perspective(1200px) rotateX(0) rotateY(0)';
                    }
                }, 3800);
            }
        }, 2200);
    }

    // ============================================
    // 5. BACK-TO-TOP BUTTON (appears on scroll)
    // ============================================
    const backToTopBtn = document.createElement('button');
    backToTopBtn.id = 'back-to-top';
    backToTopBtn.setAttribute('aria-label', 'Scroll back to top of page');
    backToTopBtn.innerHTML = `
        <span class="material-symbols-outlined" style="font-size: 26px; line-height: 1;">arrow_upward</span>
    `;
    document.body.appendChild(backToTopBtn);

    let scrollTimeout;
    window.addEventListener('scroll', () => {
        clearTimeout(scrollTimeout);
        scrollTimeout = setTimeout(() => {
            if (window.scrollY > 650) {
                backToTopBtn.classList.add('visible');
            } else {
                backToTopBtn.classList.remove('visible');
            }
        }, 80);
    }, { passive: true });

    backToTopBtn.addEventListener('click', (e) => {
        e.preventDefault();
        window.scrollTo({
            top: 0,
            behavior: prefersReducedMotion ? 'auto' : 'smooth'
        });
    });

    // Keyboard support
    backToTopBtn.addEventListener('keydown', (e) => {
        if (e.key === 'Enter' || e.key === ' ') {
            e.preventDefault();
            window.scrollTo({ top: 0, behavior: prefersReducedMotion ? 'auto' : 'smooth' });
        }
    });

    // ============================================
    // 6. ENHANCED GLASS CARDS - 3D tilt on hover (extra polish)
    // ============================================
    const cards = document.querySelectorAll('.glass-card');
    
    cards.forEach(card => {
        if (prefersReducedMotion) return;
        
        let cardRaf = null;
        
        card.addEventListener('mousemove', (e) => {
            if (cardRaf) cancelAnimationFrame(cardRaf);
            
            cardRaf = requestAnimationFrame(() => {
                const rect = card.getBoundingClientRect();
                const x = ((e.clientX - rect.left) / rect.width - 0.5) * 6;
                const y = ((e.clientY - rect.top) / rect.height - 0.5) * -6;
                
                card.style.transform = `translateY(-10px) rotateX(${y}deg) rotateY(${x}deg)`;
            });
        });
        
        card.addEventListener('mouseleave', () => {
            if (cardRaf) cancelAnimationFrame(cardRaf);
            card.style.transition = 'transform 0.7s cubic-bezier(0.23, 1, 0.32, 1), box-shadow 0.4s ease';
            card.style.transform = 'translateY(0) rotateX(0) rotateY(0)';
            
            setTimeout(() => {
                card.style.transition = 'transform 0.4s cubic-bezier(0.23, 1, 0.32, 1), box-shadow 0.4s ease, background 0.3s ease';
            }, 700);
        });
    });

    // ============================================
    // 7. HERO CONTENT STAGGERED ENTRANCE
    // ============================================
    const heroText = document.querySelector('.md\\:col-span-7');
    if (heroText && !prefersReducedMotion) {
        heroText.style.opacity = '0';
        heroText.style.transform = 'translateY(40px)';
        
        setTimeout(() => {
            heroText.style.transition = 'opacity 1.1s cubic-bezier(0.23, 1, 0.32, 1), transform 1.1s cubic-bezier(0.23, 1, 0.32, 1)';
            heroText.style.opacity = '1';
            heroText.style.transform = 'translateY(0)';
        }, 420);
    }

    // ============================================
    // 8. CONSOLE EASTER EGG + INIT LOG
    // ============================================
    console.log(
        '%c[André Karrlein Portfolio v3] ✨ Cool animations loaded successfully. Scroll, hover, tilt, enjoy! (Reduced motion: ' + prefersReducedMotion + ')',
        'color: #A855F7; font-family: monospace; font-size: 9px'
    );
    
    // Optional: Konami code for fun (up up down down left right left right b a)
    // (kept minimal, not intrusive)
});