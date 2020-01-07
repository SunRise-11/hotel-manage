package master

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"minornine.com/hotel/src/master/db"
	"minornine.com/hotel/src/master/models"
	"minornine.com/hotel/src/shared"
)

const (
	ok = "OK"
)

// HandleHealth returns OK if the server is healthy.
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ok)
}

// HandleListGameServers lists all game servers for the given ID.
//
// Query parameters:
//	- gameId: The game ID to filter by.
func HandleListGameServers(w http.ResponseWriter, r *http.Request) {
	gid := shared.GameIDType(r.URL.Query().Get("gameId"))
	servers := db.DbGetGameServersByGameId(gid)
	var response models.ListServersResponse
	response.Servers = servers
	json.NewEncoder(w).Encode(response)
}

// HandleGetGameServer returns a particular server by ID.
//
// Path parameters:
//	- id: The server to lookup.
func HandleGetGameServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request: "+err.Error(), http.StatusBadRequest)
		return
	}
	server, exists := db.DbGetGameServerById(models.ServerIDType(id))
	if !exists {
		http.Error(w, "Server by that ID not found.", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(server)
}

// HandleCreateGameServer adds a game server to the database.
func HandleCreateGameServer(w http.ResponseWriter, r *http.Request) {
	config := context.Get(r, ConfigContextKey).(*Config)
	session := context.Get(r, SessionContextKey).(Session)
	server, err := models.DecodeAndValidateGameServer(r.Body, false)
	fillImplicitGameServerFields(&server, r, session)
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check if game definition allows server creation.
	def, ok := config.GetGameDefinition(server.GameID)
	if !ok {
		http.Error(w, fmt.Sprintf("No definition for game ID '%v', and this master server doesn't allow undefined games.", server.GameID), http.StatusBadRequest)
		return
	}
	if def.HostPolicy != shared.HostPolicy_ANY && def.HostPolicy != shared.HostPolicy_CLIENTS_ONLY {
		http.Error(w, fmt.Sprintf("Game '%v' doesn't allow client server creation.", server.GameID), http.StatusBadRequest)
		return
	}

	// Db call.
	server, err = db.DbInsertGameServer(server)
	if err != nil {
		http.Error(w, "Failed to create server: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(server)
}

// HandleSpawnGameServer requests a game server spawn from a spawner instance.
func HandleSpawnGameServer(w http.ResponseWriter, r *http.Request) {
	config := context.Get(r, ConfigContextKey).(*Config)
	// session := context.Get(r, SessionContextKey).(Session)
	gid := shared.GameIDType(r.URL.Query().Get("gameId"))

	// Check if game definition allows spawning.
	def, ok := config.GetGameDefinition(gid)
	if !ok {
		http.Error(w, fmt.Sprintf("No definition for game ID '%v', and this master server doesn't allow undefined games.", gid), http.StatusBadRequest)
		return
	}
	if def.HostPolicy != shared.HostPolicy_ANY && def.HostPolicy != shared.HostPolicy_SPAWN_ONLY {
		http.Error(w, fmt.Sprintf("Game '%v' doesn't allow spawning.", gid), http.StatusBadRequest)
		return
	}

	spawner, err := GetAvailableSpawnerForGame(gid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	server, err := SpawnServerForGame(spawner, gid)
	if err != nil {
		http.Error(w, "Failed to spawn server (internal error).", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(server)
}

// HandleUpdateGameServer updates an existing game server.
//
// Path parameters:
//	- id: The server to lookup.
func HandleUpdateGameServer(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, SessionContextKey).(Session)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	serverId := models.ServerIDType(id)
	existingServer, exists := db.DbGetGameServerById(serverId)
	if !exists {
		http.Error(w, "Server by that ID not found.", http.StatusNotFound)
		return
	}
	if existingServer.SessionID != session.ID {
		http.Error(w, "Not authorized to modify that server.", http.StatusForbidden)
		return
	}
	newServer, err := models.DecodeAndValidateGameServer(r.Body, true)
	if err != nil {
		http.Error(w, "Failed to parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	existingServer.Merge(newServer)
	newServer, err = db.DbUpdateGameServerById(serverId, existingServer)
	if err != nil {
		http.Error(w, "Failed to update server: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(existingServer)
}

// HandleDeleteGameServer deletes an existing game server.
//
// Path parameters:
//	- id: The server to delete.
func HandleDeleteGameServer(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, SessionContextKey).(Session)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// TODO: Full error object should not be returned in production mode
		http.Error(w, "Failed to parse request URL: "+err.Error(), http.StatusBadRequest)
		return
	}
	serverId := models.ServerIDType(id)
	existingServer, exists := db.DbGetGameServerById(serverId)
	if !exists {
		http.Error(w, "Server by that ID not found.", http.StatusNotFound)
		return
	}
	if existingServer.SessionID != session.ID {
		http.Error(w, "Not authorized to modify that server.", http.StatusForbidden)
		return
	}

	err = db.DbDeleteGameServerById(serverId)
	if err != nil {
		http.Error(w, "Failed to delete server: "+err.Error(), http.StatusInternalServerError)
	}
}

func fillImplicitGameServerFields(server *models.GameServer, r *http.Request, session Session) {
	server.SessionID = session.ID
	if len(server.Host) < 1 {
		clientAddr := r.Header.Get("X-Forwarded-For")
		if len(clientAddr) < 1 {
			clientAddr = r.RemoteAddr
		}
		clientAddr = strings.Split(clientAddr, ":")[0]
		// If the host was not provided, assume the host is the request origin.
		log.Printf("Host not specified, inferring from request origin: %v", r.RemoteAddr)
		server.Host = r.RemoteAddr
	}
}
