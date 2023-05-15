package controllers

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

func GetCurrentOrLast(c *gin.Context) {
	config := &oauth2.Config{
		ClientID:     "c4a1d44eea4b4b00a45dabe89d3d7238",
		ClientSecret: "58fc5b44408140cda5f0145ff01ea62f",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		},
	}

	token := &oauth2.Token{
		AccessToken:  "",
		RefreshToken: "AQDXM8i5ctzhTu3y7TvfMmbNTYLUS4Qf3nGvKDmukvTI-eGXM26_5EJ_r9Oww-gIh29WN_DspgVuktTt5XVU9Rks9ftCcF8R0GO57EqJ0S2VneoTLMVULIB-BAcn4dPbmF4",
	}

	client := config.Client(context.Background(), token)

	sp := spotify.NewClient(client)

	if token.Expiry.Before(time.Now()) {
		// Refresh the access token
		tokenSource := config.TokenSource(context.Background(), token)
		newToken, err := tokenSource.Token()
		if err != nil {
			log.Fatal(err)
		}

		// Update the access token with the refreshed token
		token = newToken
	}

	current, err := sp.PlayerCurrentlyPlaying()
	if err != nil {
		log.Fatal(err)
	}

	if current.Playing {
		c.JSON(200, gin.H{"name": current.Item.Name, "artist": current.Item.Artists[0].Name, "imgUrl": current.Item.Album.Images[0].URL})
	} else {
		recents, err := sp.PlayerRecentlyPlayed()
		if err != nil {
			log.Fatal(err)
		}

		last, err := sp.GetTrack(spotify.ID(recents[0].Track.ID))
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200, gin.H{"name": last.Name, "artist": last.Artists[0].Name, "imgUrl": last.Album.Images[0].URL})
	}
}
