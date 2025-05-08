package service

import (	
	"rvkc/util"
	"rvkc/models"
	"rvkc/context_error"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)


type RoleService struct {
	serviceRole	GenericService[models.Role]
	log     	*logrus.Logger
}


func NewRoleService(
	serviceRole GenericService[models.Role],
) *RoleService {
	
	return &RoleService{
		serviceRole: 	serviceRole,
		log:     		util.GetLogger(),
	}
}


func (r *RoleService) GetById(id string) (*models.Role, error) {
	return r.serviceRole.GetBy("id = ?", id)
}


func (r *RoleService) GetByName(name string) (*models.Role, error) {
	return r.serviceRole.GetBy("name = ?", name)
}


func (r *RoleService) GetRolesByNameList(ctx *gin.Context, roles []*models.Role) ([]*models.Role, error) {
	var roleList []*models.Role
	for _, current := range roles {
		role, err := r.GetByName(current.Name)
		if err != nil {
			context_error.RoleNotFoundError(ctx)
			return nil, err
		}
		roleList = append(roleList, role)
	}
	return roleList, nil
}


func (r *RoleService) GetRolesByAccount(ctx *gin.Context, id string) ([]*models.Role, error) {
	joins := []string{
		"JOIN account_role ar ON ar.role_id = role.id",
	}
	
	roles, err := r.serviceRole.FindWithJoinsAndFilters(
		joins,
		"ar.account_id = ?",
		id,
	)

	if err == gorm.ErrRecordNotFound {
		context_error.RoleNotFoundError(ctx)
		ctx.Abort()
		return nil, err
	}

	if err != nil {
		context_error.RoleSearchError(ctx)
		ctx.Abort()
		return nil, err
	}

	return roles, nil
}