<h1 align="center">
    <picture>
      <img width="350" alt="roapi.go" src="./assets/images/roapi.png">
    </picture>
  <br>
  <a href="https://github.com/jaxron/roapi.go/blob/main/LICENSE.md">
    <img src="https://img.shields.io/github/license/jaxron/roapi.go?style=flat-square&color=008ae6">
  </a>
  <a href="https://github.com/jaxron/roapi.go/actions/workflows/ci.yml">
    <img src="https://img.shields.io/github/actions/workflow/status/jaxron/roapi.go/ci.yml?style=flat-square&color=008ae6">
  </a>
  <a href="https://github.com/jaxron/roapi.go/issues">
    <img src="https://img.shields.io/github/issues/jaxron/roapi.go?style=flat-square&color=008ae6">
  </a>
</h1>

<p align="center">
  <em><b>RoAPI.go</b> is a powerful and modular API wrapper for Roblox, written in <a href="https://golang.org/">Go</a>, empowering developers to effortlessly integrate with its platform services.</em>
</p>

---

> [!WARNING]
> This library is currently in **early development** and is **not ready for production use**. It covers only a very small fraction of the Roblox API at this time. Progress can be tracked [here](https://github.com/jaxron/roapi.go/issues/1).

# 🚀 Features

RoAPI.go offers features that prioritize flexibility and reliability. Key features include:

- **Advanced Client:**
  - [Circuit breaker](https://learn.microsoft.com/en-us/azure/architecture/patterns/circuit-breaker) for fault tolerance
  - [Retry mechanism](https://learn.microsoft.com/en-us/azure/architecture/patterns/retry) with exponential backoff
  - [Rate limiting](https://learn.microsoft.com/en-us/azure/architecture/patterns/rate-limiting-pattern) to prevent API throttling
  - Request deduplication via `singleflight`
  - Dynamic proxy rotation for distributed traffic
  - All features configurable and toggleable
- **Beginner-Friendly:**
  - Simple request construction using builders
  - No need to understand Roblox's API in-depth
- **Robust Authentication:**
  - Cookie-based authentication supported with rotation
  - Automatic CSRF token handling and refresh
- **Easy to Troubleshoot:**
  - Detailed error messages with root cause and response body
  - Configurable logging modes for debugging

> [!NOTE]
> RoAPI.go is an independently developed project and is not affiliated with Roblox Corporation. It is neither endorsed by nor sponsored by Roblox Corporation, and "Roblox" is a registered trademark of Roblox Corporation.
>
> If you encounter any issues related to this project, please report them directly on our [GitHub Issues page](https://github.com/jaxron/roapi.go/issues). Please do not contact Roblox Corporation for support regarding this library.

# 📄 License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.
