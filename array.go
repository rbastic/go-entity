package entity

type Array []*Entity

func makeEmptyChildrenMap() map[string]Array {
	return make(map[string]Array)
}
