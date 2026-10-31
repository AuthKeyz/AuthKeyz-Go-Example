package main

import (
	AuthKeyzApp "AuthKeyz/AuthKeyz"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Input(message string) string {
	fmt.Print(message)

	var input string
	fmt.Scanln(&input)
	return input
}

func UnixToReadable(unixTimestamp string) string {
	timestamp, err := strconv.ParseInt(unixTimestamp, 10, 64)
	if err != nil {
		return unixTimestamp // Return original if parsing fails
	}
	t := time.Unix(timestamp, 0).UTC()
	return t.Format("2006-01-02 15:04:05")
}

func printUserData() {
	fmt.Println("\nUser Data:")
	fmt.Println("   Username: ", AuthKeyzApp.Username)
	fmt.Println("   IP Address: ", AuthKeyzApp.IP)
	fmt.Println("   HWID: ", AuthKeyzApp.HWID)
	fmt.Println("   Created At: ", UnixToReadable(AuthKeyzApp.CreatedDate))
	fmt.Println("   Last Login At: ", UnixToReadable(AuthKeyzApp.LastLogin))
	fmt.Println("   Subscription Expiry: ", UnixToReadable(AuthKeyzApp.Expires))
	fmt.Println("   Subscription: ", AuthKeyzApp.Subscription)
}

func showMenu() {
	fmt.Println("[1] Login")
	fmt.Println("[2] Register")
	fmt.Println("[3] Upgrade")
	fmt.Println("[4] License Only Login")
}

func main() {
	AuthKeyzApp.Api(
		"",    // -- Application Name
		"", // -- Owner ID
		"1.0",        // -- Application Version
		"",           // -- Token Path (PUT NULL OR LEAVE BLANK IF YOU DON'T WANT TO USE TOKEN SYSTEM)
	)

	done := false
	for !done {
		if AuthKeyzApp.LockoutActive() {
			fmt.Printf("Locked out. Try again in %d ms\n", AuthKeyzApp.LockoutRemainingMs())
			AuthKeyzApp.CloseDelay()
			os.Exit(1)
		}

		showMenu()
		ans := Input("\nChoose your option: ")

		switch ans {
		case "1":
			username := Input("Input username: ")
			password := Input("Input password: ")
			AuthKeyzApp.Login(username, password)
			printUserData()
			AuthKeyzApp.ResetLockout()
			done = true
		case "2":
			username := Input("Input username: ")
			password := Input("Input password: ")
			license := Input("Input license: ")
			AuthKeyzApp.Register(username, password, license)
			printUserData()
			AuthKeyzApp.ResetLockout()
			done = true
		case "3":
			username := Input("Input username: ")
			license := Input("Input license: ")
			AuthKeyzApp.Upgrade(username, license)
			printUserData()
			AuthKeyzApp.ResetLockout()
			done = true
		case "4":
			license := Input("Input license: ")
			AuthKeyzApp.License(license)
			printUserData()
			AuthKeyzApp.ResetLockout()
			done = true
		default:
			fmt.Println("Invalid option")
			AuthKeyzApp.BadInputDelay()
		}
	}

	fmt.Println("\nExiting application in 10 seconds...")
	time.Sleep(10 * time.Second)
	os.Exit(0)
}
