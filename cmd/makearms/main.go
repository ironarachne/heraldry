package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/ironarachne/heraldry"
)

type Arms struct {
	FileName string
	Blazon   string
}

type OutputData struct {
	Arms []Arms
}

func main() {
	numberOfResults := flag.Int("results", 1, "Number of results to generate")
	filePath := flag.String("path", ".", "Directory to dump output into")
	createHtmlIndex := flag.Bool("html", false, "Create an HTML index of output")

	const htmlIndexTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Heraldic Device Index</title>
	<style type="text/css">
	.container {
		display: flex;
		flex-wrap: wrap;
	}
	
	.blazon {
		font-size: 1.5rem;
		font-style: italic;
	}

	.device {
		margin: 1rem;
		text-align: center;
		width: 400px;
	}

	img {
		margin: 0 auto;
		height: 420px;
		width: 320px;
	}
	</style>
</head>
<body>
    <h1>Heraldic Device Index</h1>
    <div id="heraldic-devices" class="container">
        {{range .Arms}}
        <div class="device">
            <img src="{{ .FileName }}">
            <p class="blazon">{{ .Blazon }}</p>
        </div>
        {{end}}
    </div>
</body>
</html>
	`

	flag.Parse()

	fileName := ""
	iteration := ""
	arms := Arms{}

	outputData := OutputData{}

	for i := 0; i < *numberOfResults; i++ {
		device := heraldry.Generate()
		iteration = strconv.Itoa(i)
		fileName = "device-" + iteration + ".svg"
		heraldry.RenderToSvg(device, *filePath+"/"+fileName, 320, 420)

		arms.FileName = fileName
		arms.Blazon = heraldry.RenderToBlazon(device)
		outputData.Arms = append(outputData.Arms, arms)
	}

	if *createHtmlIndex {
		writer, err := os.Create(*filePath + "/index.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		t, err := template.New("htmlIndex").Parse(htmlIndexTemplate)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = t.Execute(writer, outputData)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer writer.Close()
	}
}
