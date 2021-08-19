package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

// Worker functions for hash

// SingleHash: crc32(data) + "~" + crc32(md5(data))
func SingleHashRoutine(dataRaw interface{}, out chan interface{}, wg_global *sync.WaitGroup, mu *sync.Mutex) {
	defer wg_global.Done()

	wg := &sync.WaitGroup{} // wait group to syncronize calcualtions
	var crc32, md5, crc32md5 string

	dataInt := dataRaw.(int)
	data := strconv.Itoa(dataInt)

	// Parallel calculation of single hashes
	wg.Add(1)
	go func(data string) {
		defer wg.Done()

		crc32 = DataSignerCrc32(data)
	}(data)

	wg.Add(1)
	go func(data string) {
		defer wg.Done()

		// Lock md5 call as parallel call causes overheat
		mu.Lock()
		md5 = DataSignerMd5(data)
		mu.Unlock()

		crc32md5 = DataSignerCrc32(md5)
	}(data)

	wg.Wait()
	result := crc32 + "~" + crc32md5

	// Output for better understanding what is happenning
	fmt.Println("func SingleHash(): data =", data)
	fmt.Println("func SingleHash(): crc32(data) =", crc32)
	fmt.Println("func SingleHash(): md5(data) =", md5)
	fmt.Println("func SingleHash(): crc32(md5(data)) =", crc32md5)
	fmt.Println("func SingleHash(): result =", result)

	out <- result
}

func SingleHash(in, out chan interface{}) {
	wg_global := &sync.WaitGroup{} // wait for all goroutines to end
	mu := &sync.Mutex{}

	for dataRaw := range in {
		wg_global.Add(1)
		go SingleHashRoutine(dataRaw, out, wg_global, mu)
	}

	wg_global.Wait()
}

// MultiHash: crc32(th + data), th = 0..5
func MultiHashRoutine(dataRaw interface{}, out chan interface{}, wg_global *sync.WaitGroup) {
	defer wg_global.Done()

	wg := &sync.WaitGroup{} // wait group to syncronize calcualtions

	data := dataRaw.(string)
	fmt.Println("func MultiHash(): data =", data)

	var result string
	tmpResults := make([]string, 6)
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(data string, th int) {
			defer wg.Done()

			tmpResult := DataSignerCrc32(strconv.Itoa(th) + data)
			fmt.Println("func MultiHash(): th =", th, "crc32(th+data)=", tmpResult)
			tmpResults[th] = tmpResult
		}(data, i)
	}

	wg.Wait()
	for _, val := range tmpResults {
		result += val
	}

	fmt.Println("func MultiHash(): result =", result)
	out <- result
}

func MultiHash(in, out chan interface{}) {
	wg_global := &sync.WaitGroup{}

	for dataRaw := range in {
		wg_global.Add(1)
		go MultiHashRoutine(dataRaw, out, wg_global)
	}

	wg_global.Wait()
}

// CombineResults: get all results, sort them, appends using "_" and outputs
func CombineResults(in, out chan interface{}) {
	var combinedResult string
	allResults := []string{}

	// Collect all results into slice
	for data := range in {
		data_str := data.(string)
		allResults = append(allResults, data_str)
	}

	// Sort and combine them
	resultNumber := len(allResults)
	sort.StringSlice(allResults).Sort()
	for _, val := range allResults[0 : resultNumber-1] {
		combinedResult += val + "_"
	}

	// Insert last result and send to output channel
	combinedResult += allResults[resultNumber-1]
	fmt.Println("func CombineResult(): result =", combinedResult)
	out <- combinedResult
}

// Create a goroutine with the given worker and close related channel
// when the worker stops working
func LaunchWorker(worker job, in, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	worker(in, out)
	// Close the output channel as it is not needed anymore for writing
	if out != nil {
		close(out)
	}
}

// Execute workers of type job as unix pipeline - each functions's out channel
// serves as input channel to the next one
func ExecutePipeline(workers ...job) {
	// Created waiting group
	wg := &sync.WaitGroup{}

	// Allocate required resources
	workersAmount := len(workers)
	pipelineChannels := make([]chan interface{}, workersAmount+1)
	for i := 1; i < workersAmount; i++ {
		pipelineChannels[i] = make(chan interface{})
	}

	// Launch all workers
	for i := 0; i < workersAmount; i++ {
		wg.Add(1)
		go LaunchWorker(workers[i], pipelineChannels[i], pipelineChannels[i+1], wg)
	}

	wg.Wait()
}

func main() {
	inputData := []int{0, 1, 1, 2, 3, 5, 8}

	hashSignJobs := []job{
		job(func(in, out chan interface{}) {
			for _, fibNum := range inputData {
				out <- fibNum
			}
		}),
		job(SingleHash),
		job(MultiHash),
		job(CombineResults),
		job(func(in, out chan interface{}) {
			dataRaw := <-in
			_, ok := dataRaw.(string)
			if !ok {
				panic("Can't convert result to string")
			}
		}),
	}

	ExecutePipeline(hashSignJobs...)
}
