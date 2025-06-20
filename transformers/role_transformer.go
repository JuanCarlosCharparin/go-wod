package transformers

import (
	"wod-go/models"
	"wod-go/dto"
)

func TransformRole(role models.Role) dto.RoleResponse {

	return dto.RoleResponse{
		Id: role.Id,
		Name: role.Name,
	}
}