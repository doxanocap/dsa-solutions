package _887

type DSU struct {
	parent map[int]int
	dist   map[int]int
}

func (d *DSU) find(x int) (int, int) {
	if _, exists := d.parent[x]; !exists {
		d.parent[x] = x
		d.dist[x] = 0
		return x, 0
	}

	if d.parent[x] == x {
		return x, 0
	}

	root, rootDist := d.find(d.parent[x])

	d.parent[x] = root
	d.dist[x] = rootDist ^ d.dist[x]

	return d.parent[x], d.dist[x]
}

func (d *DSU) union(u, v, w int) bool {
	rootU, distU := d.find(u)
	rootV, distV := d.find(v)

	if rootU != rootV {
		d.parent[rootU] = rootV
		d.dist[rootU] = distU ^ distV ^ w
		return true
	}

	if distU^distV^w == 0 {
		return true
	}

	return false
}

func numberOfEdgesAdded(n int, edges [][]int) int {
	dsu := &DSU{
		parent: make(map[int]int),
		dist:   make(map[int]int),
	}

	count := 0
	for _, e := range edges {
		if dsu.union(e[0], e[1], e[2]) {
			count++
		}
	}
	return count
}
