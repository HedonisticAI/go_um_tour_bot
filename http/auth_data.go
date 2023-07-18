package get_pair

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/HedonisticAI/go_um_tour_bot/map_char"
)

type User struct {
	Username     string
	Pwd          string
	TournamentID string
}

func (User *User) Draft(Char *map_char.Content_Set, Map *map_char.Content_Set) string {
	var data Tour
	var ret string
	h := http.Client{}
	url := "https://api.challonge.com/v1/tournaments/" + User.TournamentID + "/matches.json"
	log.Print("making tour req:" + url)
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(User.Username, User.Pwd)
	res, err := h.Do(req)
	if err != nil {
		log.Panic(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	log.Println(string(body) + "\n")
	if err != nil {
		log.Panic(err.Error())
	}
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Print(err.Error())
		return "Error while unmarchalling json"
	}
	for index := range data {
		if data[index].Match.State != "complete" {
			if data[index].Match.Player1ID != 0 {
				ret += fmt.Sprintf("For pair %s - %s\n", User.get_name(data[index].Match.Player1ID), User.get_name(data[index].Match.Player2ID))
				ret += fmt.Sprintf("Characters: %s\n", Char.Take_rand(7))
				ret += fmt.Sprintf("Maps: %s\n\n", Map.Take_rand(5))
			}
		}
	}
	return ret + "."
}

func (User *User) get_name(id int) string {

	var result ParticipantData
	h := http.Client{}
	url := "https://api.challonge.com/v1/tournaments/" + User.TournamentID + "/participants/" + strconv.Itoa(id) + ".json"
	log.Print("making user request:" + url)
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(User.Username, User.Pwd)
	res, err := h.Do(req)
	if err != nil {
		log.Panic(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err.Error())
		return err.Error()
	}

	log.Println(string(body) + "\n")
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return result.Participant.Name
}
