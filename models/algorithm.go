package models

import "github.com/TvGelderen/algo-alcove/database"

type Algorithm struct {
	Id          int32
	TextId      string
	Name        string
	Type        string
	Explanation string
}

func ToModel(dbModel database.Algorithm) Algorithm {
    return Algorithm{
        Id: dbModel.ID,
        TextId: dbModel.TextID,
        Name: dbModel.Name,
        Type: dbModel.Type,
        Explanation: dbModel.Explanation,
    }
}
