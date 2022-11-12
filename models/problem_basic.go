package models

import (
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	ID         uint   `gorm:"primarykey;" json:"id"`
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`   // 问题表的唯一标识
	Title      string `gorm:"column:title;type:varchar(255);" json:"title"`        // 文章标题
	Content    string `gorm:"column:content;type:text;" json:"content"`            // 文章正文
	MaxRuntime int    `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"` // 最大运行时长
	MaxMem     int    `gorm:"column:max_mem;type:int(11);" json:"max_mem"`         // 最大运行内存
	PassNum    int64  `gorm:"column:pass_num;type:int(11);" json:"pass_num"`       // 通过次数
	SubmitNum  int64  `gorm:"column:submit_num;type:int(11);" json:"submit_num"`   // 提交次数
}

func (table *Problem) TableName() string {
	return "problem"
}

func GetProblemList(keyword string) *gorm.DB {
	return DB.Model(new(Problem)).Where("title like ? OR content like ? ", "%"+keyword+"%", "%"+keyword+"%")
}
