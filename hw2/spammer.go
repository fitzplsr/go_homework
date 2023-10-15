package main

import (
	"log"
	"sort"
	"strconv"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	channels := make([]chan interface{}, len(cmds)+1)
	for i := 0; i < len(cmds)+1; i++ {
		channels[i] = make(chan interface{})
	}

	wg := sync.WaitGroup{}

	for number, currentCmd := range cmds {
		wg.Add(1)

		go func(command func(in, out chan interface{}), in, out chan interface{}) {
			defer func() {
				wg.Done()
				close(out)
			}()

			command(in, out)
		}(currentCmd, channels[number], channels[number+1])
	}

	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	var users = &sync.Map{}
	wg := sync.WaitGroup{}

	for mail := range in {
		wg.Add(1)

		go func(email string) {
			defer wg.Done()

			user := GetUser(email)
			_, ok := users.Load(user)
			if !ok {
				users.Store(user, struct{}{})
				out <- user
			}
		}(mail.(string))
	}

	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	wg := sync.WaitGroup{}

	for user := range in {
		wg.Add(1)

		users := make([]User, 1)
		users[0] = user.(User)
		secondUser, ok := <-in
		if ok {
			users = append(users, secondUser.(User))
		}

		go func(users []User) {
			defer wg.Done()

			messages, err := GetMessages(users...)
			if err != nil {
				log.Println(err.Error())
				return
			}

			for _, message := range messages {
				out <- message
			}
		}(users)
	}

	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	const maxHasSpamConnections = 5

	for message := range in {
		messageExist := true
		var wg sync.WaitGroup

		for i := 0; i < maxHasSpamConnections; i++ {
			if !messageExist {
				break
			}

			wg.Add(1)

			go func(id MsgID) {
				defer wg.Done()

				hasSpam, err := HasSpam(id)
				if err != nil {
					log.Println(err.Error())
					return
				}

				out <- MsgData{id, hasSpam}
			}(message.(MsgID))

			if i < maxHasSpamConnections-1 {
				message, messageExist = <-in
			}
		}

		wg.Wait()
	}
}

type ResultsSorter []MsgData

func (a ResultsSorter) Len() int {
	return len(a)
}

func (a ResultsSorter) Less(i, j int) bool {
	if a[i].HasSpam == a[j].HasSpam {
		return a[i].ID < a[j].ID
	}
	return a[i].HasSpam && !a[j].HasSpam
}

func (a ResultsSorter) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func CombineResults(in, out chan interface{}) {
	result := make([]MsgData, 0)
	for res := range in {
		result = append(result, res.(MsgData))
	}

	sort.Sort(ResultsSorter(result))

	for _, res := range result {
		out <- strconv.FormatBool(res.HasSpam) + " " + strconv.FormatUint(uint64(res.ID), 10)
	}
}
