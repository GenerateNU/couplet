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
  - Afterwards, install all node dependencies with the command `npm install` in the `frontend/` directory. This needs to be re-run if dependencies change

### General
- [Task](https://taskfile.dev/installation/), a tool for running useful development tasks. A spiritual successor to [Make](https://www.gnu.org/software/make/). Our Taskfiles make it easy for developers to build and run our application quickly and consistently
- [pre-commit](https://pre-commit.com/), a tool for running shared Git hooks on pre-commit (required for development)
  - Afterwards, install all Git hooks with the command `pre-commit install --hook-type commit-msg --hook-type pre-push` in the root directory. This needs to be re-run if hooks change

If everything was successful, you can now compile and run the project!

## Running the Project
Use `task` to build and execute the project, as targets have already been defined for efficient development. Consider investigating the Taskfiles to learn how everything works!

> [!NOTE]
> This won't work until you have the necessary tools installed on your system

### Back End
1. Create a `.env` file in the root directory that defines values `DB_USER` and `DB_PASSWORD`, and optionally `PORT` (defaults to 8080)
2. Build and run the back end, either locally or in Docker container. The database must be running for the back end to run
   - Local database and local back end: Start PostgreSQL on your machine and initialize it with `createdb couplet` and `createuser ${DB_USER} --password ${DB_PASSWORD}`, then run `task backend:run`
   - Docker database and local back end: Run `sudo docker-compose up database` and `task backend:run`
   - Docker database and Docker back end: Run `sudo docker-compose up --build`
3. Access the back end at `localhost:${PORT}`, or `localhost:8080` if a port was not specified

Running locally will provide faster build times and slightly better performance, while Docker will be more consistent/reliable and won't conflict with other installations and processes on your machine. The choice is yours!

Bonus: If running the back end in Docker, you can view the API documentation with a nice UI at `localhost:80`

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
