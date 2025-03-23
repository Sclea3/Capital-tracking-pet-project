# CapitalParser

## Project Description

**CapitalParser** is a Go-based service designed to fetch and analyze market data using the HTTP API provided by
Capital.com. The project collects real-time data, computes key market metrics (such as average spread, volatility, and
anomaly detection), and optionally sends notifications via Telegram. It can also store data for historical analysis.

## Features

- **Configuration & Logging:** Load settings from environment files and monitor the application with a configurable
  logger.
- **Data Collection (Collector):** Fetch real-time or historical market data from Capital.com.
- **Analytics (Analyzer):** Compute essential market indicators (e.g., average spread, volatility) and detect anomalies.
- **Data Storage:** Persist raw and processed data in a database (SQLite, PostgreSQL, or any supported DB).
- **Handlers (Optional HTTP/REST):** Provide an interface (e.g., REST endpoints) to interact with collected or analyzed
  data.
- **Notifications (Notifier):** Send alerts about significant market movements or anomalies, potentially via Telegram.
- **Graceful Shutdown:** Properly handle system signals to cleanly shut down the application.

### Folder Overview

- **app/**  
  Contains the primary entry point (`main.go`) that initializes and runs the service.

- **cmd/**  
  Can hold additional command-line tools or separate entry points (if needed).

- **config/**  
  Could include shared configuration files, environment loading, or constants for the entire project.

- **internal/**
    - **analyzer/**: Business logic for analyzing market data (e.g., volatility, spread, anomalies).
    - **collector/**: Responsible for fetching data (real-time or historical) from external APIs.
    - **handlers/**: Contains HTTP handler functions (if you provide a REST API or similar).
    - **config/**: Internal logic for configuration, possibly separate from the global `config/` folder.
    - **notifier/**: Handles notifications, such as sending alerts to Telegram.
    - **storage/**: Database layer, including models and CRUD operations.

- **pkg/**
    - **logger/**: Shared or reusable logging functionality.

- **.env.example**  
  Template for environment variables (API keys, DB credentials, etc.).

- **LICENSE**  
  License file for your project.

- **README.md**  
  This file, describing the project.

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
2. **Navigate to the project directory:**
    ```bash
    cd CapitalParser```
3. **Configure the application:**

- Copy .env.example to .env and update the environment variables with your own API keys, database credentials, etc.
4. **Build and run the application:**
    ```bash
    go build -o capital-parser ./app
5. **Run the application:**
    ```bash
    ./capital-parser
   
## License

This project is licensed under the MIT License. See the LICENSE file for details.