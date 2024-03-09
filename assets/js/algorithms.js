function initAglorithmsPage() {
    const sideNav = document.getElementById('side-nav');

    if (sideNav) {
        const links = sideNav.querySelectorAll('.nav-link');

        for (const link of links) {
            if (window.location.href.includes(link.href)) {
                link.classList.add("active");
            } else {
                link.classList.remove("active");
            }
        }
    }
}

initAglorithmsPage();
