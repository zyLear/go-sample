package main

import (
	"fmt"
	"sort"
	"sync"
)

type person struct {
	name  string
	score int64
	age   int64
}

func main() {
	for i := 0; i < 30; i++ {
		result, err := BlockSubmitAndWaitAllFinish([]TaskReturnValue[[]any]{
			func() ([]any, error) {
				return []any{"1", "2f"}, nil
			},
			func() ([]any, error) {
				return []any{"3", 4.3}, nil
			},
			func() ([]any, error) {
				return []any{"5", 6}, nil
			},
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v\n", result)
	}

	t := &person{}

	err := BlockSubmitAndWaitAllFinishWithValue(t, []TaskWithValue[*person]{
		func(t *person) error {
			t.name = "name"
			return nil
		}, func(t *person) error {
			t.score = 11
			return nil
		}, func(t *person) error {
			t.age = 22
			return nil
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", t)

	for i := 0; i < 30; i++ {
		result, err := BlockSubmitAndWaitAllFinishWithPriority([]TaskWithPriority[[]any]{
			{func() ([]any, error) {
				return []any{"1", "2f"}, nil
			},
				int64(1),
			},
			{func() ([]any, error) {
				return []any{"3", 4.3}, nil
			},
				int64(3),
			},
			{func() ([]any, error) {
				return []any{"5", 6}, nil
			},
				int64(5),
			},
			{func() ([]any, error) {
				return []any{"1", "bbbb"}, nil
			},
				int64(1),
			},
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v\n", result)
	}
}

type TaskReturnValue[T any] func() (T, error)

type TaskWithValue[T any] func(t T) error

type TaskWithPriority[T any] struct {
	Task     TaskReturnValue[T]
	priority int64
}

type taskResult[T any] struct {
	Err  error
	Data T
}

func run(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		fn()
	}()
}

func BlockSubmitAndWaitAllFinishWithPriority[T any](tasks []TaskWithPriority[T]) ([]T, error) {
	priorityTaskMap := make(map[int64][]TaskReturnValue[T])

	for _, task := range tasks {
		priorityTaskMap[task.priority] = append(priorityTaskMap[task.priority], task.Task)
	}
	priorityList := make([]int64, 0)
	for priority, _ := range priorityTaskMap {
		priorityList = append(priorityList, priority)
	}
	sort.Slice(priorityList, func(i, j int) bool {
		return priorityList[i] < priorityList[j]
	})
	result := make([]T, 0)
	for _, priority := range priorityList {
		priorityTasks := priorityTaskMap[priority]
		data, err := BlockSubmitAndWaitAllFinish(priorityTasks)
		if err != nil {
			return nil, err
		}
		result = append(result, data...)
	}
	return result, nil
}

func BlockSubmitAndWaitAllFinish[T any](tasks []TaskReturnValue[T]) ([]T, error) {
	if len(tasks) == 0 {
		return nil, nil
	}
	resultChan := make(chan *taskResult[T], len(tasks))
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(tasks))
	for _, task := range tasks {
		// 使用同名变量保留当前遍历状态
		task := task
		run(func() {
			defer waitGroup.Done()
			data, err := task()
			resultChan <- &taskResult[T]{
				Data: data,
				Err:  err,
			}
		})
	}
	waitGroup.Wait()
	close(resultChan)
	result := make([]T, 0, len(tasks))
	for taskResult := range resultChan {
		if taskResult.Err != nil {
			return nil, taskResult.Err
		}
		result = append(result, taskResult.Data)
	}
	return result, nil
}

func BlockSubmitAndWaitAllFinishWithValue[T any](t T, tasks []TaskWithValue[T]) error {
	if len(tasks) == 0 {
		return nil
	}
	resultChan := make(chan error, len(tasks))
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(tasks))
	for _, task := range tasks {
		// 使用同名变量保留当前遍历状态
		task := task
		run(func() {
			defer waitGroup.Done()
			err := task(t)
			if err != nil {
				resultChan <- err
			}
		})
	}
	waitGroup.Wait()
	close(resultChan)
	for err := range resultChan {
		if err != nil {
			return err
		}
	}
	return nil
}
