SHELL := /bin/bash

time : target/release/problem70
	time target/release/problem70

target/release/problem70 : src/*
	cargo build  --release

run :
	cargo run

.PHONY: run time
