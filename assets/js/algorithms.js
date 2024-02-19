function initAglorithmsPage() {
    const sideNav = document.getElementById('side-nav');
    const algorithmContent = document.getElementById('algorithm-content')

    if (sideNav && algorithmContent) {
        const links = sideNav.querySelectorAll('.nav-link');

        const updateUrl = event => {
            const newUrl = window.location.origin + event.target.dataset.href;
            window.history.pushState({ path: newUrl }, '', newUrl);

            updateHxGet();
            updateLinks();
        }

        const updateHxGet = () => {
            algorithmContent.setAttribute('hx-get', `/api${window.location.pathname}`);
            htmx.process(algorithmContent);
        }

        const updateLinks = () => {
            for (const link of links) {
                if (link.innerText == "About") {
                    if (link.dataset.href == window.location.pathname) {
                        link.classList.add("active");
                    } else {
                        link.classList.remove("active");
                    }
                } else if (window.location.href.includes(link.dataset.href)) {
                    link.classList.add("active");
                } else {
                    link.classList.remove("active");
                }
            }
        }

        updateHxGet();
        updateLinks();

        for (const link of links) {
            link.addEventListener('click', updateUrl);
        }
    }
}

initAglorithmsPage();

function codeFetched(event) {
    code = event.target.querySelector('code');
    if (code) {
        hljs.highlightElement(code);
    }
}
