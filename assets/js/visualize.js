let arr = [];

function initializeVisualize() {
    for (let i = 0; i < 100; i++) {
        arr.push(Math.floor(Math.random() * 101));
    }

    const container = document.getElementById('sorting-visualize-container');
    if (!container) return;

    for (let i = 0; i < arr.length; i++) {
        const element = document.createElement('div');
        element.style.width = '1%';
        element.style.height = `${arr[i]}%`;
        element.style.background = '#c0caf5';
        element.style.margin = '0 0.5px';

        container.appendChild(element);
    }
}

initializeVisualize();
