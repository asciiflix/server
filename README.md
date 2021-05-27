# ACIIflix Server

## Development
To start up and run the application with your newest changes, without debugging capabilities, run:
```bash
sudo docker-compose up -d --build
```

If you want to run the application with a debugger, you first have to change the build target (of the go application server) from prod to debug. ``target: prod`` -> ``target: debug``

Then start the docker containers as explained above. When you start the debugging container, this only opens up the delve debugger. To start the application, open up a debugging session.

If you are using vscode, use the run config supplied in [./.vscode/lauch.json](./.vscode/lauch.json). With other IDEs, such as Goland, open up a remote debugging session at port ``2345``.
