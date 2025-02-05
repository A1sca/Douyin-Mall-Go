# ===============  auth 模块 ===================
.PHONY: gen-auth-client gen-auth-server gen-auth-http

gen-auth-client: 
	@cd rpc_gen && \
	cwgo client \
	-I ../idl \
	--type RPC \
	--service auth \
	--module github.com/A1sca/Douyin-Mall-Go/rpc_gen \
	--idl ../idl/auth.proto

gen-auth-server:
	@cd app/auth && \
	cwgo server \
	-I ../../idl \
	--type RPC \
	--service auto \
	--module github.com/A1sca/Douyin-Mall-Go/app/auth \
	--pass "-use github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen" \
	--idl ../../idl/auth.proto

gen-auth-http:
	@echo "gen-auth-http 待实现"

# ===============  user 模块 ===================
.PHONY: gen-user-client gen-user-server gen-user-http

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

gen-user-http:
	@echo "gen-user-http 待实现"


# ===============  order 模块 ===================
.PHONY: gen-order-client gen-order-server gen-order-http

gen-order-client:
	@cd rpc_gen && \
	cwgo client \
	-I ../idl \
	--type RPC \
	--service order \
	--module github.com/A1sca/Douyin-Mall-Go/rpc_gen \
	--idl ../idl/order.proto

gen-order-server:
	@cd app/order && \
	cwgo server \
	-I ../../idl \
	--type RPC \
	--service order \
	--module github.com/A1sca/Douyin-Mall-Go/app/order \
	--pass "-use github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen" \
	--idl ../../idl/order.proto

gen-order-http:
	@cd app/api && \
	cwgo server \
	-I ../../idl \
	--type HTTP \
	--service api \
	--module github.com/A1sca/Douyin-Mall-Go/app/api \
	--idl ../../idl/api/api_order.proto
