# Test Workspace with include feature

This is related to https://github.com/dagger/dagger/pull/13882

- `modules` is where a module is authored. The `dagger.toml` contains the go-sdk, and this is where we can develop the module
- `shared` is a place where the module is simply installed into the `dagger.toml`. This represent the set of modules that we want to share
- `project` is somewhere else we want to use the modules shared by `shared`, and developped inside `modules`. But here it's only about using them

To try it:

1. First cd into project:

    ```console
    $ cd project
    ```

2. Have a look at the `dagger.toml`, it should only have  `include` instruction

    ```console
    $ cat dagger.toml
    ```

    ```toml
    [[include]]
    source = "../shared"
    ```

3. Use the available module

    ```console
    $ dagger module list
    NAME       SOURCE
    greetings  modules/greetings

    $ dagger functions greetings
    Name                 Description
    base-image-address   Image the build runs on.
    container            A container with the source mounted, ready to build on.
    say-hello            An OCI-compatible container, also known as a Docker container.

    $ dagger api call greetings say-hello --name plop stdout
    TRACE  ✔ PASSED
    greetings say-hello --name plop stdout  1.6s
    
    ✔ connect 0.2s
    ✔ load workspace: . 0.1s
    ✔ parsing command line arguments 0.0s
    
    ✔ greetings(
      ┆ ws: currentWorkspace: Workspace!
      ): Greetings! 0.1s
      ✔ uploading /Users/yves/dev/src/github.com/eunomie/test-ws 0.0s ████████████ 483 B
    ✔ .sayHello(
      ┆ ws: currentWorkspace: Workspace!
      ┆ name: "plop"
      ): Container! 1.2s
    ✔ .stdout: String! 0.0s
    
    hello plop
    ```
