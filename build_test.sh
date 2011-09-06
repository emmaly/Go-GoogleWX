#!/bin/sh -e

gomake install
8g test.go && 8l -o test test.8
