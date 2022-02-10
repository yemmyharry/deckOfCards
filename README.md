<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="https://images.unsplash.com/photo-1623522264952-8dff960ec8f2?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8N3x8YWNlJTIwb2YlMjBzcGFkZXN8ZW58MHx8MHx8&auto=format&fit=crop&w=500&q=60" alt="Logo" width="80" height="80">
  </a>
<h3 align="center">Deck Of Cards API</h3>
  </p>
</div>

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#testing">Testing</a></li>
  </ol>
</details>


## About The Project
An API to create a deck of cards, open a deck of cards, draw one or multiple cards.

## Getting Started

### Prerequisites

Before this project can be run locally you will need to install [Go](https://golang.org/doc/install)

### Installation

To utilize the project, the following needs to be done:
1. Clone this repository
2. Install the dependencies, using the following command:
```
go mod tidy
```

## Usage

1. To run the project locally, use the following command:
```
go run main.go
```
2. To create a deck of cards, use Postman to make a request to the following URL:
```
http://localhost:3000/api/create
```
3. To open a deck of cards, use Postman to make a request to the following URL:
```
http://localhost:3000/api/open/{deck_id}
```
4. To draw a card or multiple cards, use Postman to make a request to the following URL:
```
http://localhost:3000/api/draw/{deck_id}?count={any number}
```

## Testing
Tests can be run using the following command:
```
go test ./...
```

