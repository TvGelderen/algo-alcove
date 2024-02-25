package seeds

import (
	"context"
	"fmt"

	"github.com/TvGelderen/algo-alcove/database"
	"github.com/TvGelderen/algo-alcove/models"
)

func seedBubbleSort(db *database.Queries) {
	algorithm := models.Algorithm{
		TextId:   "bubble-sort",
		Name:     "Bubble Sort",
		Type:     models.AlgoritmhTypeSorting,
		Position: 0,
		Explanation: "<p>Bubble Sort, though considered one of the simplest sorting algorithms, holds a unique place in the world of data organization. Devised as a basic introduction to sorting, it provides a straightforward approach to arranging elements in ascending or descending order.</p>" +
			"<p>At its core, Bubble Sort works by repeatedly stepping through the list, comparing adjacent elements, and swapping them if they are in the wrong order. The process is akin to the bubbling up of larger elements to their correct positions. While conceptually clear and easy to implement, Bubble Sort's simplicity comes at a cost - its time complexity is quadratic, making it less efficient for large datasets compared to more sophisticated algorithms.</p>" +
			"<p>The algorithm's name stems from the way smaller elements gradually rise to their proper places like bubbles in a liquid. Despite its elementary nature, Bubble Sort is often used in educational settings to introduce the concept of sorting algorithms. It serves as a hands-on demonstration of the basic principles involved in organizing data.</p>" +
			"<p>One notable feature of Bubble Sort is its stability. In a stable sorting algorithm, the relative order of equal elements remains unchanged, and Bubble Sort excels in preserving this order. This makes it suitable for scenarios where maintaining the initial order of identical elements is crucial.</p>",
	}

	id, err := models.AddAlgorithmToDB(db, algorithm)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	algorithmCode := models.AlgorithmCode{
		AlgorithmId: id,
		Language:    "javascript",
		Code: `function bubbleSort(arr) {
    var i, j;
    var len = arr.length;
    var isSwapped = false;
 
    for (i = 0; i < len; i++) {
        isSwapped = false;
 
        for (j = 0; j < (len - i - 1); j++) {
            if (arr[j] > arr[j + 1]) {
                var temp = arr[j]
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
                isSwapped = true;
            }
        }
 
        if (!isSwapped) {
            break;
        }
    }
}`,
	}

	err = models.AddAlgorithmCodeToDB(db, algorithmCode)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

    err = db.UpdateAlgorithmHasCode(context.Background(), database.UpdateAlgorithmHasCodeParams{
        ID: id,
        HasCode: true,
    })
}
