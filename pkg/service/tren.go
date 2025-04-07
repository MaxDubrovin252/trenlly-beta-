package service

import (
	"trenlly"
	"trenlly/pkg/repository"

	"github.com/sirupsen/logrus"
)

type TrenService struct {
	repo repository.Trening
}

func NewTreningService(repo repository.Trening) *TrenService {
	return &TrenService{repo: repo}
}
func (s *TrenService) CreateTren(tren trenlly.Trening, userId int) (int, *trenlly.Result, error) {

	result, err := NewResult(tren.Weight, tren.Height, tren.Time)

	if err != nil {
		logrus.Error(err)
		return 0, result, err
	}

	id, err := s.repo.CreateTren(tren, userId)

	if err != nil {
		return 0, result, err
	}

	return id, result, nil

}

func (s *TrenService) Delete(userId, trenId int) error {
	return s.repo.Delete(userId, trenId)
}
func (s *TrenService) UpdateTren(userId, trenId int, input trenlly.UpdateTrenItem) error {

	if err := input.Validate(); err != nil {
		logrus.Error(err.Error())

	}
	return s.repo.UpdateTren(userId, trenId, input)
}

func (s *TrenService) GetAll(userId int) ([]trenlly.Trening, error) {
	return s.repo.GetAll(userId)
}

func NewResult(weight, height, time int) (*trenlly.Result, error) {
	var result trenlly.Result

	switch weight {

	case 97, 98, 99, 100:
		{
			if height >= 190 {
				result.Recomendation = "no recomendation,if you want add cardio"
			}

			if height <= 180 {
				result.Recomendation = "you have a little overweiht"
				result.Solution = "add basketball or swim in your life"

			}
		}
	case 90, 91, 92, 93, 94, 95, 96:
		if height >= 185 {
			result.Recomendation = "you have a perfect body"
			result.Solution = "if you want upgrade your body or discipline and self-defense add a boxing in your life"
			result.Terms = "6-10 mouhts"
		}

		if height <= 175 && height > 160 {
			result.Recomendation = "oh you have a overweight, go to the gym or cadrio"
			result.Solution = "on this week buy a subscription for gym or swimming pool"
			result.Terms = "6-12 monts"

		}

		if height <= 160 {
			result.Recomendation = "very fast go to doctor you have problem with overweight"
			result.Solution = "drink more water, put away cola or each other from your diet"
			result.Terms = "1-5 years"
		}
	case 80, 81, 82, 83, 84, 85, 86, 87, 88, 89:
		if height >= 175 && height < 185 {
			result.Recomendation = "you have really good form"
			result.Solution = "add any cardio or boxing in your sport routine"
			result.Terms = "5-8 months"
		}

		if height <= 160 {
			result.Recomendation = "your height is low and your have overweight"
			result.Solution = "add cardio and tren more"
			result.Terms = "6-12 months"
		}

		if height > 185 {
			result.Recomendation = "you have good body structure"
			result.Solution = "if you want add cardio"

		}

	case 70, 71, 72, 73, 74, 75, 76, 77, 78, 79:
		if height >= 180 {
			result.Recomendation = "you have really good form"
			result.Solution = "add any cardio or boxing in your sport routine"
			result.Terms = "5-8 months"
		}

		if height < 180 {
			result.Recomendation = "you have good form but if youn want you can upgrade this"
			result.Solution = "add any another sport in your life for example:boxing or swim"
			result.Terms = "5-10 months"
		}

		if height <= 170 && height > 160 {
			result.Recomendation = "you have a normal body"
			result.Solution = "if you want a upgrade your body go to the gym"
			result.Terms = "1 years"
		}

		if height < 160 {
			result.Recomendation = "add a basketball or swim in your life"
			result.Terms = "7-15 months"
		}

	case 60, 61, 62, 63, 64, 65, 66, 67, 68, 69:
		if height >= 180 {
			result.Recomendation = "you have a problem with weight"
			result.Solution = "visit a doctor at this week"

		}

		if height <= 170 && height > 160 {
			result.Recomendation = "you have a normal body"
			result.Solution = "if you want a upgrade your body go to the gym"
			result.Terms = "1 years"
		}

		if height < 160 {
			result.Recomendation = "add a basketball or swim in your life"
			result.Terms = "7-15 months"
		}

	}
	return &trenlly.Result{
		Recomendation: result.Recomendation,
		Solution:      result.Solution,
		Terms:         result.Terms,
		Comment:       result.Comment,
	}, nil
}

func (s *TrenService) GetById(userId, trenId int) (trenlly.Trening, error) {
	return s.repo.GetById(userId, trenId)
}
