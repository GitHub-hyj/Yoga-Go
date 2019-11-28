package yogacommand

import (
	"github.com/GitHub-hyj/Yoga-Go/internal/yogaclient"
	"github.com/GitHub-hyj/Yoga-Go/yogaliner"
)

func RunLogin(email, password string) (token string, err error) {
	line := yogaliner.NewLiner()
	defer line.Close()

	//if email == "" {
	//	email, err = line.State.Prompt("请输入邮箱, 回车键提交 > ")
	//	if err != nil {
	//		return
	//	}
	//}
	//if password == "" {
	//	// liner 的 PasswordPrompt 不安全, 拆行之后密码就会显示出来了
	//	fmt.Printf("请输入密码(输入的密码无回显, 确认输入完成, 回车提交即可) > ")
	//	password, err = line.State.PasswordPrompt("")
	//	if err != nil {
	//		return
	//	}
	//}

	email = "932761407@qq.com"
	password = "tank123"

	yogaclient.Login("/api/user/login", map[string]string{"email": email, "password": password})

	return "123", nil
}
