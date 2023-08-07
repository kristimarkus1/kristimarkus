#!/bin/bash

curl -s "https://01.kood.tech/assets/superhero/all.json" | jq -r '.[] | select(.id==170) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'
