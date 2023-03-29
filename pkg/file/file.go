package file

import (
	"awesomeProject/pkg/model"
	"awesomeProject/pkg/storage"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func OpenFileCsv(ms *storage.MemStorage) {
	csvFile, err := os.Open("cities.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for i := 0; ; i++ {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal("err!=nil ", error)
		}
		idint, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println("ошибка парсинга", err)
		}
		populationint, err := strconv.Atoi(line[4])
		if err != nil {
			fmt.Println("ошибка парсинга", err)
		}
		poundationint, err := strconv.Atoi(line[5])
		if err != nil {
			fmt.Println("ошибка парсинга", err)
		}
		var k model.CityStruct = model.CityStruct{idint, line[1], line[2], line[3], populationint, poundationint}
		ms.Records[i] = &k
	}

}

func CloseFileCsv(ms *storage.MemStorage) {
	fmt.Println("Сохраняем файл")
	csvFile, _ := os.Create("cities.csv")
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, record := range ms.Records {
		var s []string
		s = append(s, strconv.Itoa(record.Id))
		s = append(s, record.Name)
		s = append(s, record.Region)
		s = append(s, record.District)
		s = append(s, strconv.Itoa(record.Population))
		s = append(s, strconv.Itoa(record.Foundation))
		writer.Write(s)
	}
	writer.Flush()
}
