package mst

import "sort"

type Edge[N comparable] struct {
	Node1  N
	Node2  N
	Weight int
}

type edges[N comparable] []Edge[N]

func (egs edges[N]) Len() int { return len(egs) }
func (egs edges[N]) Less(i, j int) bool {
	return egs[i].Weight < egs[j].Weight
}
func (egs edges[N]) Swap(i, j int) {
	egs[i], egs[j] = egs[j], egs[i]
}

type edgeGroup[N comparable] struct {
	m   map[N]bool // included nodes map
	egs edges[N]   // edges
}

func (eg *edgeGroup[N]) add(e Edge[N]) {
	eg.m[e.Node1] = true
	eg.m[e.Node2] = true
	eg.egs = append(eg.egs, e)
}
func (eg *edgeGroup[N]) merge(other *edgeGroup[N]) {
	for k := range other.m {
		eg.m[k] = true
	}
	eg.egs = append(eg.egs, other.egs...)
}

func FindMST[N comparable](nodes []N, egs edges[N]) edges[N] {
	sort.Sort(egs)

	groups := []*edgeGroup[N]{}
	for _, e := range egs {
		// check both nodes are included in a group
		// then ignore the edge
		var ignore bool
		for _, g := range groups {
			if g.m[e.Node1] && g.m[e.Node2] {
				ignore = true
				break
			}
		}
		if ignore {
			continue
		}

		// if a group has one of edge nodes, include this edge
		var addedgroup *edgeGroup[N]
		for _, g := range groups {
			if g.m[e.Node1] || g.m[e.Node2] {
				g.add(e)
				addedgroup = g
				break
			}
		}

		// there is no group has one of edge nodes, make a new group
		if addedgroup == nil {
			eg := &edgeGroup[N]{
				m: make(map[N]bool),
			}
			eg.add(e)
			groups = append(groups, eg)
			continue
		}

		// if two groups has one of edge nodes each then merge it
		for i, g := range groups {
			if addedgroup == g {
				continue
			}
			if g.m[e.Node1] || g.m[e.Node2] {
				// merge two groups
				addedgroup.merge(g)
				groups = append(groups[:i], groups[i+1:]...)
				break
			}
		}
	}

	if len(groups) != 1 {
		// if there are two or more groups then it means couldn't make single MST
		return nil
	}
	return groups[0].egs
}
