[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">mc-bedrock-runner</h3>

  <p align="center">
  Minecraft Bedrock server process wrapper with attached RCON server
    <br />
    <a href="https://github.com/saulmaldonado/mc-bedrock-runner"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/saulmaldonado/mc-bedrock-runner/tree/main/example/mc-server.yml">View Example</a>
    ·
    <a href="https://github.com/saulmaldonado/mc-bedrock-runner/issues">Report Bug</a>
    ·
    <a href="https://github.com/saulmaldonado/mc-bedrock-runner/issues">Request Feature</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
    <li><a href="#author">Author</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

A Linux Bedrock server process wrapper that adds an RCON server for remote command execution and adds SIGTERM/SIGINT handling for graceful server termination.

### Built With

- [go-mc](https://github.com/Tnze/go-mc/net)

<!-- GETTING STARTED -->

## Getting Started

### Prerequisites

- Minecraft Bedrock server software

  - Linux server are only supported at this time
  - [Bedrock server software](https://www.minecraft.net/en-us/download/server/bedrock)

### Installation

#### Download and extract mc-bedrock-runner binary

Download latest version

[mc-bedrock-runner.tar.gz](https://github.com/saulmaldonado/mc-bedrock-runner/releases/download/latest/mc-bedrock-runner.tar.gz)

Extract

```sh
tar -xzf mc-bedrock-runner.tar.gz
```

<!-- USAGE EXAMPLES -->

## Usage

Server wrapper takes two optional flags and a required argument, the path to the Bedrock server executable (file is usually `bedrock_server`)

### Run Locally In Command Line

```
mc-bedrock-runner [flags] ./bedrock_server

  --password string
        RCON authentication password (default "minecraft")

  --port uint
        RCON server port (default 25575)
```

```sh
mc-bedrock-runner --port 25575 --password minecraft ./bedrock_server
```

### Run Locally With Docker

```sh
docker run -it --name mc -p 19132:19132/udp -p 25575:25575 saulmaldonado/mc-bedrock-runner --port 25575 --password minecraft bedrock_server
```

## License

Distributed under the MIT License. See [LICENSE](./LICENSE) for more information.

<!-- ACKNOWLEDGEMENTS -->

## Acknowledgements

- [Tnze/go-mc](https://github.com/Tnze/go-mc/net)

## Author

### Saul Maldonado

- 🐱 Github: [@saulmaldonado](https://github.com/saulmaldonado)
- 🤝 LinkedIn: [@saulmaldonado4](https://www.linkedin.com/in/saulmaldonado4/)
- 🐦 Twitter: [@saul_mal](https://twitter.com/saul_mal)
- 💻 Website: [saulmaldonado.com](https://saulmaldonado.com/)

## Show your support

Give a ⭐️ if this project helped you!

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/saulmaldonado/mc-bedrock-runner.svg?style=for-the-badge
[contributors-url]: https://github.com/saulmaldonado/mc-bedrock-runner/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/saulmaldonado/mc-bedrock-runner.svg?style=for-the-badge
[forks-url]: https://github.com/saulmaldonado/mc-bedrock-runner/network/members
[stars-shield]: https://img.shields.io/github/stars/saulmaldonado/mc-bedrock-runner.svg?style=for-the-badge
[stars-url]: https://github.com/saulmaldonado/mc-bedrock-runner/stargazers
[issues-shield]: https://img.shields.io/github/issues/saulmaldonado/mc-bedrock-runner.svg?style=for-the-badge
[issues-url]: https://github.com/saulmaldonado/mc-bedrock-runner/issues
[license-shield]: https://img.shields.io/github/license/saulmaldonado/mc-bedrock-runner.svg?style=for-the-badge
[license-url]: https://github.com/saulmaldonado/mc-bedrock-runner/blob/master/LICENSE.txt
