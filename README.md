# Scrape-Google

A simple Google search scraper built in Go that extracts search results from Google using `goquery`.

## Features

- Scrapes Google search results from different country domains.
- Uses random user-agents to avoid detection.
- Supports proxy usage for anonymity.
- Configurable number of pages and search results per page.
- Parses and extracts the title, URL, and description from search results.

## Installation

### Prerequisites

- Go 1.16 or later installed. [Download Go](https://go.dev/dl/)
- Internet connection for making HTTP requests.

### Clone the Repository

```sh
git clone https://github.com/yourusername/scrape-google.git
cd scrape-google
