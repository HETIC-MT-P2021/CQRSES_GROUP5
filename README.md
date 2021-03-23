![License](https://img.shields.io/github/license/HETIC-MT-P2021/CQRSES_GROUP5)
![golang](https://img.shields.io/github/languages/top/HETIC-MT-P2021/CQRSES_GROUP5)
![golang-version](https://img.shields.io/github/go-mod/go-version/HETIC-MT-P2021/CQRSES_GROUP5)
![commit](https://img.shields.io/github/last-commit/HETIC-MT-P2021/CQRSES_GROUP5)
![build-CI](https://img.shields.io/github/workflow/status/HETIC-MT-P2021/CQRSES_GROUP5/CI)

GOQRS
===============
This school project is an API with no GUI, using the **CQRS Pattern** 
and an **EventSourcing** implementation to handle Authentification and manage
Orders and OrderLines, for a restaurant or a coffee for exemple.

We used **Docker, Compose, RabbitMQ and Go language** in order to make a scalable
and fast application, allowing evolutions.  

Usage
===============

`git config core.hooksPath .githooks`

> Configure GitHooks

`cp docker-compose.yaml.dist docker-compose.yaml`

> Docker configuration override, don't forget to add the Token, SQL and RBMQ variables

` docker-compose up --build`

> Run the project

Consumer repo
===============

https://github.com/HETIC-MT-P2021/CQRSES_GROUP5_CONSUMER

Resources
===============

### Order

| Field     | Type           | Editable | Description      |
| --------- | -------------- | -------- | ---------------- |
| id        | int            | no       | Order ID         |
| reference | string         | yes      | Order reference  |
| customer  | string         | yes      | Customer name    |
| time      | timestamp(UTC) | no       | Order created on |

### Order Lines


| Field    | Type            | Editable | Description              |
| -------- | --------------- | -------- | ------------------------ |
| id       | int             | no       | Order Line ID            |
| meal     | string          | yes      | Type of meal ordered     |
| quantity | int             | yes      | Quantity of meal ordered |
| price    | int             | no       | Meal price (single unit) |
| order_id | timestamp (UTC) | no       | Order ID                 |

Doc
===============

You can find in the [doc](https://github.com/HETIC-MT-P2021/CQRSES_GROUP5/tree/develop/doc) folder
different files allowing everyone to understand how this app is structured and help you to use it:
* [The global architecture](https://github.com/HETIC-MT-P2021/CQRSES_GROUP5/blob/develop/doc/Architecture.png)
* [The entities used](https://github.com/HETIC-MT-P2021/CQRSES_GROUP5/blob/develop/doc/Entities.png)
* [Functional Specs](https://github.com/HETIC-MT-P2021/CQRSES_GROUP5/blob/develop/doc/Functional%20Specs.docx)
* [Technical Specs](https://github.com/HETIC-MT-P2021/CQRSES_GROUP5/blob/develop/doc/Technical%20Architecture.docx)
* [A Postman Collection](https://github.com/HETIC-MT-P2021/CQRSES_GROUP5/blob/develop/doc/goqrs.postman_collection.json) 

### Generated Doc

`golds ./...`

> Start a local doc server

Or

`golds -gen -dir=generated -nouses ./...`

> Generate static HTML doc pages

`golds -dir=generated`

> View the generated doc


Authors
===============

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/Araknyfe">
        <img src="https://github.com/Araknyfe.png" width="150px;"/><br>
        <b>Athénaïs Dussordet</b>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/AlexandreLch">
        <img src="https://github.com/AlexandreLch.png" width="150px;"/><br>
        <b>Alexandre Lellouche</b>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/Traineau">
        <img src="https://github.com/Traineau.png" width="150px;"/><br>
        <b>Thomas Raineau</b>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/SteakBarbare">
        <img src="https://github.com/SteakBarbare.png" width="150px;"/><br>
        <b>Corto Dufour</b>
      </a>
    </td>
  </tr>
</table>