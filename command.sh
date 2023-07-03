#!/bin/sh

if [ "$1" = "daemon" ]; then
  clear
  CompileDaemon -command="./restapi"
elif [ "$1" = "build" ]; then
  go build main.go
elif [ "$1" = "start" ]; then
  go run main.go
elif [ "$1" = "migrate" ]; then
  go run ./migrate/migrate.go
elif [ "$1" = "seed" ]; then
  go run ./seed/seeder.go
elif [ "$1" = "-h" ] || [ "$2" = "-h" ] || [ "$1" = "help" ] || [ "$2" = "help" ]; then
  clear
  echo "usege : ./command.sh [command]"
  echo ""
  echo "command:"
  echo "  - daemon // to run Go with CompileDaemon"
  echo "  - build // to run Go build"
  echo "  - start // to run Go start"
  echo "  - migrate // to run Go migrate Database"
  echo "  - seed // to run Go seed data"
  echo "  - help // see other command"
else
  echo "Invalid comand. Usage: ./command.sh"
  echo ""
  echo "./command.sh help"
fi
