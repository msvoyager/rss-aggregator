package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/msvoyager/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID  uuid.UUID `json:"feedid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("ERROR PARSING JSON: %v", err))
		return
	}

	follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldnt create feed follow: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedFollowToFeedFollow(follow))
}

// func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

// 	feeds, err := apiCfg.DB.GetFeeds(r.Context())

// 	if err != nil {
// 		respondWithError(w, 400, fmt.Sprintf("couldnt get feeds: %v", err))
// 		return
// 	}

// 	respondWithJson(w, 201, databaseFeedsToFeeds(feeds))

// }
