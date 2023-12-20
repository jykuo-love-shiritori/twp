package buyer

type tagSet struct {
	Set map[int32]bool
}

func NewTagSet(tags []int32) *tagSet {
	t := &tagSet{}
	t.Set = make(map[int32]bool)
	for _, tag := range tags {
		t.Set[tag] = true
	}
	return t
}
func (t *tagSet) Intersect(other *tagSet) bool {
	for k := range t.Set {
		if other.Set[k] {
			return true
		}
	}
	return false
}
