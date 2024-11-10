# Facebook UID Extraction Bot in Go

This project is a **Facebook UID Extraction Bot** developed in **Go**. It uses Selenium WebDriver to automate data extraction from Facebook profiles and exports the profile information into a CSV file. With Go’s concurrency capabilities, the bot can process thousands of profiles efficiently, making it an excellent tool for anyone learning about web scraping, automation, and data management in Go.

> **Note:** This bot is intended for educational purposes only. It will not work for Meta Custom Audiences due to Facebook’s API and usage policies.

## Table of Contents

- [Features](#features)
- [How It Works](#how-it-works)
- [Requirements](#requirements)
- [Setup and Installation](#setup-and-installation)
- [Running the Bot](#running-the-bot)
- [Project Structure](#project-structure)
- [Code Walkthrough](#code-walkthrough)
- [Limitations](#limitations)
- [Future Improvements](#future-improvements)
- [License](#license)

---

## Features

- **High-Volume UID Extraction**: The bot can handle and process thousands of profile links simultaneously.
- **CSV Export**: Extracted data is automatically organized into a CSV file.
- **Selenium Integration**: Uses Selenium WebDriver for automated web interaction with Facebook pages.
- **Concurrency**: Built with Go’s concurrency model to increase efficiency.

## How It Works

1. **Initialize WebDriver**: Starts a Chrome WebDriver using Selenium.
2. **Open Profile Links**: The bot navigates to each Facebook profile link in the provided list.
3. **Extract UID**: Scrapes the user ID from each profile using the unique `meta` tag (`al:android:url`).
4. **Store Results**: Writes the extracted profile link and UID to a CSV file for easy access and analysis.

## Requirements

- **Go** (version 1.15+ recommended)
- **Selenium** (latest stable release)
- **ChromeDriver** (matching version to your Chrome browser)
- **Google Chrome** (latest version)

## Setup and Installation

1. **Install Go**: Ensure Go is installed and properly set up in your PATH. You can download it [here](https://golang.org/dl/).
2. **Install Selenium and ChromeDriver**:
   - Download ChromeDriver from [here](https://chromedriver.chromium.org/downloads), and place it in an accessible directory.
   - Make sure to adjust the path to `chromedriver.exe` in the code as necessary.
3. **Clone the Repository**:
   ```bash
   git clone https://github.com/Kaiftech/Facebook-UID-Extractor-Bot-in-Go.git
   cd Facebook-UID-Extractor-Bot-in-Go
   ```
4. **Install Required Packages**:
   ```bash
   go get -u github.com/tebeka/selenium
   ```

## Running the Bot

1. **Add Facebook Profile Links**: Update the `profileLinks` array in `main.go` with the profile links you wish to scrape.
2. **Start the Bot**:
   ```bash
   go run main.go
   ```
3. **View Output**: The output file `facebook_profiles.csv` will contain profile links and corresponding UIDs.

## Project Structure

```
├── main.go             # Main bot logic
└── facebook_profiles.csv  # Output CSV file with extracted UIDs
```

## Code Walkthrough

### main.go

1. **Bot Configuration**:
   - Sets browser options to disable specific features (e.g., infobars) to avoid bot detection.
   
2. **Initialize WebDriver**:
   - Launches Selenium Chrome WebDriver at the specified port and configures Chrome options for automation.

3. **CSV File Creation**:
   - Creates a CSV file to store extracted data with headers (`Profile Link` and `Profile UID`).

4. **Data Extraction**:
   - Visits each Facebook profile link, extracts the UID from the `al:android:url` meta tag, and writes it to the CSV file.

5. **Helper Functions**:
   - `scrapeProfileUID`: Opens a profile link and extracts the UID from a specific meta tag.
   - `waitForElement`: Waits for a particular HTML element to load, ensuring data extraction proceeds smoothly.

### Sample Code Snippets

#### Initializing WebDriver
```go
caps := selenium.Capabilities{"browserName": "chrome"}
chromeArgs := []string{
    "--disable-blink-features=AutomationControlled",
    "--disable-infobars",
    "--start-maximized",
}
caps["goog:chromeOptions"] = map[string]interface{}{"args": chromeArgs}
```

#### Extracting UID
```go
metaTag, err := wd.FindElement(selenium.ByCSSSelector, "meta[property='al:android:url']")
content, err := metaTag.GetAttribute("content")
var uid string
_, err = fmt.Sscanf(content, "fb://profile/%s", &uid)
```

#### Writing to CSV
```go
record := []string{profileLink, uid}
if err := writer.Write(record); err != nil {
    log.Printf("Could not write record to CSV: %v", err)
}
```

## Limitations

- **Facebook’s Policy Restrictions**: Due to Facebook’s restrictions, this bot is only intended for educational and personal use.
- **Custom Audience Limitations**: This bot will not work for Meta Custom Audiences.
- **IP Blocking**: Repeated automated requests may result in IP blocking by Facebook.

## Future Improvements

- **Error Handling**: Adding more sophisticated error handling and retry mechanisms.
- **Multi-Browser Support**: Extending support to other browsers like Firefox.
- **Enhanced Data Extraction**: Expanding the bot to extract additional profile data.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details. 

---

Enjoy experimenting with Go and Selenium, and happy coding!
