package service

type CreateProgramRequest struct {
	Program_id   uint32 `form:"program_id" binding:"oneof=0 1"`
	Program_name string `form:"program_name" binding:"max=100"`
	Content      string `form:"content" binding:"required,min=10,max=500"`
	Ptype        string `form:"ptype" binding:"required,min=2,max=50"`
	Answer       string `form:"answer" binding:"required,min=1,max=100"`
	Difficulty   string `form:"difficulty" binding:"required,min=1,max=50"`
}

func (svc *Service) CreateProgram(param *CreateProgramRequest) error {
	return svc.dao.CreateProgram(param.Program_id, param.Program_name, param.Content, param.Ptype, param.Answer, param.Difficulty)
}
