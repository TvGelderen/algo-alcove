package seeds

import (
	"context"
	"fmt"

	"github.com/TvGelderen/algo-alcove/database"
	"github.com/TvGelderen/algo-alcove/models"
)

func seedPathFindingAlgorithms(db *database.Queries) {
	algorithms := []models.Algorithm{
		{
			TextId:   "breadth-first-search",
			Name:     "Breadth-First Search",
			Type:     models.AlgoritmhTypePathFinding,
			Position: 0,
			Explanation: "<p>Breadth-First Search (BFS) stands as a foundational algorithm in computer science, renowned for its ability to systematically explore and traverse graphs and trees layer by layer. Unlike its counterpart, Depth-First Search, BFS prioritizes breadth, making it an invaluable tool in numerous applications, ranging from network routing to puzzle solving.</p>" +
				"<p>The core concept of BFS involves exploring all neighbors of a given node before moving on to the next level of nodes. This systematic approach ensures that BFS uncovers the shortest path from a source node to all other nodes in an unweighted graph. A queue is typically employed to maintain the order of exploration, ensuring nodes are visited in the order they are discovered.</p>" +
				"<p>One of the notable applications of BFS is its role in finding the shortest path in maze-solving scenarios. By exploring adjacent cells layer by layer, BFS guarantees that the shortest path is identified before moving on to longer routes. This characteristic makes BFS particularly useful in scenarios where finding the quickest route is of paramount importance.</p>" +
				"<p>BFS is also crucial in network analysis, helping to determine the optimal routing of information. By examining neighboring nodes at each layer, BFS aids in identifying the most efficient path for data transmission, a critical aspect in the design and optimization of computer networks.</p>" +
				"<p>Despite its efficiency in certain scenarios, BFS may exhibit limitations in terms of memory usage. Storing all nodes at a particular level in a queue can consume significant memory, making it less suitable for graphs with a vast number of nodes. Additionally, BFS may not perform optimally in scenarios with weighted graphs, as it doesn't account for edge weights during traversal.</p>" +
				"<p>In conclusion, Breadth-First Search stands as a powerful algorithm, excelling in scenarios where finding the shortest path is paramount. Its systematic exploration and layer-by-layer approach make BFS a valuable tool in various applications, offering a balanced mix of simplicity and effectiveness in graph traversal.</p>",
		},
		{
			TextId:   "depth-first-search",
			Name:     "Depth-First Search",
			Type:     models.AlgoritmhTypePathFinding,
			Position: 1,
			Explanation: "<p>Depth-First Search (DFS) is a fundamental algorithm in computer science that explores and traverses graphs and trees by diving as deeply as possible into a branch before backtracking. This algorithm is widely employed in various applications, from maze-solving to network analysis, owing to its simplicity and versatility.</p>" +
				"<p>DFS operates on the principle of recursion or utilizing a stack to keep track of nodes to be visited. Starting from a designated node, DFS explores as far as possible along each branch before backtracking. This process continues until all nodes have been visited or a specific condition is met.</p>" +
				"<p>One of the key advantages of DFS lies in its memory efficiency. Unlike Breadth-First Search, which requires a queue to store all nodes at a particular level, DFS only needs to store a single path, making it well-suited for scenarios with limited memory resources.</p>" +
				"<p>The algorithm is commonly used to solve problems involving connectivity, such as identifying connected components in a graph. By traversing the graph in a depth-first manner, DFS can efficiently identify groups of interconnected nodes.</p>" +
				"<p>Furthermore, DFS plays a pivotal role in topological sorting, a critical task in project scheduling and task prioritization. By employing DFS, one can establish a linear ordering of nodes that preserves the precedence relationships in a directed acyclic graph (DAG).</p>" +
				"<p>Despite its many strengths, DFS does have limitations. It may not always find the shortest path between two nodes in a weighted graph, as it doesn't consider edge weights during traversal. Additionally, DFS can encounter issues with infinite loops if not carefully implemented, emphasizing the need for proper termination conditions.</p>" +
				"<p>In conclusion, Depth-First Search is a powerful algorithm with widespread applications, offering a balance between simplicity and efficiency. Its ability to explore intricate structures, coupled with its memory-friendly nature, makes DFS a valuable tool in the toolkit of any computer scientist or programmer.</p>",
		},
		{
			TextId:   "dijkstras-algorithm",
			Name:     "Dijkstra's Algorithm",
			Type:     models.AlgoritmhTypePathFinding,
			Position: 2,
			Explanation: "<p>Dijkstra's algorithm, conceived by Dutch computer scientist Edsger W. Dijkstra in 1956, stands as a landmark contribution to graph theory and is widely utilized for finding the shortest path in weighted graphs. This algorithm has played a pivotal role in diverse applications, from network routing to robotics and geographical information systems.</p>" +
				"<p>The primary objective of Dijkstra's algorithm is to determine the shortest path from a designated source node to all other nodes in a graph with non-negative edge weights. It employs a greedy approach, iteratively selecting the node with the smallest tentative distance and updating its neighbors' distances accordingly.</p>" +
				"<p>The algorithm maintains two sets of nodes: one with tentative distances yet to be finalized and another with the settled nodes. It starts by setting the tentative distance to the source node as zero and all other nodes as infinity. In each iteration, it selects the node with the smallest tentative distance, explores its neighbors, and updates their distances if a shorter path is found. This process continues until all nodes are settled.</p>" +
				"<p>Dijkstra's algorithm is renowned for its optimality in finding the shortest path. Its guarantee of optimality is attributed to the fact that once a node is settled, its distance is final and will not be revisited, ensuring the algorithm always selects the shortest path at each step.</p>" +
				"<p>The application of Dijkstra's algorithm extends across various domains. In computer networking, it aids in determining the optimal route for data transmission. In robotics, it facilitates path planning for autonomous vehicles, ensuring they navigate efficiently through their environment.</p>" +
				"<p>While Dijkstra's algorithm excels in scenarios with non-negative edge weights, it may encounter challenges in graphs with negative weights. For such cases, alternative algorithms like the Bellman-Ford algorithm are more suitable. Nevertheless, Dijkstra's algorithm remains a cornerstone in the toolkit of algorithms, providing an effective and reliable method for finding the shortest path in a variety of practical scenarios.</p>",
		},
		{
			TextId:   "a-star-algorithm",
			Name:     "A* Algorithm",
			Type:     models.AlgoritmhTypePathFinding,
			Position: 3,
			Explanation: "<p>The A* algorithm, a popular pathfinding and graph traversal algorithm, has proven to be a versatile tool in artificial intelligence, robotics, and game development. Created by Peter Hart, Nils Nilsson, and Bertram Raphael in 1968, A* combines the best of both breadth-first search and greedy algorithms, offering a reliable solution for finding the shortest path in weighted graphs.</p>" +
				"<p>At its core, A* uses a heuristic approach to estimate the cost of reaching the goal from a particular node. This heuristic function, often denoted as h(n), guides the algorithm in selecting the most promising nodes, enabling it to explore the graph efficiently. A* considers both the cost of reaching a node from the start (g(n)) and the estimated cost to the goal (h(n)), aiming to minimize the sum f(n) = g(n) + h(n).</p>" +
				"<p>One of the distinguishing features of A* is its optimality and efficiency in finding the shortest path. By incorporating the heuristic function, A* intelligently narrows down its search, prioritizing nodes that are more likely to lead to an optimal solution. This makes A* particularly well-suited for applications where finding the most efficient route is crucial, such as in GPS navigation systems.</p>" +
				"<p>A* has found widespread use in various fields, including robotics and video game development. In robotics, A* helps autonomous systems plan their paths, avoiding obstacles and reaching destinations with minimum effort. In game development, A* is employed for character movement, ensuring non-player characters navigate game environments effectively.</p>" +
				"<p>Despite its strengths, A* is not without considerations. The quality of the heuristic function greatly influences the algorithm's performance, and poorly designed heuristics can lead to suboptimal solutions. Additionally, A* may struggle in scenarios with dynamic environments or changing costs, as it doesn't adapt to real-time changes.</p>" +
				"<p>Despite its strengths, A* is not without considerations. The quality of the heuristic function greatly influences the algorithm's performance, and poorly designed heuristics can lead to suboptimal solutions. Additionally, A* may struggle in scenarios with dynamic environments or changing costs, as it doesn't adapt to real-time changes.</p>",
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
