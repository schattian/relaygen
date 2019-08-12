# gqlgen-relay
GraphQL Relay support for gqlgen

## Usage
1. Install `gen`

```bash
go get github.com/clipperhouse/gen
```

2. Add relay
```bash
gen add github.com/hookactions/gqlgen-relay
```

3. Add to code

3a. Go code
```go
package model

// +gen relayNode relayCursor
type User struct {
	FirstName, LastName string
}
```

3b. schema.graphql

```graphql
interface Node {
    id: ID!
}

type PageInfo {
    hasNextPage: Boolean!
    hasPreviousPage: Boolean!
    startCursor: String
    endCursor: String
}

type User {
    firstName: String!
    lastName: String!
}

type UserEdge {
    node: User
    cursor: String
}

type UserConnection {
    edges: [UserEdge]
    pageInfo: PageInfo!
    totalCount: Int
}
```

3c. gqlgen.yaml

```yaml
...
models:
  # existing config
  # ...
  # New
  PageInfo:
    model: github.com/hookactions/gqlgen-relay.PageInfo
  Node:
    model: github.com/hookactions/gqlgen-relay.Node
  User:
    model: github.com/your/package/model.User
  UserEdge:
    model: github.com/your/package/model.UserEdge
  UserConnection:
    model: github.com/your/package/model.UserConnection
```
