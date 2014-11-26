package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/ajstarks/svgo"
)

var width = 800
var height = 400

var startTime = now()

func drawPoint(osvg *svg.SVG, point int, process int) {
	seconds := now()
	difference := (int64(seconds) - int64(startTime)) / 100000

	pointLocation := int(difference)
	pointLocationV := 0
	color := "#000000"
	switch {
	case process == 1:
		pointLocationV = 60
		color = "#cc6666"
	default:
		pointLocationV = 180
		color = "#66cc66"
	}

	osvg.Rect(pointLocation, pointLocationV, 3, 5, "fill:"+color+";stroke:none;")
	time.Sleep(150 * time.Millisecond)
}

func makeChart(rw io.Writer) *svg.SVG {
	outputSVG := svg.New(rw)

	outputSVG.Start(width, height)
	outputSVG.Rect(10, 10, width-20, 100, "fill:#eeeeee;stroke:none;")
	outputSVG.Rect(10, 130, width-20, 100, "fill:#eeeeee;stroke:none;")

	times := make([]string, width/100+1)
	for i := 0; i <= width; i++ {
		switch {
		case i%4 == 0:
			outputSVG.Circle(i, 377, 1, "fill:#cccccc;stroke:none;")
		case i%10 != 0:
			outputSVG.Rect(i, 0, 1, height, "fill:#cccccc")
		case i%10 == 0:
			outputSVG.Rect(i, 0, 1, height, "fill:#dddddd")
		}

		if i%100 == 0 {
			times[i/100] = strconv.FormatInt(int64(i), 10)
		}

	}

	for i, timeText := range times {
		outputSVG.Text(i*100, 380, timeText, "text-anchor:middle;font-size:10px;fill:#000000")
	}

	outputSVG.Text(20, 30, "Process 1 Timeline", "text-anchor:start;font-size:12px;fill:#333333;")
	outputSVG.Text(20, 150, "Process 2 Timeline", "text-anchor:start;font-size:12px;fill:#333333;")

	return outputSVG
}

func endWithLabel(osvg *svg.SVG, label string) {
	osvg.Text(650, 360, label, "text-anchor: start;font-size:12px;fill:#333333;")
	osvg.End()
}

func withChart(rw io.Writer, label string, action func(*svg.SVG)) {
	chart := makeChart(rw)
	action(chart)
	endWithLabel(chart, label)
}

func drawBasicPoints(chart *svg.SVG) {
	for i := 0; i < 100; i++ {
		go drawPoint(chart, i, 1)
		drawPoint(chart, i, 2)
		fmt.Println("Drawn", i)
	}
}

func visualize(rw http.ResponseWriter, req *http.Request, visualization func()) {
	startTime = now()
	fmt.Println("request to /visualize")
	rw.Header().Set("Content-Type", "image/svg+xml")
	visualization()
}

func visualizeBasic(rw http.ResponseWriter, req *http.Request) {
	visualize(rw, req, func() {
		withChart(rw, "Run without goroutines", drawBasicPoints)
	})
}

func main() {
	runtime.GOMAXPROCS(2)
	http.Handle("/visualize", http.HandlerFunc(visualizeBasic))

	err := http.ListenAndServe(":1900", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
