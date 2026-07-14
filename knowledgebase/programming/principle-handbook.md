# Data Structures
## Array and String

### Array vs String
Array: fixed or dynamic collection of same-type elements, contiguous memory, O(1) index access. String: character array, often immutable, same big-O rules as arrays.

### Time Complexity
| Op | Array | String |
|---|---|---|
| Access | O(1) | O(1) |
| Search | O(n) | O(n) |
| Insert/Remove (end) | O(1) | O(n) |
| Insert/Remove (middle) | O(n) | O(n) |
| Substring search | — | O(n·m), O(n+m) with KMP |

### Key Concepts
- **Contiguous memory:** elements stored sequentially, enables O(1) random access
- **Fixed size (array) vs dynamic (list):** arrays fixed at creation; dynamic arrays resize when capacity exceeded (amortized O(1) append)
- **Immutability (strings):** strings are immutable in most languages; concatenation creates new string (O(n))
- **String encoding:** ASCII, UTF-8, UTF-16; affects character access and length calculation
- **In-place operations:** modifying array/string without extra space (space optimization trade-off)
- **Two-pointer patterns:** for palindromes, reversals, container problems
## Linked List
### Basics
A linked list is a linear structure of nodes, each holding data plus a pointer to the next node, stored in non-contiguous memory.

**Types:**
- **Singly:** each node has a next pointer only
- **Doubly:** each node has next and prev pointers
- **Circular:** tail points back to head

### Time Complexity
| Op | Singly | Doubly |
|---|---|---|
| Access/Search | O(n) | O(n) |
| Insert/Delete at head | O(1) | O(1) |
| Insert/Delete at tail | O(n)* | O(1) with tail pointer |
| Insert/Delete in middle | O(n) to find + O(1) to link | same |

*Can be O(1) with tail pointer if you already have reference to tail.

### Key Concepts
- **Non-contiguous memory:** nodes scattered in memory, no random access by index
- **Pointer-based navigation:** traverse by following pointers, cost of traversal O(n)
- **Dynamic allocation:** no fixed size, can grow/shrink freely (except circular which wraps)
- **Node structure:** data + pointer(s); doubly-linked adds space overhead but enables backward traversal
- **Reference tracking:** must maintain head pointer; losing it means losing entire list

