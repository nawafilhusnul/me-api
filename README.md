## Project Description

This is a backend project that consists of services for content management on my personal portfolio website and for the portfolio website itself.

### Folder Structure
```bash
- src/
|   |- project/
|   |   |- delivery/
|   |   |   |- http/
|   |   |   |   |- project.go
|   |   |- usecase/
|   |   |   |- project.go
|   |   |   |- project_test.go
|   |   |- repository/
|   |   |   |- firestore
|   |   |   |   |- project.go
|   |   |   |   |- helper.go
|   |- domain/
|   |   |- project.go
|   |   |- response.go
|   |   |- error.go
|   |   |- mocks/
- docs/
|   |- public/
|   |   |- swagger.json
|   |   |- swagger.yaml
|   |   |- index.html
|   |   |- docs.go
|   |   |- 404.html
- global/
|   |- helper.go
- main.go
- go.mod
- README.md
```

### Features (To be completed soon)

- Create, Read, Update, Delete operations on Project entity.

### Installation (To be completed soon)

To install this project locally, follow these steps:

1. Have you golang version >= go1.21.1. You may check by this command, otherwise download your golang first.
```shell
go version
```

2. Clone this project, or you may fork this project also if you want to.
```
git clone https://github.com/nawafilhusnul/me-api.git
```

3. Change directory to project.
```shell
cd me-api
```

4. Install all dependencies.
```shell
go mod install
```

5. Ensures that the `go.mod` file matches the source code in the module.
```shell
go mod tidy
```

6. Run the project by simply execute this command.
```shell
go run main.go
```

7. Or if you have [air](https://github.com/cosmtrek/air) installed.

```shell
air
```

### Demo

To use the demo, follow these steps:

1. Visit the demo website on [https://me-dashboard-vk4z.onrender.com](https://me-dashboard-vk4z.onrender.com). Thanks for [render.com](https://render.com) for free hosting.
2. Find the API docs by hitting [https://me-dashboard-vk4z.onrender.com/docs](https://me-dashboard-vk4z.onrender.com/docs).
3. Feel free explore the various features and functionalities of the project.
