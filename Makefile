.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && \
	cwgo server \
	-I ../../idl \
	--type HTTP \
	--service frontend \
	--module github.com/A1sca/Douyin-Mall-Go/app/frontend \
	--idl ../../idl/frontend/home.proto
