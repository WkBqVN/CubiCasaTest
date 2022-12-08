package database

import (
	"cubicasa/models"
	"database/sql"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

// UserRepository db must not nil , if nil all service will return false
// --99 number marker failed
// can return a number marker(not only bool) for check what type of error raise
// member must have team id to create
type UserRepository struct {
	db     *gorm.DB
	config Config
}

func (userRepo *UserRepository) InitDB(dataBase string) {
	userRepo.config = CreateConfig(dataBase)
	userRepo.db = ConnectToDB(userRepo.config)
}
func (userRepo *UserRepository) GetAllUser() ([]models.User, bool) {
	if userRepo.db == nil {
		return nil, false
	}
	var users []models.User
	result := userRepo.db.Find(&users)
	if result.Error != nil {
		return nil, false
	} else {
		return users, true
	}
}

func (userRepo *UserRepository) GetUserById(id int) (models.User, bool) {
	var user models.User
	if userRepo.db == nil {
		return user, false
	}
	result := userRepo.db.First(&user, id)
	if result.Error != nil {
		return user, false
	} else {
		return user, true
	}
}

func (userRepo *UserRepository) UpdateUser(id int, user models.User) (int, bool) {
	if ValidateId(id) {
		userData, isExist := userRepo.GetUserById(id)
		if isExist != true {
			return id, false
		}
		if !validateUser(user) {
			if user.TeamId == 0 {
				user.TeamId = userData.TeamId
			}
		}
	} else {
		return -99, false
	}
	if userRepo.db == nil {
		return -99, false
	}
	result := userRepo.db.Where("user_id = ?", id).Updates(&user)
	if result.Error != nil {
		return id, false
	} else {
		return id, true
	}
}

func (userRepo *UserRepository) CreateUser(user models.User) (models.User, bool) {
	if userRepo.db == nil || !validateUser(user) {
		return user, false
	}
	if user.TeamId == 0 || user.TeamId > 100 {
		return user, false
	}
	result := userRepo.db.Create(&user)
	if result.Error != nil {
		return user, false
	} else {
		return user, true
	}
}

func (userRepo *UserRepository) DeleteUser(id int) (int, bool) {
	if userRepo.db == nil {
		return -99, false
	}
	_, isExist := userRepo.GetUserById(id)
	if isExist != true {
		return id, false
	}
	result := userRepo.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return id, false
	} else {
		return id, true
	}
}

/*
cause design was need team id on first team create so join user(no team_Id) to new team function will not work
,so I changed to move user between team logic still be same (not enough time to change design)
*/

func (userRepo *UserRepository) MoveUser(userId int, teamId int) bool {
	resultData, isDone := userRepo.GetUserById(userId)
	if isDone == true {
		resultData.TeamId = uint(teamId)
		_, isDoneUser := userRepo.UpdateUser(userId, resultData)
		if isDoneUser == true {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// SearchUsername /*
func (userRepo *UserRepository) SearchUsername(userName string) (map[string]string, bool) {
	// userName can't have special char
	if !MatchRegex(userName) {
		return nil, false
	}
	userNameQuery := "%" + userName + "%"
	rows, err := userRepo.db.Raw("select users.user_name,users.role,teams.team_name, hubs.hub_name from testdata.users"+
		" left join testdata.teams on  testdata.users.team_id  = testdata.teams.team_id "+
		" left join testdata.hubs on testdata.teams.hub_id = testdata.hubs.hub_id"+
		" where testdata.users.user_name like ? ", userNameQuery).Rows()
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	mapResult := map[string]string{
		"userName": "",
		"role":     "",
		"teamName": "",
		"hubName":  "",
	}
	if err != nil {
		return nil, false
	} else {
		var userNameRaw string
		var roleRaw string
		var teamNameRaw string
		var hubNameRaw string
		for rows.Next() {
			err := rows.Scan(&userNameRaw, &roleRaw, &teamNameRaw, &hubNameRaw)
			if err != nil {
				return nil, false
			}
		}
		mapResult["userName"] = userNameRaw
		mapResult["role"] = roleRaw
		mapResult["teamName"] = teamNameRaw
		mapResult["hubName"] = hubNameRaw
		if mapResult["userName"] == "" {
			return nil, false
		} else {
			return mapResult, true
		}
	}
}

func validateUser(user models.User) bool {
	// update is bind on param path not put in json
	if user.TeamId == 0 {
		return false
	}
	if user.UserID != 0 {
		return false
	}
	if (strings.Trim(user.UserName, " ") == "" || len(user.UserName) > 20 || !MatchRegex(user.UserName)) &&
		(strings.Trim(user.Role, " ") == "" || len(user.UserName) > 20 || !MatchRegex(user.Role)) {
		return false
	}
	return true
}

func MatchRegex(inputString string) bool {
	matched, _ := regexp.MatchString("[a-zA-Z0-9]", inputString)
	return matched
}

func ValidateId(id int) bool {
	if id < 0 || id > 100 {
		return false
	}
	return true
}
