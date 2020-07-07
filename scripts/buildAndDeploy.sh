#!/bin/bash

cd web
yarn build

cd build
tar zcvf q.tar.gz *

scp q.tar.gz wall:
