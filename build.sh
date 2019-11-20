#!/bin/sh

cd frontend
docker build -t koala-pos-server-ui .
cd ..

cd customer-frontend
docker build -t koala-pos-customer-ui .
cd ..

cd backend
go run mage.go buildProd
cd ..
