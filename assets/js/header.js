const header = document.querySelector('header');

header.addEventListener('load', event => {
    console.log(event);
});

const menuIcon = document.querySelector('.menu-icon-wrapper');

if (menuIcon && header) {
    const resizeListener = () => {
        if (window.innerWidth > 768) {
            header.dataset.open = "false";
            window.removeEventListener('resize', resizeListener);
        }
    }

    menuIcon.addEventListener('click', () => {
        if (header.dataset.open == "true") {
            header.dataset.open = "false";
            window.removeEventListener('resize', resizeListener);
        } else {
            header.dataset.open = "true";
            window.addEventListener('resize', resizeListener);
        }
    });
}

const navLinks = document.getElementById('nav-links');

if (navLinks) {
    for (const child of navLinks.children) {
        if (child.innerText == "Home") {
            if (window.location.href == child.href) {
                child.classList.add("active");
            }
        } else if (window.location.href.includes(child.href)) {
            child.classList.add("active");
        }
    }
}
