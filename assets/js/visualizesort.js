let arr;
let sortContainer;
let sorting;
let reloading;
let sortingDelay = 1;

function initVisualizeSort() {
    arr = [];
    arrElements = [];
    reloading = true;

    for (let i = 0; i < 100; i++) {
        arr.push(Math.floor(Math.random() * 100) + 1);
    }

    sortContainer = document.getElementById('sorting-visualize-container');
    if (!sortContainer) return;

    const sortingDelayInput = document.getElementById('sorting-delay');
    if (sortingDelayInput) {
        sortingDelayInput.addEventListener('change', () => {
            sortingDelay = parseInt(sortingDelayInput.value);
        })
    }

    sortContainer.innerHTML = '';
    for (let i = 0; i < arr.length; i++) {
        const element = document.createElement('div');
        element.classList.add('sort-bar');
        element.style.height = `${arr[i]}%`;

        sortContainer.appendChild(element);
        arrElements.push(element);
    }

    if (!sorting) {
        reloading = false;
    }
}

initVisualizeSort();

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
            break;
        case 'quick-sort':
            await quickSort(0, arr.length - 1);
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
            await sortSleep();

            if (arr[j] > arr[j + 1]) {
                swap(j, j + 1);
                swapped = true;
            }

            await sortSleep();
            deselect();
        }

        if (!swapped) {
            break;
        }
    }
}

async function insertionSort() {
    let i, j, key;
    let len = arr.length;

    for (i = 1; i < len; i++) {
        select(i);

        key = arr[i];
        j = i - 1;

        while (j >= 0 && arr[j] > key) {
            if (reloading) {
                reloading = false;
                return;
            }

            select(j);
            await sortSleep();

            arr[j + 1] = arr[j];

            sortContainer.insertBefore(sortContainer.children[j + 1], sortContainer.children[j]);

            await sortSleep();
            deselect();

            j = j - 1;
        }

        arr[j + 1] = key;

        deselect();
    }
}

async function selectionSort() {
    let i, j, min_idx;
    let len = arr.length;

    for (i = 0; i < len - 1; i++) {
        min_idx = i;

        for (j = i + 1; j < len; j++) {
            if (reloading) {
                reloading = false;
                return;
            }

            if (arr[j] < arr[min_idx]) {
                min_idx = j;
            }
        }

        select(i);
        select(min_idx);
        await sortSleep();

        swap(min_idx, i);

        await sortSleep();
        deselect();
    }
}

async function merge(l, m, r) {
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
        select(l + i);
        select(m + 1 + j);
        select(k);
        await sortSleep();

        if (L[i] <= R[j]) {
            arr[k] = L[i];
            if (k < l + 1) {
                sortContainer.insertBefore(sortContainer.children[l + i], sortContainer.children[k]);
            }
            i++;
        } else {
            arr[k] = R[j];
            if (k < m + 1 + j) {
                sortContainer.insertBefore(sortContainer.children[m + 1 + j], sortContainer.children[k]);
            }
            j++;
        }
        k++;

        await sortSleep();
        deselect();
    }

    while (i < n1) {
        arr[k] = L[i];
        if (k < l + 1) {
            select(l + i);
            select(k);
            await sortSleep();

            sortContainer.insertBefore(sortContainer.children[l + i], sortContainer.children[k]);

            await sortSleep();
            deselect();
        }
        i++;
        k++;
    }

    while (j < n2) {
        arr[k] = R[j];
        if (k < m + 1 + j) {
            select(m + 1 + j);
            select(k);
            await sortSleep();

            sortContainer.insertBefore(sortContainer.children[m + 1 + j], sortContainer.children[k]);

            await sortSleep();
            deselect();
        }
        j++;
        k++;
    }
}

async function mergeSort(l, r) {
    if (l >= r) {
        return;
    }

    let m = l + parseInt((r - l) / 2);

    if (reloading) {
        reloading = false;
        return;
    }

    await mergeSort(l, m);
    await mergeSort(m + 1, r);
    await merge(l, m, r);
}

async function partition(low, high) {
    let pivot = arr[high];
    let i = low - 1;

    for (let j = low; j < high; j++) {
        if (arr[j] < pivot) {
            i++;
            select(i);
            select(j);
            await sortSleep();

            swap(i, j);

            await sortSleep();
            deselect();
        }
    }

    select(high);
    select(i + 1);
    await sortSleep();

    swap(high, i + 1);

    await sortSleep();
    deselect();

    return i + 1;
}

async function quickSort(low, high) {
    if (low < high) {
        let pi = await partition(low, high);

        await quickSort(low, pi - 1);
        await quickSort(pi + 1, high);
    }
}

const sortSleep = async () => await new Promise(r => setTimeout(r, sortingDelay / 2));

function swap(i, j) {
    const temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;

    if (i < j) {
        sortContainer.insertBefore(sortContainer.children[i], sortContainer.children[j]);
        sortContainer.insertBefore(sortContainer.children[j], sortContainer.children[i]);
    } else {
        sortContainer.insertBefore(sortContainer.children[j], sortContainer.children[i]);
        sortContainer.insertBefore(sortContainer.children[i], sortContainer.children[j]);
    }
}

function select(i) {
    sortContainer.children[i].classList.add('selected');
}

function deselect() {
    const selected = sortContainer.querySelectorAll('.selected');
    for (const el of selected) {
        el.classList.remove('selected');
    }
}
