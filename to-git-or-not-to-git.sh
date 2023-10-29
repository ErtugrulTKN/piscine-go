#!/bin/bash
curl https://talent.uniworkhub.com/assets/superhero/all.json |  jq '.[] | select(.id==170).name'.power'.gender'