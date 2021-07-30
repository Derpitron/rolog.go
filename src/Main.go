package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/DisgoOrg/disgohook"
	"github.com/joho/godotenv"
	"github.com/shirou/gopsutil/v3/process"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

var (
	hook, _ = disgohook.NewWebhookClientByToken(nil, nil, goDotEnvVariable("HOOK"))
	placeId string
	reset   = false
)

func GetProcessByName(targetProcessName string) *process.Process {
	processes, _ := process.Processes()

	for _, proc := range processes {
		name, _ := proc.Name()

		if name == targetProcessName {
			return proc
		}
	}

	return nil
}

func UpdateRobloxPresence() {
	roblox := GetProcessByName("RobloxPlayerBeta.exe")

	for roblox == nil {
		roblox = GetProcessByName("RobloxPlayerBeta.exe")

		if reset == false {
			reset = true

			fmt.Println("reset client activity")
		}
	}

	reset = false

	args, _ := roblox.Cmdline()

	placePattern := regexp.MustCompile(`placeId=(\d+)`)
	placeMatch := placePattern.FindStringSubmatch(args)[1]

	if placeMatch != placeId {
		placeId = placeMatch

		_, _ = hook.SendContent("`Started playing:`\nhttps://www.roblox.com/games/" + placeId + "/-")

		fmt.Println("set activity: " + placeId)
	}
}

func main() {
	for true {
		UpdateRobloxPresence()

		time.Sleep(time.Second * 5)
	}
}
