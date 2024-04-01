let graphContainer = null;
let edgeContainer = null;
let searching = false;
let nodes = [];
let startNode = 1;
let endNode = 10;
let searchDelay = 1;

// -o-o-
// o-o-o
// o-o-o
// -o-o-

const graph = [
    [0, 0, 1, 0, 0, 0, 2, 0, 0],
    [3, 0, 0, 0, 4, 0, 0, 0, 5],
    [6, 0, 0, 0, 7, 0, 0, 0, 8],
    [0, 0, 9, 0, 0, 0, 10, 0, 0],
];

class Node {
    constructor(id, x, y, element) {
        this.id = id;
        this.x = x;
        this.y = y;
        this.element = element;
        this.edges = [];
    }

    addEdge(node) {
        this.edges.push(node);
    }
}

function initializeVisualizePathFind() {
    graphContainer = document.getElementById('path-finding-visualize-container');
    edgeContainer = document.getElementById('edge-container');
    if (!graphContainer || !edgeContainer) return;

    graphContainer.style.gridTemplateColumns = `repeat(${graph[0].length}, 50px)`;
    graphContainer.style.rowGap = '100px';
    
    for (let y = 0; y < graph.length; y++) {
        for (let x = 0; x < graph[y].length; x++) {
            const el = document.createElement('div');
            if (graph[y][x] !== 0) {
                el.classList.add('node');
                el.id = nodes.length;
                el.innerHTML = `<span>${nodes.length + 1}</span>`;
                const node = new Node(nodes.length, x, y, el);
                nodes.push(node);
            }
            graphContainer.appendChild(el);
        }
    }

    const containerRect = graphContainer.getBoundingClientRect();

    for (let i = 0; i < nodes.length; i++) {
        for (let j = 0; j < nodes.length; j++) {
            if (i === j || 
                nodes[i].y <= nodes[j].y - 2 || 
                nodes[i].y >= nodes[j].y + 2 ||
                (nodes[i].y === nodes[j].y && Math.abs(i - j) > 1)
            ) {
                continue;
            }

            if (!nodes[i].edges.includes(nodes[j]) && Math.random() > 0.5) {
                const edge = document.createElementNS('http://www.w3.org/2000/svg', 'line');
                console.log(nodes[i].element.getBoundingClientRect());
                console.log(containerRect);

                const x1 = nodes[i].element.getBoundingClientRect().x - containerRect.left + 25;
                const y1 = nodes[i].element.getBoundingClientRect().y - containerRect.top + 25;
                const x2 = nodes[j].element.getBoundingClientRect().x - containerRect.left + 25;
                const y2 = nodes[j].element.getBoundingClientRect().y - containerRect.top + 25;

                edge.setAttribute("x1", `${x1}`);
                edge.setAttribute("y1", `${y1}`);
                edge.setAttribute("x2", `${x2}`);
                edge.setAttribute("y2", `${y2}`);
                edgeContainer.appendChild(edge);

                nodes[i].addEdge(nodes[j]);
                nodes[j].addEdge(nodes[i]);
            }
        }
    }
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
