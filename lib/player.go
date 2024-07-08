package lib

type Player int

const (
	ComputerPlayer Player = 1
	HumanPlayer    Player = 2
)

func (p Player) ToCell() cell {
	switch p {
	case ComputerPlayer:
		return ComputerCell
	case HumanPlayer:
		return HumanCell
	default:
		return EmptyCell
	}
}
