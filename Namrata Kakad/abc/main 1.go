package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/**
{
    "erase_block": "0x8000",
    "logic_block": "0x200",
    "offset": "0x000000",     
    "partiton": [
        {
            "plain_file": [
                {
                    "file": "bird.wav",
                    "local_file": "./pack/bird.wav"
                },
                {
                    "file": "closingcar.wav",
                    "local_file": "./pack/closingcar.wav"
                },
                {
                    "file": "nm0114.wav",
                    "local_file": "./pack/nm0114.wav"
                }
            ]
        }
    ],
    "size": "0x800000",
    "type": "FBD2"
}

**/

type FileJSON struct {
	FileLocation  string `json:"file,omitempty"`
	LocalLocation string `json:"local_file,omitempty"`
}

type PartitionJSON struct {
	Files []FileJSON `json:"plain_file"`
}

type FileStruct struct {
	EraseBlock string          `json:"erase_block"`
	LogicBlock string          `json:"logic_block"`
	Offset     string          `json:"offset"`
	Partition  []PartitionJSON `json:"partiton"`
	Size       string          `json:"size"`
	Type       string          `json:"type"`
}

func main() {

	var dataLocation string
	var outputLocation string
	flag.StringVar(&dataLocation, "i", "./", "Folder Location")
	flag.StringVar(&outputLocation, "o", "./", "Output Location")
	flag.Parse()

	file, err := os.Create(outputLocation + "output.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var FileStructData FileStruct

	FileStructData.EraseBlock = "0x8000"
	FileStructData.LogicBlock = "0x200"
	FileStructData.Offset = "0x000000"
	FileStructData.Size = "0x800000"
	FileStructData.Type = "FBD2"

	var partitionFiles PartitionJSON

	err = filepath.Walk(dataLocation,
		func(path string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			partitionFiles.Files = append(partitionFiles.Files, FileJSON{FileLocation: path, LocalLocation: fileInfo.Name()})
			// fmt.Println(path)
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	FileStructData.Partition = append(FileStructData.Partition, partitionFiles)

	jsonData, err := json.MarshalIndent(FileStructData, "", "\t")
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	file.Write(jsonData)
	fmt.Printf("json data: %s\n", jsonData)
}
