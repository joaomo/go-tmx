package tmx

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
)

// Read a tmx file given its path
func Read(filepath string) (*Tmx, int, error) {
	byteValue, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, 0, err
	}
	var tmxData Tmx
	if err := xml.Unmarshal(byteValue, &tmxData); err != nil {
		return nil, 0, err
	}
	return &tmxData, len(byteValue), nil
}

func main() {

	wordPtr := flag.String("filename", "file.tmx", "File to split")
	MaxSizePtr := flag.Int("max-size", 50*1024*1024, "Max size of each the file")
	output_prefix := flag.String("out_prefix", "part_", "File to split")

	flag.Parse()

	tmx_file, tmx_file_size, err := Read(*wordPtr)

	if err == nil {
		split_xml(tmx_file, *output_prefix, tmx_file_size, *MaxSizePtr)
		fmt.Println("word:", *wordPtr)
	} else {
		fmt.Println("Error:", err)

	}

	fmt.Println("tail:", flag.Args())
}
