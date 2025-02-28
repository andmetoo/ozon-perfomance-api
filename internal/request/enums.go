package request

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal --names

// ContentType
// ENUM(application/json=0)
type ContentType string
