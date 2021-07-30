package robloxgamelog

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"time"
	"regexp"
	"net/http"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/nickname32/discordhook"
)

godotenv.Load()

var (
	placeID string
	reset = false
)

type MarketPlaceInfo struct {
	Name	string	`json:"Name"`
}

func getProcessByName(targetProcessName string) *process.Process {
	processes, _ := process.Processes()

	for _, proc := range processes {
		name, _ := proc.Name()
		if (name == targetProcessName) {
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

func UpdateRobloxPresence() {
	roblox := GetProcessByName("RobloxPlayerBeta.exe")

	for (roblox == nil) {
		roblox = GetProcessByName("RobloxPlayerBeta.exe")

		if (reset == false) {
			reset = true

			client.Logout()
			fmt.Println("reset client activity")
		}
	}

	err := client.Login("823294557155754005")

	if (err != nil) {
		fmt.Println(err)
	}

	reset = false

	args, _ := roblox.Cmdline()

	placePattern := regexp.MustCompile(`placeId=(\d+)`)
	placeMatch := placePattern.FindStringSubmatch(args)[1]

	// timePattern := regexp.MustCompile(`launchtime=(\d+)`)
	// timeMatch := timePattern.FindStringSubmatch(args)[1]

	// startTime, _ := strconv.ParseInt(timeMatch, 10, 64)

	now := time.Now()

	if (placeMatch != placeId) {
		placeId = placeMatch
		place := GetPlaceInfoByPlaceId(placeId)