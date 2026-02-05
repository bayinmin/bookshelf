package attack

var Introspection = map[string]interface{}{
	"variables": map[string]interface{}{},
	"query": `
        {
          __schema {
            queryType { name }
            mutationType { name }
            subscriptionType { name }
            types { ...FullType }
            directives {
              name
              description
              locations
              args { ...InputValue }
            }
          }
        }

        fragment FullType on __Type {
          kind
          name
          description
          fields(includeDeprecated: true) {
            name
            description
            args { ...InputValue }
            type { ...TypeRef }
            isDeprecated
            deprecationReason
          }
          inputFields { ...InputValue }
          interfaces { ...TypeRef }
          enumValues(includeDeprecated: true) {
            name
            description
            isDeprecated
            deprecationReason
          }
          possibleTypes { ...TypeRef }
        }

        fragment InputValue on __InputValue {
          name
          description
          type { ...TypeRef }
          defaultValue
        }

        fragment TypeRef on __Type {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
                ofType {
                  kind
                  name
                  ofType {
                    kind
                    name
                    ofType {
                      kind
                      name
                      ofType {
                        kind
                        name
                      }
                    }
                  }
                }
              }
            }
          }
        }
    `,
}

var MutationIntrospection = map[string]interface{}{
	"variables": map[string]interface{}{},
	"query": `
        {
          __schema {
            mutationType { name }
            types { ...FullType }
            directives {
              name
              description
              locations
              args { ...InputValue }
            }
          }
        }

        fragment FullType on __Type {
          kind
          name
          description
          fields(includeDeprecated: true) {
            name
            description
            args { ...InputValue }
            type { ...TypeRef }
            isDeprecated
            deprecationReason
          }
          inputFields { ...InputValue }
          interfaces { ...TypeRef }
          enumValues(includeDeprecated: true) {
            name
            description
            isDeprecated
            deprecationReason
          }
          possibleTypes { ...TypeRef }
        }

        fragment InputValue on __InputValue {
          name
          description
          type { ...TypeRef }
          defaultValue
        }

        fragment TypeRef on __Type {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
                ofType {
                  kind
                  name
                  ofType {
                    kind
                    name
                    ofType {
                      kind
                      name
                    }
                  }
                }
              }
            }
          }
        }
    `,
}

var SubscriptionIntrospection = map[string]interface{}{
	"variables": map[string]interface{}{},
	"query": `
        {
          __schema {
            subscriptionType { name }
            types { ...FullType }
            directives {
              name
              description
              locations
              args { ...InputValue }
            }
          }
        }

        fragment FullType on __Type {
          kind
          name
          description
          fields(includeDeprecated: true) {
            name
            description
            args { ...InputValue }
            type { ...TypeRef }
            isDeprecated
            deprecationReason
          }
          inputFields { ...InputValue }
          interfaces { ...TypeRef }
          enumValues(includeDeprecated: true) {
            name
            description
            isDeprecated
            deprecationReason
          }
          possibleTypes { ...TypeRef }
        }

        fragment InputValue on __InputValue {
          name
          description
          type { ...TypeRef }
          defaultValue
        }

        fragment TypeRef on __Type {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
                ofType {
                  kind
                  name
                  ofType {
                    kind
                    name
                    ofType {
                      kind
                      name
                    }
                  }
                }
              }
            }
          }
        }
    `,
}

var SimpleIntrospection = map[string]interface{}{
	"variables": map[string]interface{}{},
	"query": `
        {
          __schema {
            types {
              name
              kind
              fields {
                name
                type { name kind }
              }
            }
          }
        }
    `,
}

var SimpleMutationIntrospection = map[string]interface{}{
	"variables": map[string]interface{}{},
	"query": `
        {
          __schema {
            mutationType { name }
            types {
              name
              kind
              fields {
                name
                type { name kind }
              }
            }
          }
        }
    `,
}

var ListAvilableQueries = map[string]interface{}{
	"variables": map[string]interface{}{},
	"query": `
        {
          __schema {
            queryType { name }
            mutationType { name }
            subscriptionType { name }
            types { ...FullType }
            directives {
              name
              description
              locations
              args { ...InputValue }
            }
          }
        }

        fragment FullType on __Type {
          kind
          name
          description
          fields(includeDeprecated: true) {
            name
            description
            args { ...InputValue }
            type { ...TypeRef }
            isDeprecated
            deprecationReason
          }
          inputFields { ...InputValue }
          interfaces { ...TypeRef }
          enumValues(includeDeprecated: true) {
            name
            description
            isDeprecated
            deprecationReason
          }
          possibleTypes { ...TypeRef }
        }

        fragment InputValue on __InputValue {
          name
          description
          type { ...TypeRef }
          defaultValue
        }

        fragment TypeRef on __Type {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
                ofType {
                  kind
                  name
                  ofType {
                    kind
                    name
                    ofType {
                      kind
                      name
                    }
                  }
                }
              }
            }
          }
        }
    `,
}

var ListSchemaTypes = map[string]interface{}{
	"variables": map[string]interface{}{},
	"query": `
        {
          __schema {
            types {
              name
              kind
              description
              fields {
                name
                type { name kind }
              }
            }
          }
        }
    `,
}

var GraphqlQueries = map[string]interface{}{
	"Introspection":               Introspection,
	"MutationIntrospection":       MutationIntrospection,
	"SubscriptionIntrospection":   SubscriptionIntrospection,
	"SimpleIntrospection":         SimpleIntrospection,
	"SimpleMutationIntrospection": SimpleMutationIntrospection,
	"ListAvilableQueries":         ListAvilableQueries,
	"ListSchemaTypes":             ListSchemaTypes,
}
