# Shopping List App

The Shopping List App is a simple yet efficient application written in Go that allows users to manage their shopping lists using GraphQL. It provides a clean API to perform CRUD operations on shopping list items.

## Table of Contents
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [API](#api)
  - [Queries](#queries)
  - [Mutations](#mutations)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or higher
- Git

### Installation

1. Clone the repository:
```sh
git clone https://github.com/your_username/shopping-list-app.git
```

2. Navigate into the project directory:
```sh
cd shopping-list-app
```

3. Install the dependencies:
```sh
go get -v -t -d ./...
```

4. Run the application:
```sh
go run main.go
```

The server will start, and you can send GraphQL requests to `http://localhost:8080/shoppinglist`.

## Usage

Send GraphQL requests to `http://localhost:8080/shoppinglist` to interact with the shopping list.

## API

### Queries

- `shoppingItems`: Fetch all shopping items.
- `shoppingItem(id: Int)`: Fetch a single shopping item by its ID.

### Mutations

- `addShoppingItem(itemName: String, description: String, count: Int)`: Add a new item to the shopping list.
- `updateShoppingItem(id: Int, itemName: String, description: String, count: Int, purchased: Boolean)`: Update an existing item.
- `deleteShoppingItem(id: Int)`: Remove an item from the shopping list by its ID.

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

dan.smiledev@email.com

[https://github.com/HugeSmileDev/shopping-list-app](https://github.com/HugeSmileDev/ShoppingList-Backend)
