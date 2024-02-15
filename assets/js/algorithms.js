const sideNav = document.getElementById('side-nav');

if (sideNav) {
    const links = sideNav.querySelectorAll('a');   

    for (const link of links) {
        if (link.innerText == "About") {
            if (link.href == window.location.href) {
                link.classList.add("active");
            }
        } else {
            if (window.location.href.includes(link.href)) {
                link.classList.add("active");
            }
        }
    }
}
