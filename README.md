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

## Requirements

If you use docker you will only need:

- Docker;
- Docker-Compose;

Refer to [Docker-Setup](#docker-setup) to install with docker.

To run this project, you will also need to install the following dependencies on your system:

- [go](https://golang.org/doc/install)

## Usage

```git config core.hooksPath .githooks```
> Configure GitHooks

```cp docker-compose.yaml.dist docker-compose.yaml```
> Docker configuration override, don't forget to add the Token and SQL variables

``` docker-compose up --build```
> Run the project

## Contributing

- Your branch should have a name that reflects it's purpose.

- Each commit must follow the [Commit Conventions](https://github.com/HETIC-MT-P2021/CQRSES_GROUP5/blob/develop/CONTRIBUTING.md)

## Resources
