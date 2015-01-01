package model

type Api struct {
	treeRoot      *TreeNode
	people        []Person
	skillHoldings *SkillHoldings
}

func NewApi() *Api {
	return &Api{
		people:        make([]Person, 0),
		skillHoldings: NewSkillHoldings()}
}
