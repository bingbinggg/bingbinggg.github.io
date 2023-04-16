package controller

import (
	"GIN/gin_demo/common/rsp"
	"GIN/gin_demo/entity"
	"GIN/gin_demo/service"
	"bufio"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)

func CreateSpot(c *gin.Context) {
	//定义一个Spot变量
	var spot entity.Spot
	//将请求中的body数据解析到Spot结构变量中
	if err := c.ShouldBind(&spot); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := service.CreateSpot(&spot); err != nil {
		rsp.Error(c, "新增景点失败")
	} else {
		rsp.Success(c, "新增景点成功", spot)
	}
}

func GetSpotList(c *gin.Context) {
	if todoList, err := service.GetAllSpot(); err != nil {
		rsp.Error(c, "显示景点列表失败")
	} else {
		rsp.Success(c, "显示景点列表成功", todoList)
	}
}

// 显示依据评分降序得到Spot集合
func GetSpotOrder(c *gin.Context) {
	todoList, err := service.GetPartSpot()
	if err != nil {
		rsp.Error(c, "显示景区排序列表失败")
		return
	}

	//返回显示列表的页面
	c.HTML(http.StatusOK, "spotDisplay.html", gin.H{
		"data": todoList,
	})
}

// 显示地址得到的Spot排序集合
func GetSpotOrder1(c *gin.Context) {
	if todoList, err := service.GetPartSpot1("福建"); err != nil {
		rsp.Error(c, "显示景区地址排序列表失败")
	} else {
		rsp.Success(c, "显示景区地址排序列表成功", todoList)
	}
}

func UpdateSpot(c *gin.Context) {
	//得到URL上的id信息
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		rsp.Error(c, "景点更新，景点ID不合法")
		return
	}
	var spot entity.Spot
	spot, err = service.GetSpotById(i)
	if err != nil {
		rsp.Error(c, "景点更新，景点不存在")
		return
	}

	//将请求中的body数据解析到User结构变量中
	if err := c.ShouldBind(&spot); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := service.UpdateSpot(&spot); err != nil {
		rsp.Error(c, "更新景点信息失败")
		return
	} else {
		rsp.Success(c, "更新景点信息成功", spot)
	}
}

func DeleteSpotById(c *gin.Context) {
	id := c.Param("id")

	//将id传给service层的UpdateUser方法，进行User的删除
	if err := service.DeleteSpotById(id); err != nil {
		rsp.Error(c, "删除景点失败")
	} else {
		rsp.Success(c, "删除景点成功", id)
	}
}

// 景区信息录入接口
func InfoInterface(c *gin.Context) {
	c.HTML(http.StatusOK, "spotinfo.html", nil)
}

// 景区信息录入实现
func Info(c *gin.Context) {
	//将请求中的body数据解析到User结构变量中
	var reqUser entity.Spot
	if err := c.ShouldBind(&reqUser); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	//获取景区信息
	name := reqUser.Name
	address := reqUser.Address
	mobile := reqUser.Mobile
	text := reqUser.Text

	newUser := &entity.Spot{
		Name:    name,
		Address: address,
		Mobile:  mobile,
	}

	//创建该景区
	if err := service.CreateSpot(newUser); err != nil {
		rsp.Error(c, "景区信息录入，失败")
		return
	}
	id := newUser.Id

	//创建文本文件存储描述信息
	filePath := "D:/Code/GoProject/GoLand/src/gin/gin_demo/file/spot/spot_text/"
	filePath += strconv.Itoa(id) + ".txt"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		rsp.Error(c, "景区信息录入，描述文件打开失败")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			rsp.Error(c, "景区信息录入，描述文件关闭失败")
			return
		}
	}(file)

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(text)
	if err != nil {
		rsp.Error(c, "景区信息录入，文件输入失败")
		return
	}
	err = writer.Flush()
	if err != nil {
		rsp.Error(c, "景区信息录入,文本输入失败")
		return
	}
	newUser.Text = strconv.Itoa(id) + ".txt"

	//从请求中读取图片文件
	f, err := c.FormFile("picture")
	if err != nil {
		rsp.Error(c, "景区信息录入，景图读取失败")
		return
	} else {
		//将文件保存在服务器
		dst := path.Join("./file./spot/spot_image/", strconv.Itoa(id)+".jpg")

		if err := c.SaveUploadedFile(f, dst); err != nil {
			rsp.Error(c, "景区信息录入，景图保存服务器错误")
			return
		}

		newUser.Picture = strconv.Itoa(id) + ".jpg"
	}
	rsp.Success(c, "景区信息录入，成功", newUser)
}
