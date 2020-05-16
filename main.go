package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/tmdvs/Go-Emoji-Utils"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	fptr := flag.String("fpath", "history.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)

	var userName string
	messageCount := 0
	lolCount := 0
	lmaoCount := 0
	search := "..."
	var userMessages []string
	var receivedUserMessages []string
	var allMessages []string
	profanitiesMessages := []string{"fuck", "merde", "putain", "ass"}
	sentReceivedWords := []string{"amen", "akpe", "merci", "nagode", "imela", "thanks", "thank you", "alhamdulillah", "shukran"}
	emojisCount := 0
	profanitiesCount := 0
	receivedEmojisCount := 0
	receivedAngryEmojiCount := 0
	sentReceivedWordsCount := 0

	fmt.Print("Enter user name : ")
	fmt.Scanf("%s", &userName)

	for s.Scan() {
		fmt.Println("Search in progress", search)

		//Count Messages sent
		//if strings.Contains(s.Text(), "Olevie KOUAMI") {
		if strings.Contains(s.Text(), userName) {

			//Count lol
			if strings.Contains(s.Text(), "lol") {
				lolCount++
			}

			//Count lmao
			if strings.Contains(s.Text(), "lmao") {
				lmaoCount++
			}

			userMessages = append(userMessages, s.Text())

			messageCount++
		}else{
			receivedUserMessages = append(receivedUserMessages, s.Text())
		}

		allMessages = append(allMessages, s.Text())
	}
	//Count emojis
	input := strings.Join(userMessages, " ")
	result := emoji.FindAll(input)
	emojisCount = len(result)

	//Count profanities
	sort.Strings(profanitiesMessages)
	var itemOccurrences = mapForEach(profanitiesMessages, func(it string) int {
		return strings.Count(strings.ToLower(input), strings.ToLower(it))
	})
	for _, element := range itemOccurrences {
		profanitiesCount = profanitiesCount + element
	}

	//Count received Emojis
	receivedInput := strings.Join(receivedUserMessages, " ")
	receivedResult := emoji.FindAll(receivedInput)
	receivedEmojisCount = len(receivedResult)

	//Count received Angry Emoji
	receivedAngryResult, _ := emoji.Find("😡", receivedInput)
	receivedAngryEmojiCount = receivedAngryResult.Occurrences

	//Count sent Received Words ...
	allInput := strings.Join(allMessages, " ")
	sort.Strings(sentReceivedWords)
	var itemWordOccurrences = mapForEach(sentReceivedWords, func(it string) int {
		return strings.Count(strings.ToLower(allInput), strings.ToLower(it))
	})

	countItemByItem := make(map[string]int)
	sort.Strings(sentReceivedWords)
	for index, element := range itemWordOccurrences {
		sentReceivedWordsCount = sentReceivedWordsCount + element
		countItemByItem[sentReceivedWords[index]] = element
	}

	fmt.Println("Total number of messages sent", messageCount)
	fmt.Println("Total number of times the user sent \"lol\"", lolCount)
	fmt.Println("Total number of times the user sent \"lmao\"", lmaoCount)
	fmt.Println("Total number of times the user sent emojis", emojisCount)
	fmt.Println("Total number of profanities the user sent", profanitiesCount)
	fmt.Println("Total number of times the user received emojis", receivedEmojisCount)
	fmt.Println("Total number of times the user received the angry 😡 emoji", receivedAngryEmojiCount)
	fmt.Println("Total number of times the user sent and recieved the words \"amen\", \"akpe\", \"merci\", \"nagode\", \"imela\", \"thanks\", \"thank you\", \"alhamdulillah\", \"shukran\"", sentReceivedWordsCount)

	//Count sent Received Words ... item by item
	keys := make([]string, 0, len(countItemByItem))
	for k := range countItemByItem {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println("Total number of times the user sent and recieved the word", k, countItemByItem[k])
	}

	//Error
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func mapForEach(arr []string, fn func(it string) int) []int {
	var newArray []int
	for _, it := range arr {
		// We are executing the method passed
		newArray = append(newArray, fn(it))
	}
	return newArray
}
