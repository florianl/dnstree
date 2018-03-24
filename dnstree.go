package dnstree

import "fmt"

type tree struct {
	rank     int
	branches map[rune]*tree
}

var root *tree

func init() {
	root = new(tree)
	root.rank = -1
	root.branches = make(map[rune]*tree)
}

func reverse(dns string) []rune {
	runes := []rune(dns)
	lr := len(runes)

	// Reverse the runes array
	for i := lr/2 - 1; i >= 0; i-- {
		j := lr - 1 - i
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

func (t *tree) insert(dns []rune, rank int) error {
	lr := len(dns)

	if branch, ok := t.branches[dns[0]]; ok {
		if lr == 1 {
			if branch.rank == -1 {
				branch.rank = rank
				return nil
			} else {
				return fmt.Errorf("Element already exists")
			}
		}
		branch.insert(dns[1:], rank)
	} else {
		var branch = new(tree)
		branch.branches = make(map[rune]*tree)
		t.branches[dns[0]] = branch
		if lr == 1 {
			branch.rank = rank
		} else {
			branch.rank = -1
			branch.insert(dns[1:], rank)
		}

	}

	return nil
}

// Insert a dns element in the tree
func Insert(dns string, rank int) {
	runes := reverse(dns)
	root.insert(runes, rank)
}

func (t *tree) search(dns []rune) (int, error) {
	var lr = len(dns)
	if lr == 0 {
		if t.rank == -1 {
			return -1, fmt.Errorf("Not found")
		} else {
			return t.rank, nil
		}
	}
	if branch, ok := t.branches[dns[0]]; ok {
		return branch.search(dns[1:])
	}
	return -1, fmt.Errorf("Not found")
}

// Search for a dns element in the tree
func Search(dns string) (int, error) {
	runes := reverse(dns)
	return root.search(runes)
}
