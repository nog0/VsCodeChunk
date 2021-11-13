package main
import ("fmt"
        "sort")
var drinkinga =map[string]string{
	"咖啡":"cofee","可乐":"coke",
	"红酒":"Wine","酒":"Alcohol",
	"白兰地":"Brandy","白酒":"Liquor",
}
func main(){
	drinkingb:=make(map[string]string,len(drinkinga))
for i,k:=range drinkinga{
	drinkingb[k]=i
}
for i,k:=range drinkingb{
	fmt.Printf("Key: %v, Value: %v / ", i, k)
}//键值对调，打印drinkingb
  	fmt.Println()
	i:=0
   	drinkingsli:=make([]string,len(drinkingb))
   	for v,_ :=range drinkingb{
   	drinkingsli[i]=v
   	i++
	}
	sort.Strings(drinkingsli)
	for i,k:=range drinkingsli{
		fmt.Printf("Key: %v, Value: %v / ",drinkingsli[i],drinkingb[k])
	}//排序map，打印
}