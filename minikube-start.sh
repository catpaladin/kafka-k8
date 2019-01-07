#!/usr/bin/env bash

minikube start \
    --memory=8192 --cpus=4 --kubernetes-version=v1.13.1 \
    --extra-config=apiserver.authorization-mode=RBAC
