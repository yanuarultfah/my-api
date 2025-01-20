package workerpool

import (
	"context"
	"fmt"
	"log"
	usersdb "my-api/datasource/pgsql/users_db"
	"my-api/utils/errors"
)

const (
	query = `INSERT INTO domain (GlobalRank,
    TldRank,
    Domain,
    TLD,
    RefSubNets,
    RefIPs,
    IDN_Domain,
    IDN_TLD,
    PrevGlobalRank,
    PrevTldRank,
    PrevRefSubNets,
    PrevRefIPs) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);`
)

// func dispatchWorkers(jobs <-chan []interface{}, wg *sync.WaitGroup) {
// 	for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
// 		go func(workerIndex int, jobs <-chan []interface{}, wg *sync.WaitGroup) {
// 			counter := 0
// 			for job := range jobs {
// 				// runjob, saveErr := service.WorkersService.SaveCsv()
// 				 DoTheJob(workerIndex, counter, job)
// 				wg.Done()
// 				counter++
// 			}
// 		}(workerIndex, jobs, wg)
// 	}
// }

func (csvWorker *CsvWorker) DoTheJob(workerIndex int,
	counter int,
	values []interface{}) *errors.RestErr {
	// var (
	// 	workerIndex int
	// 	counter     int
	// 	values      []interface{}
	// )
	for {
		var outerError error
		func(outerError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()
			conn, err := usersdb.Client.Conn(context.Background())
			// query := fmt.Sprintf("INSERT INTO domain (%s) VALUES (%s)",
			// 	strings.Join(dataHeaders, ","),
			// 	strings.Join(generateQuestionsMark(len(dataHeaders)), ","),
			// )

			// stmt, err := usersdb.Client.Prepare(query)
			// if err != nil {
			// 	log.Fatal(err.Error())
			// }
			// defer stmt.Close()

			_, err = conn.ExecContext(context.Background(), query, values...)
			if err != nil {
				log.Fatal(err.Error())
			}
			err = conn.Close()
			if err != nil {
				log.Fatal(err.Error())
			}
			// saveCsv := stmt.QueryRow(csvWorker.GlobalRank, csvWorker.TldRank, csvWorker.Domain, csvWorker.Tld, csvWorker.RefSubNets, csvWorker.RefIPs, csvWorker.IDN_Domain, csvWorker.IDN_TLD, csvWorker.PrevGlobalRank, csvWorker.PrevTldRank, csvWorker.PrevRefSubNets, csvWorker.PrevRefIPs)
			// if saveCsv != nil {
			// 	log.Fatal(err.Error())
			// }
		}(&outerError)
		if outerError == nil {
			break
		}
	}
	if counter%100 == 0 {
		log.Println("=> worker", workerIndex, "inserted", counter, "data")
	}

	return nil
}

// func generateQuestionsMark(n int) []string {
// 	s := make([]string, 0)
// 	for i := 0; i < n; i++ {
// 		var newval = "$" + strconv.Itoa(i)
// 		s = append(s, newval)
// 	}
// 	return s
// }
