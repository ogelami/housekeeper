There are two ways the path to configuration and plugins can be specified

# During building
```
SYSCONFDIR=$(pwd) LIBDIR=$(pwd)/bin make all
```

# During runtime
```
HOUSEKEEPER_CONFIGURATION_PATH=$(pwd) HOUSEKEEPER_PLUGIN_PATH=$(pwd)/bin bin/housekeeper
```