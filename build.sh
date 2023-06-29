#!/bin/bash

help(){
    echo "Usage : bash build.sh [OPTIONS]"
    echo ""
    echo "Build the project"
    echo ""
    echo "Options: "
    echo -e "\t-h display help"
    echo -e "\t-d generate docker image"
    echo ""
    echo "Sample : bash build.sh -d"
    exit 0
}

#Check if script is running throught /bin/sh (SHLVL=1) or /bin/bash (SHLVL=2)
if [ "$SHLVL" -lt "2" ] ; then
    echo "ERROR : Consider running this script using /bin/bash not /bin/sh. \"/bin/bash build.sh -h\" for help"
    exit 1
fi

declare docker=false

declare red='\033[0;31m'
declare yellow='\033[0;33m'
declare default='\033[0m'
declare cyan='\033[0;36m'

while getopts ":hd" option; do
   case $option in
      h) # display Help
         help
         exit;;
      d) # docker
         echo -e "${yellow}App and Docker image will be generated"${default}""
         docker=true
         ;;
      n) # native
         echo -e "${yellow}Native Docker image will be generated locally"${default}""
         native=true
         ;;
      \?) # exclude
         echo -e "${red}Error: Invalid option. Use -h for help${default}"
         exit;;
   esac
done

echo -e "${cyan}Building project...${default}"

echo "Generate package folder..."
# clean package
rm -dr package/

mkdir package

# copying resources
cp -r docker/* package/

#Go version needed : 1.20
go mod tidy
#build linux binary
GOOS=linux GOARCH=amd64 go build -o bin/go-api-amd64-linux src/main.go

# possible to build for linux, windows and macOS on different architectures
#GOOS=windows GOARCH=386 go build -o bin/app-386.exe app.go
# 64-bit
#GOOS=darwin GOARCH=amd64 go build -o bin/app-amd64-darwin app.go
# 32-bit
#GOOS=darwin GOARCH=386 go build -o bin/app-386-darwin app.go
# 64-bit
#GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux app.go
# 32-bit
#GOOS=linux GOARCH=386 go build -o bin/app-386-linux app.go

if [ "$docker" = true ]; then
    echo -e "${yellow}Building docker image...${default}"

    docker build --no-cache --build-arg VERSION=master -t valentinconan/go-api:master .
fi

echo -e "${cyan}Build project done${default}"




