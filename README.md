# Go REST Project (go-rest-project)

Sample Go web application to showcase REST API development.

## Features & Technologies

TODO

## Architecture

TODO

## API Schema

| Method   | Resource                   | Security              | Description                       |
|:--------:|:---------------------------|:---------------------:|:----------------------------------|
| `GET`    | `/events`                  | Public                | Get all events                    |
| `GET`    | `/events/{id}`             | Public                | Get event by ID                   |
| `POST`   | `/events`                  | Authenticated         | Create a new event                |
| `PUT`    | `/events/{id}`             | Authenticated & Owner | Update an event                   |
| `DELETE` | `/events/{id}`             | Authenticated & Owner | Delete an event                   |
| `POST`   | `/events/{id}/register`    | Authenticated         | Register user for an event        |
| `DELETE` | `/events/{id}/register`    | Authenticated         | Unregister user from an event     |
| `POST`   | `/signup`                  | Public                | Sign up a new user                |
| `POST`   | `/login`                   | Public                | Authenticate a user               |

## License

This project is licensed under the [MIT License](LICENSE).
