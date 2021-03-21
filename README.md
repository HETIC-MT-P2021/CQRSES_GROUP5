![License](https://img.shields.io/github/license/HETIC-MT-P2021/CQRSES_GROUP5)
![golang](https://img.shields.io/github/languages/top/HETIC-MT-P2021/CQRSES_GROUP5)
![golang-version](https://img.shields.io/github/go-mod/go-version/HETIC-MT-P2021/CQRSES_GROUP5)
![commit](https://img.shields.io/github/last-commit/HETIC-MT-P2021/CQRSES_GROUP5)
![build-CI](https://img.shields.io/github/workflow/status/HETIC-MT-P2021/CQRSES_GROUP5/CI)

## Authors

[Athénais Dussordet](https://github.com/Araknyfe)

[Alexandre Lellouche](https://github.com/AlexandreLch)

[Thomas Raineau](https://github.com/Traineau)

[Corto Dufour](https://github.com/SteakBarbare)

## Usage

```git config core.hooksPath .githooks```
> Configure GitHooks

```cp docker-compose.yaml.dist docker-compose.yaml```
> Docker configuration override, don't forget to add the Token, SQL and RBMQ variables

``` docker-compose up --build```
> Run the project

## Resources

### Order

| Field                 | Type            | Editable | Description                            |
| --------------------- | --------------- | -------- | -------------------------------------- |
| id                    | int             | no       | Order ID                               |
| reference             | string          | yes      | Order reference                        |
| customer              | string          | yes      | Customer name                          |
| time                  | timestamp(UTC)  | no       | Order created on                       |


### Order Lines
| Field                 | Type            | Editable | Description                            |
| --------------------- | --------------- | -------- | -------------------------------------- |
| id                    | int             | no       | Order Line ID                          |
| meal                  | string          | yes      | Type of meal ordered                   |
| quantity              | int             | yes      | Quantity of meal ordered               |
| price                 | int             | no       | Meal price (single unit)               |
| order_id              | timestamp (UTC) | no       | Order ID                               |
