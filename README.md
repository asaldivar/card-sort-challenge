# Card Sort

Coding challenge for Happy Returns. Implements topological sort on a directed acyclic graph to sort a series of boarding cards with origin and destination information in linear time and space, O(|V|+|E|) where |V| is the number of nodes and |E| is the number of edges.

Steps:
 - Construct an adjacency list to represent the graph
 - Use a map `visited` to store the state of nodes (unvisited, visited and fully explored, visited but not yet fully explored)
 - Call the `topoSort` function on each node, which will recursively call `topoSort` on each of the node's neighbors
 - If there are any cycles (a node was revisited before being marked as fully explored), we return an error
 - The first node to be fully explored is the sink, this node then gets appended to the `output` slice
 - As exploration of each node completes, `output` will be appended to in the reverse order of the itinerary (last to first)

### Alternate Implementation with Reverse Graph

Note that the requirements of the problem guarantee that there are not multiple paths to a final destination (no cities can be repeated). 

As such, an alternative way to implement the sort would simply be to construct the reverse graph using a map. The node missing from the map is the origin, and can be used to trace the rest of the itinerary. This algorithm also runs in linear time and space.
