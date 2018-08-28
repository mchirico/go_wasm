#!/bin/bash
docker build -t us.gcr.io/mchirico/gosslcorsserver:latest .
docker push us.gcr.io/mchirico/gosslcorsserver:latest
