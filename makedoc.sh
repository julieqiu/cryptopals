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
go run ./cmd/makedoc $1 $2
