# blog
**blog** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## Defining blog post
```
ignite scaffold type post title body creator id:uint

```
## Implement CRUD operations
### Scaffold create message
```
ignite scaffold message create-post title body --response id:uint

```
### Scaffold update messsage

```
ignite scaffold message update-post title body id:uint
```

### Scaffold delete message
```
ignite scaffold message delete-post id:uint

```

### Scaffolding query message
```
ignite scaffold query show-post id:uint --response post:Post
ignite scaffold query list-post --response post:Post --paginated
```

## Build the chain
```
ignite chain build
```

## Interacting with blog chain
```
blogd tx blog create-post hello world --from alice --chain-id blog
blogd q blog show-post 0
blogd q blog list-post
blogd tx blog update-post "Hello" "Cosmos" 0 --from alice --chain-id blog
blogd query blog list-post
blogd tx blog delete-post 0 --from alice  --chain-id blog


```


### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/rohitroyrr8/blog@latest! | sudo bash
```
`rohitroyrr8/blog` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
