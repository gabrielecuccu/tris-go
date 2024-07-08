package lib

type Field [3][3]cell

func NewEmptyField() *Field {
	return &Field{
		{EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell},
	}
}
