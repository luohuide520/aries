package form

import "aries/model"

// 登录表单
type LoginForm struct {
	Username string `json:"username" binding:"required,min=3,max=30" label:"用户名"` // 用户名
	Pwd      string `json:"pwd" binding:"required,min=6,max=20" label:"密码"`       // 密码
	//Captcha  string `form:"captcha" binding:"required" label:"验证码"`               // 验证码
}

// 注册表单
type RegisterForm struct {
	Username  string `json:"username" binding:"required,min=3,max=30" label:"用户名"`                // 用户名
	Pwd       string `json:"pwd" binding:"required,min=6,max=20" label:"密码"`                      // 密码
	SecondPwd string `json:"second_pwd" binding:"required,min=6,max=20,eqfield=Pwd" label:"确认密码"` // 确认密码
	Email     string `json:"email" binding:"required,max=30,email" label:"邮箱"`                    // 邮箱
	//UserImg    string `form:"user_img" binding:"required,max=255" label:"头像"`                       // 头像
}

// 绑定登录表单到实体结构
func (form LoginForm) BindToModel() model.User {
	return model.User{
		Username: form.Username,
		Pwd:      form.Pwd,
	}
}

// 绑定注册表单到实体结构
func (form RegisterForm) BindToModel() model.User {
	return model.User{
		Username: form.Username,
		Pwd:      form.Pwd,
		Email:    form.Email,
		//UserImg:  form.UserImg,
	}
}

/*// 定义返回错误信息
func (form LoginForm) GetError(err validator.ValidationErrors) string {
	for _, fieldError := range err {
		if fieldError.Field() == "Username" {
			switch fieldError.Tag() {
			case "required":
				return "请输入用户名"
			}
		}
		if fieldError.Field() == "Pwd" {
			switch fieldError.Tag() {
			case "required":
				return "请输入密码"
			}
		}
	}
	return err.Error()
}*/
