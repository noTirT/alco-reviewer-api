#!/bin/bash
echo "Stopping background ngrok process"
kill -9 $(pgrep ngrok)
echo "ngrok stopped"
docker compose down
