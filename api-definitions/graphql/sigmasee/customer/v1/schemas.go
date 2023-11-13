package schema

import (
	_ "embed"
)

//go:embed schema.graphql
var Schema string

//go:embed ent.graphql
var EntgoSchema string
