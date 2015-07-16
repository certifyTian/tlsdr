package main

import (
	_ "fmt"
	"os"
	"log"
	"io/ioutil"
	"flag"
	"container/list"
)

func readData(input *string) ([]byte, error) {
	var bytes []byte
	var err error

	if *input == "-" {
		log.Println("reading bytes from stdin")
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		log.Println("reading bytes from '" + *input + "'")
		bytes, err = ioutil.ReadFile(*input)
	}

	return bytes, err
}

func parseData(bytes []byte) (list.List) {
	log.Println("Parsing data")
	// TODO Tian
	return list.List{}
}

func analyzeData(data list.List) (list.List) {
	log.Println("Analyzing data")
	// TODO Rahul
	return data
}

func visualizeData(data list.List, output string, format string) {
	log.Println("Visualizing data to " + output + " as " + format)
	// TODO Yuan
}

// This is main
func main() {
	input := flag.String("i", "", "PCAP source; - interpreted as STDIN; otherwise file name")
	output := flag.String("o", "-", "Output directory; - interpreted as STDOUT. Otherwise dir name. STDOUT only works for txt")
	help := flag.Bool("h", false, "Display this help")
	format := flag.String("f", "txt", "Format of output. Valid values are (txt|html|json)")
	log.Println("input was: " + *input)

	flag.Parse()

	if *help {
		log.Println("User demanded help")
		flag.Usage()
		os.Exit(0)
	}

	if *input == "" {
		log.Println("User didn't provide input")
		flag.Usage()
		os.Exit(1)
	}

	if *format != "txt" && *format != "html" && *format != "json" {
		log.Println("Invalid format")
		flag.Usage()
		os.Exit(2)
	}

	if *output == "-" && *format != "txt" {
		log.Println("txt is the only supported format for STDOUT")
		flag.Usage()
		os.Exit(3)
	}

	bytes, err := readData(input)
	if err == nil {
		p := parseData(bytes)
		a := analyzeData(p)
		visualizeData(a, *output, *format)
	}

}