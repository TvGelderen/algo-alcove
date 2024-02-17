package models

import (
	"context"

	"github.com/TvGelderen/algo-alcove/database"
)

type AlgorithmType int16

const (
	AlgoritmhTypeSorting AlgorithmType = iota
	AlgoritmhTypePathFinding
	AlgoritmhTypeOther
)

type AlgorithmName struct {
	TextId string
	Name   string
	Type   int16
}

func ToAlgorithmName(dbModel database.GetAlgorithmNamesRow) AlgorithmName {
	return AlgorithmName{
		TextId: dbModel.TextID,
		Name:   dbModel.Name,
		Type:   dbModel.Type,
	}
}

type Algorithm struct {
	Id          int32
	TextId      string
	Name        string
	Type        AlgorithmType
	Position    int16
	Explanation string
}

func ToAlgorithm(dbModel database.Algorithm) Algorithm {
	return Algorithm{
		Id:          dbModel.ID,
		TextId:      dbModel.TextID,
		Name:        dbModel.Name,
		Type:        AlgorithmType(dbModel.Type),
		Position:    dbModel.Position,
		Explanation: dbModel.Explanation,
	}
}

func AddAlgorithmToDB(db *database.Queries, model Algorithm) (int32, error) {
	return db.CreateAlgorithm(context.Background(), database.CreateAlgorithmParams{
		TextID:      model.TextId,
		Name:        model.Name,
		Type:        int16(model.Type),
		Explanation: model.Explanation,
	})
}

type AlgorithmCode struct {
	Id          int32
	AlgorithmId int32
	Language    string
	Code        string
}

func ToAlgorithmCode(dbModel database.AlgorithmCode) AlgorithmCode {
	return AlgorithmCode{
		Id:          dbModel.ID,
		AlgorithmId: dbModel.AlgorithmID,
		Language:    dbModel.Language,
		Code:        dbModel.Code,
	}
}

func AddAlgorithmCodeToDB(db *database.Queries, model AlgorithmCode) error {
	return db.CreateAlgorithmCode(context.Background(), database.CreateAlgorithmCodeParams{
		AlgorithmID: model.AlgorithmId,
		Language:    model.Language,
		Code:        model.Code,
	})
}
