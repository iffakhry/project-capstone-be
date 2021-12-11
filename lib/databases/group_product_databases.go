package databases

import (
	"final-project/config"
	"final-project/models"
)

// function database untuk menambahkan user baru (registrasi)
func CreateGroupProduct(group *models.GroupProduct) (interface{}, error) {
	if err := config.DB.Create(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}
