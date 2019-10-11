package main
import "fmt"

func main(){
    fmt.Println("hey!")
    wellcome()               // this function will come from wellcome file 
                             //because if we declare functions in the same 
                             //package then we also able to use them into the same package level
}
