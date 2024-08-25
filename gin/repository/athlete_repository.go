package repository

import (
	"errors"
	"gin/models"
)

// Mock
var athletes = map[int]models.Athlete{
	1: {ID: 1, Name: "John Doe"},
}

func GetAthleteNameByID(id int) (string, error) {
	athlete, exists := athletes[id]
	if !exists {
		return "", errors.New("athlete not found")
	}
	return athlete.Name, nil
}
