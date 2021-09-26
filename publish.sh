if [ -z "$1" ]
  then
    echo "No set specified."
    exit 1
fi
if [ -z "$2" ]
  then
    echo "No challenge specified"
    exit 1
fi
u="github.com/julieqiu/cryptopals@"$(go list -m -f {{.Version}} github.com/julieqiu/cryptopals@$(git rev-parse HEAD)) &&
curl -X POST "https://pkg.go.dev/fetch/$u" &&
open "https://pkg.go.dev/$u/set$1/challenge$2"
