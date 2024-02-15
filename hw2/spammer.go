package main

import (
	"log"
	"reflect"
	"sort"
	"strconv"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	channels := make([]chan interface{}, len(cmds)+1)
	for i := range channels {
		channels[i] = make(chan interface{})
	}

	wg := sync.WaitGroup{}

	for index, currentCmd := range cmds {
		wg.Add(1)

		go func(command func(in, out chan interface{}), in, out chan interface{}) {
			defer func() {
				wg.Done()
				close(out)
			}()

			command(in, out)
		}(currentCmd, channels[index], channels[index+1])
	}

	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	var users = &sync.Map{}
	wg := sync.WaitGroup{}

	for value := range in {
		wg.Add(1)
		mail, ok := value.(string)
		if !ok {
			log.Printf("[SelectUsers()] Wrong type\tExpected: string Actually: %v", reflect.TypeOf(value))
			return
		}
		go func(email string) {
			defer wg.Done()

			user := GetUser(email)
			_, ok := users.Load(user)
			if !ok {
				users.Store(user, struct{}{})
				out <- user
			}
		}(mail)
	}

	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	wg := sync.WaitGroup{}

	for value := range in {
		wg.Add(1)

		users := make([]User, 1, 1)
		user, ok := value.(User)
		if !ok {
			log.Printf("[SelectUsers()] Wrong type\tExpected: User Actually: %v", reflect.TypeOf(value))
			return
		}
		users[0] = user
		secondValue, ok := <-in
		if ok {
			secondUser, ok := secondValue.(User)
			if !ok {
				log.Printf("[SelectUsers()] Wrong type\tExpected: User Actually: %v", reflect.TypeOf(secondValue))
				return
			}
			users = append(users, secondUser)
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

	for value := range in {
		messageExist := true
		var wg sync.WaitGroup

		for i := 0; i < HasSpamMaxAsyncRequests; i++ {
			if !messageExist {
				break
			}

			msgId, ok := value.(MsgID)
			if !ok {
				log.Printf("[SelectUsers()] Wrong type\tExpected: MsgId Actually: %v", reflect.TypeOf(value))
				return
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
			}(msgId)

			if i < HasSpamMaxAsyncRequests-1 {
				value, messageExist = <-in
			}
		}

		wg.Wait()
	}
}

func CombineResults(in, out chan interface{}) {
	result := make([]MsgData, 0)
	for res := range in {
		msgData, ok := res.(MsgData)
		if !ok {
			log.Printf("[SelectUsers()] Wrong type\tExpected: MsgData Actually: %v", reflect.TypeOf(res))
			return
		}
		result = append(result, msgData)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].HasSpam == result[j].HasSpam {
			return result[i].ID < result[j].ID
		}
		return result[i].HasSpam && !result[j].HasSpam
	})

	for _, res := range result {
		out <- strconv.FormatBool(res.HasSpam) + " " + strconv.FormatUint(uint64(res.ID), 10)
	}
}
