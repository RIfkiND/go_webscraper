set -e

command_exist(){
    command -v "$1" >/dev/null 2>&1
}

echo "Instaling dependencies ..."

if [! -f .env]; then
echo ".env not found .create new"
cat << EOL > .env
#Default end

APP_ENV=development
APP_PORT=8080
SCRAPER_LOG_LEVEL=info
DATA