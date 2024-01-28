package handlers

import (
    "encoding/json"
    "net/http"
	"strings"
    "github.com/gorilla/mux"
    "github.com/x6txy/tsis1/api"
	"strconv"
)

func GetPlayers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(api.Players)
}

func GetPlayerByName(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    playerName := params["last_name"]


    playerName = strings.ToLower(playerName)

    for _, player := range api.Players {
        if strings.ToLower(player.LastName) == playerName {
            json.NewEncoder(w).Encode(player)
            return
        }
    }

    http.Error(w, "Player not found", http.StatusNotFound)
}

func GetPlayerById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    playerId := params["id"]

    for _, player := range api.Players {
        if strconv.Itoa(player.Id) == playerId {
            json.NewEncoder(w).Encode(player)
            return
        }
    }

    http.Error(w, "Player not found", http.StatusNotFound)
}

func GetPlayersByPosition(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    position := params["position"]

    position = strings.ToLower(position)
    var matchingPlayers []api.Player

    for _, player := range api.Players {
        if strings.ToLower(player.Position) == position {
            matchingPlayers = append(matchingPlayers, player)
        }
    }

    if len(matchingPlayers) > 0 {
        json.NewEncoder(w).Encode(matchingPlayers)
    } else {
        http.Error(w, "Players with the specified position not found", http.StatusNotFound)
    }
}

func GetPlayersByNation(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    nation := params["nation"]

    nation = strings.ToLower(nation)
    var matchingPlayers []api.Player

    for _, player := range api.Players {
        if strings.ToLower(player.Nation) == nation {
            matchingPlayers = append(matchingPlayers, player)
        }
    }

    if len(matchingPlayers) > 0 {
        json.NewEncoder(w).Encode(matchingPlayers)
    } else {
        http.Error(w, "Players from the specified nation not found", http.StatusNotFound)
    }
}


func HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    

    healthStatus := "App is healthy!\n\n"


    appDescription := "This is a simple web API providing information about football players. It allows users to retrieve player details, perform searches, and check the overall health of the application.\n"
    authorInfo := "\n\nAuthor: Bahauddin\n"


    healthCheckResponse := healthStatus + appDescription + authorInfo


    w.Write([]byte(healthCheckResponse))
}
