package main
import(
	"errors"
	"fmt"
	"io/fs"
	"os"
)
func main(){
	if _,err := os.Open("non-existing");err != nil{
		if errors.Is(err,fs.ErrNotExist){
			fmt.Println("File does not exist")
		}else {
			fmt.Println(err)
		}
	}
}