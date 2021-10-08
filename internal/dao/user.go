package dao

import "webDemo/internal/model"

func (d *Dao) CreateUser(name string, passwd string, id uint32, privilege uint32) error {
	user := model.User{
		Name:      name,
		Password:  passwd,
		ID:        0,
		Privilege: privilege,
	}
	return user.CreateUser(d.engine)
}

func (d *Dao) UpdateUser(name string, passwd string, id uint32, privilege uint32) error {
	user := model.User{
		Name:      name,
		Password:  passwd,
		ID:        0,
		Privilege: privilege,
	}
	return user.Update(d.engine)
}

func (d *Dao)CountUser(name string) (int, error){
	user := model.User{
		Name: name,
	}
	return user.Count(d.engine)
}

func (d *Dao) GetUserPasswd(name string) (string, error){
	user := model.User{
		Name: name,
	}
	return user.GetPasswd(d.engine)
}
