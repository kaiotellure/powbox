package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/chromedp/chromedp"
)

func main() {
    // Create a context
    ctx, cancel := chromedp.NewExecAllocator(
        context.Background(),
        chromedp.NoDefaultBrowserCheck, // Avoid default checks
        chromedp.Flag("headless", false), // Non-headless mode (visible browser)
        chromedp.Flag("disable-gpu", false),
    )
    defer cancel()

    // Create a new browser context
    ctx, cancel = chromedp.NewContext(ctx)
    defer cancel()

    // URL to scrape
    url := "https://google.com/search?q=inurl:netflix" // Replace with your target URL

    // Variable to store the anchor links
    var links []string

    // Run the browser tasks
    err := chromedp.Run(ctx,
        // Navigate to the URL
        chromedp.Navigate(url),
        // Wait until the page is fully loaded (adjust selector as needed)
        chromedp.WaitVisible("body", chromedp.ByQuery),
        // Extract all <a> tags' href attributes
        chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(a => a.href);`, &links),
    )
    if err != nil {
        log.Fatalf("Failed to run chromedp: %v", err)
    }

    // Print the collected links
    fmt.Println("Found anchor links:")
    for i, link := range links {
        fmt.Printf("%d: %s\n", i+1, link)
    }

    // Keep the browser open for a few seconds to observe (optional)
    time.Sleep(5 * time.Second)
}