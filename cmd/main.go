package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"time"
)

func main() {
	x, y := robotgo.Location()

	evChan := hook.Start()
	defer hook.End()

	noActivityDuration := 5 * time.Second
	timer := time.NewTimer(noActivityDuration)
	defer timer.Stop()
	mouseToggleChan := make(chan bool)

	distance := 1
	xn, yn := x+distance, y+distance

	// 鼠标移动控制
	go func() {
		for {
			select {
			case <-mouseToggleChan:
				robotgo.Move(xn, yn)
				fmt.Println("during...mouse now:(", xn, ",", yn, ")")
				robotgo.Move(x, y)
				fmt.Println("during...mouse now:(", x, ",", y, ")")
			}
		}
	}()
	// 事件捕获监听
	go func() {
		for {
			select {
			case e := <-evChan:
				switch e.Kind {
				case
					hook.KeyDown,
					hook.KeyUp,
					hook.MouseDown,
					hook.MouseUp,
					hook.MouseMove,
					hook.MouseWheel,
					hook.MouseDrag:
					timer.Reset(noActivityDuration)
				}
			case <-timer.C:
				// 获取当前鼠标位置
				x, y = robotgo.Location()
				xn, yn = x+distance, y+distance
				fmt.Println("waiting...mouse now:(", x, ",", y, ")")
				mouseToggleChan <- true // Start mouse toggling
			}
		}
	}()

	fmt.Println("program started...")
	select {}
}
