package spotify

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

func New() *spotify.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		Scopes: []string{
			spotifyauth.ScopePlaylistReadCollaborative,
			spotifyauth.ScopePlaylistReadPrivate,
		},
		TokenURL: spotifyauth.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(context.Background(), token)
	return spotify.New(httpClient)
}

func UserPlaylists(ctx context.Context, client *spotify.Client, user string) ([]string, error) {
	playlists, err := client.GetPlaylistsForUser(ctx, user)
	if err != nil {
		return nil, err
	}

	playlistNames := []string{}
	for i, playlist := range playlists.Playlists {
		if i < 3 {
			p := fmt.Sprintf("<img src=%q alt=%q height=%d/>", playlist.Images[0].URL, playlist.ID.String(), 320)
			playlistNames = append(playlistNames, p)
		}
	}

	return playlistNames, nil
}
