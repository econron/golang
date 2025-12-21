package nezumi

import (
	"net/http"
	"fmt"
	"os"
)

// エラーが出たということがログにしか出ない。エラーハンドリングになってない。
func BadErrorHandling() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println("error:", err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://www.yahoo.co.jp", "https://badshoteconronhogehoge"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}

type Result struct {
	Error error
	Response *http.Response
}
type ErrorResults struct {
	Errors []error
}
func GoodErrorHandling() ErrorResults {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
					case <-done:
						return
					case results <- result:
				}
			}
		}()
		return results
	}
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://www.yahoo.co.jp", "https://badshoteconronhogehoge"}
	errorResults := ErrorResults{}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
			errorResults.Errors = append(errorResults.Errors, result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
	return errorResults
}

// どこかのステージでcloseしない限り、チャネルのfor rangeが無限に回り続ける。
func PipelinePattern() {
	// 可変長引数を受け取りintStreamを生成
	generator := func(done <-chan interface{}, integers ...int) <- chan int {
		intStream := make(chan int, len(integers))
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}
	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
					case <-done:
						return
					case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(
		done <-chan interface{},
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
					case <-done:
						return
					case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	// 1, doneチャネルを作成
	done := make(chan interface{})
	// 1, それに対するcloseを作成
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4, 5)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipeline {
		fmt.Println(v)
	}
}

type MyError struct {
	Inner error
	Message string
	StackTrace string
	Misc map[string]interface{}
}

func wrapError(err error, message string, stackTrace string, misc map[string]interface{}) MyError {
	return MyError{
		Inner: err,
		Message: message,
		StackTrace: stackTrace,
		Misc: misc,
	}
}

func (err MyError) Error() string {
	return err.Message
}

type LowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{error: err}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}

type IntermediateErr struct {
	error
}