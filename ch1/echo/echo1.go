package main

/**
** echo1.exe 1 2 3 4
**
**/
import (
	"fmt"
	"os"
)

func main(){
	var s, sep string
	for i:=1;i<len(os.Args);i++{
		s +=sep + os.Args[i]
		sep=""
	}
	fmt.Println(s)
}