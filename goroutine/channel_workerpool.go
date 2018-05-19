package main

import (
    "time"
    "sync"
    "math/rand"
    "fmt"
)

type Job struct {
    id     int
    number int
}

type Result struct {
    job          Job
    sumOfNumbers int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func main() {
    numberOfJobs := 200
    go createJobs(numberOfJobs)
    done := make(chan bool)
    go printResults(done)
    numberOfWorkers := 10
    createWorkerPool(numberOfWorkers)
    <- done
}

func digits(number int) int {
    sum := 0
    no := number
    for no != 0 {
        digit := no % 10
        sum += digit
        no /= 10
    }
    time.Sleep(1 * time.Second)
    return sum
}

func createJobs(numberOfJobs int) {
    for i := 0; i < numberOfJobs; i++ {
        randomNumber := rand.Intn(999)
        job := Job{i, randomNumber}
        jobs <- job
    }
    close(jobs)
}

func createWorker(wg *sync.WaitGroup) {
    for job := range jobs {
        result := Result{job, digits(job.number)}
        results <- result
    }
    wg.Done()
}

func createWorkerPool(numberOfWorkers int) {
    var wg sync.WaitGroup
    for i := 0; i < numberOfWorkers; i++ {
        wg.Add(1)
        go createWorker(&wg)
    }
    wg.Wait()
    close(results)
}

func printResults(done chan bool) {
    for result := range results {
        fmt.Printf("Job id %d, input number %d, result %d\n", result.job.id, result.job.number, result.sumOfNumbers)
    }
    done <- true
}
