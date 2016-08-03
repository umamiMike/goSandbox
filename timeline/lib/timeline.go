package timeline

import "fmt"

type timeline struct {
	events :[]
}

func main() {

	fmt.Println("Hello")
	myTimeline := timeline{}

	myTimeline.Event = "I am a string"

	fmt.Println(myTimeline.Event)
}
