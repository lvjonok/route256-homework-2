#!/bin/bash

DBSTRING="host=localhost user=root password=root dbname=rootdb sslmode=disable"

goose postgres "$DBSTRING" up