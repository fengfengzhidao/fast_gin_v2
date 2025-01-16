package flags

import (
	"fast_gin/models"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type User struct {
}

func (User) Create() {
	var user models.UserModel
	fmt.Println("请选择角色： 1 管理员 2 普通用户")
	_, err := fmt.Scanln(&user.RoleID)
	if err != nil {
		fmt.Println("输入错误", err)
		return
	}
	if user.RoleID != 1 && user.RoleID != 2 {
		fmt.Println("用户角色输入错误", err)
		return
	}
	fmt.Println("请输入用户名")
	fmt.Scanln(&user.Username)

	fmt.Println("请输入密码")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd())) // 读取用户输入的密码
	if err != nil {
		fmt.Println("读取密码时出错:", err)
		return
	}
	fmt.Println("请再次输入密码")
	rePassword, err := terminal.ReadPassword(int(os.Stdin.Fd())) // 读取用户输入的密码
	if err != nil {
		fmt.Println("读取密码时出错:", err)
		return
	}
	if string(password) != string(rePassword) {
		fmt.Println("两次密码不一致")
		return
	}

	fmt.Println(user.RoleID)
	fmt.Println(user.Username)
	fmt.Println(string(password))

}
func (User) List() {

}
