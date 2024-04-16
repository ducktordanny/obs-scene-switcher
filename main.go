package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/joho/godotenv"
)

func printScenes(client *goobs.Client) {
	sceneList, err := client.Scenes.GetSceneList()
	if err != nil {
		panic(err)
	}
	scenes := sceneList.Scenes
	for _, scene := range scenes {
		fmt.Println(scene.SceneName)
	}
}

func switchSceneTo(client *goobs.Client, sceneName string) {
	params := scenes.NewSetCurrentProgramSceneParams().WithSceneName(sceneName)
	client.Scenes.SetCurrentProgramScene(params)
}

func main() {
	var sceneName string
	flag.StringVar(&sceneName, "name", "", "Name of the scene to switch to.")
	flag.StringVar(&sceneName, "n", "", "Alias for 'name'.")
	flag.Parse()

	// TODO: We may not need this in the end
	err := godotenv.Load(".env")
	if err != nil {
		panic("Could not load local .env file.")
	}

	obsWebsocketPassword := os.Getenv("OBS_WS_PW")
	if obsWebsocketPassword == "" {
		panic("Empty environment variable")
	}
	client, err := goobs.New("localhost:4455", goobs.WithPassword(obsWebsocketPassword))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()

	if sceneName != "" {
		switchSceneTo(client, sceneName)
		return
	}
	printScenes(client)
}
