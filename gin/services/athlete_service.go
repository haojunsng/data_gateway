package services

import (
	"gin/repository"
)

func GetAthleteName(id int) (string, error) {
	return repository.GetAthleteNameByID(id)
}
