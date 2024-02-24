function initAglorithmsPage() {
    const sideNav = document.getElementById('side-nav');
    const algorithmContent = document.getElementById('algorithm-content')

    if (sideNav && algorithmContent) {
        const links = sideNav.querySelectorAll('.nav-link');

        const handleClick = event => {
            pushState(event.target.href);
            updateLinks();
        }

        const updateLinks = () => {
            for (const link of links) {
                if (window.location.href.includes(link.href)) {
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
            console.log('popstate');
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
