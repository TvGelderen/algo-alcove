let arr;
let visualizeContainer;
let sorting;
let reloading;

function initializeVisualize() {
    arr = [];
    arrElements = [];
    reloading = true;

    for (let i = 0; i < 100; i++) {
        arr.push(Math.floor(Math.random() * 101));
    }

    visualizeContainer = document.getElementById('sorting-visualize-container');
    if (!visualizeContainer) return;

    visualizeContainer.innerHTML = '';
    for (let i = 0; i < arr.length; i++) {
        const element = document.createElement('div');
        element.style.width = '1%';
        element.style.height = `${arr[i]}%`;
        element.style.background = '#c0caf5';
        element.style.margin = '0 0.5px';

        visualizeContainer.appendChild(element);
        arrElements.push(element);
    }
    
    if (!sorting) {
        reloading = false;
    }
}

initializeVisualize();

async function sort(event) {
    if (sorting) {
        return;
    }

    sorting = true;

    switch (event.target.dataset.algorithm) {
        case 'bubble-sort':
            await bubbleSort();
            break;
    }

    sorting = false;
}

async function bubbleSort() {
    var i, j, swapped;
    var len = arr.length;
 
    for (i = 0; i < len; i++) {
        swapped = false;
 
        for (j = 0; j < (len - i - 1); j++) {
            if (reloading) {
                reloading = false;
                return;
            }

            if (arr[j] > arr[j + 1]) {
                await swap(j, j + 1); 
                swapped = true;
            }
        }
 
        if (!swapped) {
            break;
        }
    }
}

async function swap(i, j) {
    select(i);
    select(j);

    await new Promise(r => setTimeout(r, 10));

    const temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;

    visualizeContainer.insertBefore(visualizeContainer.children[i], visualizeContainer.children[j + 1]); 
    visualizeContainer.insertBefore(visualizeContainer.children[j], visualizeContainer.children[i + 1]); 

    await new Promise(r => setTimeout(r, 10));

    deselect(i);
    deselect(j);
}

function select(i) {
    visualizeContainer.children[i].style.background = '#00ff00';
}

function deselect(i) {
    visualizeContainer.children[i].style.background = '#c0caf5';
}
