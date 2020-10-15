# Squid CLI
A little utility to grab content from the web, as well as profile websites.

### HTTP GET (url command)

![](https://i.imgur.com/4mArxGu.png)

Sends a GET request to specified link. In this case, a request was sent to [linktree.edavalos.com/links](https://linktree.edavalos.com/links)

### Ping (profile command)
Sends a specified ammount of requests to a website, returning the delay for each one and packet sizes. Below are some example results.

<details>
  <summary>Results from a small website: <a href="https://linktree.edavalos.com/">linktree.edavalos.com</a></summary>
  <img src="https://i.imgur.com/dNoyxJU.png"></img>
</details>

<details>
  <summary>Results from a medium website: <a href="https://www.google.com/">google.com</a></summary>
  <img src="https://i.imgur.com/176gE9F.png"></img>
</details>

<details>
  <summary>Results from a large website: <a href="https://github.com/mtxrii">github.com/mtxrii</a></summary>
  <img src="https://i.imgur.com/VtGcsQP.png"></img>
</details>

<details>
  <summary>Results from an invalid url</summary>
  <img src="https://i.imgur.com/xKo2Vbr.png"></img>
</details>

# Usage & Compilation
To compile this yourself, you'll need [Go](https://golang.org/doc/install).

Once installed, navigate to the root directory and run
```
go build
```

Alternatively you can use the makefile. If you have GNU/Make or are on Unix, simply run
```
make
```

### Help
There are three commands:
* `squid url <full url>` - sends a get request and prints the results
* `squid profile <ammount> <full url>` - sends specified number of requests to specified url and profiles speeds
* `squid help (cmd)` - pulls up help menu
