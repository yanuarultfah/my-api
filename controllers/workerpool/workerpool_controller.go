package workerpool

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"my-api/domain/workerpool"
	"my-api/service"
	"my-api/utils/errors"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const totalWorker = 100

var dataHeaders = make([]string, 0)

func Simpan(c *gin.Context) {
	start := time.Now()
	csvReader, csvFile, err := OpenCsvFile()
	if err != nil {
		log.Fatal(errors.NewInternalServerError("error read file"))
	}
	defer csvFile.Close()

	jobs := make(chan []interface{})
	wg := new(sync.WaitGroup)
	go DispatchWorkers(jobs, wg)
	ReadCsvFilePerLineThenSendToWorker(csvReader, jobs, wg)

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}

func DispatchWorkers(jobs <-chan []interface{}, wg *sync.WaitGroup) {
	var wp workerpool.CsvWorker
	for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
		go func(workerIndex int, wg *sync.WaitGroup) {
			counter := 0
			for job := range jobs {
				_, saveErr := service.WorkersService.SaveCsv(wp, workerIndex, counter, job)
				if saveErr != nil {
					// fmt.Println("Ok..... dispatch worker")
					return
				}
				// doTheJob(workerIndex, counter, job)
				// println(job)
				wg.Done()
				counter++
			}
		}(workerIndex, wg)
	}
	// fmt.Println(wp, "success dispatch")
}

func OpenCsvFile() (*csv.Reader, *os.File, *errors.RestErr) {
	log.Println("=> open csv file")
	stmtcsvdata, err := os.Open("majestic_million.csv")
	if err != nil {
		return nil, nil, errors.NewInternalServerError(err.Error())
	}
	// if errreader := csv.NewReader(stmtcsvdata); errreader != nil {
	// 	return nil,nil,errors.NewBadRequestError("error open csv")
	// }
	reader := csv.NewReader(stmtcsvdata)
	return reader, stmtcsvdata, nil

}

func ReadCsvFilePerLineThenSendToWorker(csvReader *csv.Reader, jobs chan<- []interface{}, wg *sync.WaitGroup) {
	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if len(dataHeaders) == 0 {
			dataHeaders = row
			continue
		}

		rowOrdered := make([]interface{}, 0)
		for _, each := range row {
			rowOrdered = append(rowOrdered, each)
		}

		wg.Add(1)
		jobs <- rowOrdered
	}
	close(jobs)
}
