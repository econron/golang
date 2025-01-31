package main

import (
	"fmt"
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

func main() {
	controller := NewController()

	controller.Action("save_jobs", saveJobs)

	job := map[string]string{
		"title": "Software Engineer",
		"company": "Google",
	}
	result := controller.Execute("save_jobs", job)
	fmt.Println(result)
}