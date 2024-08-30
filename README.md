# Cat API Go Project

This project is a Go-based backend service that interacts with The Cat API to provide various cat-related functionalities. It uses the Beego web framework and offers endpoints for fetching random cat images, breed information, favoriting images, and voting on cats.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Configuration](#configuration)
4. [Running the Project](#running-the-project)
5. [API Endpoints](#api-endpoints)
6. [Project Structure](#project-structure)
7. [Contributing](#contributing)
8. [License](#license)

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Go (version 1.16 or later)
- Git

## Installation

### 1. Install Go

If you haven't installed Go, follow these steps:

1. Visit the official Go downloads page: https://golang.org/dl/
2. Download the appropriate installer for your operating system.
3. Follow the installation instructions for your OS:
   - Windows: Run the MSI installer and follow the prompts.
   - macOS: Open the package file and follow the prompts.
   - Linux: Extract the archive to `/usr/local`:
     ```
     tar -C /usr/local -xzf go1.x.x.linux-amd64.tar.gz
     ```
4. Add Go to your PATH:
   - For bash, add the following to your `~/.bashrc` or `~/.bash_profile`:
     ```
     export PATH=$PATH:/usr/local/go/bin
     export GOPATH=$HOME/go
     export PATH=$PATH:$GOPATH/bin
     ```
   - For other shells, add the equivalent to your shell's configuration file.
5. Verify the installation by opening a new terminal and running:
   ```
   go version
   ```

### 2. Clone the Repository

Clone this repository to your local machine:

```
git clone https://github.com/Speak2/beego_backend
cd cat-api-go-project
```

### 3. Install Dependencies

This project uses Beego v2 and other dependencies. Install them using the following commands:

```
go get github.com/beego/beego/v2
go get github.com/astaxie/beego/logs
go mod tidy
```

## Configuration

### 1. API Key

This project requires an API key from The Cat API. To obtain one:

1. Visit https://thecatapi.com/
2. Sign up for an account
3. Generate an API key from your dashboard

### 2. Configuration File

Create a `conf/app.conf` file in the project root with the following content:

```ini
appname = cats_backend
httpport = 8080
runmode = dev
cat_api_key = YOUR_CAT_API_KEY
```

Replace `YOUR_CAT_API_KEY` with the API key you obtained from The Cat API.

## Running the Project

To run the project, use the following command from the project root directory:

```
bee run
```

If you don't have `bee` installed, you can install it with:

```
go get github.com/beego/bee/v2
```

Alternatively, you can run the project directly with:

```
go run main.go
```

The server will start, and you should see output indicating that it's running on `http://localhost:8080`.

## API Endpoints

The project provides the following API endpoints:

- `GET /api/random-cat`: Fetch a random cat image
- `GET /api/breeds`: Get a list of all cat breeds
- `GET /api/breed-images`: Get images for a specific breed (requires `breed_id` query parameter)
- `POST /api/favorites`: Add a cat image to favorites
- `POST /api/votes`: Vote on a cat image
- `GET /api/get_favorites`: Get a list of favorite cat images
- `DELETE /api/delete_favorite/:favoriteId`: Delete a favorite cat image

For detailed usage of these endpoints, refer to the controller files in the `controllers` directory.

## Project Structure

The project follows a standard Beego directory structure:

```
cat-api-go-project/
├── conf/
│   └── app.conf
├── controllers/
│   ├── breeds_controller.go
│   ├── favorites_controller.go
│   ├── random_cat.go
│   └── voting_controller.go
├── routers/
│   └── router.go
├── main.go
├── go.mod
├── go.sum
└── README.md
```

- `conf/`: Contains configuration files
- `controllers/`: Contains the logic for handling API requests
- `routers/`: Defines the routing for the application
- `main.go`: The entry point of the application

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[Specify your license here]

