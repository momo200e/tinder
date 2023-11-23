package repo

import (
	"fmt"
	"math"
	"sync"
	"tinder/domain"

	avl "github.com/emirpasic/gods/trees/avltree"
)

// 使用平衡樹是為了高效的身高查詢，增/刪/查的時間複雜度O(log n)
// Map[UserName] = user info
// AVL樹節點組成：key:身高, value=[]*domain.User
type repository struct {
	Users       map[string]*domain.User
	maleAVL     *avl.Tree
	femaleAVL   *avl.Tree
	*sync.Mutex // TODO: 待優化：改成分多個讀寫鎖
}

func NewUserRepository() domain.Repository {
	return &repository{
		Users:     map[string]*domain.User{},
		maleAVL:   avl.NewWithIntComparator(),
		femaleAVL: avl.NewWithIntComparator(),
		Mutex:     &sync.Mutex{},
	}
}

func (repo *repository) AddUser(user domain.User) {
	repo.Lock()
	defer repo.Unlock()

	var avlTree *avl.Tree
	switch user.Gender {
	case domain.Male:
		avlTree = repo.maleAVL
	case domain.Female:
		avlTree = repo.femaleAVL
	default:
		fmt.Println("invalid gender")
		return
	}

	result, _ := avlTree.Get(int(user.Height))

	// 更新userId
	// insert user
	repo.Users[user.Name] = &user

	if result == nil {
		avlTree.Put(int(user.Height), []*domain.User{repo.Users[user.Name]})
	} else {
		userList := result.([]*domain.User)
		userList = append(userList, repo.Users[user.Name])
		avlTree.Put(int(user.Height), userList)
	}
}

func (repo *repository) DeleteUserByName(name string) {
	repo.Lock()
	defer repo.Unlock()

	user := repo.Users[name]

	var avlTree *avl.Tree
	switch user.Gender {
	case domain.Male:
		avlTree = repo.maleAVL
	case domain.Female:
		avlTree = repo.femaleAVL
	default:
		fmt.Println("invalid gender")
		return
	}

	result, _ := avlTree.Get(int(user.Height))
	if result == nil {
		fmt.Println("not found user height")
		return
	} else {
		userList := result.([]*domain.User)
		// 將user從slice中刪除
		if len(userList) == 1 {
			avlTree.Remove(int(user.Height))
		} else {
			for i, u := range userList {
				if u.Name == user.Name {
					// 高效刪除
					userList[i] = userList[len(userList)-1]
					userList = userList[:len(userList)-1]
					break
				}
			}
			avlTree.Put(int(user.Height), userList)
		}

		// 將user從map中移除
		delete(repo.Users, name)

	}
}

func (repo *repository) GetUserByName(name string) *domain.User {
	return repo.Users[name]
}

func (repo *repository) UpdateUserRemainDatesDecrByName(name string) (uint8, error) {
	repo.Lock()
	defer repo.Unlock()

	if repo.Users[name] != nil {
		repo.Users[name].RemainDates--
		return repo.Users[name].RemainDates, nil
	} else {
		return 0, fmt.Errorf("user not found")
	}
}

// FindUsersByGenderAndHeight 根據性別和身高查找用戶。
// 如果 isFindGreater 為 false，則查找身高小於 height 的用戶。
// count 參數限制了返回的用戶數量。
//
// 參數:
//     gender: 用於篩選的性別（男/女）。
//     height: 用於篩選的身高。
//     isFindGreater: 決定是否查找身高大於 height 的用戶。
//     count: 限制返回的用戶數量。設定0則回傳全部
//
// 返回:
//     符合條件的用戶列表。
func (repo *repository) FindUsersByGenderAndHeight(gender domain.Gender, height uint8, isFindGreater bool, count int) []*domain.User {
	if count == 0 {
		// 預設回傳數無限制
		count = math.MaxInt
	}
	var rootNode *avl.Node
	if gender == domain.Male {
		rootNode = repo.maleAVL.Root
	} else if gender == domain.Female {
		rootNode = repo.femaleAVL.Root
	} else {
		fmt.Println("invalid gender")
		return nil
	}

	if rootNode == nil {
		return nil
	}

	matchedUsers := []*domain.User{}

	// reverse = true 代表root左邊或右邊都以遍歷完
	reverse := false
	node := rootNode
	for count > 0 {
		if isFindGreater {
			// 找出身高 > height的user
			if node.Key.(int) > int(height) {
				users := node.Value.([]*domain.User)
				if count < len(users) {
					// 結束
					matchedUsers = append(matchedUsers, users[:count-1]...)
					return matchedUsers
				} else {
					matchedUsers = append(matchedUsers, users...)
					count -= len(users)
					if count == 0 {
						// 結束
						return matchedUsers
					}
				}
			}

			// 先找root右邊
			if !reverse {
				// 換下個節點
				node = node.Next()
				if node == nil {
					// 翻轉開始遍歷root另一側
					reverse = true
					node = rootNode.Prev()
					if node == nil {
						// 全部查完
						// 結束
						return matchedUsers
					}
				}
			} else {
				// 翻轉找另一邊
				node = node.Prev()
				if node == nil {
					// 全部查完
					// 結束
					return matchedUsers
				}
			}
		} else {
			// 找出身高 < height的user
			if node.Key.(int) < int(height) {
				users := node.Value.([]*domain.User)
				if count < len(users) {
					// 結束
					matchedUsers = append(matchedUsers, users[:count-1]...)
					return matchedUsers
				} else {
					matchedUsers = append(matchedUsers, users...)
					count -= len(users)
					if count == 0 {
						// 結束
						return matchedUsers
					}
				}
			}

			// 先找root左邊
			if !reverse {
				// 換下個節點
				node = node.Prev()
				if node == nil {
					// 翻轉開始遍歷root另一側
					reverse = true
					node = rootNode.Next()
					if node == nil {
						// 全部查完
						// 結束
						return matchedUsers
					}
				}
			} else {
				// 翻轉找另一邊
				node = node.Next()
				if node == nil {
					// 全部查完
					// 結束
					return matchedUsers
				}
			}
		}
	}

	// TODO: 待優化，不要回傳指針
	return matchedUsers
}
