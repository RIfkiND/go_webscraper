ser -e

if [-f .env];then
echo "Load env"
export $(grep -v '^#' .env|xargs)
fi


echo "Buildin Application..."

go build -o bin/scraper ./cmd/main.go

echor "Running application..."
.bin/scraper

rm -rf bin/

echo "Complete"