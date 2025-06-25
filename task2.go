package main

import (

"fmt"
"sort"
)

func mergeAndSort(arr1, arr2[]int)[]int{

	combined := append(arr1,arr2...)
	sort.Ints(combined)
	return combined

}

func main(){
	
	arr1 :=[]int{3,1,5,6,-2,75,75,675,}
	arr2 :=[]int{2,4,6,3,6656588998,-4545}

	fmt.Println(mergeAndSort(arr1,arr2))




}