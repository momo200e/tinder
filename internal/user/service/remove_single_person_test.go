package service

import (
	"testing"
	"tinder/domain"
	"tinder/internal/user/repo"
)

func TestRemoveSinglePerson(t *testing.T) {
	repo := repo.NewUserRepository()
	svc := NewService(repo)

	defaultUsers := []domain.User{
		{
			Name:        "Tom",
			Height:      180,
			Gender:      domain.Male,
			RemainDates: 1,
		},
		{
			Name:        "Lily",
			Height:      160,
			Gender:      domain.Female,
			RemainDates: 3,
		},
		{
			Name:        "Paul",
			Height:      176,
			Gender:      domain.Male,
			RemainDates: 10,
		},
		{
			Name:        "Joan",
			Height:      170,
			Gender:      domain.Female,
			RemainDates: 7,
		},
	}

	for _, u := range defaultUsers {
		repo.AddUser(u)
	}

	svc.RemoveSinglePerson("Tom")

	testTom := repo.GetUserByName("Tom")
	if testTom != nil {
		t.Errorf("Expected user Tom to be deleted, but got: %v", testTom)
	}

	svc.RemoveSinglePerson("Tom")

}
