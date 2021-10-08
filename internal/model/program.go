package model

import "github.com/jinzhu/gorm"

type Program struct {
	Program_id   uint32 `gorm:"primary_key" json:"program_id"`
	Program_name string `json:"program_name"`
	Content      string `json:"content"`
	Ptype        string `json:"ptype"`
	Answer       string `json:"answer"`
	Difficulty   string `json:"difficulty"`
}

func (program *Program) TableName() string {
	return "program"
}

func (p *Program) CreateProgram(db *gorm.DB) error { //创建题目
	return db.Create(&p).Error
}

func (p *Program) UpdateProgram(db *gorm.DB) error { //更新题目
	return db.Model(&Program{}).Where("program_id = ?", p.Program_id).Update(p).Error
}

func (p *Program) DeleteByProgram_Id(db *gorm.DB) error { //根据题目id删除题目
	return db.Where("program_id = ?", p.Program_id).Delete(&p).Error
}

func (p *Program) DeleteByProgram_Name(db *gorm.DB) error { //根据题目名字删除题目
	return db.Where("program_name = ?", p.Program_name).Delete(&p).Error
}

func (p *Program) GetContent(db *gorm.DB) error { //读取题目内容
	return db.First(&p).Error
}

func (p *Program) GetAllContent(db *gorm.DB) error { //随机选取一道题
	return db.Take(&p).Error
}
