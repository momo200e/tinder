package service

import (
	"testing"
	"tinder/domain"
	"tinder/internal/user/repo"
)

func TestQuerySinglePerson(t *testing.T) {
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
			Name:        "Ken",
			Height:      178,
			Gender:      domain.Male,
			RemainDates: 12,
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
		{
			Name:        "Lucy",
			Height:      177,
			Gender:      domain.Female,
			RemainDates: 5,
		},
		{
			Name:        "Guan",
			Height:      190,
			Gender:      domain.Female,
			RemainDates: 44,
		},
	}

	for _, u := range defaultUsers {
		repo.AddUser(u)
	}

	matches, errFormat := svc.QuerySinglePerson("Lucy", 2)
	if errFormat != nil {
		t.Errorf("Expected err format nil, but got: %v", errFormat)
	}

	if len(matches) != 2 {
		t.Errorf("Expected 2 user, got %d", len(matches))
	}

	for _, u := range matches {
		if u.Name != "Tom" && u.Name != "Ken" {
			t.Errorf("Expected match user name Tom or Ken, got %v", u.Name)
		}
	}

	userTom := repo.GetUserByName("Tom")
	if userTom != nil {
		t.Errorf("Expected user Tom to be deleted, but got %v", userTom)
	}

	userLucy := repo.GetUserByName("Lucy")
	if userLucy.RemainDates != 3 {
		t.Errorf("Expected user Lucy remaining dates to be 3, but got %v", userLucy.RemainDates)
	}

	matches, errFormat = svc.QuerySinglePerson("Lucy", 3)
	if errFormat != nil {
		t.Errorf("Expected err format nil for match 3, but got: %v", errFormat)
	}

	if len(matches) != 1 {
		t.Errorf("Expected 1 user, got %d", len(matches))
	}

	userLucy = repo.GetUserByName("Lucy")
	if userLucy.RemainDates != 2 {
		t.Errorf("Expected user Lucy remaining dates to be 2, but got %v", userLucy.RemainDates)
	}

	_, errFormat = svc.QuerySinglePerson("Lucy", 1)
	if errFormat != nil {
		t.Errorf("Expected err format nil, but got: %v", errFormat)
	}
	_, errFormat = svc.QuerySinglePerson("Lucy", 1)
	if errFormat != nil {
		t.Errorf("Expected err format nil, but got: %v", errFormat)
	}

	userLucy = repo.GetUserByName("Lucy")
	if userLucy != nil {
		t.Errorf("Expected user Lucy to be deleted, but got %v", userLucy)
	}

	_, errFormat = svc.QuerySinglePerson("Lucy", 1)
	if errFormat == nil {
		t.Errorf("Expected err format not nil, but got: nil")
	}

	matches2, errFormat := svc.QuerySinglePerson("Ken", 2)
	if errFormat != nil {
		t.Errorf("Expected err format nil for match 2, but got: %v", errFormat)
	}

	if len(matches2) != 2 {
		t.Errorf("Expected 2 user, got %d", len(matches2))
	}

	for _, u := range matches2 {
		if u.Name != "Lily" && u.Name != "Joan" && u.Name != "Lucy" {
			t.Errorf("Expected match user name Lily or Joan or Lucy, got %v", u.Name)
		}
	}

	matches3, errFormat := svc.QuerySinglePerson("Ken", 5)
	if errFormat != nil {
		t.Errorf("Expected err format nil for match 3, but got: %v", errFormat)
	}

	if len(matches3) != 2 {
		t.Errorf("Expected 2 user, got %d", len(matches3))
	}

}
