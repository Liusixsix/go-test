package controller

import (
	"gin-demo/common"
	"gin-demo/model"
	"gin-demo/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDb()
	db.AutoMigrate(model.Category{})
	return CategoryController{DB: db}
}

func (cate CategoryController) Create(c *gin.Context) {
	var requestCategory model.Category
	c.ShouldBind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(c, nil, "数据验证失败,分类名称必填")
		return
	}

	cate.DB.Create(&requestCategory)
	response.Success(c, gin.H{"category": requestCategory}, "创建成功")
}

func (cate CategoryController) Update(c *gin.Context) {
	var requestCategory model.Category
	c.ShouldBind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(c, nil, "数据验证失败,分类名称必填")
		return
	}
	caregoryId, _ := strconv.Atoi(c.Params.ByName("id"))
	var updatCategory model.Category
	if cate.DB.First(&updatCategory, caregoryId).RecordNotFound() {
		response.Fail(c, nil, "分类不存在")
		return
	}
	cate.DB.Model(&updatCategory).Update("name", requestCategory.Name)
	response.Success(c, gin.H{"cate": updatCategory}, "修改成功")
}

func (cate CategoryController) Show(c *gin.Context) {
	caregoryId, _ := strconv.Atoi(c.Params.ByName("id"))
	var category model.Category
	if cate.DB.First(&category, caregoryId).RecordNotFound() {
		response.Fail(c, nil, "分类不存在")
		return
	}
	response.Success(c, gin.H{"cate": category}, "")
}

func (cate CategoryController) Delete(c *gin.Context) {
	caregoryId, _ := strconv.Atoi(c.Params.ByName("id"))
	var category model.Category
	if cate.DB.First(&category, caregoryId).RecordNotFound() {
		response.Fail(c, nil, "分类不存在")
		return
	}
	if err := cate.DB.Delete(model.Category{}, caregoryId).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "删除成功")
}

func GetList(c *gin.Context) {
	var list []model.Pachong
	DB := common.GetDb()
	if err := DB.Find(&list).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, gin.H{
		"data": list,
	}, "查询成功")
}
