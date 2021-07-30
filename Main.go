package rbxgamelog

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/DisgoOrg/disgohook"
	"github.com/joho/godotenv"
	"github.com/shirou/gopsutil/process"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

var (
	placeID   string
	reset     = false
	hook, err = disgohook.NewWebhookByToken(nil, nil, goDotEnvVariable("HOOK"))
)

type MarketPlaceInfo struct {
	Name string `json:"Name"`
}

func getProcessByName(targetProcessName string) *process.Process {
	processes, _ := process.Processes()

	for _, proc := range processes {
		name, _ := proc.Name()
		if name == targetProcessName {
			return proc
		}
	}
	return nil
}

func GetPlaceInfoByPlaceId(placeID string) *MarketPlaceInfo {
	url := "https://api.roblox.com/marketplace/productinfo?assetId=" + placeID
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	var info *MarketPlaceInfo
	json.NewDecoder(resp.Body).Decode(&info)
	return info
}

func updateRobloxPresence() {
	roblox := GetProcessByName("RobloxPlayerBeta.exe")

	for roblox == nil {
		roblox = GetProcessByName("RobloxPlayerBeta.exe")

		if reset == false {
			reset = true

			client.Logout()
			fmt.Println("reset client activity")
		}
	}

	err := client.Login("823294557155754005")

	if err != nil {
		fmt.Println(err)
	}
	reset = false
	args, _ := roblox.Cmdline()

	placePattern := regexp.MustCompile(`placeId=(\d+)`)
	placeMatch := placePattern.FindStringSubmatch(args)[1]

	if placeMatch != placeId {
		placeId = placeMatch
		place := GetPlaceInfoByPlaceId(placeId)
		message, err := webhook.SendContent("`Played game:`\nhttps://www.roblox.com/games/" + placeId + "/- `at time:`" + time.Now())
	}
}

func main() {
	for true {
		updateRobloxPresence()
		time.Sleep(time.Second * goDotEnvVariable("INTERVAL"))
	}
}
