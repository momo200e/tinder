package service

import (
	"fmt"
	"sync"
	"tinder/domain"
)

// TODO: 待優化，目前配對查詢&刪除共用鎖，避免並發bug，但處理速度差
var serviceLock sync.Mutex

func (svc *service) AddSinglePersonAndMatch(newUser *domain.User) ([]*domain.User, *domain.ErrorFormat) {
	// 檢查是否存在用戶
	checkExistUser := svc.repo.GetUserByName(newUser.Name)
	if checkExistUser != nil {
		return nil, &domain.ErrUserAlreadyExist
	}
	// 新增用戶
	svc.repo.AddUser(*newUser)

	// 進行配對
	matches, domainErr := svc.QuerySinglePerson(newUser.Name, int(newUser.RemainDates))
	if domainErr != nil {
		return nil, domainErr
	}
	return matches, nil
}

func (svc *service) QuerySinglePerson(userName string, number int) ([]*domain.User, *domain.ErrorFormat) {
	serviceLock.Lock()
	defer serviceLock.Unlock()
	fmt.Println("進來了QuerySinglePerson", userName, number)
	user := svc.repo.GetUserByName(userName)
	if user == nil {
		return nil, &domain.ErrUserNotFound
	}

	if user.RemainDates < uint8(number) {
		return nil, &domain.ErrRemainDatesNotEnough
	}

	// 進行配對
	var findGender domain.Gender
	switch user.Gender {
	case domain.Male:
		findGender = domain.Female
	case domain.Female:
		findGender = domain.Male
	}
	isFindGreater := user.Gender == domain.Female
	fmt.Printf("findGender: %v, user.Height: %v, isFindGreater: %v, number: %v \n", findGender, user.Height, isFindGreater, number)
	matches := svc.repo.FindUsersByGenderAndHeight(findGender, user.Height, isFindGreater, number)
	fmt.Printf("matches: %+v \n", matches)
	for _, match := range matches {
		fmt.Printf("match: %+v \n", match)
		// 目前寫法不會並發，所以不會噴error
		userRemainDates, _, domainErr := svc.matchDating(user.Name, match.Name)
		if domainErr != nil {
			return nil, domainErr
		}
		if userRemainDates == 0 {
			return matches, nil
		}
	}

	return matches, nil
}

func (svc *service) RemoveSinglePerson(userName string) *domain.ErrorFormat {
	serviceLock.Lock()
	defer serviceLock.Unlock()
	if svc.repo.GetUserByName(userName) == nil {
		return &domain.ErrUserNotFound
	}
	svc.repo.DeleteUserByName(userName)
	return nil
}

// 避免match時並發
var matchLock sync.Mutex

func (svc *service) matchDating(userAName, userBName string) (uint8, uint8, *domain.ErrorFormat) {
	matchLock.Lock()
	defer matchLock.Unlock()

	fmt.Println("userAName", userAName)
	fmt.Println("userBName", userBName)
	// 約會次數扣1
	userARemainDates, err := svc.repo.UpdateUserRemainDatesDecrByName(userAName)
	if err != nil {
		return 0, 0, &domain.ErrUserNotFound
	}
	userBRemainDates, err := svc.repo.UpdateUserRemainDatesDecrByName(userBName)
	if err != nil {
		return 0, 0, &domain.ErrUserNotFound
	}
	if userARemainDates == 0 {
		// 刪除
		svc.repo.DeleteUserByName(userAName)
	}
	if userBRemainDates == 0 {
		// 刪除
		svc.repo.DeleteUserByName(userBName)
	}

	fmt.Println("userAName", userAName)
	fmt.Println("userARemainDates", userARemainDates)
	fmt.Println("userBName", userBName)
	fmt.Println("userBRemainDates", userBRemainDates)
	return userARemainDates, userBRemainDates, nil
}
