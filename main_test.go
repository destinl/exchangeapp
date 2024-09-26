package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTaskControdl(t *testing.T) {
	taskNum := 5

	wg := sync.WaitGroup{}
	wg.Add(taskNum)

	for i := 0; i < taskNum; i++ {
		go func(i int) {
			fmt.Println("info", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func TestClose(t *testing.T) {
	test := make(chan int, 10)

	go func(info chan int) {
		for {
			select {
			case val, ok := <-test:
				if !ok {
					t.Logf("Channel Closed!")
					return
				}
				t.Logf("data %d\n", val)
			}
		}
	}(test)

	go func() {
		test <- 1
		time.Sleep(1 * time.Second)
		test <- 2

		close(test)
	}()

	time.Sleep(5 * time.Second)
}

func TestCloseByAnother(t *testing.T) {
	test := make(chan int, 5)
	exit := make(chan struct{})

	go func(info chan int, exit chan struct{}) {
		for {
			select {
			case val := <-info:
				t.Logf("data %d\n", val)

			case <-exit:
				t.Logf("Task Exit!!\n")
				return
			}
		}
	}(test, exit)

	go func() {
		test <- 1
		time.Sleep(1 * time.Second)
		test <- 2
		close(exit)
	}()
	time.Sleep(5 * time.Second)
}

func TestTimeControl(t *testing.T) {
	test := make(chan int, 5)

	go func(info chan int) {
		for {
			select {
			case val := <-info:
				t.Logf("Data %d\n", val)

			case <-time.After(2 * time.Second):
				t.Logf("Time out!\n")
				return
			}
		}
	}(test)

	go func() {
		test <- 1
		time.Sleep(2 * time.Second) //>=2
		test <- 2
	}()

	time.Sleep(5 * time.Second)
}

func Test(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				t.Log("Context cancelled!")
				return
			}
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	time.Sleep(2 * time.Second)
}
