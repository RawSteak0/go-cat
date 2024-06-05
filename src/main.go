package main
import (
	"fmt"
	"os"
)

func main(){
	var args []string = os.Args[1:]
	if len(args) < 2 {
		fmt.Println("gocat: error: bad argument")
		return;
	}
}
