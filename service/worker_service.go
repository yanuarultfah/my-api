package service

import (
	"my-api/domain/workerpool"
	"my-api/utils/errors"
)

var (
	WorkersService workersServiceInterface = &workersService{}
)

type workersService struct{}

type workersServiceInterface interface {
	SaveCsv(workerpool.CsvWorker, int, int, []interface{}) (*workerpool.CsvWorker, *errors.RestErr)
}

func (w *workersService) SaveCsv(wp workerpool.CsvWorker, workerIndex int,
	counter int,
	values []interface{}) (*workerpool.CsvWorker, *errors.RestErr) {
	if err := wp.DoTheJob(workerIndex, counter, values); err != nil {
		return nil, err
	}

	return &wp, nil
}
