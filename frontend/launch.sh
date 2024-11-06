#!/bin/bash

FRONTEND_DIR="frontend"
BACKEND_DIR="backend"

read -p "do you want to compile the React(y/n): " compile_frontend

if [[ "$compile_frontend" == "y" || "$compile_frontend" == "Y" ]]; then
    echo "Compiling the React application..."
    cd "$FRONTEND_DIR"
    npm install
    npm run build
    cd ..
fi

echo "starting React..."
cd "$FRONTEND_DIR"
npm start &
frontend_pid=$!
cd ..

echo "Starting Gin..."
cd "$BACKEND_DIR"
go run main.go &
backend_pid=$!
cd ..

trap "kill $frontend_pid $backend_pid" EXIT

wait $frontend_pid
wait $backend_pid
