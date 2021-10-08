package service

//import "webDemo/pkg/errcode"

type CreateUserRequest struct {
	Name      string `form:"name" binding:"max=100"`
	Password  string `form:"password" binding:"max=100"`
	Privilege uint32 `form:"privilege,default=0" binding:"oneof=0 1"`
	Id        uint32
}
type LoginRequest struct {
	Name     string `form:"name" binding:"max=100"`
	Password string `form:"password" binding:"max=100"`
}
type UpdateRequest struct {
	Name      string `form:"name" binding:"max=100"`
	Password  string `form:"password" binding:"max=100"`
	Privilege uint32 `form:"privilege,default=0" binding:"oneof=0 1"`
}

func (svc *Service) CreateUser(param *CreateUserRequest) error {
	// single, err := svc.dao.CountUser(param.Name)
	// if err != nil {
	// 	return err
	// }
	// if single == 0 {
	// 	return svc.dao.CreateUser(param.Name, param.Password, 0, param.Privilege)
	// } else {
	// 	return errcode.ErrorUserExist
	// }
	return svc.dao.CreateUser(param.Name, param.Password, 0, param.Privilege)
}

func (svc *Service) LoginUser(param *LoginRequest) (string, error) {
	return svc.dao.GetUserPasswd(param.Name)
}

func (svc *Service) UpdateUser(param *UpdateRequest) error {
	return svc.dao.UpdateUser(param.Name, param.Password, 0, param.Privilege)
}
