# Relaygen

GraphQL Relay support for gqlgen. Using code generation assists the boilerplates creation of nodes, edges & pagination types.

## Installation

Dependencies:

- go 1.13

`go get -u github.com/sebach1/relaygen`

## Usage

> For further details about usage, use the relaygen CLI help cmd `relaygen --help`

Notice that the default output is given to the **STDOUT**, so you can preview any generation omitting the `> <FILE>` boilerplate.

<br>

1. Add the base boilerplate to support relay dep injections on  golang models:

`relaygen -base > <FILENAME>.go`

2. Add the schema interface & relay definitions:
  
`relaygen -pkg <PKG_NAME> -base -sdl > <FILENAME>.graphql`

3. Add the golang entity boilerplates as needed:

`relaygen -pkg <PKG_NAME> -name <ALIAS_OF_ENTITY> -type <TYPE_OF_ENTITY> > <FILENAME>.go`

4. Add the entity to your schema definition:

`relaygen -pkg <PKG_NAME> -name <ALIAS_OF_ENTITY> -type <TYPE_OF_ENTITY> -sdl > <FILENAME>.go`

5. Manual step to avoid magic configs. Add the needed config to your `gqlgen.yaml` config file.

```yaml
models:
  PageInfo:
    model: <BASE_DEF_PATH>/relay.PageInfo  # Notice that "relay" is the default value of the flag when generating the base
  Node:
    model: <BASE_DEF_PATH>/relay.Node
  <ALIAS_OF_ENTITY>Edge:
    model: <ENTITY_RELAY_PATH>/<PKG>.<ALIAS_OF_ENTITY>Edge
  <ALIAS_OF_ENTITY>Connection:
    model: <ENTITY_RELAY_PATH>/<PKG>.<ALIAS_OF_ENTITY>Connection
```
