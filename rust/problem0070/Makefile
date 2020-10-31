SHELL := /bin/bash

time : target/release/problem0070
	time target/release/problem0070

target/release/problem70 : src/*
	cargo build  --release

run :
	cargo run

.PHONY: run time
