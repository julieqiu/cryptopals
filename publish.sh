u="github.com/julieqiu/cryptopals@"$(go list -m -f {{.Version}} github.com/julieqiu/cryptopals@$(git rev-parse HEAD)) &&
curl -X POST "https://pkg.go.dev/fetch/$u" &&
open "https://pkg.go.dev/$u/set$1/c$2"
