package models

import (
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
	Explanation string
}

func ToAlgorithm(dbModel database.Algorithm) Algorithm {
	return Algorithm{
		Id:          dbModel.ID,
		TextId:      dbModel.TextID,
		Name:        dbModel.Name,
		Type:        AlgorithmType(dbModel.Type),
		Explanation: dbModel.Explanation,
	}
}
