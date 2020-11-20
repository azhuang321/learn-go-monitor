package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Task struct {
	gorm.Model
	User   User   `gorm:"ForeignKey:UserId;AssociationForeignKey:ID"`
	UserID int    `gorm:"column:userId;not null"`
	Type   string `sql:"type:ENUM('TICKER','OTHER')"`
	Status string `sql:"type:ENUM('ENABLE','DISABLE')"`
	Rules  string `gorm:"column:rules;type:varchar(200);not null"`
}

// Insert 新增任务
func (t *Task) Insert() (taskId uint, err error) {
	result := DB.Create(&t)
	taskId = t.ID
	if result.Error != nil {
		err = result.Error
	}
	return
}

// FindOne 查询任务信息
func (t *Task) FindOne(condition map[string]interface{}) (*Task, error) {
	var taskInfo Task
	result := DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name,email,avatar,status")
	}).Select("id,userId,type,type,status,rules").Where(condition).First(&taskInfo)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	if taskInfo.ID > 0 {
		return &taskInfo, nil
	}
	return nil, nil
}

// UpdateOne 更新任务
func (t *Task) UpdateOne(taskId uint, data map[string]interface{}) (*Task, error) {
	err := DB.Model(&Task{}).Where("id = ?", taskId).Update(data).Error
	if err != nil {
		return nil, err
	}
	var updTask Task
	err = DB.Select([]string{"id", "userId", "type", "status", "rules"}).First(&updTask, taskId).Error
	if err != nil {
		return nil, err
	}
	return &updTask, nil
}

// DeleteOne 删除任务
func (t *Task) DeleteOne(taskId uint) error {
	if err := DB.Select([]string{"id"}).First(&t, taskId).Error; err != nil {
		return err
	}
	if err := DB.Delete(&t).Error; err != nil {
		return err
	}
	return nil
}

// Query 无分页查询
func (t *Task) Query(query map[string]interface{}) ([]*Task, error) {
	var tasks []*Task
	err := DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name,email,avatar,status")
	}).Select("id, userId, type, status, rules").Where(query).Find(&tasks).Error
	return tasks, errors.WithStack(err)
}

// Search 分页数据查询
func (t *Task) Search(query interface{}, page int, pageSize int) ([]*Task, error) {
	var tasks []*Task
	err := DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name,email,avatar,status")
	}).Select("id, userId, type, status, rules").Offset(pageSize * (page - 1)).Limit(pageSize).Where(query).Find(&tasks).Error
	return tasks, errors.WithStack(err)
}

// Count 分页总数查询
func (t *Task) Count(query interface{}) (int, error) {
	var count int
	err := DB.Model(&Task{}).Where(query).Count(&count).Error
	return count, errors.WithStack(err)
}
