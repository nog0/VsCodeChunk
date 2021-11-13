package main
import "fmt"
func main() {
	up:=func(skillName string){
		fmt.Println("阿巴阿巴阿巴！", skillName)
	}
	op:=func(skillName string) {
		fmt.Println("尝尝我的厉害吧！", skillName)
	}
	var str1 string
	var str string
	fmt.Println("输入烂话")
	fmt.Println("up：阿巴阿巴阿巴！")
	fmt.Println("op：尝尝我的厉害吧！")
	fmt.Println()
	fmt.Scanf("%s",&str1)
	fmt.Println("开始犯病")
	fmt.Scanf("%s",&str)
	if str1=="up"{
		ReleaseSkill(str,up )
	}else if str1=="op"{
		ReleaseSkill(str,op )
	}
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