### Key Techniques
- **Fast & slow pointers:** cycle detection (Floyd's), find middle node, find nth-from-end
- **Dummy head node:** simplifies edge cases when head itself changes (insert/delete at start)
- **Reverse in place:** track prev/curr/next pointers; common for "reverse list" or "reverse in groups"
- **Two-pointer merge:** merging two sorted lists
- **Runner technique:** detect palindrome list (reverse second half, compare)

### Gotchas
- Losing the head reference — always keep a pointer to head before mutating
- Null pointer checks before dereferencing `.next`
- Off-by-one when counting nodes for "nth from end" problems
- Memory overhead per node (extra pointer) vs array
- No O(1) random access — can't index like `list[i]`
## Stack
### Basics
A stack is a linear structure following Last-In-First-Out (LIFO) order — the last element added is the first one removed.

**Core operations:**
- **Push:** add element to top
- **Pop:** remove and return top element
- **Peek/Top:** view top element without removing
- **isEmpty:** check if stack is empty

### Time Complexity
| Op | Stack (array or linked-list based) |
|---|---|
| Push | O(1) |
| Pop | O(1) |
| Peek/Top | O(1) |
| Search | O(n) |

### Key Concepts
- **LIFO order:** last element added is first one out; opposite of queue
- **Top element:** all operations work on the top only; middle elements not accessible
- **Implementation choices:** array-based or linked-list based; affects resizing behavior and overhead
- **Call stack:** underlies recursion and function calls; system stack for tracking execution context
- **Undo/redo patterns:** natural fit for tracking state history in applications

### Gotchas
- Popping or peeking an empty stack — always check `isEmpty` first
- Array-based stack may need resizing (amortized O(1), occasional O(n))
- Linked-list based stack avoids resizing but has extra pointer overhead per node
- Not suited for random access — no indexing like arrays
## Matrix
### Basics
A matrix is a 2D array of m rows and n columns, stored in contiguous memory. Elements accessed by [row][col] indices. Stored in row-major order (C/C++/Java) or column-major order (Fortran/MATLAB).

### Time Complexity
| Op | Complexity |
|---|---|
| Access by index | O(1) |
| Search | O(m·n) |
| Insert/Delete | O(m·n) (resize and shift) |

### Key Techniques
- **Traversals:** row-wise, column-wise, diagonal, spiral, zigzag patterns
- **In-place rotation:** 90/180 degrees using layer-by-layer or transpose-then-flip approach
- **DFS/BFS on grid:** treating matrix as graph, finding connected components, island problems
- **Dynamic programming:** path counting, shortest path, max/min problems on grids
- **Matrix flipping:** horizontal or vertical reflection
- **Matrix multiplication:** O(m·n·p) for m×n and n×p matrices

### Gotchas
- Index bounds: rows [0, m-1], cols [0, n-1] — off-by-one errors common
- Row-major vs column-major storage affects cache performance and iteration order
- In-place operations may overwrite needed values — track carefully or use temporary storage
- Visiting cells in 2D DP/BFS — must mark visited to avoid infinite loops
- Rectangular matrices (m ≠ n) require explicit bounds checking for both dimensions
## Hashmap
### Basics
A hashmap is an unordered collection of key-value pairs. Uses a hash function to map keys to array indices, providing fast average-case lookups. Built on a hash table (array) with collision resolution.

**Collision handling approaches:**
- **Chaining:** each bucket stores linked list of entries with same hash
- **Open addressing:** find another empty slot (linear probing, quadratic probing, double hashing)

### Time Complexity
| Op | Average | Worst Case |
|---|---|---|
| Insert | O(1) | O(n) |
| Delete | O(1) | O(n) |
| Search/Lookup | O(1) | O(n) |
| Space | — | O(n) |

*Worst case occurs with poor hash function or high collision rate*

### Key Concepts
- **Hash function:** maps keys uniformly to array indices [0, capacity-1]
- **Load factor:** ratio of entries to capacity; high load factor triggers resizing
- **Resizing/Rehashing:** double capacity when load factor exceeds threshold, rehash all entries (O(n))
- **Immutable keys:** strings, numbers, tuples; mutable keys (lists, dicts) cause lookup failures

### Gotchas
- Hash collisions reduce performance — quality hash function is critical
- Load factor too high causes O(n) degradation; too low wastes memory
- Mutable objects as keys break lookup after modification
- Resizing is O(n) but amortized to O(1) per insert
- Order not guaranteed in most implementations (Python 3.7+ preserves insertion order)
- Memory overhead for unused slots, especially with low load factor
## Heap
### Basics
A heap is a complete binary tree that satisfies the heap property. **Min-heap:** each parent ≤ its children. **Max-heap:** each parent ≥ its children. Usually implemented as an array where index i has children at 2i+1 and 2i+2, parent at (i-1)/2.

### Time Complexity
| Op | Complexity |
|---|---|
| Insert | O(log n) |
| Delete min/max | O(log n) |
| Peek/Get min/max | O(1) |
| Build heap from array | O(n) |
| Search | O(n) |

### Key Concepts
- **Complete binary tree:** all levels filled except possibly last, filled left to right; enables array representation
- **Heap property:** only parent-child relationship guaranteed, not full sorting
- **Array indexing:** implicit structure; parent(i) = (i-1)/2, left(i) = 2i+1, right(i) = 2i+2
- **Heapify:** restore heap property after mutation (heapify-up or heapify-down)
- **Priority queue:** natural application; extract highest/lowest priority efficiently

### Key Techniques
- **Insert:** add element at end, bubble up (heapify-up) to restore property
- **Delete min/max:** remove root, move last element to root, bubble down (heapify-down)
- **Heapify-up (sift-up):** move element up while it violates heap property with parent
- **Heapify-down (sift-down):** move element down, swap with smaller/larger child, repeat
- **Build heap:** O(n) construction by heapifying from bottom-up (faster than n inserts)
- **Heap sort:** build max-heap, repeatedly extract max to get sorted array

### Gotchas
- Not a sorted structure — only guarantees top element is min/max, rest unordered
- Array indexing: 0-based vs 1-based affects parent/child formulas; off-by-one common
- Deletion removes root only, not arbitrary element (use different strategy for that)
- Build heap O(n) is better than inserting n times O(n log n)
- Heap stored as array means no random access by key; must scan O(n) to find specific value

# Tree & Graph
## Binary Tree
### General
A binary tree is a hierarchical structure where each node has at most 2 children (left and right). A null tree is empty; otherwise it has a root with a left and/or right subtree.

**Types:**
- **Full binary tree:** every node has 0 or 2 children
- **Complete binary tree:** all levels filled except possibly last, filled left to right
- **Perfect binary tree:** all internal nodes have 2 children and all leaves at same level
- **Balanced binary tree:** height difference between left and right subtrees ≤ 1

### Time Complexity (General Tree)
| Op | Complexity |
|---|---|
| Traversal (in/pre/post-order) | O(n) |
| Search | O(n) |
| Insert/Delete | O(n) to find + O(1) to link |

### Key Concepts
- **Recursive structure:** each subtree is itself a binary tree; enables recursive solutions
- **Depth vs height:** depth = distance from root to node; height = distance from node to furthest leaf
- **Level-order:** nodes at same distance from root
- **Balanced vs unbalanced:** affects search/insert/delete performance
- **Tree properties:** node count, height bounds, relationship between levels

### Key Techniques - Traversals
- **In-order (left, root, right):** gives sorted sequence for BST
- **Pre-order (root, left, right):** useful for copying tree, creating new tree with same structure
- **Post-order (left, right, root):** useful for deletion, computing node values bottom-up
- **Level-order (BFS):** breadth-first, uses queue, processes by levels

### Gotchas
- Recursive solutions hit stack overflow on deep trees (use iterative with explicit stack)
- Off-by-one in height/depth calculations
- Null children must be checked before recursing
- Confusing in-order vs pre-order vs post-order traversal order
- Tree construction from traversals requires at least 2 traversals (e.g., in-order + pre-order)

### Breadth-first Search
Traverse tree level-by-level using a queue. Each level processed before moving to next level.

**Time/Space:** O(n) time, O(w) space where w = maximum width (most nodes at same level)

**Implementation:** use queue, enqueue root, repeatedly dequeue and enqueue children, collect values

### Binary Search Tree
A binary tree where for each node: all values in left subtree < node value < all values in right subtree. This property holds recursively for all subtrees.

| Op | Average | Worst (Unbalanced) |
|---|---|---|
| Search | O(log n) | O(n) |
| Insert | O(log n) | O(n) |
| Delete | O(log n) | O(n) |
| Find min/max | O(log n) | O(n) |
| In-order traversal | O(n) | O(n) |

**Key Properties:**
- In-order traversal yields sorted sequence
- Left subtree height ≠ right subtree height in general (unbalanced possible)
- No duplicate handling (implementation-dependent)

**Key Operations:**
- **Search:** compare with root, go left if smaller, right if larger
- **Insert:** find correct leaf position, add new node maintaining BST property
- **Delete:** three cases — leaf (remove), one child (replace with child), two children (find in-order successor/predecessor)
- **Find min:** go left until no left child
- **Find max:** go right until no right child

**Gotchas:**
- Unbalanced trees degrade to O(n) performance (use AVL or Red-Black for guaranteed balance)
- Deletion with two children requires careful handling of successor/predecessor
- In-order traversal requires going left-root-right carefully to avoid revisiting nodes
- Duplicate handling must be specified (allow, reject, or count)
## Graph
### General
A graph is a data structure consisting of vertices (nodes) and edges connecting them. Graphs can be directed or undirected, weighted or unweighted, connected or disconnected.

**Representations:**
- **Adjacency list:** map/dict of vertex to list of neighbors; space O(V+E), good for sparse graphs
- **Adjacency matrix:** 2D array where matrix[i][j] = weight/presence of edge; space O(V²), good for dense graphs
- **Edge list:** list of (u, v, weight) tuples; space O(E)

**Graph types:**
- **Directed vs undirected:** edges have direction or not
- **Weighted vs unweighted:** edges have weights/costs or not
- **Cyclic vs acyclic (DAG):** contains cycles or directed acyclic
- **Connected/strongly connected:** can reach all nodes; all pairs mutually reachable (directed)

### Time Complexity
| Op | Adjacency List | Adjacency Matrix |
|---|---|---|
| Space | O(V + E) | O(V²) |
| Add edge | O(1) | O(1) |
| Remove edge | O(degree) | O(1) |
| Check edge (u,v) | O(degree) | O(1) |
| DFS/BFS | O(V + E) | O(V²) |

### Key Concepts
- **Vertex/Node and Edge:** fundamental units; edge connects two vertices
- **In-degree/Out-degree:** number of incoming/outgoing edges (directed graphs)
- **Path and cycle:** sequence of edges; cycle returns to starting vertex
- **Connected component:** maximal subset of vertices where all pairs have paths
- **Strongly connected component (SCC):** directed graph; all vertices mutually reachable
- **Bipartite graph:** vertices partitionable into two sets with edges only between sets
- **Topological order:** valid ordering of vertices in DAG (dependent ordering)

### Key Techniques
- **DFS (Depth-First Search):** recursive or stack-based, explores deep before backtracking
- **BFS (Breadth-First Search):** queue-based, explores all neighbors before going deeper
- **Shortest path:** Dijkstra's (weighted, non-negative), Bellman-Ford (weighted, handles negative), BFS (unweighted)
- **Minimum spanning tree:** Kruskal's (sort edges, union-find), Prim's (grow tree from vertex)
- **Topological sort:** Kahn's algorithm (in-degree based) or DFS post-order
- **Cycle detection:** DFS color marking (white/gray/black), union-find for undirected
- **Connected components:** DFS/BFS from unvisited, find SCCs with Kosaraju's or Tarjan's

### Gotchas
- Directed vs undirected changes algorithm logic (edge has one direction vs two)
- Self-loops and multi-edges must be handled explicitly
- Weighted graphs require priority queue for shortest path (Dijkstra's)
- Negative edge weights break Dijkstra's; use Bellman-Ford instead
- DFS/BFS must track visited to avoid infinite loops
- Adjacency matrix O(V²) space is prohibitive for large sparse graphs

### Breadth-first Search
Traverse graph level-by-level exploring all neighbors before going deeper. Finds shortest path in unweighted graphs.

**Algorithm:** enqueue source, mark visited, repeatedly dequeue and enqueue unvisited neighbors

**Time:** O(V + E) — visit each vertex and edge once
**Space:** O(V) — queue holds at most V vertices

**Applications:**
- Shortest path in unweighted graphs
- Connected components
- Level-order exploration
- Bipartiteness checking (2-coloring)

# Algorithmic Techniques
## Two Pointers
### Basics
Two pointers is a technique using two pointers to iterate through a data structure (usually sorted arrays, linked lists, or strings) to solve problems in a single or multiple passes. Pointers move based on problem conditions, avoiding nested loops.

**Common patterns:**
- **Converging pointers:** one starts at beginning, one at end; move towards center
- **Same direction, different speeds:** both start at beginning but move at different rates (e.g., slow/fast)
- **Sliding window variant:** both move in same direction maintaining a window or range

### Time Complexity
| Scenario | Complexity |
|---|---|
| Converging (sorted array) | O(n) |
| Same direction (cycle detection) | O(n) |
| Two pointers + sorting | O(n log n) |

**Space:** O(1) typically, or O(n) if output array needed

### Key Concepts
- **Sorted assumption:** most effective on sorted data; pointers can skip sections intelligently
- **Single pass:** avoids nested O(n²) loops by using pointer positions
- **Boundary conditions:** handle correctly when pointers meet or reach ends
- **Pointer movement logic:** increment/decrement based on comparison results
- **Opposite directions:** converging pointers work well for symmetric problems (palindromes, containers)

### Key Techniques
- **Converging pointers:** left = 0, right = n-1, move based on comparison (e.g., sum of pair)
- **Slow and fast:** both start at beginning; slow increments by 1, fast by 2 (cycle detection, find middle)
- **Two sorted arrays merge:** pointer for each array, advance smaller value
- **Remove duplicates:** two pointers to track write position and read position
- **Palindrome check:** compare from ends moving inward
- **Container with most water:** two pointers find max area between lines
- **Reverse in place:** swap values moving from ends towards center

### Common Problems
- **Two sum:** sorted array, find two numbers summing to target
- **Container with most water:** find max area between two lines
- **Trapping rain water:** accumulate water based on heights
- **Palindrome check:** validate string is same forwards/backwards
- **Merge sorted arrays:** combine two sorted lists
- **Remove element:** move unwanted elements to end
- **Intersection/union:** find common/distinct elements between arrays

### Gotchas
- Requires sorted or specific structure; doesn't work on arbitrary arrays
- Off-by-one at boundaries when pointers converge or reach ends
- Pointer movement direction must match problem requirements
- Equal pointers don't always mean process complete (check logic)
- In-place modifications may break symmetry; track carefully
- Slow/fast pointer timing: fast moves 2x speed only works on linked lists for cycle detection
## Sliding Window
### Basics
Sliding window is a technique maintaining a contiguous subsequence (window) that slides through data. Window expands/shrinks based on conditions, tracking some state (sum, count, frequencies) to solve problems efficiently without nested loops.

**Window types:**
- **Fixed-size window:** window size k remains constant, slides one element at a time
- **Variable-size window:** window grows/shrinks dynamically to maintain a condition or find optimal range

### Time Complexity
| Scenario | Complexity |
|---|---|
| Fixed-size window | O(n) |
| Variable-size window | O(n) single pass; O(n log n) with binary search |
| Nested with sorting | O(n log n) |

**Space:** O(min(n, charset size)) for state tracking (hash map, counter)

### Key Concepts
- **Window boundaries:** left and right pointers define the range
- **Invariant:** maintain a property within window (sum, character counts, etc.)
- **Expansion:** increase window size by moving right pointer
- **Contraction:** decrease window size by moving left pointer
- **State tracking:** track metrics within window (hash map for frequencies, running sum)
- **Single pass:** left pointer never moves backward; each element seen once → O(n)

### Key Techniques
- **Fixed-size window:** initialize window with k elements, slide by removing left and adding right
- **Variable-size window (expand-until-fail):** expand right until condition fails, then shrink left to restore
- **Two-pointer coordination:** left and right move independently based on conditions
- **State updates:** add element when expanding, remove when contracting; keep counter/sum consistent
- **Optimization:** avoid recomputing state; update incrementally as window changes

### Common Problems
- **Longest substring without repeating characters:** expand right, shrink left when duplicate found
- **Minimum window substring:** expand right to include all chars, shrink left to minimize window
- **Maximum sum of k consecutive elements:** fixed window, slide and track sum
- **Longest substring with k distinct characters:** expand right adding chars, shrink left to maintain k distinct
- **Sliding window maximum:** find max in each k-sized window (use deque for O(n))
- **Permutation in string:** fixed-size window matching character frequencies
- **Find all anagrams:** fixed-size window comparing character counts

### Gotchas
- Empty window at start (left == right); logic must handle it
- Shrinking window must update state correctly before removing elements
- Off-by-one in window size (inclusive vs exclusive boundaries)
- Hash map cleanup: don't forget to remove keys when count reaches zero
- Fixed vs variable window logic differs significantly; don't mix patterns
- Repeated characters: need to track all occurrences, not just presence
- Right pointer reaches end before condition met: handle incomplete windows
## Intervals
### Basics
Interval problems deal with ranges defined by start and end points. Common patterns include:
- **Merging intervals:** combine overlapping or adjacent ranges
- **Scheduling:** detect conflicts or find available slots
- **Overlap detection:** determine if intervals intersect
- **Insertion:** add new interval maintaining merged state
- **Covering:** find minimum intervals to cover a range

### Time Complexity
| Scenario | Complexity |
|---|---|
| Merge intervals (sort + linear merge) | O(n log n) |
| Insert interval (with merge) | O(n) |
| Meeting rooms (detect conflicts) | O(n log n) |
| Interval scheduling | O(n log n) |
| Interval tree queries | O(log n) per query, O(n) space |

**Space:** O(n) for output, O(1) if modifying in-place

### Key Concepts
- **Interval representation:** pair (start, end); closed [a,b] or half-open [a,b)
- **Overlapping definition:** [1,3] and [3,5] may or may not overlap depending on problem (inclusive vs exclusive boundaries)
- **Sorting intervals:** typically by start time; if ties, by end time (ascending or descending)
- **Timeline sweep:** process events in order (start/end markers) to detect overlaps or compute availability
- **Greedy selection:** early-ending intervals often optimal for scheduling

### Key Techniques
- **Sorting intervals:** sort by start time, then by end time for consistent ordering
- **Merging overlaps:** iterate through sorted intervals, expand current interval if next overlaps, else append
- **Overlap detection:** check if two intervals overlap; conditions vary by boundary type
- **Interval tree:** balanced tree supporting range queries; insertion/deletion/query O(log n)
- **Segment tree:** array-based tree for range updates and queries
- **Timeline/event sweep:** create events for starts/ends, sort, sweep through detecting state changes
- **Greedy scheduling:** select intervals by earliest end time to maximize count

### Common Problems
- **Merge intervals:** combine overlapping ranges into minimal set
- **Insert interval:** add new interval to merged list, merging as needed
- **Meeting rooms I:** detect if all meetings can be attended (no conflicts)
- **Meeting rooms II:** find minimum rooms needed for all meetings
- **Interval scheduling:** find maximum non-overlapping intervals
- **Video stitching/Leap frog:** find minimum intervals to cover a range
- **Non-overlapping intervals:** remove minimum intervals for no overlaps
- **Employee free time:** find common free slots among schedules
- **Sky line problem:** merge intervals with heights to visualize outline

### Gotchas
- **Boundary conditions:** off-by-one errors with [a,b] vs [a,b); clarify if boundaries are inclusive/exclusive
- **Overlapping definition:** [1,3] and [3,5] overlap (touch) or not; problem must specify
- **Sorting stability:** if start times equal, must sort by end consistently (e.g., ascending)
- **Empty result:** after merge/filter, output may be empty or single interval
- **Merge direction:** direction of insertion/traversal affects logic (left-to-right vs reverse)
- **Greedy pitfalls:** earliest start is NOT optimal for maximizing count; use earliest end
- **In-place modification:** modifying input array affects space complexity claims

## Binary Search
### Basics
Binary search is a divide-and-conquer algorithm that efficiently finds a target value in a sorted array by repeatedly dividing the search space in half. Maintains two pointers (left and right) defining the active range, calculates a middle index, and eliminates half of the remaining elements with each comparison.

**Key principle:** on each iteration, the search space is divided by 2, exploiting the sorted property to prune impossible regions.

**Prerequisite:** the array must be sorted (or at least have a monotonic property).

### Time Complexity
| Scenario | Complexity |
|---|---|
| Best case | O(1) — target found immediately at mid |
| Average case | O(log n) — typical search |
| Worst case | O(log n) — element not found or requires full halving |
| Space | O(1) iterative; O(log n) recursive (call stack) |

### Key Concepts
- **Sorted array requirement:** binary search only works on sorted or monotonic data structures
- **Monotonic property:** not limited to sorted arrays; any monotonicity works (e.g., "first occurrence where condition true")
- **Search space reduction:** each iteration reduces candidates by half; log n iterations maximum
- **Boundary conditions:** handling inclusive/exclusive boundaries crucial for correctness
- **Mid calculation:** mid = left + (right - left) / 2 prevents integer overflow (vs (left + right) / 2)
- **Pointer invariants:** maintain what left and right represent (inclusive bounds, inclusive-exclusive, etc.)

### Key Techniques
- **Left/right pointer initialization:** left = 0, right = n-1 (inclusive) or right = n (exclusive)
- **Mid calculation:** mid = left + (right - left) / 2 avoids overflow in languages with fixed integer sizes
- **Comparison and pointer update:**
  - If arr[mid] == target: found, return mid
  - If arr[mid] < target: eliminate left half, left = mid + 1
  - If arr[mid] > target: eliminate right half, right = mid - 1
- **Loop termination:** continue while left <= right (inclusive) or left < right (exclusive)
- **Search boundaries:** find first/last occurrence by continuing search after finding target
- **Rotated array search:** identify which half is sorted, then determine which half contains target

### Common Problems
- **Basic search:** find target in sorted array
- **First/last occurrence:** find leftmost or rightmost position of target element
- **Search in rotated sorted array:** handle rotation point to determine valid search direction
- **Find closest element:** locate element closest to target value
- **Search range:** find all indices within a range
- **Peak element:** find local maximum in mountain-like array
- **Find k-th smallest:** use binary search on answer space
- **Median of two sorted arrays:** binary search partition point
- **Capacity-based search:** find minimum capacity that satisfies condition (e.g., koko eating bananas)

### Gotchas
- **Off-by-one errors:** confusing inclusive vs exclusive boundaries; wrong mid update logic
- **Infinite loops:** incorrect pointer updates (e.g., mid = (left + right) / 2 then left = mid) on single-element range
- **Integer overflow:** mid = (left + right) / 2 overflows in some languages; use mid = left + (right - left) / 2
- **Boundary condition mistakes:** using < instead of <= or wrong update direction leads to wrong result
- **Assuming uniqueness:** code may assume no duplicates; duplicates require special handling
- **Forgetting monotonicity:** binary search requires sorted or monotonic property; doesn't work on arbitrary arrays
- **Post-loop boundary check:** after loop exits with left > right, verify if element found or where insertion point is
- **Mid always rounds down:** mid = (left + right) / 2 truncates for odd ranges; matters for "find closest" problems
## Kadane's Algorithm
### Basics
Kadane's algorithm finds the maximum sum of a contiguous subarray in linear time. It uses a greedy approach combined with dynamic programming: at each position, decide whether to extend the current subarray or start a new one. The key insight is that we only need to track the maximum sum ending at the current position (max_current) and the global maximum seen so far (max_global).

**How it works:**
- Iterate through array from left to right
- At each element, max_current = max(element, max_current + element)
- Update max_global = max(max_global, max_current)
- Return max_global after single pass

**Core idea:** at each step, either extend the existing subarray or start fresh; if extending adds a positive value, keep going; if negative, consider starting fresh.

### Time Complexity
| Scenario | Complexity |
|---|---|
| Standard maximum subarray | O(n) time, O(1) space |
| Maximum product subarray | O(n) time, O(1) space |
| Circular maximum subarray | O(n) time, O(n) space (variant) |
| Minimum subarray sum | O(n) time, O(1) space (variant) |

**Space:** O(1) — only two variables needed, or O(n) for circular variant requiring total_sum tracking

### Key Concepts
- **max_current (max ending here):** maximum subarray sum that includes current element; reset or carry forward based on next element
- **max_global (max so far):** best answer found so far; updated whenever max_current exceeds it
- **Contiguous requirement:** subarray must be contiguous; gaps break the sequence
- **Negative number handling:** negative numbers can lower max_current; if sum becomes negative and next element is positive, may reset to that element
- **Greedy choice:** at each step, choose to extend or reset based on which gives better result
- **Single pass:** no need to revisit data; information tracked through two variables

### Key Techniques
- **Standard Kadane's:** maintain max_current and max_global; max_current = max(nums[i], max_current + nums[i])
- **Track indices:** store start/end indices of max subarray by tracking when max_current resets
- **Maximum product subarray:** track both max and min ending here (min can become positive with negative number)
- **Minimum subarray sum:** use same logic but track minimum; useful for finding worst-case subarray
- **Circular array variant:** find max(total_sum - min_subarray, max_subarray); handles wrapping subarrays
- **Modified variants:** find subarray with at least k elements, find subarrays with sum >= target, etc.

### Common Problems
- **Maximum subarray sum:** return the maximum sum of any contiguous subarray
- **Maximum sum subarray with at least k elements:** extend base algorithm with length constraint
- **Maximum product subarray:** track both max and min simultaneously (negative × negative = positive)
- **Circular maximum subarray:** max of (standard max, total_sum - min_subarray); be careful of all-negative case
- **Maximum subarray sum with k subarrays:** combine Kadane's with DP to find k non-overlapping subarrays with max total
- **Subarray sum equals target:** find if any contiguous subarray sums to target (use Kadane's variant or prefix sum + hash map)
- **Best time to buy and sell stock:** special case; max profit = max(prices[j] - prices[i]) where j > i; equivalent to max subarray on price differences

### Gotchas
- **All negative numbers:** algorithm returns the single least negative number, not 0; handle empty subarray requirement carefully
- **Single element array:** answer is that element itself; max_current and max_global both initialized to first element
- **Empty subarray handling:** Kadane's assumes at least one element; if empty subarrays allowed (sum = 0), modify initialization
- **Integer overflow:** with large arrays, sum can overflow; use long/int64 or BigInteger depending on language
- **Negative infinity initialization:** max_global should start at negative infinity or first element, not 0 (breaks all-negative case)
- **Circular array trap:** don't confuse circular variants; all-negative array needs special handling (max element, not 0 or min subarray)
- **Product variant complexity:** negative numbers make min tracking essential; forgetting min_current causes incorrect results
- **Index tracking:** if storing indices, must update carefully when resetting max_current; off-by-one common
- **Tied answers:** multiple subarrays may have same max sum; algorithm returns one, but problem may require all or specific one
## Bit Manipulation
### Basics
Bit manipulation is operating directly on binary representations of numbers using bitwise operations. Numbers are stored as sequences of bits (0s and 1s) in binary; bit operations are fundamental CPU instructions, executing in O(1) time. Useful for optimization, flag storage, efficient encoding, and solving problems with inherent binary structure.

**Core bitwise operations:**
- **AND (&):** 1 only if both bits are 1; used to isolate or mask bits
- **OR (|):** 1 if at least one bit is 1; used to set bits
- **XOR (^):** 1 if bits differ; used to toggle or compare bits
- **NOT (~):** flips all bits (complement)
- **Left shift (<<):** multiply by 2^k; shift n << k multiplies n by 2^k
- **Right shift (>>):** divide by 2^k; shift n >> k divides n by 2^k (logical or arithmetic)
- **Bit test:** check if bit at position i is set: (n & (1 << i)) != 0

### Time Complexity
| Op | Complexity |
|---|---|
| AND, OR, XOR, NOT | O(1) |
| Left/Right shift | O(1) |
| Count set bits (naive) | O(k) where k = number of bits in n |
| Count set bits (Brian Kernighan) | O(number of 1s) |
| Check power of 2 | O(1) |
| Operations on n-bit numbers | O(log n) for n-bit integers |

*Most bit operations are O(1) CPU instructions; complexity depends on total bit width, typically O(log n) for n*

### Key Concepts
- **Binary representation:** numbers stored as sequences of bits; LSB (Least Significant Bit) at position 0, MSB (Most Significant Bit) at highest position
- **Bit masks:** patterns of bits used to isolate or set specific bit ranges; e.g., mask = 1 << i extracts bit at position i
- **Bit shifting:** left shift (<<) multiplies by 2^k, right shift (>>) divides by 2^k; used for fast multiplication/division and bit positioning
- **Two's complement:** standard representation for signed integers; most significant bit is sign bit; negative numbers represented as ~x + 1
- **Endianness:** byte order (big-endian vs little-endian); bit ordering within bytes and across bytes can be non-intuitive
- **Bit position:** usually 0-based from LSB (rightmost); position i refers to bit (n >> i) & 1
- **Signed vs unsigned shifts:** unsigned always fills with 0; signed (arithmetic shift) fills with sign bit
- **All-ones pattern:** (1 << n) - 1 gives n bits all set to 1; useful for masking lower n bits

### Key Techniques
- **Isolate a bit:** use AND with mask; (n & (1 << i)) extracts bit at position i
- **Set a bit:** use OR; n | (1 << i) sets bit at position i to 1
- **Clear a bit:** use AND with complement; n & ~(1 << i) clears bit at position i
- **Toggle a bit:** use XOR; n ^ (1 << i) flips bit at position i
- **XOR cancels:** a ^ a = 0; a ^ 0 = a; XOR is self-inverse, useful for finding single numbers
- **AND isolates:** a & a = a; a & 0 = 0; AND with mask extracts specific bits
- **Power of 2 check:** n > 0 && (n & (n - 1)) == 0 checks if n is a power of 2 (only one bit set)
- **Count set bits:** iterate through bits or use Brian Kernighan (n & (n-1)) to clear lowest set bit in O(popcount) time
- **Get rightmost set bit:** n & -n in two's complement; negative number has all bits flipped + 1
- **Clear lowest set bit:** n & (n - 1); subtracting 1 flips all bits from lowest set bit rightward
- **Check bit i set:** (n >> i) & 1 or (n & (1 << i)) != 0

### Common Problems
- **Single Number:** XOR all elements; duplicates cancel out, leaving single unpaired number
- **Single Number II:** elements appear 3 times except one; count bits modulo 3 for each position
- **Single Number III:** two unpaired numbers; partition by first differing bit using XOR result
- **Missing Number:** XOR all indices with array elements; missing number cancels when others XOR
- **Hamming Distance:** count differing bits between two numbers using (x ^ y) and counting set bits
- **Reverse Bits:** build result by extracting bits from end and placing at start using shifts and masks
- **Power of 2:** check (n & (n-1)) == 0 and n > 0; only one bit should be set
- **Number of 1 Bits (Hamming Weight):** count set bits using Brian Kernighan or built-in popcount
- **Majority Element:** most frequent bit at each position determines result; use bit counting
- **Bitwise AND of Numbers Range:** find common prefix in binary representations by shifting both boundaries
- **UTF-8 Validation:** check bit patterns of continuation bytes and lead bytes using masks
- **Sum of Two Integers:** implement addition using XOR (sum without carry) and AND with shift (carry)

### Gotchas
- **Signed vs unsigned overflow:** left shift on negative numbers is undefined in many languages; right shift of negative is arithmetic (sign-extends) in some languages
- **Endianness confusion:** bit order within bytes may differ from expectations on different architectures; be explicit about bit positions
- **Bit position counting:** easy to confuse 0-based (standard in code) vs 1-based; position 0 is the rightmost (LSB)
- **Shift direction:** left (<<) increases value, right (>>) decreases; confusing when thinking about "shifting to lower positions"
- **Sign bit and negative numbers:** two's complement means ~n = -(n+1); NOT operation on positive gives negative
- **AND/OR precedence:** often need parentheses; & has higher precedence than comparison operators in C/Java
- **Off-by-one in masks:** (1 << n) - 1 gives n bits; (1 << (n+1)) - 1 gives n+1 bits
- **Clearing bits with 0 mask:** n & 0 always gives 0; must use complemented mask n & ~mask
- **Overflow in bit counting:** shifting 1 too far left overflows; max safe shift is bit width - 1
- **Platform-dependent behavior:** bit width (32 vs 64), sign extension, and undefined shift amounts vary by language/platform

## Trie
### Basics
A Trie (prefix tree) is a tree structure where each node represents a character, and paths from root to leaf form strings. Commonly used for efficiently storing and searching strings with shared prefixes.

Trie structure consists of:
- **Root node:** starting point representing empty string or word prefix start
- **Node per character:** each node stores a single character
- **Children map:** each node contains a map/array mapping characters to child nodes
- **End-of-word marker:** boolean flag on each node indicating word termination
- **Implicit string storage:** words are formed by concatenating characters along the path from root

### Time Complexity
| Op | Complexity |
|---|---|
| Insert | O(m) |
| Search | O(m) |
| Delete | O(m) |
| Prefix search | O(m) |
| Space | O(alphabet_size × N) |

*m = word length; N = number of words; alphabet_size typically 26 (lowercase English) to 256 (bytes)*

### Key Concepts
- **Prefix tree:** words with common prefixes share the same path segment, reducing redundant storage
- **Node structure:** character identifier + children map + end-of-word boolean flag
- **Path invariant:** every word is represented by exactly one unique path from root to marked end node
- **Shared prefixes:** core advantage; "cat", "car", "card" share the "ca" prefix path
- **Alphabet size impact:** affects both space (array size) and speed; sparse alphabets favor hash maps over arrays
- **Traversal validation:** null checks required at each step; missing checks cause null pointer exceptions

### Key Techniques
- **Insert word:** iterate through characters; create missing nodes; mark final character as end-of-word
- **Search word:** follow character path from root; valid only if full path exists and end-of-word is true
- **Prefix search:** traverse to end of prefix; if complete path exists, prefix is valid entry point
- **Autocomplete:** traverse to prefix end node, then DFS collecting all words rooted at that subtree
- **Find all words:** DFS from root, collecting paths that terminate at marked end-of-word nodes
- **Delete word:** traverse word path, unmark end-of-word flag; optionally prune childless nodes bottom-up

### Common Problems
- **Autocomplete/search suggestion:** suggest words matching user-typed prefix
- **Spell checker:** validate word existence; find typo candidates via Trie proximity
- **Longest common prefix:** find where Trie paths diverge (first branching node)
- **Word search in grid (II):** combine Trie with 2D grid DFS; use Trie to validate words during traversal
- **Replace words:** identify and replace matching strings using Trie dictionary lookup
- **IP address validation:** Trie-like structure for efficient IP prefix routing
- **Lexicographic sorting:** traverse Trie in order, collecting words in sorted sequence

### Gotchas
- **Space overhead:** single character per node; sparse alphabets (e.g., only 5 chars used) waste array slots — use hash maps for sparse cases
- **End-of-word marker essential:** "car" and "card" both exist in Trie; unmarking "car" corrupts search — must mark both
- **Null checks mandatory:** missing child existence checks before recursion cause crashes; validate at every level
- **Character encoding:** ASCII vs Unicode affects alphabet size and memory cost; UTF-8 requires careful handling
- **Case sensitivity:** decide early ("Cat" vs "cat" vs both); affects prefix behavior and search correctness
- **Memory per node:** allocating full array for sparse alphabets wastes memory; profile and consider hash maps
- **DFS depth in autocomplete:** unbounded DFS from prefix can be slow on large subtrees; implement result limits
- **Deletion pruning complexity:** removing nodes after delete requires careful bookkeeping; only prune truly childless nodes

## Backtracking
### Basics
Backtracking is a technique for exploring all possible candidates to solve problems by building solutions incrementally, undoing (backtracking) choices when a path proves invalid.

**Pattern:**
- Build candidate solution step-by-step
- If current path violates constraints, backtrack (undo last choice)
- Recursively try remaining alternatives
- Repeat until valid solution found or all paths exhausted

### Time Complexity
| Scenario | Complexity |
|---|---|
| Worst case (no pruning) | O(N^M) where N = choices per step, M = depth |
| With pruning | O(N^M) — exponential but reduced by early termination |
| Space | O(M) recursion stack depth |

*Actual complexity depends heavily on pruning strategy and problem constraints*

### Key Concepts
- **Decision tree:** tree of all possible choices; backtracking explores it depth-first
- **Choice-explore-unchoose pattern:** select option, recurse, undo state before trying next option
- **Constraint satisfaction:** backtrack when constraints violated (pruning)
- **Pruning:** eliminate invalid branches early to reduce search space
- **State management:** track current partial solution; must restore state on backtrack
- **Base cases:** solution found (collect result) or no valid choices (backtrack)

### Key Techniques
- **Recursive backtracking function:** takes current state (partial solution), recursively tries choices
- **Base case:** check if valid solution (all cells filled, all constraints met); if so, record/return
- **Recursive case:** iterate through valid choices, recursively explore each, undo choice after return
- **Pruning:** check constraints before recursing; skip invalid branches immediately
- **State restoration:** undo choice (remove from set, restore value, decrement counter) before trying next option
- **Memoization:** cache invalid states to avoid recomputing same subproblems (with appropriate state representation)

### Common Problems
- **N-Queens:** place N queens on NxN board with no conflicts
- **Permutations:** generate all orderings of elements
- **Combinations:** select k elements from n without regard to order
- **Word search:** find if word exists in 2D grid of letters
- **Sudoku solver:** fill grid with 1-9 satisfying row/column/box constraints
- **Parentheses generation:** generate all valid combinations of n pairs of parentheses
- **Subsets:** generate all subsets of a set
- **Rat in maze:** find path from start to end avoiding obstacles
- **Graph coloring:** assign colors to vertices minimizing conflicts

### Gotchas
- **Forgetting to undo state:** if you don't restore state after recursion, subsequent iterations see corrupted data
- **Infinite recursion:** missing base case or incorrect termination condition causes stack overflow
- **Over-pruning:** too aggressive pruning skips valid solutions
- **Under-pruning:** insufficient pruning checks still explores invalid branches (inefficient)
- **Exponential time complexity:** without good pruning, backtracking explores all O(N^M) branches
- **Stack overflow on deep recursion:** very deep problems may exceed stack limits (use iterative alternative)
- **Forgetting to handle all constraints:** incomplete pruning logic allows invalid solutions to be returned

## Divide and Conquer
### Basics
Divide and conquer breaks a problem into smaller subproblems, solves them recursively, and combines results. Applicable to many problem types: sorting (Merge Sort, Quick Sort), searching (Binary Search), array problems (max subarray), and geometric problems.

**Three steps:**
- **Divide:** break problem into disjoint subproblems
- **Conquer:** solve subproblems recursively; base case solves directly if small enough
- **Combine:** merge subproblem solutions into overall solution

### Time Complexity
| Scenario | Complexity |
|---|---|
| Balanced split (T(n) = aT(n/b) + f(n)) | Use Master Theorem; often O(n log n) |
| Merge Sort | O(n log n) time, O(n) space |
| Quick Sort | O(n log n) average, O(n²) worst case |
| Binary Search | O(log n) |
| Unbalanced split | O(n²) or worse if one subproblem is nearly entire problem |

### Key Concepts
- **Subproblem independence:** divided subproblems don't overlap; avoid redundant computation (vs dynamic programming)
- **Optimal substructure:** optimal solution built from optimal solutions to subproblems
- **Recurrence relation:** T(n) = time to divide + time to solve subproblems + time to combine
- **Master Theorem:** closed-form solution for recurrences of form T(n) = aT(n/b) + f(n)
- **Split balance:** equal or nearly equal splits lead to O(n log n); highly unbalanced degrades to O(n²)
- **Combine complexity:** efficient merging crucial; O(n) combine time with O(log n) depth gives O(n log n) overall

### Key Techniques
- **Recursive decomposition:** identify how to split problem into independent subproblems of same type
- **Base case:** directly solve small subproblems (e.g., single element, size < k)
- **Combine step:** efficiently merge subproblem solutions (merge two sorted lists in O(n), combine left/right maximums)
- **Balanced split:** divide roughly equally to achieve O(log n) recursion depth
- **Iterative variant:** some divide-and-conquer problems can be solved iteratively to avoid recursion overhead (e.g., iterative merge sort)

### Common Problems
- **Merge Sort:** divide array in half, recursively sort both, merge sorted halves; O(n log n) guaranteed
- **Quick Sort:** partition around pivot, recursively sort left/right; O(n log n) average, O(n²) worst case
- **Binary Search:** divide search space in half, recursively search relevant half
- **Maximum subarray (Kadane alternative):** divide array, find max subarray in left/right/middle, combine
- **Closest pair of points:** divide points by x-coordinate, find closest pair in each half, check across boundary
- **Strassen's matrix multiplication:** divide matrices into blocks, reduce multiplications from 8 to 7
- **Count inversions:** divide array, count inversions in halves, count across boundary (using merge)

### Gotchas
- **Overlapping subproblems:** if subproblems overlap, divide-and-conquer is inefficient; use dynamic programming instead
- **Incorrect combine logic:** merging results incorrectly loses correctness (e.g., max of two halves ≠ overall max if max spans boundary)
- **Unbalanced split:** one subproblem nearly entire problem causes O(n²) instead of O(n log n)
- **Stack overflow on deep recursion:** very deep recursion trees may exceed stack; convert to iteration if needed
- **Off-by-one in boundaries:** splitting into [left, mid] and [mid+1, right] vs [left, mid-1] and [mid, right] affects results
- **Combine step complexity overlooked:** if combine is O(n²), overall becomes O(n²) even with O(log n) depth
- **Base case edge cases:** ensure base case correctly handles empty, single element, or very small inputs

# Advanced Topics
## 1D Dynamic Programming
## Multidimensional Dynamic Programming
### Basics
Multidimensional dynamic programming extends 1D DP concepts to problems involving multiple independent dimensions. Instead of a single recurrence relation dp[i], multidimensional DP uses tables like dp[i][j] (2D), dp[i][j][k] (3D), or higher. Each dimension typically represents a different constraint or resource (rows, columns, weights, items, time steps, positions). The core principle remains: break the problem into overlapping subproblems, store intermediate results, and build up to the final answer.

**Common structures:**
- **2D grid DP:** dp[i][j] where i = row, j = column; often solving path-counting or optimization on grids
- **2D knapsack variants:** dp[i][w] where i = item index, w = weight/capacity; or dp[w][v] for weight-value tradeoffs
- **3D DP:** dp[i][j][k] adding another dimension (e.g., two strings for LCS, item and two knapsacks, state + time)
- **Higher dimensions:** rarely needed but follow the same principles (exponential state growth)

### Time Complexity
| Problem Class | Time | Space |
|---|---|---|
| 2D grid path/matrix DP | O(m·n) | O(m·n) or O(min(m,n)) optimized |
| 0/1 Knapsack 2D | O(n·w) where n = items, w = capacity | O(n·w) or O(w) optimized |
| Knapsack variants (bounded, unbounded) | O(n·w) to O(n·w·q) | O(w) to O(n·w) |
| Longest Common Subsequence (LCS) | O(m·n) where m, n = string lengths | O(m·n) or O(min(m,n)) optimized |
| Edit Distance | O(m·n) | O(m·n) or O(min(m,n)) optimized |
| Unique Paths on grid | O(m·n) | O(m·n) or O(n) optimized |
| Max/Min Path Sum | O(m·n) | O(m·n) or O(n) optimized |
| Coin Change (multiple coins) | O(amount·coin_count) | O(amount·coin_count) or O(amount) |

### Key Concepts
- **Multiple state dimensions:** recurrence depends on two or more independent variables; each dimension corresponds to a constraint (capacity, string position, grid location, time step)
- **State definition:** clearly specify what dp[i][j][...] represents (e.g., "max profit using items 0..i-1 with capacity j")
- **Recurrence relation:** transition between states typically considers choices in current dimension and carries forward prior state (dp[i][j] = max/min of dp[i-1][...], dp[i][j-1], etc.)
- **Boundary initialization:** handle base cases for each dimension carefully (e.g., dp[0][j] = 0 for all j, dp[i][0] = 0 or some default)
- **State ordering dependency:** order of dimension matters; usually fill table in specific order (left-to-right, top-to-bottom) to ensure dependencies are computed first
- **State meaning consistency:** each cell must unambiguously represent the subproblem; mixing inclusive/exclusive indexing causes errors

### Key Techniques
- **Bottom-up 2D tabulation:** create 2D array, initialize base cases, fill row-by-row or column-by-column ensuring dependencies are met
- **Space optimization (rolling arrays):** if recurrence only depends on previous row/column, use 1D array and update in-place or swap between two arrays
- **Space optimization (single 1D array):** careful in-place update with backward iteration to avoid overwriting needed values
- **Path reconstruction:** track parent pointers (previous state) during computation to reconstruct actual solution, not just optimal value
- **Distinct ways/counting:** instead of max/min, count number of ways to achieve goal; initialize dp[0][0] = 1, multiply or add appropriately
- **2D iteration orders:** row-major (iterate rows, then columns), column-major (iterate columns, then rows), diagonal order, or level-by-level; choose order based on dependencies
- **Multiple constraint combination:** problem may involve two independent knapsack-like constraints (e.g., weight and value limits); nest loops for each constraint

### Common Problems
- **0/1 Knapsack:** dp[i][w] = max value using items 0..i-1 with weight limit w; recurrence considers taking or skipping current item
- **Bounded Knapsack:** each item has limited quantity; extend knapsack with nested loop over counts
- **Unbounded Knapsack:** items can be used unlimited times; recurrence is dp[w] = max(dp[w], dp[w-item_weight] + item_value)
- **Coin Change (minimum coins):** dp[amount] = min coins to make amount; similar to unbounded knapsack
- **Longest Common Subsequence (LCS):** dp[i][j] = length of LCS of str1[0..i-1] and str2[0..j-1]; recurrence matches or skips characters
- **Edit Distance (Levenshtein):** dp[i][j] = min edits to transform str1[0..i-1] to str2[0..j-1]; consider insert, delete, replace
- **Unique Paths on grid:** dp[i][j] = number of paths to cell (i, j) from top-left; recurrence is dp[i][j] = dp[i-1][j] + dp[i][j-1]
- **Maximum Path Sum in grid:** dp[i][j] = max sum path from start to (i, j); recurrence is dp[i][j] = cell_value + max(dp[i-1][j], dp[i][j-1])
- **Minimum Path Sum in grid:** dp[i][j] = min sum path from start to (i, j); similar to max but takes minimum
- **2D Partition/Range DP:** dp[i][j] = optimal partitioning/solution for range [i, j]; recurrence tries all split points k
- **Two Strings DP (LCS, Edit Distance):** general class of problems comparing two sequences; 2D table naturally represents alignment
- **Item + Constraint DP (subset sum, partition):** dp[i][s] or dp[i][s][k] where i = item index, s = current sum/constraint, k = additional parameter

### Gotchas
- **Index bounds checking:** multidimensional arrays require careful bounds; off-by-one errors in rows, columns, or other dimensions are common
- **Space optimization losing path info:** if using 1D rolling array to save space, original 2D table info is lost; path reconstruction becomes harder
- **State ordering/dependency:** if table filled in wrong order, recurrence depends on not-yet-computed values; verify order before coding
- **Initialization correctness:** base cases must be correct for all dimensions; missing or wrong initialization cascades through entire table
- **Confusing inclusive vs exclusive indexing:** decide whether dp[i] means "using items 0 to i-1" or "using items 0 to i"; inconsistency breaks recurrence
- **Dimension size miscalculation:** if problem has capacity w up to 10^6 and many items, O(n·w) space/time may be prohibitive; consider optimizations
- **Recurrence mixing up dimensions:** wrong order of parameters in recurrence (e.g., dp[i][j] vs dp[j][i]) leads to incorrect answers
- **Carrying unnecessary state:** extra dimensions increase time/space; verify each dimension is necessary
- **In-place updates:** when using 1D array with space optimization, updating in wrong order (forward vs backward) can overwrite needed values
- **Tied or multiple optimal solutions:** algorithm finds one optimal value; if problem requires all solutions or specific lexicographic ordering, requires additional tracking

## Math