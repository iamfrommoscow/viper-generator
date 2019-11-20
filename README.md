# viper generator

## Installation
1. If you don't have installed Go, you can install it __[here](https://golang.org/doc/install)__
2. Then clone this repo ```git clone https://github.com/iamfrommoscow/viper-generator.git```
3. Go to the project path ```cd viper-generator```

## How to use
To generate code, you need to use command ```go run .``` with flags

You can use flags to:

1. Name your module and write the path where to save your new module ```go run . -moduleName=NewModule -path=modules/```
2. Remove Interactor or Router if you don't want to use it in your module ```go run . -interactor=false -router=false -moduleName=NewModule -path=modules/```