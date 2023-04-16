package service

import (
	"GIN/gin_demo/dao"
	"GIN/gin_demo/entity"
)

// 新建spot信息
func CreateSpot(spot *entity.Spot) (err error) {
	if err = dao.SqlSession.Create(spot).Error; err != nil {
		return err
	}
	return
}

// 获取spot集合
func GetAllSpot() (spotList []*entity.Spot, err error) {
	if err := dao.SqlSession.Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 得到Spot排序集合
func GetPartSpot() (spotList []*entity.Spot, err error) {
	if err := dao.SqlSession.Order("score desc").Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 依据地址得到Spot排序集合
func GetPartSpot1(address string) (spotList []*entity.Spot, err error) {
	if err := dao.SqlSession.Where("address=?", address).Order("score desc").Find(&spotList).Error; err != nil {
		return nil, err
	}
	return
}

// 根据id查询景点spot
func GetSpotById(id int) (spot entity.Spot, err error) {
	if err := dao.SqlSession.Where("id=?", id).First(&spot).Error; err != nil {
		return spot, err
	}
	return
}

// 根据id删除spot
func DeleteSpotById(id string) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.Spot{}).Error
	return
}

// 更新景点信息
func UpdateSpot(spot *entity.Spot) (err error) {
	err = dao.SqlSession.Save(spot).Error
	return
}
