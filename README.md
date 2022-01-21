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
