[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/aydinnyunus/wallet-tracker">
  </a>

<h3 align="center">Wallet Tracker CLI</h3>

  <p align="center">
    <br />
    <a href="https://github.com/aydinnyunus/wallet-tracker"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    ·
    <a href="https://github.com/aydinnyunus/wallet-tracker/issues">Report Bug</a>
    ·
    <a href="https://github.com/aydinnyunus/wallet-tracker/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
         <ul>
            <li><a href="#track-wallet">Track Wallet</a></li>
            <li><a href="#track-wallet-with-network">Track Wallet with Network</a></li>
            <li><a href="#detect-exchanges-on-exit-nodes">Detect Exchanges on Exit Nodes</a></li>
           <li><a href="#start-neodash">Start Neodash</a></li>
           <li><a href="#get-exchange-wallet">Get Exchange Wallet</a></li>
         </ul>
   </li>
    <li><a href="#downloads">Downloads</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>


<!-- GETTING STARTED -->

## Getting Started

General information about setting up Wallet Tracker CLI locally


## Installation

```bash
go get github.com/aydinnyunus/wallet-tracker
```

<!-- USAGE EXAMPLES -->

## Usage

### Track Wallet

After you have install requirements , you can simply track the wallet via:

```shell
   $ wallet-tracker tracker track --wallet 37oTUqiViE3YySs8xxAtKgTzQgoVuSVbse
```

### Track Wallet with Network

If you want to specify network ( you don't need that for now ) use this command:

```shell
   $ wallet-tracker tracker track --wallet 37oTUqiViE3YySs8xxAtKgTzQgoVuSVbse --network BTC
```

### Detect Exchanges on Exit Nodes

If you want to Detect Exchanges on Exit Nodes use this command:

```shell
   $ wallet-tracker tracker track --wallet 37oTUqiViE3YySs8xxAtKgTzQgoVuSVbse --detect-exchanges
```

### Start Neodash

If you want to visualize Wallets and Transactions using Neo4J database use this command:

```shell
   $ wallet-tracker neodash start
```

### Get Exchange Wallet

If you want to get exchange wallets use this command:

```shell
   $ wallet-tracker redis get --exchanges uniswap --limit 3
```

## Build

Basic building process like the following would suffice.

```shell
   $ go build -o wallet-tracker cmd/wallet-tracker/main.go
```

## Downloads

### Tarball

1. Download [latest-release] for your operating system/architecture
2. Unzip binary and place it somewhere in your path
3. Make it executable


<!-- ROADMAP -->

## Roadmap

See the [open issues](https://github.com/aydinnyunus/wallet-tracker/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to be learned, inspire, and create. Any
contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the Apache License 2.0 License. See `LICENSE` for more information.



<!-- CONTACT -->

## Contact

[<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/linkedin.png" title="LinkedIn">](https://linkedin.com/in/yunus-ayd%C4%B1n-b9b01a18a/)       [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/github.png" title="Github">](https://github.com/aydinnyunus/WhatsappBOT)     [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/instagram-new.png" title="Instagram">](https://instagram.com/aydinyunus_/) [<img target="_blank" src="https://img.icons8.com/bubbles/100/000000/twitter.png" title="LinkedIn">](https://twitter.com/aydinnyunuss)




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/usestrix/cli.svg?style=for-the-badge

[contributors-url]: https://github.com/aydinnyunus/wallet-tracker/graphs/contributors

[forks-shield]: https://img.shields.io/github/forks/usestrix/cli.svg?style=for-the-badge

[forks-url]: https://github.com/aydinnyunus/wallet-tracker/network/members

[stars-shield]: https://img.shields.io/github/stars/usestrix/cli?style=for-the-badge

[stars-url]: https://github.com/aydinnyunus/wallet-tracker/stargazers

[issues-shield]: https://img.shields.io/github/issues/usestrix/cli.svg?style=for-the-badge

[issues-url]: https://github.com/aydinnyunus/wallet-tracker/issues

[license-shield]: https://img.shields.io/github/license/usestrix/cli.svg?style=for-the-badge

[license-url]: https://github.com/aydinnyunus/wallet-tracker/blob/master/LICENSE.txt

[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555

[linkedin-url]: https://linkedin.com/in/aydinnyunus

[product-screenshot]: data/images/base_command.png

[latest-release]: https://github.com/aydinnyunus/wallet-tracker/releases
