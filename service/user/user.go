package user

import (
	"errors"
	"gameday/db/model"
	"gameday/global"
	"time"
)

func (us *UserService) UpdateUsername(user *model.User) error {
	err := global.GameDB.Model(&model.User{}).Where("hashcode = ?", user.Hashcode).Updates(map[string]interface{}{
		"username": user.Username,
		"status":   1,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) GetUserByHash(user *model.User) (*model.User, error) {
	if err := global.GameDB.Model(&model.User{}).Where("hashcode", user.Hashcode).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetSuccessQuestionByUserId(s *model.Complete) ([]uint, error) {
	//var score []*model.Score
	var completeIDs []uint
	if err := global.GameDB.
		Model(&model.Complete{}).
		Where("user_id = ?", s.UserId).
		Pluck("question_id", &completeIDs).Error; err != nil {
		return nil, err
	}

	return completeIDs, nil
}

func (us *UserService) UserScore(s *model.Complete) (uint, error) {
	var score uint
	if err := global.GameDB.Model(&model.Complete{}).
		Select("COALESCE(SUM(score), 0)"). // 默认值，如果字段sum(score)为NULL则设置默认值为0
		Where("user_id = ?", s.UserId).
		Scan(&score).
		Error; err != nil {
		return 0, err
	}

	return score, nil
}

func (us *UserService) StopGame(GameId uint) (*model.Game, error) {
	var game *model.Game
	if err := global.GameDB.Where("id = ?", GameId).Find(&game).Error; err != nil {
		return nil, err
	}

	return game, nil
}

func (us *UserService) PostAnswer(request *model.Question) error {
	var quest model.Question
	if err := global.GameDB.Where("id = ?", request.ID).Find(&quest).Error; err != nil {
		return err
	}

	if request.Answer == quest.Answer {
		return nil

	}

	return errors.New("密码答案错误")
}

func (us *UserService) StoreComplete(userId, questionId, grade uint) error {
	if err := global.GameDB.Create(&model.Complete{
		QuestionId: questionId,
		Score:      grade,
		UserId:     userId,
	}).Error; err != nil {
		return err
	}

	return nil
}

type result struct {
	Username string `json:"username"`
	Count    int    `json:"count"`
	Sum      int    `json:"sum"`
}

func (us *UserService) RankingList(user *model.User) ([]result, error) {

	var results []result

	// select `user`.username, COUNT(*) ,SUM(complete.score)
	// from `user` left join `complete`
	// on `user`.id = complete.user_id
	// where game_id = 9
	//group by `user`.id

	if err := global.GameDB.Table("user").
		Select("`user`.username, IFNULL(COUNT(complete.user_id), 0) as `count` ,SUM(complete.score) as `sum`").
		Joins("LEFT JOIN complete ON `user`.id = complete.user_id").
		Where("game_id = ?", user.GameID).
		Group("`user`.id").
		Order("`sum` DESC").
		Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

type result2 struct {
	model.Complete
	model.Question
	CreatedAt time.Time `json:"created_at""`
}

func (us *UserService) UserThing(complete *model.Complete) ([]result2, error) {
	var results []result2
	//select question.title, complete.score, complete.created_at
	//	from `complete` inner join `question`
	//	on `complete`.question_id = question.id
	//	where `complete`.user_id = 126
	if err := global.GameDB.Table("complete").
		Select("question.title as title, complete.score as score, complete.created_at created_at").
		Joins("INNER JOIN question ON complete.question_id = question.id").
		Where("complete.user_id = ?", complete.UserId).
		Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

type result3 struct {
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	Score     uint      `json:"score"`
	Title     string    `json:"title"`
}

func (us *UserService) Event(user *model.User) ([]result3, error) {
	var results []result3
	//select complete.created_at as created_at,
	//	`user`.username as username,
	//	complete.score as score,
	//	question.title as title
	//	from complete
	//	inner join `user`
	//	on complete.user_id = `user`.id
	//	inner join question
	//	on complete.question_id = question.id
	//	Where `user`.game_id = 9
	if err := global.GameDB.Table("complete").
		Select("complete.created_at as created_at,`user`.username as username, complete.score as score, question.title as title").
		Joins("inner join `user` on complete.user_id = `user`.id").
		Joins("inner join question on complete.question_id = question.id").
		Where("`user`.game_id = ?", user.GameID).
		Scan(&results).
		Error; err != nil {
		return nil, err
	}
	return results, nil
}
