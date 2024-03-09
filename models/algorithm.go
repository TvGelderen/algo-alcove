package models

type AlgorithmType int16

const (
	AlgoritmhTypeSorting AlgorithmType = iota
	AlgoritmhTypePathFinding
	AlgoritmhTypeOther
)

func (agorithmType AlgorithmType) AlgorithmTypeToString() string {
    switch agorithmType{
        case AlgoritmhTypeSorting:
            return "Sorting"
        case AlgoritmhTypePathFinding:
            return "Path Finding"
        default:
            return ""
    }
}

func (agorithmType AlgorithmType) AlgorithmTypeToPathString() string {
    switch agorithmType{
        case AlgoritmhTypeSorting:
            return "sorting"
        case AlgoritmhTypePathFinding:
            return "pathfinding"
        default:
            return ""
    }
}

type AlgorithmName struct {
	TextId   string
	Name     string
	Type     AlgorithmType
}

type Algorithm struct {
	Id          int32
	TextId      string
	Name        string
	Type        AlgorithmType
	Position    int16
	Explanation string
	HasCode     bool
}

type AlgorithmCode struct {
	Id          int32
	AlgorithmId int32
	Language    string
	Code        string
}
