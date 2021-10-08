package dao

import (
	"webDemo/internal/model"
)

func (p *Dao) CreateProgram(program_id uint32, program_name string, content string, ptype string, answer string, difficulty string) error {
	program := model.Program{
		Program_id:   program_id,
		Program_name: program_name,
		Content:      content,
		Ptype:        ptype,
		Answer:       answer,
		Difficulty:   difficulty,
	}
	return program.CreateProgram(p.engine)
}

func (p *Dao) UpdateProgram(program_id uint32, program_name string, content string, ptype string, answer string, difficulty string) error {
	program := model.Program{
		Program_id:   program_id,
		Program_name: program_name,
		Content:      content,
		Ptype:        ptype,
		Answer:       answer,
		Difficulty:   difficulty,
	}
	return program.UpdateProgram(p.engine)
}
