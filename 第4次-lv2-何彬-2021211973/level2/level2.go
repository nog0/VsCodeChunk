package main

import (
	"fmt"
	"time"
)

func main(){
	go func() {
		for {now := time.Now()
			now1 := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, now.Location())
			if(now1.Sub(now)>=0){
				timer:=time.NewTimer(now1.Sub(now))
				<-timer.C
				fmt.Println("早八算什么，早六才是吾辈应起之时")
			}else{
				next:=now.Add(time.Hour*24)
				next2:=time.Date(next.Year(),next.Month(),next.Day(),6,0,0,0,now.Location())
				timer:=time.NewTimer(next2.Sub(now))
				<-timer.C
				fmt.Println("早八算什么，早六才是吾辈应起之时")
			}
		}
	}()
	go func() {
		for {now := time.Now()
			now1 := time.Date(now.Year(), now.Month(), now.Day(), 2, 0, 0, 0, now.Location())
			if(now1.Sub(now)>=0){
				timer:=time.NewTimer(now1.Sub(now))
				<-timer.C
				fmt.Println("谁能比我卷！")
			}else{
				next:=now.Add(time.Hour*24)
				next2:=time.Date(next.Year(),next.Month(),next.Day(),2,0,0,0,now.Location())
				timer:=time.NewTimer(next2.Sub(now))
				<-timer.C
				fmt.Println("谁能比我卷！")
			}
		}
	}()
	for{
		timer:=time.NewTimer(time.Second)
		<-timer.C
		fmt.Println("芜湖！起飞！")
	}
}