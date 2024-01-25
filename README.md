# couplet
A mobile app that connects people through shared interests in events rather than superficial swipes

![backend-ci](https://github.com/GenerateNU/couplet/actions/workflows/backend-ci.yaml/badge.svg)

## Set Up Your Development Environment
First, understand the tech stack:

- The back end is written in [Go](https://go.dev/) and utilizes [ogen](https://ogen.dev/) to generate API routes from [OpenAPI](https://www.openapis.org/) specification. The back end stores data and serves API consumers, such as the mobile app
- The database is [PostgreSQL](https://www.postgresql.org/), and it simply stores data and gives it to the back end upon request
- The front end is [React Native](https://reactnative.dev/) written with [TypeScript](https://www.typescriptlang.org/) and uses [Expo](https://expo.dev/) as a build tool. Users on iOS and Android will use the mobile app to interact with our core service while the app makes requests to our back end

Before we can compile and run our application, we need to install several languages, package managers, and various tools.
The installation process can vary by tool and operating system, so follow the provided installation instructions for each item below

### Back End
- [Go](https://go.dev/doc/install), our primary backend language
  - Afterwards, install all go dependencies with the command `go mod download` in the `backend/` directory. This needs to be re-run if dependencies change
- [golangci-lint](https://golangci-lint.run/usage/install/#local-installation), a powerful Go linter (required for development)
- [PostgreSQL](https://www.postgresql.org/download/), our SQL database
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/), used to build our back end in isolated containers

### Front End
- [Node](https://nodejs.org/en/learn/getting-started/how-to-install-nodejs), our frontend package manager
  - Afterwards, install all Node dependencies with the command `npm install` in the `frontend/` directory. This needs to be re-run if dependencies change

### General
- [Task](https://taskfile.dev/installation/), a tool for running useful development tasks and a spiritual successor to [Make](https://www.gnu.org/software/make/). Our Taskfiles make it easy for developers to build and run our application quickly and consistently
- [pre-commit](https://pre-commit.com/), a tool for running shared Git hooks on pre-commit (required for development)
  - Afterwards, install all Git hooks with the command `pre-commit install --hook-type commit-msg --hook-type pre-push` in the root directory. This needs to be re-run if hooks change

If everything was successful, you can now compile and run the project!

## Development
Use `task` to check, test, build, and execute the project locally, as targets have already been defined for efficient development. Consider investigating the Taskfiles to learn how everything works!

> [!NOTE]
> This won't work until you have the necessary tools installed on your system

### Back End
The back end has two components, the PostgreSQL database and the Go server. These components can each be run either on your local machine or in a Docker container, and the choice is yours!

> [!NOTE]
> Running locally will provide faster build times and better performance, while Docker will be more consistent/reliable and won't conflict with other installations and processes on your machine. Running the database in Docker and the server locally is a nice middle-ground

#### Layers
1. The API endpoints defined in `openapi.yaml` are exposed to the public, and accessible over HTTP
2. The Handler defines methods for each API endpoint that validate inputs and delegate logic to the Controller
3. The Controller executes business logic, performing operations and manipulating the database. The meat and potatoes of our code will live here

#### Environment Variables
Regardless of how you choose to run them, the two components must be configured to speak to each other, which is accomplished with environment variables. If you have never used them, they are commonly specified with this format `VAR1=value VAR2=value ./executable`. It can be annoying to specify multiple values every time, such as when we configure the server to find and connect to the database, so the root Taskfile and Docker have been configured to read values from an `.env` file in the root directory. The back end should error and let you know if you are missing required variables, but you can also look at `.env.sample` to see what variables you need to define in your environment

#### Run Option 1: Local Database and Local Server
1. Install PostgreSQL and start it (this is system dependent)
2. Initialize the server's credentials by running `createuser ${DB_USER} --password ${DB_PASSWORD}`. Replace those fields with whatever you want, but make sure to add them to your `.env` file
3. Initialize the actual database by running `createdb ${DB_NAME}` and add the name to your `.env` file
4. Define `DB_HOST=localhost` in your `.env` file
5. Define the `DB_PORT` in your `.env` file as the port PostgreSQL is running on (5432 by default)
6. Run the server with `task backend:run`, optionally defining the server port as `PORT` in the `.env` file
7. Access the back end server at `localhost:${PORT}`, or `localhost:8080` if a port was not specified

#### Run Option 2: Docker Database and Local Server
1. Add the database credentials `DB_USER` and `DB_PASSWORD` to your `.env` file. Use whatever values you want, but make sure to add them to your `.env` file
2. Add `DB_NAME=couplet`, `DB_HOST=localhost`, and `DB_PORT=5432` to your `.env` file so the server can connect to the database in Docker
3. Run the database in Docker with `sudo docker-compose up database`
4. Run the server with `task backend:run`, optionally defining the server port as `PORT` in the `.env` file
5. Access the back end server at `localhost:${PORT}`, or `localhost:8080` if a port was not specified

#### Run Option 3: Docker Database and Docker Server
1. Add the database credentials `DB_USER` and `DB_PASSWORD` to your `.env` file. Use whatever values you want, but make sure to add them to your `.env` file
2. Add `DB_NAME=couplet`, `DB_HOST=localhost`, and `DB_PORT=5432` to your `.env` file so the server can connect to the database in Docker
3. Optionally define the server port as `PORT` in the `.env` file
4. Run the database and the server in Docker with `sudo docker-compose up --build`
5. Access the back end server at `localhost:${PORT}`, or `localhost:8080` if a port was not specified
Bonus: If running the back end server in Docker, you can view the API documentation with a simple UI at `localhost:80`

### Front End
1. Download the Expo Go app on your phone from either the App Store or the Google Play store
2. From the `frontend/` directory, run `npx expo start`
3. Scan the QR code on your phone and wait for the app to load
4. If you want to run the app on your computer, you will need to make sure you spin up the relevant emulator. This is either an Android Studio emulator if you want to run on Android, or an XCode simulator if you want to run on iOS
5. To run on android, press a. To run on iOS, press i

## Contributing
- Nobody is allowed to push to `main`. Open a new branch, push to it, and open a pull request to get your changes into `main`
- Your code must pass formatter and linter checks to make valid commits
- You must write *useful* tests for your pull requests to be accepted
- Your code must pass GitHub Actions checks (formatters, linters, and tests) for your pull requests to be accepted
- At least one TL must review and accept your PR before merging
- Read other developers' pull requests and give feedback! This is how we all improve and build a great product
- Be kind in code critique. Describe issues as properties of the code, not the developer (Good: "the **code** does not solve the issue", bad: "**you** failed to solve the problem)
- Give actionable feedback. Everyone wants to improve, but that requires advice on how to improve (and often an explanation of why improvement is necessary)
