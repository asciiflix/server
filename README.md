
# Asciiflix Server

Imagine this, you just spent an entire summer on netflix. 
This sucks, netflix, prime, etc. are all far too addictive and in many cases impossible to use with german bandwidth.
What we need, is a bad version of netflix with low bandwidth usage. 
Asciiflix can provide all of these features!

<!-- TODO: Add Demo Image -->

## Table of Contents
- [Contributing](#contributing) 
- [License](#license)



## Contributing

Contributions are always welcome!

If you want to solve an issue or add a feature, just hit us up with a new pull request.

### Building

If you want to build your own container image simply use the docker-compose file:

```bash
docker-compose build --build-arg VERSION=YOURVERSION
```

### Deployment

If you want to start this app on your own server, you can simply start the docker-compose file:

```bash
docker-compose up -d
```

The ASCIIflix Server will listen on port 8080 with the default settings.

### Development

To start and run the application with your newest changes and debugging capabilities, you can simply open this project in VSCode with the .devcontainer configs. <br>
Of course you have to install docker on your dev host. After installing docker you can enter the Container Mode:

1. Press `CTRL + SHIFT + P`
2. Type `Remote-Containers: Reopen in Container`, but keep in mind your current VSCode project should be the asciiflix server project
3. VSCode will download a docker image, this can take some minutes. After the download is finished, VSCode reopens the project in a docker-container. Now you can develop, debug and do some dev stuff..

*Note: in order to load the config you have to create a ``config.env`` in the root dir. An example can be found in ``config.env.example``, which can simply be renamed.*


## License

[GPL v3](https://www.gnu.org/licenses/gpl-3.0.en.html)
