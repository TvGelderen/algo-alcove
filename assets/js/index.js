let previousPage = null;
let pageContainer = null;

document.addEventListener('DOMContentLoaded', () => {
    previousPage = getCurrentPage();
    pageContainer = document.getElementById('page-content');

    initHeader();
});

function getCurrentPage() {
    return window.location.pathname.split('/')[1];
}

function updateUrl(path) {
    const newUrl = window.location.origin + path;
    pushState(newUrl);
}

function pushState(url) {
    window.history.pushState({ path: url }, '', url);
}

function updatePageContent() {
    if (pageContainer) {
        pageContainer.setAttribute('hx-get', `/pages${window.location.pathname}`);
        htmx.process(pageContainer);
    }
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

    const handleClick = event => {
        pushState(event.target.href);
        updateLinks();
    }

    const updateLinks = () => {
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

    updateLinks();
    updatePageContent();

    for (const link of links) {
        link.addEventListener('click', handleClick);
    }

    window.addEventListener('popstate', () => {
        let currentPage = getCurrentPage();

        if (previousPage !== currentPage) {
            previousPage = currentPage;
            updatePageContent();
            updateLinks();
        }
    });
}
