let arr;
let visualizeContainer;
let sorting;
let reloading;
let sortingDelay = 1;

function initializeVisualizeSort() {
    arr = [];
    arrElements = [];
    reloading = true;

    for (let i = 0; i < 100; i++) {
        arr.push(Math.floor(Math.random() * 101));
    }

    visualizeContainer = document.getElementById('sorting-visualize-container');
    if (!visualizeContainer) return;

    const sortingDelayInput = document.getElementById('sorting-delay');
    if (sortingDelayInput) {
        sortingDelayInput.addEventListener('change', () => {
            sortingDelay = parseInt(sortingDelayInput.value);
        })
    }

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

initializeVisualizeSort();

async function sort(event) {
    if (sorting) {
        return;
    }

    sorting = true;

    switch (event.target.dataset.algorithm) {
        case 'bubble-sort':
            await bubbleSort();
            break;
        case 'insertion-sort':
            await insertionSort();
            break;
        case 'selection-sort':
            await selectionSort();
            break;
        case 'merge-sort':
            await mergeSort(0, arr.length - 1);
            console.log(arr);
            break;
    }

    sorting = false;
}

async function bubbleSort() {
    let i, j, swapped;
    let len = arr.length;
 
    for (i = 0; i < len; i++) {
        swapped = false;
 
        for (j = 0; j < (len - i - 1); j++) {
            if (reloading) {
                reloading = false;
                return;
            }

            select(j);
            select(j + 1);

            await sleep();

            if (arr[j] > arr[j + 1]) {
                swap(j, j + 1); 
                swapped = true;
            }

            await sleep();

            deselect(j);
            deselect(j + 1);
        }
 
        if (!swapped) {
            break;
        }
    }
}

async function insertionSort()
{
    let i, j, key;
    let len = arr.length;

    for (i = 1; i < len; i++)
    {
        select(i);

        key = arr[i];
        j = i - 1;

        while (j >= 0 && arr[j] > key) 
        {
            if (reloading) {
                reloading = false;
                return;
            }

            select(j);

            await sleep();

            arr[j + 1] = arr[j];

            visualizeContainer.insertBefore(visualizeContainer.children[j + 1], visualizeContainer.children[j]); 

            await sleep();

            deselect(j);
            deselect(j + 1);

            j = j - 1;
        }

        arr[j + 1] = key;

        deselect(i);
    }
}

async function selectionSort()
{
    let i, j, min_idx;
    let len = arr.length;
  
    for (i = 0; i < len - 1; i++)
    {
        select(i);

        min_idx = i;
        for (j = i + 1; j < len; j++) {
            if (reloading) {
                reloading = false;
                return;
            }

            select(j);

            if (arr[j] < arr[min_idx]) {
                await sleep();
                deselect(i);
                min_idx = j;
            } else {
                deselect(j);
            }
        }
  
        await sleep();

        swap(min_idx, i);

        await sleep();

        deselect(i);
        deselect(min_idx);
    }
}

async function merge(l, m, r)
{
    let n1 = m - l + 1;
    let n2 = r - m;
 
    let L = new Array(n1);
    let R = new Array(n2);
 
    for (let i = 0; i < n1; i++) {
        L[i] = arr[l + i];
	}
    for (let j = 0; j < n2; j++) {
        R[j] = arr[m + 1 + j];
	}
 
    let i = 0;
    let j = 0;
    let k = l;
 
    while (i < n1 && j < n2) {
        if (L[i] <= R[j]) {
            arr[k] = L[i];
            visualizeContainer.insertBefore(visualizeContainer.children[l + i], visualizeContainer.children[k]);
            visualizeContainer.insertBefore(visualizeContainer.children[k], visualizeContainer.children[l + i]);
            i++;
        } else {
            arr[k] = R[j];
            visualizeContainer.insertBefore(visualizeContainer.children[m + 1 + j], visualizeContainer.children[k]); 
            visualizeContainer.insertBefore(visualizeContainer.children[k], visualizeContainer.children[m + 1 + j]); 
            j++;
        }
        k++;
    }
 
    while (i < n1) {
        arr[k] = L[i];
        visualizeContainer.insertBefore(visualizeContainer.children[l + i], visualizeContainer.children[k + 1]); 
        visualizeContainer.insertBefore(visualizeContainer.children[k], visualizeContainer.children[l + i]); 
        i++;
        k++;
    }

    while (j < n2) {
        arr[k] = R[j];
        visualizeContainer.insertBefore(visualizeContainer.children[m + 1 + j], visualizeContainer.children[k + 1]); 
        visualizeContainer.insertBefore(visualizeContainer.children[k], visualizeContainer.children[m + 1 + j]); 
        j++;
        k++;
    }
}

async function mergeSort(l, r){
    if (l >= r) {
        return;
    }

    let m = l + parseInt((r - l) / 2);

    await mergeSort(l, m);
    await mergeSort(m + 1, r);
    await merge(l, m, r);
}

const sleep = async () => await new Promise(r => setTimeout(r, sortingDelay / 2));

function swap(i, j) {
    const temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;

    visualizeContainer.insertBefore(visualizeContainer.children[i], visualizeContainer.children[j + 1]); 
    visualizeContainer.insertBefore(visualizeContainer.children[j], visualizeContainer.children[i + 1]); 
}

function select(i) {
    visualizeContainer.children[i].style.background = '#00ff00';
}

function deselect(i) {
    visualizeContainer.children[i].style.background = '#c0caf5';
}
