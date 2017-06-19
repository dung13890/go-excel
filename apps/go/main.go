/**
 * Created by Mr.Dung on 13/06/2017.
 */

package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//write()
	//read()
	multipleRead()
}

func write() {
	start := time.Now()
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	for i := 1; i < 1000000; i++ {
		row = sheet.AddRow()
		row.AddCell().SetInt(i)
		row.AddCell().SetInt((rand.Intn(100000)))
	}
	err = file.Save("test3.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("Time: %v\n", time.Since(start))
}

func read() {
	start := time.Now()
	total := 0
	excelFileName := "test2.xlsx"
	xlFile, _ := xlsx.OpenFile(excelFileName)
	for _, row := range xlFile.Sheets[0].Rows {
		val, _ := row.Cells[1].Int()
		fmt.Println(val)
		total += val
	}

	fmt.Printf("Total:%v Time:%v\n", total, time.Since(start))
}

func multipleRead() {
	start := time.Now()
	totalFirst := 0
	totalSecond := 0
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		excelFileName := "test3.xlsx"
		xlFile, _ := xlsx.OpenFile(excelFileName)
		for _, row := range xlFile.Sheets[0].Rows {
			fmt.Println("1")
			val, _ := row.Cells[1].Int()
			totalFirst += val
		}
	}()

	go func() {
		defer wg.Done()
		excelFileName := "test2.xlsx"
		xlFile, _ := xlsx.OpenFile(excelFileName)
		for _, row := range xlFile.Sheets[0].Rows {
			fmt.Println("2")
			val, _ := row.Cells[1].Int()
			totalSecond += val
		}
	}()
	wg.Wait()
	fmt.Printf("Total First: %v\nTotal Second: %v\nTotal: %v\nTime: %v\n", totalFirst, totalSecond, totalFirst+totalSecond, time.Since(start))
}
