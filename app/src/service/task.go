package service

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"monitor/extend/code"
	"monitor/models"
)

// TaskService用户服务逻辑
type TaskService struct {
	TaskID uint
	UserId int
	Type   string
	Rules  string
}

// QueryByID 通过任务Id查询任务信息
func (t *TaskService) QueryByID() (task *models.Task, err error) {
	taskModel := &models.Task{}
	condition := map[string]interface{}{
		"id": t.TaskID,
	}
	task, err = taskModel.FindOne(condition)
	return
}

type TaskItem struct {
	TaskId    uint
	UserId    uint
	UserName  string
	Email     string
	TaskType  string
	Status    string
	Token     string
	Operator  string
	WarnPrice float64
}
type TaskRuleParam struct {
	Token     string
	Operator  string
	WarnPrice float64
}

// QueryByType 通过类型查询相关任务列表
func (t *TaskService) QueryByType() ([]TaskItem, error) {
	taskModel := &models.Task{}
	condition := map[string]interface{}{
		"type":   t.Type,
		"status": "ENABLE",
	}
	tasks, err := taskModel.Query(condition)
	if err != nil {
		return nil, err
	}
	resList := make([]TaskItem, len(tasks))
	for i, v := range tasks {
		log.Debug().Msgf("rule: %v, type: %T", v.Rules, v.Rules)
		rule := TaskRuleParam{}
		err := json.Unmarshal([]byte(v.Rules), &rule)
		if err != nil {
			log.Error().Msgf(err.Error())
			return nil, err
		}
		resList[i] = TaskItem{
			TaskId:    v.ID,
			UserId:    v.User.ID,
			UserName:  v.User.UserName,
			Email:     v.User.Email,
			TaskType:  v.Type,
			Status:    v.Status,
			Token:     rule.Token,
			Operator:  rule.Operator,
			WarnPrice: rule.WarnPrice,
		}
	}
	return resList, nil
}

// StoreTask 添加任务
func (t *TaskService) StoreTask() (taskId uint, err error) {
	log.Info().Msg(" enter StoreTask service ")
	task := &models.Task{
		UserID: t.UserId,
		Type:   t.Type,
		Status: "ENABLE",
		Rules:  t.Rules,
	}
	taskId, err = task.Insert()
	return
}

// UpdateInfo 更新任务信息
func (t *TaskService) UpdateInfo(taskId uint) (*models.Task, *code.Code) {
	taskModel := &models.Task{}
	updateTask, err := taskModel.UpdateOne(taskId, map[string]interface{}{
		"type":  t.Type,
		"rules": t.Rules,
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, code.ServiceInsideError
	}
	return updateTask, nil
}

func (t *TaskService) RemoveTask(taskID uint) error {
	taskModel := &models.Task{}
	err := taskModel.DeleteOne(taskID)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskService) QueryByPage(condition interface{}, page, pageSize int) ([]TaskItem, int, error) {
	taskModel := &models.Task{}
	taskList, err := taskModel.Search(condition, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	resList := make([]TaskItem, len(taskList))
	for i, v := range taskList {
		log.Debug().Msgf("rule: %v, type: %T", v.Rules, v.Rules)
		rule := TaskRuleParam{}
		err := json.Unmarshal([]byte(v.Rules), &rule)
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, 0, err
		}

		resList[i] = TaskItem{
			TaskId:    v.ID,
			UserId:    v.User.ID,
			UserName:  v.User.UserName,
			Email:     v.User.Email,
			TaskType:  v.Type,
			Status:    v.Status,
			Token:     rule.Token,
			Operator:  rule.Operator,
			WarnPrice: rule.WarnPrice,
		}
	}

	count, err := taskModel.Count(condition)
	if err != nil {
		return nil, 0, err
	}
	return resList, count, nil
}
