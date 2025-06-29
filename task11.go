package main
import(
	"fmt"
	"math"
)


func powerFunc(exp int) func(int)int{

	return func(x int)int{
		return int(math.Pow(float64(x),float64(exp)))
	}
}

func main(){


	square := powerFunc(2)
	fmt.Println(square(4))

	krug := powerFunc(3)
	fmt.Println(krug(4))

}