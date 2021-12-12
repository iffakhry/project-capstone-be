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
	fee := 4500

	group.NameGroupProduct = "Group " + strconv.Itoa(len_group+1)
	group.CapacityGroupProduct = 1
	group.AdminFee = fee
	group.TotalPrice = fee + 55000
	group.Duration = duration.Format("01-02-2006")
	group.Status = "Available"

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
	UpdateGroupProductCapacity(id)
	query := config.DB.Table("group_products").Select("*").Where("group_products.id = ?", id).Find(&group)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	return group, nil
}

func UpdateGroupProductCapacity(id int) (interface{}, error) {
	group := models.GroupProduct{}
	config.DB.Find(&group, id)
	if group.CapacityGroupProduct >= 5 {
		query := config.DB.Model(&group).Where("group_products.id = ?", id).Update("status", "Not Available")
		if query.Error != nil {
			return nil, query.Error
		}
		return group, nil
	}
	capacity := group.CapacityGroupProduct + 1
	fmt.Println("lihat capacity", capacity)
	query1 := config.DB.Model(&group).Where("group_products.id = ?", id).Update("capacity_group_product", capacity)
	if query1.Error != nil {
		return nil, query1.Error
	}
	return group, nil
}

func GetGroupProductByAvailable(str string) (interface{}, error) {
	group := []models.GetGroupProduct{}
	// UpdateGroupProductCapacity(id)
	query := config.DB.Table("group_products").Select("*").Where("group_products.status = ?", str).Find(&group)
	if query.Error != nil || query.RowsAffected < 1 {
		return nil, query.Error
	}
	return group, nil
}
