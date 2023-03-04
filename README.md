## Program is unmaintained. This may or may not work, use at your own risk.

# RoLog.go
RoLog is a program written in Golang for Windows which takes the Roblox experiences you join and posts its experience page to a Discord Webhook. Remember the games you've played before, or tell your community what you're playing right now!

![image](https://user-images.githubusercontent.com/46134462/127735187-d3c6ae1c-7e6d-44d6-bb55-e77a4ead7d1a.png)

# Installation
1. Install [Golang](https://golang.org/dl/) for your system
2. Open a terminal, and run 
```
git clone https://github.com/Derpitron/rolog.go.git
```
3. When the command finishes, run
```
cd rolog.go
```
4. Rename `.env-example` to `.env`, then fill in the `channel-id`/`webhook-id` of your webhook. 

<details>
  <summary>How to find this ID</summary>
    To find this, take your webhook's link, e.g: `https://discord.com/api/webhooks/870936928793534504/IS_NTyJX7Kx7EP3tuJDXdvon8bJLO13QIF9YglKwj-JH39y_4j_yQcG3zFR2wfiAJi-Y`. Then, remove the `https://discord.com/api/webhooks/` part of the link. You are now left with `870936928793534504/IS_NTyJX7Kx7EP3tuJDXdvon8bJLO13QIF9YglKwj-JH39y_4j_yQcG3zFR2wfiAJi-Y`. Put this as the value of the `HOOK` variable.
</details>

5. Once you have made sure of the webhook, there are two commands you can run to compile the program.

- To compile the program with a console window, run:
```
go build main.go
```  
- To compile the program without a console window, run:
```
go build -ldflags -H=windowsgui main.go
```  
Either way, both will compile a file called `main.exe`. 

# Information
  
- You **must** keep the compiled `.exe` in the same folder as your `.env` file, otherwise the program will not run. If you need to move the `.exe` somewhere else, you may make a shortcut and move it, e,g to your `shell:startup` folder.  

- If you want to stop the program, close the console window. If you've compiled the program without the console window and want to stop it, open Task Manager, find the name of your `.exe` file and End Task.

- The program must be running in order to send the link! It must be running on the system in which you're playing the experience, so you cannot host it on a server or service like Heroku!

# Attributions

- [fuzzydragon](https://github.com/fuzzydragon) for a way to get the PlaceID from the running Roblox Process. Go check out their [Roblox Discord Rich Presence](https://github.com/fuzzydragon/Roblox-Game-Presence-For-Discord) Program written in Golang, it works wonderfully!

- [The DisgoOrg team](https://github.com/DisgoOrg) for clearing up my noobish questions with their Disgohook library and helping me out with tidying up the code. They have a lot of cool projects written in Golang, for example [a Discord Webhook Client](https://github.com/DisgoOrg/disgohook), a [Discord API Wrapper](https://github.com/DisgoOrg/disgo), and much more!
