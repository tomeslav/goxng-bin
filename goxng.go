package main

import (
	"fmt"
	"github.com/tomeslav/goxng"
	"flag"
)

func main() {

	var duration = flag.Int("duration", 100, "Display duration for each image, in millisecond")

	flag.Parse()
	
	imgs := flag.Args()
	if len(imgs) < 2 {
		fmt.Println("Please specify images filenames")
		return
	}
	
	xml, err := goxng.GetXng(imgs, *duration)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(xml)
}
