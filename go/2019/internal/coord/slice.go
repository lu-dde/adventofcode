package coord

//Slice is just an array of Coords with funcs for sorting
type Slice []Pair

func (cs Slice) Len() int {
	return len(cs)
}
func (cs Slice) Less(i, j int) bool {
	return cs[i].less(cs[j])
}
func (cs Slice) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
