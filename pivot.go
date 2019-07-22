package main

import (
	"fmt"
	"os"
	//"string"
	"encoding/csv"
)


func main() {
	fmt.Println(os.Args[1])
	//table := [][]string{}
	csvFile, err := os.Open(os.Args[1])
         if err != nil {
                 fmt.Println(err)
         }

	reader := csv.NewReader(csvFile)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
	 }
	 fmt.Println(csvData)
	 m := len(csvData)
	 n := len(csvData[1])
	 for i:=0; i< n; i++{
		 for j :=0 ; j<m ;j++ {
			fmt.Printf(csvData[j][i] + ",")
		}
		fmt.Println()
	 }
}

