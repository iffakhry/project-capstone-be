package databases

import (
	"final-project/config"
	"final-project/models"
	"fmt"
	"strconv"
	"time"
)

// function database untuk menambahkan user baru (registrasi)
func CreateGroupProduct(group *models.GroupProduct) (interface{}, error) {
	duration := time.Now().AddDate(0, 0, 14)
	//mengambil banyaknya jumlah data yang ada pada group
	_, len_group, _ := GetAllGroupProduct()

	group.NameGroupProduct = "Group " + strconv.Itoa(len_group+1)
	group.CapacityGroupProduct = 1
	group.AdminFee = 4500
	group.TotalPrice = 6500 + 55000
	group.Duration = duration.Format("01-02-2006")

	if err := config.DB.Create(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func GetAllGroupProduct() (interface{}, int, error) {
	group := []models.GetGroupProduct{}
	query := config.DB.Table("group_products").Select("*").Find(&group)
	if query.Error != nil {
		return nil, 0, query.Error
	}
	return group, len(group), nil
}

func GetGroupProductById(id int) (interface{}, error) {
	group := models.GetGroupProduct{}
	query := config.DB.Table("group_products").Select("*").Where("group_products.id = ?", id).Find(&group)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	return group, nil
}

func UpdateGroupProductCapacity(id int) (interface{}, error) {
	group := models.GroupProduct{}
	config.DB.Find(&group, id)
	capacity := group.CapacityGroupProduct + 1
	fmt.Println("lihat capacity", capacity)
	query := config.DB.Model(&group).Where("group_products.id = ?", id).Update("capacity_group_product", capacity)
	if query.Error != nil {
		return nil, query.Error
	}
	return group, nil
}
