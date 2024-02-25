package seeds

import (
	"context"
	"fmt"

	"github.com/TvGelderen/algo-alcove/database"
	"github.com/TvGelderen/algo-alcove/models"
)

func seedSortingAlgorithms(db *database.Queries) {
	algorithms := []models.Algorithm{
		{
			TextId:   "insertion-sort",
			Name:     "Insertion Sort",
			Type:     models.AlgoritmhTypeSorting,
			Position: 1,
			Explanation: "<p>In the realm of sorting algorithms, Insertion Sort stands out for its simplicity and elegance. Like a meticulous dance of elements, this algorithm efficiently arranges a list by iteratively placing unsorted elements into their correct positions within a sorted section of the array.</p>" +
				"<p>Insertion Sort operates by dividing the array into two segments – the sorted and the unsorted. The algorithm takes one element at a time from the unsorted segment and inserts it into its proper place within the sorted segment. This process continues until the entire array is sorted. While Insertion Sort may not be the fastest algorithm for large datasets, its efficiency becomes apparent with partially ordered or small sets of data.</p>" +
				"<p>One of the key advantages of Insertion Sort lies in its adaptive nature. It performs exceptionally well when given partially sorted data, as it minimizes the number of comparisons and swaps. This makes it a favorable choice for scenarios where the dataset is likely to have already been partially ordered.</p>" +
				"<p>The step-by-step nature of Insertion Sort makes it a valuable tool for educational purposes. It provides a clear and intuitive introduction to the concepts of sorting algorithms, making it an ideal starting point for those new to the field.</p>" +
				"<p>Insertion Sort's simplicity is reflected not only in its algorithmic steps but also in its minimal space requirements. It is an in-place sorting algorithm, meaning it sorts the data without requiring additional memory, making it memory-efficient for applications with tight constraints.</p>",
		},
		{
			TextId:   "selection-sort",
			Name:     "Selection Sort",
			Type:     models.AlgoritmhTypeSorting,
			Position: 2,
			Explanation: "<p>In the intricate world of sorting algorithms, Selection Sort emerges as a modest yet effective choreographer of order. Its simple and intuitive approach to arranging elements makes it a valuable player in the array of sorting techniques, demonstrating that sometimes, simplicity can be the key to success.</p>" +
				"<p>Selection Sort operates by dividing the array into two segments – the sorted and the unsorted. The algorithm repeatedly selects the smallest (or largest) element from the unsorted segment and swaps it with the first unsorted element. This process continues until the entire array is sorted. While Selection Sort's time complexity is quadratic, making it less efficient for large datasets, its straightforward implementation and minimal resource requirements contribute to its appeal.</p>" +
				"<p>One of the notable characteristics of Selection Sort is its in-place nature. The algorithm sorts the data without requiring additional memory, making it memory-efficient and suitable for scenarios with limited resources.</p>" +
				"<p>Selection Sort's simplicity extends to its minimalistic code, making it an accessible choice for educational purposes. It serves as an excellent introductory sorting algorithm, allowing learners to grasp the fundamental concepts of sorting without the complexity of more advanced methods.</p>" +
				"<p>While not the fastest algorithm in terms of time complexity, Selection Sort shines in scenarios where the dataset is small or nearly sorted. Its uncomplicated nature makes it suitable for applications where a quick and straightforward sorting solution is preferred over more complex algorithms.</p>",
		},
		{
			TextId:   "merge-sort",
			Name:     "Merge Sort",
			Type:     models.AlgoritmhTypeSorting,
			Position: 3,
			Explanation: "<p>Merge Sort, a sophisticated algorithm revered for its efficiency and stability, orchestrates a harmonious symphony of divide-and-conquer principles to elegantly sort arrays. Developed by John von Neumann in 1945, this algorithm has stood the test of time, becoming a staple in the world of sorting.</p>" +
				"<p>At its core, Merge Sort divides the array into two halves, recursively sorts each half, and then merges them back together. This divide-and-conquer approach ensures that the algorithm breaks down the problem into more manageable sub-problems, tackling them with precision before seamlessly combining the results. The efficiency of Merge Sort lies in its consistent O(n log n) time complexity, making it a reliable choice for sorting large datasets.</p>" +
				"<p>One of the distinguishing features of Merge Sort is its stability. A stable sorting algorithm preserves the relative order of equal elements, ensuring that identical values maintain their initial positions in the sorted array. This stability is crucial in scenarios where maintaining the original order of equivalent elements is essential.</p>" +
				"<p>Merge Sort's elegance extends to its suitability for external sorting – a technique employed when the dataset is too large to fit into memory. By dividing the data into smaller chunks, sorting them individually, and then merging the sorted chunks, Merge Sort efficiently manages vast datasets without compromising on speed or stability.</p>" +
				"<p>The clarity of Merge Sort's recursive approach makes it an excellent teaching tool for understanding the principles of divide and conquer. It offers a visually intuitive representation of how breaking down a complex problem into smaller, more manageable parts can lead to a robust solution.</p>",
		},
		{
			TextId:   "quick-sort",
			Name:     "Quick Sort",
			Type:     models.AlgoritmhTypeSorting,
			Position: 4,
			Explanation: "<p>Enter Quick Sort, a swift and efficient maestro in the world of sorting algorithms. Renowned for its speed and versatility, Quick Sort employs a divide-and-conquer strategy to orchestrate a symphony of elements into a harmonious order.</p>" +
				"<p>At its core, Quick Sort operates by selecting a 'pivot' element from the array and partitioning the other elements into two sub-arrays – those less than the pivot and those greater. The algorithm then recursively applies the same process to each sub-array, eventually leading to a fully sorted array. This divide-and-conquer strategy, coupled with its clever pivot selection, bestows Quick Sort with an average time complexity of O(n log n), making it a formidable contender in the realm of sorting algorithms.</p>" +
				"<p>One of the standout features of Quick Sort is its adaptability. Unlike some sorting algorithms that struggle with partially sorted data, Quick Sort excels in such scenarios. Its ability to quickly rearrange and conquer partial disorder contributes to its efficiency, making it a popular choice in various applications.</p>" +
				"<p>The efficiency of Quick Sort extends beyond its time complexity; it is an in-place sorting algorithm, meaning it doesn't require additional memory to function. This makes it particularly advantageous in situations where memory constraints are a consideration.</p>" +
				"<p>Despite its remarkable speed, Quick Sort is not immune to the worst-case scenario of O(n^2) time complexity, which occurs when the pivot selection consistently leads to unbalanced partitions. However, this drawback is often mitigated by using randomized pivot selection or other optimization techniques.</p>" +
				"<p>In the grand performance of sorting algorithms, Quick Sort emerges as a swift and adaptable maestro. Its ability to efficiently navigate through both ordered and disordered arrays, coupled with its minimal memory requirements, solidifies Quick Sort's position as a key player in the diverse repertoire of sorting techniques.</p>",
		},
		{
			TextId:   "heap-sort",
			Name:     "Heap Sort",
			Type:     models.AlgoritmhTypeSorting,
			Position: 5,
			Explanation: "<p>In the grand symphony of sorting algorithms, Heap Sort takes center stage as a regal conductor, orchestrating a meticulous arrangement of elements with efficiency and precision. Born from the elegance of heap data structures, this algorithm elegantly achieves a sorted array through a series of systematic comparisons and transformations.</p>" +
				"<p>Heap Sort's foundation lies in the concept of a binary heap – a specialized tree-based data structure. The algorithm first transforms the array into a max-heap, where each parent node is greater than or equal to its child nodes. Once the max-heap is established, the algorithm repeatedly extracts the maximum element from the heap and places it at the end of the array. This process is repeated until the entire array is sorted.</p>" +
				"<p>The efficiency of Heap Sort stems from its optimal time complexity of O(n log n), making it a strong contender for large datasets. The inherent structure of the heap facilitates efficient traversal and comparison, ensuring a swift and effective sorting process.</p>" +
				"<p>One notable characteristic of Heap Sort is its stability, a quality that ensures equal elements maintain their original order in the sorted array. This stability can be crucial in scenarios where the initial sequence of identical elements holds significance.</p>" +
				"<p>Heap Sort's in-place nature further enhances its appeal. As it doesn't require additional memory, it is a memory-efficient solution for applications with constraints on space usage.</p>" +
				"<p>While Heap Sort may not be as intuitive as some other sorting algorithms, its systematic and efficient approach makes it a valuable tool in various domains. The regal conductor that is Heap Sort gracefully handles the complexities of large datasets, earning its place as a distinguished player in the realm of sorting algorithms.</p>",
		},
	}

	for _, algorithm := range algorithms {
		_, err := db.CreateAlgorithm(context.Background(), database.CreateAlgorithmParams{
			TextID:      algorithm.TextId,
			Name:        algorithm.Name,
			Type:        int16(algorithm.Type),
			Position:    algorithm.Position,
			Explanation: algorithm.Explanation,
		})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
