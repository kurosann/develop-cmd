package doctor

import (
	"context"
	"develop-cmd/pkg/C"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

func Doctor() *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "检查gitlab环境变量",
		Run: func(c *cobra.Command, args []string) {
			// 检查glab是否安装
			if _, err := exec.LookPath("glab"); err != nil {
				fmt.Println("glab未安装")
				installGlab(c.Context())
				return
			}
			fmt.Println("glab已安装")

			bs, _ := C.CmdOutByte(c.Context(), "glab", "auth", "status")

			if strings.Contains(string(bs), "Logged in to imageharbor.xyz") {
				fmt.Println("glab已认证")
			} else {
				fmt.Println("glab认证失败")
				auth(c.Context())
			}

			fmt.Println("gitlab环境变量检查通过")
		},
	}
}

func installGlab(ctx context.Context) {
	fmt.Println("正在安装glab")
	// 查看操作系统
	os := runtime.GOOS
	if os == "darwin" {
		// 查看brew是否安装
		if _, err := exec.LookPath("brew"); err != nil {
			fmt.Println("brew未安装，请手动安装glab")
			fmt.Println("https://github.com/profclems/glab/releases")
			return
		}
		bs, err := C.CmdOutByte(ctx, "brew", "install", "glab")
		if err != nil {
			fmt.Println("glab安装失败")
			return
		}
		fmt.Println(string(bs))
	} else if os == "linux" && runtime.GOARCH == "amd64" {
		bs, err := C.CmdOutByte(ctx, "sudo", "apt-get", "install", "glab")
		if err != nil {
			fmt.Println("glab安装失败")
			return
		}
		fmt.Println(string(bs))
	} else {
		fmt.Println("不支持的操作系统，请手动安装glab，并配置环境变量")
		fmt.Println("https://github.com/profclems/glab/releases")
		return
	}
	fmt.Println("glab安装成功")
}

func auth(ctx context.Context) {
	fmt.Println("正在认证glab")
	fmt.Println("请输入gitlab hostname:")
	var hostname string
	fmt.Scanln(&hostname)
	fmt.Printf("请访问 https://%s/-/user_settings/personal_access_tokens 获取gitlab token\n", hostname)
	fmt.Println("请输入gitlab token:")
	var token string
	fmt.Scanln(&token)
	bs, err := C.CmdOutByte(ctx, "glab", "auth", "login", "--hostname", hostname, "--token", token)
	if err != nil {
		fmt.Println("glab认证失败")
		return
	}
	fmt.Println(string(bs))
}
