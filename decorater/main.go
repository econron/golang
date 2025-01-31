package main

import (
	"fmt"
	"log"
	anyfn "decorater/decorater"
)

type Controller struct {
	actions map[string]func(interface{}) string
}

func NewController() *Controller {
	return &Controller{actions: make(map[string]func(interface{}) string)}
}

func (c *Controller) Action(name string, fn func(interface{}) string) {
	c.actions[name] = fn
}

func (c *Controller) Execute(name string, param interface{}) string {
	if fn, exists := c.actions[name]; exists {
		return fn(param)
	}
	return "Action not found"
}

func saveJobs(job interface{}) string {
	jobData, ok := job.(map[string]string) // 型アサーション
	if !ok {
		return "Invalid job data"
	}
	return fmt.Sprintf("Saved job: %s at %s", jobData["title"], jobData["company"])
}

type tekitou struct {
	Val1 string
	Val2 float64
}

func tekitouFunc(tekitoumap map[uint64]*tekitou, tekitoustring string) map[uint64]float64 {
	fmt.Printf("tekitoustring %s", tekitoustring)
	ret := map[uint64]float64{}
	for i,v := range tekitoumap {
		fmt.Printf("tekitou struct Val1: %s", v.Val1)
		ret[i] = v.Val2
	}
	return ret
}

func main() {
	log.Println("Start fn(string)interface{} controller")
	controller := NewController()

	controller.Action("save_jobs", saveJobs)

	job := map[string]string{
		"title": "Software Engineer",
		"company": "Google",
	}
	result := controller.Execute("save_jobs", job)
	log.Println(result)
	log.Println("End fn(string)interface{} controller")

	log.Println("Start any fn controller")
	controller2 := anyfn.NewController()
	controller2.Action("save_jobs", saveJobs)
	controller2.Action("tekitou", tekitouFunc)

	result2 := controller2.Execute("save_jobs", map[string]string{
		"title": "Software Engineer",
		"company": "Apple",
	})

	tm := map[uint64]*tekitou{}
	tm[1] = &tekitou{
		Val1: "tekitou1 val1",
		Val2: 1.1,
	}
	tm[2] = &tekitou {
		Val1: "tekitou2 val1",
		Val2: 2.2,
	}
	result3 := controller2.Execute("tekitou", tm, "tekitou func 2nd arg")
	
	log.Println(result2...)
	log.Println(result3...)

	log.Println("End any fn controller")
}