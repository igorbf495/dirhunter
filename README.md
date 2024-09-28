# DirHunter

DirHunter is an open-source directory brute force tool for web servers written in Go. It allows you to discover hidden directories on a web application by providing a target URL and a wordlist for the search.

## Features

- **Multi-threading:** Perform directory searches using multiple threads for improved performance.
- **Command-line parameters:** Simple and easy to use, allowing you to specify the target URL and wordlist directly in the terminal.
- **Supports different HTTP status codes:** Displays HTTP 200 (OK) and 403 (Forbidden) responses, making it easy to identify found and forbidden directories.


## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/dirhunter.git
   cd dirhunter
   ```

2. **Build the project:**

    ```bash
   go build -o dirhunter
   ```

## Usage

   
    ./dirhunter -url <target URL> -wordlist <wordlist file> [-threads <number of threads>]
   

## Parameters

- **url:** Target URL for brute-forcing (e.g., http://example.com).
- **wordlist:** Path to the wordlist file (e.g., wordlist.txt).
- **threads:** Number of threads (optional, default is 10).

## Example

    
    ./dirhunter -url http://arrozdoce.com -wordlist wordlist.txt -threads 20
    

## Contact

- **Author:** Igor Batista
- **Email:** [fel.hacking@gmail.com](mailto:fel.hacking@gmail.com)
- **GitHub:** [github.com/igorbf495](https://github.com/igorbf495)
