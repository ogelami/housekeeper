# Specifying configuration path
There are two ways the path to configuration can be specified

## During building
```
SYSCONFDIR=$(pwd) make all
```

## During runtime
```
HOUSEKEEPER_CONFIGURATION_PATH=$(pwd) bin/housekeeper
```

# Prerequisite
* A MQTT server running

# Building

## housekeeper and all plugins
```
git clone https://github.com/ogelami/housekeeper
cd housekeeper
make all
```

## Building the webserver content and serving

The website is using ReactJS and is located in /web more information on how to build that part of the project can be found [here](web/README.md).

```
cd webserver
yarn
yarn build
```
