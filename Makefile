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
