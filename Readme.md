There are two ways the path to configuration and plugins can be specified

# Specifying config and plugin path 

## During building
```
SYSCONFDIR=$(pwd) LIBDIR=$(pwd)/bin make all
```

## During runtime
```
HOUSEKEEPER_CONFIGURATION_PATH=$(pwd) HOUSEKEEPER_PLUGIN_PATH=$(pwd)/bin bin/housekeeper
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

The website is using ReactJS and is located in /webserver

```
cd webserver
yarn
yarn build
```
