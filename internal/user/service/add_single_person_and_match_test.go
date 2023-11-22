package service

import (
	"fmt"
	"testing"
	"tinder/domain"
	"tinder/internal/user/repo"
)

func TestAddSinglePersonAndMatch(t *testing.T) {
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
	fmt.Println("---------------------mock done------------------------")

	newUser := domain.User{Name: "Lucy", Height: 177, Gender: domain.Female, RemainDates: 5}

	matches, errFormat := svc.AddSinglePersonAndMatch(&newUser)
	if errFormat != nil {
		t.Errorf("Expected err format nil, but got: %v", errFormat)
	}

	if len(matches) != 1 {
		t.Errorf("Expected 1 user, got %d", len(matches))
	}

	if matches[0].Name != "Tom" {
		t.Errorf("Expected match user name Tom, got %v", matches[0].Name)
	}

	newUser2 := domain.User{Name: "Jason", Height: 168, Gender: domain.Male, RemainDates: 3}

	matches2, errFormat := svc.AddSinglePersonAndMatch(&newUser2)
	if errFormat != nil {
		t.Errorf("Expected err format nil, but got: %v", errFormat)
	}

	if len(matches2) != 1 {
		t.Errorf("Expected 1 user, got %d", len(matches2))
	}

	if matches2[0].Name != "Lily" {
		t.Errorf("Expected match user name Lily, got %v", matches2[0].Name)
	}

	userLucy := repo.GetUserByName("Lucy")
	if userLucy == nil {
		t.Error("Expected user Lucy, got noting")
	}

	if userLucy.RemainDates != 4 {
		t.Errorf("Expected user Lucy remaining dates 4, got %d", userLucy.RemainDates)
	}

	userTom := repo.GetUserByName("Tom")
	if userTom != nil {
		t.Errorf("Expected user Tom to be deleted, got %s", userTom.Name)
	}

	newUser3 := domain.User{Name: "Lucy", Height: 177, Gender: domain.Female, RemainDates: 5}

	_, errFormat = svc.AddSinglePersonAndMatch(&newUser3)
	if errFormat != &domain.ErrUserAlreadyExist {
		t.Errorf("Expected err format to be ErrUserAlreadyExist, but got: %v", errFormat)
	}
}
