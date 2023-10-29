#!/bin/bash
curl https://talent.uniworkhub.com/assets/superhero/all.json |  jq ".[] | select(.id==$HERO_ID ).connections.relatives" | tr -d '"'
