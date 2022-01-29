### Cloning project
```
# clone housekeeper
git clone https://github.com/ogelami/housekeeper

# clone housekeeper including react web-front
git clone --recurse-submodules -j8 https://github.com/ogelami/housekeeper
```

### Building
```
# fetch and build housekeeper
make all

# building the web-front
cd web
yarn
yarn build
```
