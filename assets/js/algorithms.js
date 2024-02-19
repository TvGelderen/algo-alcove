function initAglorithmsPage() {
    const sideNav = document.getElementById('side-nav');
    const algorithmContent = document.getElementById('algorithm-content')

    if (sideNav && algorithmContent) {
        const links = sideNav.querySelectorAll('.nav-link');

        const handleClick = event => {
            updateUrl(event.target.dataset.href);
            updateAlgorithmContent();
            updateLinks();
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

        const updateAlgorithmContent = () => {
            algorithmContent.setAttribute('hx-get', `/api${window.location.pathname}`);
            htmx.process(algorithmContent);
        }

        updateAlgorithmContent();
        updateLinks();

        for (const link of links) {
            link.addEventListener('click', handleClick);
        }

        window.addEventListener('popstate', () => {
            updateAlgorithmContent();
            updateLinks();
        });
    }
}

initAglorithmsPage();

function codeFetched(event) {
    code = event.target.querySelector('code');
    if (code) {
        hljs.highlightElement(code);
    }
}
