#!/bin/bash

ls -1 | tail -n +2 | awk 'NR % 2==1'
