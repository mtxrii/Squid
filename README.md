# Squid CLI
A little utility to grab content from the web, as well as profile websites.

### HTTP GET (url command)

![](https://i.imgur.com/4mArxGu.png)

Sends a GET request to specified link. In this case, a request was sent to [linktree.edavalos.com/links](https://linktree.edavalos.com/links)

### Ping (profile command)

![](https://i.imgur.com/dNoyxJU.png)
![](https://i.imgur.com/176gE9F.png)
![](https://i.imgur.com/VtGcsQP.png)
![](https://i.imgur.com/xKo2Vbr.png)

Profiles a website. The examples above show speeds and packet sizes from a small website: [linktree.edavalos.com](https://linktree.edavalos.com/) (cloudflare wrangler), a medium one: [google.com](https://www.google.com/), and a large one: [github.com](https://github.com/mtxrii), as well as an invalid url.

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
* squid url <full url> - sends a get request and prints the results
* squid profile <ammount> <full url> - sends specified number of requests to specified url and profiles speeds
* squid help (cmd) - pulls up help menu
