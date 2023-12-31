package game

import (
	"errors"
	"fmt"
	"gameday/db/model"
	"gameday/global"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gameGroup struct {
}

// CreateGame 创建比赛，根据比赛名创建ID
func (g *gameGroup) CreateGame(game *model.Game) (*model.Game, error) {
	// 判断数据库是否有
	if affectedRow := global.GameDB.Where("name = ?", game.Name).First(game).RowsAffected; affectedRow == 0 {
		var users []model.User

		// 创建用户hash
		for i := 0; i < int(game.People); i++ {
			user := &model.User{}
			code := uuid.New()
			user.Hashcode = code.String()
			users = append(users, *user)
		}

		// 创建数据库记录
		game.Hashes = users
		if err := global.GameDB.Create(game).Error; err != nil {
			return nil, err
		}

		return game, nil
	}
	return nil, errors.New("重复创建比赛")
}

func (g *gameGroup) StartGame(gameId uint) error {
	err := global.GameDB.Table("game").Where("id = ?", gameId).Update("status", 1).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *gameGroup) StopGame(gameId uint) error {
	err := global.GameDB.Table("game").Where("id = ?", gameId).Update("status", 2).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *gameGroup) DoneGame(gameId uint) error {
	err := global.GameDB.Table("game").Where("id = ?", gameId).Update("status", 3).Error
	if err != nil {
		return err
	}

	return nil
}

// LastGame 查找最后一场比赛
func (g *gameGroup) LastGame() (*model.Game, error) {
	var game *model.Game
	// SELECT * FROM game order by id desc limit 1
	if err := global.GameDB.Preload("Hashes").Last(&game).Order("gameid").Error; err != nil {
		return nil, err
	}

	return game, nil
}

// GetHash 返回此比赛的所有hash
func (g *gameGroup) GetHash(gameId uint) (*model.Game, error) {
	var game *model.Game
	err := global.GameDB.Preload("Hashes").Where("id = ?", gameId).Find(&game).Error
	if err != nil {
		return nil, err
	}
	return game, nil
}

// GetAllGame 返回所有的比赛
func (g *gameGroup) GetAllGame() ([]*model.Game, error) {
	var games []*model.Game
	if err := global.GameDB.Preload("Group").Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}

// RankList 返回这个比赛的排行榜
func (g *gameGroup) RankList(GameId uint) (*model.Game, error) {
	var game *model.Game

	// 骚操作，自定义Preload, 返回排序好的关系，在这里也就是返回属于game_id的User（按照score倒序）
	err := global.GameDB.Preload("Hashes", func(db *gorm.DB) *gorm.DB {
		return db.Order("user.score DESC")
	}).Where("id = ?", GameId).Find(&game).Error
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (g *gameGroup) ReviewUserName(user *model.User) error {

	if err := global.GameDB.Table("user").Where("id = ?", user.ID).Update("status", 2).Error; err != nil {
		return err
	}

	return nil
}

func (g *gameGroup) RefuseUserName(user *model.User) error {

	if err := global.GameDB.Table("user").Where("id = ?", user.ID).Update("status", 3).Error; err != nil {
		return err
	}

	return nil
}

func (g *gameGroup) GetQuestionsByGameId(game *model.Game) (*model.Group, error) {
	var group *model.Group
	if err := global.GameDB.Model(&model.Game{}).Where("id = ?", game.ID).Find(&game).Error; err != nil {
		return nil, err
	}
	if err := global.GameDB.Model(&model.Group{}).Preload("Questions").Where("id = ?", game.GroupID).Find(&group).Error; err != nil {
		return nil, err
	}

	fmt.Println(group)
	return group, nil
}
