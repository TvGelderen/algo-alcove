document.addEventListener('DOMContentLoaded', () => {
    previousPage = getCurrentPage();
    pageContainer = document.getElementById('page-content');

    initHeader();
});

function getCurrentPage() {
    return window.location.pathname.split('/')[1];
}

function initHeader() {
    const header = document.querySelector('header');
    const menuIcon = document.querySelector('.menu-icon-wrapper');

    if (header && menuIcon) {
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

    const links = header.querySelectorAll('#nav-links .nav-link');

    for (const link of links) {
        if (link.innerText == "Home") {
            if (link.href == window.location.href) {
                link.classList.add("active");
            } else {
                link.classList.remove("active");
            }
        } else if (window.location.href.includes(link.href)) {
            link.classList.add("active");
        } else {
            link.classList.remove("active");
        }
    }
}
