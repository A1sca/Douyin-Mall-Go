.PHONY: gen-frontend gen-frontend-home gen-frontend-auth

gen-frontend-home:
	@cd app/frontend && \
	cwgo server \
	-I ../../idl \
	--type HTTP \
	--service frontend \
	--module github.com/A1sca/Douyin-Mall-Go/app/frontend \
	--idl ../../idl/frontend/home.proto

gen-frontend-auth:
	@cd app/frontend && \
	cwgo server \
	-I ../../idl \
	--type HTTP \
	--service frontend \
	--module github.com/A1sca/Douyin-Mall-Go/app/frontend \
	--idl ../../idl/frontend/auth_page.proto

gen-frontend: gen-frontend-auth gen-frontend-home

.PHONY: gen-user-client gen-user-server

gen-user-client:
	@cd rpc_gen && \
	cwgo client \
	-I ../idl \
	--type RPC \
	--service user \
	--module github.com/A1sca/Douyin-Mall-Go/rpc_gen \
	--idl ../idl/user.proto

gen-user-server:
	@cd app/user && \
	cwgo server \
	-I ../../idl \
	--type RPC \
	--service user \
	--module github.com/A1sca/Douyin-Mall-Go/app/user \
	--pass "-use github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen" \
	--idl ../../idl/user.proto
