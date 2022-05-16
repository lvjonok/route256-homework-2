#!/bin/bash

DBSTRING="host=localhost user=root password=root dbname=rootdb sslmode=disable"

# postgresql://root:root@localhost:5432/rootdb

goose postgres "$DBSTRING" up