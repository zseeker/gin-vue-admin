package repository

import "github.com/bingjian-zhu/gin-vue-admin/models"

//IRoleRepository Role接口定义
type IRoleRepository interface {
	//GetUserRoles 分页返回Articles获取用户身份信息
	GetUserRoles(userName string) (roles []models.Role)
}
