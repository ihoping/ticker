package main

import (
	"github.com/gen2brain/dlgs"
	"github.com/robfig/cron/v3"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	configTasks := &tasks{}

	data, err := readConfig("ticker.yml")
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = yaml.Unmarshal(data, configTasks)
	if err != nil {
		log.Fatalf("%s", err)
	}

	c := cron.New(cron.WithSeconds())

	for _, task := range configTasks.Tasks {
		j := job{task: task}
		_, _ = c.AddJob(task.Crontab, j)
	}

	c.Start()
	select {}
}

type task struct {
	Type    string
	Title   string
	Content string
	Crontab string `json:"crontab"`
}

type tasks struct {
	Tasks []task
}

type job struct {
	task task
}

func (j job) Run() {
	if j.task.Type == "info" {
		info(j.task.Title, j.task.Content)
	} else if j.task.Type == "err" {
		err(j.task.Title, j.task.Content)
	} else if j.task.Type == "warn" {
		warn(j.task.Title, j.task.Content)
	}
}

func readConfig(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func info(title, text string) {
	_, err := dlgs.Info(title, text)
	if err != nil {
		panic(err)
	}
}

func warn(title, text string) {
	_, err := dlgs.Warning(title, text)
	if err != nil {
		panic(err)
	}
}

func err(title, text string) {
	_, err := dlgs.Error(title, text)
	if err != nil {
		panic(err)
	}
}
