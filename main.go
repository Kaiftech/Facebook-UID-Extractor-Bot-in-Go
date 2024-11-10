package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

type BotConfig struct {
	Browser string
}

var profileLinks = []string{
	"https://www.facebook.com/zuck/",
}

func main() {
	// Start the terminal interface
	terminalInterface()
}

func terminalInterface() {
	config := BotConfig{Browser: "Chrome"}
	launchBot(config)
}

func launchBot(config BotConfig) {
	fmt.Println("Launching bot with the following configuration:")
	fmt.Printf("Browser: %s\n", config.Browser)

	// Set up Selenium WebDriver
	seleniumPath := `` // Path to ChromeDriver
	port := 8081

	svc, err := selenium.NewChromeDriverService(seleniumPath, port)
	if err != nil {
		log.Fatalf("Error starting the ChromeDriver server: %v", err)
	}
	defer svc.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeArgs := []string{
		"--disable-blink-features=AutomationControlled",
		"--disable-infobars",
		"--start-maximized",
	}
	caps["goog:chromeOptions"] = map[string]interface{}{
		"args": chromeArgs,
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Error connecting to the WebDriver server: %v", err)
	}
	defer wd.Quit()

	// Create CSV file
	file, err := os.Create("facebook_profiles.csv")
	if err != nil {
		log.Fatalf("Could not create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	headers := []string{"Profile Link", "Profile UID"}
	if err := writer.Write(headers); err != nil {
		log.Fatalf("Could not write headers to CSV: %v", err)
	}

	// Iterate over each profile link and scrape data
	for _, profileLink := range profileLinks {
		uid, err := scrapeProfileUID(wd, profileLink)
		if err != nil {
			log.Printf("Failed to scrape profile UID for %s: %v", profileLink, err)
			continue
		}

		// Write data to CSV
		record := []string{profileLink, uid}
		if err := writer.Write(record); err != nil {
			log.Printf("Could not write record to CSV: %v", err)
		}
	}

	fmt.Println("Bot completed its tasks successfully!")
}

func scrapeProfileUID(wd selenium.WebDriver, profileLink string) (string, error) {
	fmt.Printf("Loading profile: %s\n", profileLink)

	// Open the Facebook profile page
	if err := wd.Get(profileLink); err != nil {
		return "", fmt.Errorf("could not navigate to profile: %v", err)
	}

	// Wait for the page to load
	err := waitForElement(wd, selenium.ByCSSSelector, "meta[property='al:android:url']", 2*time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to find profile UID meta tag: %v", err)
	}

	// Get the UID from the meta tag
	metaTag, err := wd.FindElement(selenium.ByCSSSelector, "meta[property='al:android:url']")
	if err != nil {
		return "", fmt.Errorf("could not find profile UID meta tag: %v", err)
	}

	content, err := metaTag.GetAttribute("content")
	if err != nil {
		return "", fmt.Errorf("could not get content attribute: %v", err)
	}

	// Extract UID from content
	var uid string
	_, err = fmt.Sscanf(content, "fb://profile/%s", &uid)
	if err != nil {
		return "", fmt.Errorf("could not extract UID from content: %v", err)
	}

	return uid, nil
}

func waitForElement(wd selenium.WebDriver, by, value string, timeout time.Duration) error {
	end := time.Now().Add(timeout)
	for time.Now().Before(end) {
		_, err := wd.FindElement(by, value)
		if err == nil {
			return nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return fmt.Errorf("element %s not found after %v", value, timeout)
}
