package tmx

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func Write(filepath string, tmxfile *Tmx) {

	xml_bytes, err_marshal := xml.MarshalIndent(tmxfile, "", " ")

	if err_marshal != nil {
		panic(err_marshal)
	}
	write_error := ioutil.WriteFile(filepath, xml_bytes, 0644)
	if write_error != nil {
		panic(write_error)
	}
}

func split_xml(tmxfile *Tmx, output_prefix string, file_size int, max_size int) {
	// First file can be really large so let's make a shell
	tmx_template := generate_new_tmx_shell(tmxfile)
	fmt.Println("MaxSizePtr:", max_size)
	split_racio := file_size / max_size
	tu_nr := len(tmxfile.Body.Tu)
	max_tu := tu_nr
	if split_racio > 1 {
		max_tu = int(tu_nr / split_racio)
	}
	nr_files := (split_racio) + 1
	fmt.Println("Nr files:", nr_files)
	//fmt.Println("Tmx File:", *tmxfile)
	for i := 0; i < nr_files; i++ {
		starting_value := i * max_tu
		end_value := (i + 1) * max_tu
		if end_value > tu_nr {
			end_value = tu_nr
		}
		part_tmx := generate_new_tmx_shell(tmx_template)
		part_tmx.Body.Tu = tmxfile.Body.Tu[starting_value:end_value]
		file_name := fmt.Sprintf("%s%d_%s.%s", output_prefix, i, tmxfile.XMLName.Space, tmxfile.XMLName.Local)
		Write(file_name, part_tmx)
	}

}

func generate_new_tmx_shell(file *Tmx) *Tmx {
	New_tmx := Tmx{
		XMLName: file.XMLName,
		Text:    file.Text,
		Version: file.Version,
		Header:  file.Header,
		Body:    Body{Text: file.Body.Text},
	}

	return &New_tmx
}
