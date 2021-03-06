package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/kp54/ayb-bot/util"

	_ "github.com/joho/godotenv/autoload"
)

type Note struct {
	Text string `json:"text"`
}

func loadNotes(filename string) ([]string, error) {
	var (
		err   error
		bytes []byte
		notes []string
	)

	if bytes, err = ioutil.ReadFile(filename); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bytes, &notes); err != nil {
		return nil, err
	}

	return notes, nil
}

func constructNote(text string) Note {
	return Note{
		Text: text,
	}
}

func joinAuthorization(payload *[]byte) error {
	var (
		err error
		tmp map[string]string
	)

	json.Unmarshal(*payload, &tmp)
	tmp["i"] = os.Getenv("AYB_BOT_AUTHORIZATION_TOKEN")
	if *payload, err = json.Marshal(tmp); err != nil {
		return err
	}

	return nil
}

func postNote(note Note) ([]byte, error) {
	payload, err := json.Marshal(&note)
	if err != nil {
		return nil, err
	}

	joinAuthorization(&payload)
	fmt.Println("payload: " + string(payload))

	baseURL := os.Getenv("AYB_BOT_BASE_URL")
	destURL, err := util.JoinBaseAndPath(baseURL, "notes/create")
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(destURL, "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	token := os.Getenv("AYB_BOT_AUTHORIZATION_TOKEN")
	if token == "" {
		log.Fatal("env AYB_BOT_AUTHORIZATION_TOKEN not defined")
	}
	baseURL := os.Getenv("AYB_BOT_BASE_URL")
	if baseURL == "" {
		log.Fatal("env AYB_BOT_BASE_URL not defined")
	}
	dur := os.Getenv("AYB_BOT_POST_DURATION")
	if dur == "" {
		log.Fatal("env AYB_BOT_POST_DURATION not defined")
	}
	durI, err := strconv.Atoi(dur)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	notes, err := loadNotes("notes.json")
	if err != nil {
		log.Fatal(err)
	}
	rand.Shuffle(len(notes), func(i, j int) {
		notes[i], notes[j] = notes[j], notes[i]
	})

	fmt.Println("Initialization Finished")

	index := 0
	scheduler.Every(time.Duration(durI)).Minutes().Run(func() {
		note := constructNote(notes[index])
		resp, err := postNote(note)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("response: " + string(resp))

		if index++; len(notes) == index {
			index = 0
		}
	})
	runtime.Goexit()
}
