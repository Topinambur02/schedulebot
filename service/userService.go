package service

import (
	"log"
	"time"

	"topinambur02.com/m/v2/db"
	"topinambur02.com/m/v2/model"
)

type UserService struct {
    Db *db.Database
}

func (s *UserService) GetOrCreateUser(tgId int64, username string) (model.User, error) {
    user, err := s.Db.FindByTgID(tgId)
    
    if err != nil {
        newUser := model.User{
            Tg_id:    tgId,
            Username: username,
			Role: "USER",
			CreatedAt: time.Now(),
        }
        if err := s.Db.Create(&newUser); err != nil {
            return model.User{}, err
        }
        log.Printf("New user created: %s (ID: %d)", username, tgId)
        return newUser, nil
    } else if user.Username != username {
        user.Username = username
        if err := s.Db.Update(user); err != nil {
            return *user, err
        }
    }
    
    return *user, nil
}
