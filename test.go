package main

import "fmt"

func main() {
	var metrics [8]string
	metrics[0] = "rttPktTransmit"
	metrics[1] = "rttPktLoss"

	for i:=0;  i < len(metrics); i++ {
		fmt.Println(i)
	}

	x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,}
	
	//fmt.Println(len(x))
	
	
    out := 0
    for i := 0; i < len(x)-1-1; i++ {
		if x[i] < x[i+1] {
			fmt.Println(x[i])
		}
	}
	fmt.Println(out)
}