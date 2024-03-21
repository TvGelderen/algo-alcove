let graphContainer = null;
let searching = false;
let startNode = 1;
let endNode = 10;
let searchDelay = 1;

function initializeVisualizePathFind() {
    graphContainer = document.getElementById('path-finding-visualize-container');
}

initializeVisualizePathFind();

async function search(event) {
    if (sorting) {
        return;
    }

    searching = true;

    switch (event.target.dataset.algorithm) {
        case 'breadth-first-search':
            await bfs();
            break;
    }

    searching = false;
}

async function bfs() {

}

const searchSleep = async () => await new Promise(r => setTimeout(r, searchDelay / 2));
