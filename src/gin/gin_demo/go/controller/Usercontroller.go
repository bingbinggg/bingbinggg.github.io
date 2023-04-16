package controller

import (
	"GIN/gin_demo/common/rsp"
	"GIN/gin_demo/entity"
	"GIN/gin_demo/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"path"
	"strconv"
)

func CreateUser(c *gin.Context) {
	//定义一个User变量
	var user entity.User
	//将请求中的body数据解析到User结构变量中
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := service.CreateUser(&user); err != nil {
		rsp.Error(c, "新增用户失败")
	} else {
		rsp.Success(c, "新增用户成功", user)
	}
}

func GetUserList(c *gin.Context) {
	if todoList, err := service.GetAllUser(); err != nil {
		rsp.Error(c, "显示用户列表失败")
	} else {
		rsp.Success(c, "显示用户列表成功", todoList)
	}
}

func UpdateUser(c *gin.Context) {
	//得到URL上的id信息
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		rsp.Error(c, "用户更新，用户ID不合法")
		return
	}
	var user entity.User
	user, err = service.GetUserById(i)
	if err != nil {
		rsp.Error(c, "用户更新，用户不存在")
		return
	}

	//将请求中的body数据解析到User结构变量中
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := service.UpdateUser(&user); err != nil {
		rsp.Error(c, "更新用户信息失败")
		return
	} else {
		rsp.Success(c, "更新用户信息成功", user)
	}
}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")

	//将id传给service层的UpdateUser方法，进行User的删除
	if err := service.DeleteUserById(id); err != nil {
		rsp.Error(c, "删除用户失败")
	} else {
		rsp.Success(c, "删除用户成功", id)
	}
}

// 用户注册接口
func RegisterInterface(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

// 用户注册接口
func Register(c *gin.Context) {
	//定义一个User变量
	var user entity.User
	//将请求中的body数据解析到User结构变量中
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	//获取注册信息
	username := user.Name
	password := user.Password
	email := user.Email
	mobile := user.Mobile

	//密码验证
	if len(password) < 6 {
		rsp.Error(c, "用户注册，密码必须大于6位")
		return
	}
	//然后对密码进行加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		rsp.Error(c, "用户密码加密错误")
		return
	}

	newUser := &entity.User{
		Name:     username,
		Password: string(hashPassword),
		Email:    email,
		Mobile:   mobile,
	}

	//创建该用户
	if err := service.CreateUser(newUser); err != nil {
		rsp.Error(c, "用户注册，失败")
		return
	}
	id := newUser.Id

	//从请求中读取头像
	f, err := c.FormFile("avatar")
	if err != nil {
		rsp.Error(c, "头像上传，头像读取失败")
	} else {
		//将文件保存在服务器
		dst := path.Join("./file./avatar/", strconv.Itoa(id)+".jpg")

		if err := c.SaveUploadedFile(f, dst); err != nil {
			rsp.Error(c, "头像上传，头像保存服务器错误")
			return
		}
		user.Avatar = strconv.Itoa(id) + ".txt"
	}
	rsp.Success(c, "用户注册，成功", newUser)
}

// 登录界面
func LoginInterface(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// 用户登录接口
func Login(c *gin.Context) {

	var reqUser entity.User
	//将请求中的body数据解析到User结构变量中
	if err := c.ShouldBind(&reqUser); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	email := reqUser.Email
	password := reqUser.Password

	//查询用户
	var user entity.User
	user, err := service.GetUserByEmail(email)
	if err != nil {
		rsp.Error(c, "用户登录，用户不存在")
		return
	}
	//校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		rsp.Error(c, "用户登录，密码错误")
		return
	}
	//返回登录后的页面
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Name":     user.Email,
		"Password": password,
	})
}
