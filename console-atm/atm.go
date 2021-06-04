package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	// 用户名
	username string
	// 余额
	balance int64
}

// 构建用户实体
var user User

func main() {
	// 循环标记
	loop := true

	// 使用简写 for 循环来实现 while 效果
	for loop {
		// 输出菜单
		wrapMenu()

		// 接收控制台输入
		reader := bufio.NewReader(os.Stdin)
		// 如果只需要接收一个字符，使用 ReadRune() 函数即可
		input, _, err := reader.ReadRune()

		// 存在输入错误
		if err != nil {
			fmt.Println(err)
			return
		}

		// switch 语句中的 case 分支默认自带 break
		switch input {
		case '1':
			save()
		case '2':
			get()
		case '3':
			query()
		case '4':
			// 循环状态标记为 false
			loop = exit()
		default:
			fmt.Println("未知菜单，请重新输入")
		}
	}
}

func wrapMenu() {
	fmt.Println("------------------------------------------")
	fmt.Println("\t\t欢迎光临地府银行")
	fmt.Println("\t\t1. 存款")
	fmt.Println("\t\t2. 取款")
	fmt.Println("\t\t3. 查询")
	fmt.Println("\t\t4. 退出")
	fmt.Println("------------------------------------------")
	fmt.Print("请输入功能序号（ 仅支持数字 ）：")
}

func save() {
	// 检查用户状态
	checkUserStatus()

	fmt.Print("请输入存款金额（ 仅支持正整数，且整百的数值 ）：")
	loop := true

	var amount int64
	var err error

	for loop {
		// NewScanner 也是一种接收控制台输入的方式
		scanner := bufio.NewScanner(os.Stdin)

		input := ""
		// Scanner 是一种持续接收控制台输入的接口
		for scanner.Scan() {
			// 每次按回车，都作为一次接收，并进行输出
			input = scanner.Text()

			// 这里只接收一次，所以直接跳出
			break
		}

		// 金额转换
		amount, err = parseAmount(input)

		if !checkAmount("存款", amount, err) {
			continue
		}

		// 余额累加
		user.balance += amount

		// 跳出循环
		loop = false
	}

	fmt.Printf("成功存入 %d 元，当前余额 -> %d \n", amount, user.balance)
}

func get() {
	checkUserStatus()

	fmt.Print("请输入取款金额（ 仅支持正整数，且整百的数值 ）：")
	loop := true

	var amount int64
	var err error

	for loop {
		var input string
		// Scanln() 函数也是一种接收控制台输入的方式，碰到换行符时将停止接收
		// Scanln(params ...) 可以接收多个参数，在输入过程中，通过空格作为分隔符来为多个参数接收值
		// 函数返回两个值，第一个值是接收到的参数个数，第二个值是异常
		_, _ = fmt.Scanln(&input)

		// 金额转换
		amount, err = parseAmount(input)

		if !checkAmount("取款", amount, err) {
			continue
		}

		if user.balance-amount < 0 {
			fmt.Printf("余额不足，当前余额 -> %d ，请重新输入：", user.balance)
			continue
		}

		// 余额累减
		user.balance -= amount

		// 跳出循环
		loop = false
	}

	fmt.Printf("成功取出 %d 元，当前余额 -> %d \n", amount, user.balance)
}

// 金额转换
func parseAmount(input string) (int64, error) {
	pureInput := strings.TrimSpace(input)

	// 尝试将输入的值转换为数字，10 表示十进制，0 表示不限制位宽
	return strconv.ParseInt(pureInput, 10, 0)
}

// 金额校验
func checkAmount(label string, amount int64, err error) bool {
	if err != nil {
		fmt.Print("金额无效，请重新输入：")
		return false
	}

	if amount <= 0 {
		fmt.Printf("%s金额必须大于 0 ，请重新输入：", label)
		return false
	}

	if amount%100 != 0 {
		fmt.Printf("%s金额只能输入整百数值，请重新输入：", label)
		return false
	}

	return true
}

func query() {
	checkUserStatus()

	// 获取当前时间，以及日期格式转换
	// 格式化的规则很奇葩，不是传统的 yyyy-MM-dd HH:mm:ss 这种，也不是 %y-%m-%d 这种
	// 而是使用数字表达，具体如下：
	// 月 1, 01
	// 日 2, 02, _2
	// 时 3, 03, 15
	// 分 4, 04
	// 秒 5, 05
	// 年 06, 2006
	now := time.Now().Format("2006-01-02 03:04:05")

	fmt.Printf("尊敬的 %s 用户您好，现在是北京时间 %s ，您当前账户余额 -> %d \n", user.username, now, user.balance)
}

func exit() bool {
	if user.username == "" {
		fmt.Println("欢迎下次光临")

		// 停止循环
		return false
	}

	fmt.Printf("尊敬的 %s 用户您好，如果退出系统，您的所有余额将即刻充公，您当前余额 -> %d \n", user.username, user.balance)
	fmt.Print("确定要为地府银行做贡献吗？(Y/n) ：")

	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
		return true
	}

	// 使用 string() 函数将输入值转换为 string
	// 使用 strings.TrimSpace() 函数将输入值中的空格，Unicode 字符过滤，这样如果直接回车可以判定为 ""
	// 使用 strings.ToLower() 函数将输入值转为小写
	inputStr := strings.ToLower(strings.TrimSpace(string(input)))

	switch inputStr {
	case "":
		// 表示穿透到下一个 case
		fallthrough
	case "y":
		fmt.Println("地府银行全体员工感谢您的无私奉献，欢迎下次光临")
		return false
	case "n":
		fallthrough
	default:
		fmt.Println("欢迎继续使用本系统")
		return true
	}
}

// 检查用户状态，如果用户没有初始化，则需要先初始化
func checkUserStatus() {
	// 存在用户名
	if user.username != "" {
		return
	}

	fmt.Print("第一次登录，请输入用户名：")
	loop := true

	for loop {
		reader := bufio.NewReader(os.Stdin)
		// 如果要接收多个字符，需要使用 ReadString() 函数，传入的参数表示分隔语句的标记
		input, _ := reader.ReadString('\n')

		// 因为上述获取输入的分隔标记是 \n ，所以需要先将标记符去掉，才能获取到纯净的字符串
		username := strings.Replace(input, "\n", "", -1)
		// 移除输入内容两端可能存在的空格
		username = strings.TrimSpace(username)

		// 非空验证
		if username == "" {
			fmt.Print("用户名无效，请重新输入：")
			continue
		}

		// 长度验证
		if len(input) > 12 {
			fmt.Print("用户名的长度需要控制在 12 个字符以内，请重新输入：")
			continue
		}

		// 空格验证
		if strings.ContainsRune(username, ' ') {
			fmt.Print("用户名不能包含空格，请重新输入：")
			continue
		}

		// 初始化用户
		user.username = username
		user.balance = 0

		// 结束循环
		loop = false
	}

	fmt.Printf("欢迎 %s 用户的第一次登录，您当前余额为 %d \n", user.username, user.balance)
}
