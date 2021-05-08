package base

type Stalk struct {
	Id          int
	Name        string `description:"对象名"`
	Description string `description:"描述"`
}

func GetTree() (tree []Stalk) {
	id := 0
	for k, _ := range ActionsMap {
		id++
		stalk := Stalk{
			Id:          id,
			Name:        k,
			Description: k,
		}
		tree = append(tree, stalk)
	}
	return tree
}
