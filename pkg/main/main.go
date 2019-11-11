package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/step/saurontypes"
)

func getMessages(num int) []saurontypes.AngmarMessage {
	messages := []saurontypes.AngmarMessage{}
	repos := []struct {
		Name, Repo string
	}{
		{"craftybones", "key-val-parser"},
		{"luciferankon", "prime_number"},
		{"luciferankon", "data_analysis"},
		{"luciferankon", "tic_tac_toe_go"},
		{"luciferankon", "gauge-js"},
		{"luciferankon", "tic_tac_toe_ruby"},
		{"luciferankon", "instagram"},
		{"luciferankon", "tic-tac-toe"},
	}
	for i := 0; float64(i) < math.Min(float64(num), float64(len(repos))); i++ {
		repo := repos[i]
		url := "https://api.github.com/repos/__NAME__/__REPO__/tarball/refs/heads/master"
		url = strings.Replace(url, "__NAME__", repo.Name, 1)
		url = strings.Replace(url, "__REPO__", repo.Repo, 1)
		msg := saurontypes.AngmarMessage{
			Url:     url,
			SHA:     "master",
			Pusher:  repo.Name,
			Project: repo.Repo,
			Tasks: []saurontypes.Task{
				{Queue: "test", ImageName: "orc_sample"},
				{Queue: "lint", ImageName: "orc_sample"},
			},
		}
		messages = append(messages, msg)
	}
	return messages
}

func main() {
	redisClient := getRedisClient()
	flag.Parse()
	messages := getMessages(numberOfMessages)
	for _, message := range messages {
		m, _ := json.Marshal(message)
		redisClient.Enqueue(queueName, string(m))
		fmt.Println(message.String())
	}
}
