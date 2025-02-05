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
	--service auth \
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

# ===============  product 模块 ===================
.PHONY: gen-product-client gen-product-server gen-product-http

gen-product-client:
	@cd rpc_gen && \
	cwgo client \
	-I ../idl \
	--type RPC \
	--service product \
	--module github.com/A1sca/Douyin-Mall-Go/rpc_gen \
	--idl ../idl/product.proto

gen-product-server:
	@cd app/product && \
	cwgo server \
	-I ../../idl \
	--type RPC \
	--service product \
	--module github.com/A1sca/Douyin-Mall-Go/app/product \
	--pass "-use github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen" \
	--idl ../../idl/product.proto

gen-product-http:
	@echo "gen-product-http 待实现"

# ===============  cart 模块 ===================
.PHONY: gen-cart-client gen-cart-server gen-cart-http

gen-cart-client:
	@cd rpc_gen && \
	cwgo client \
	-I ../idl \
	--type RPC \
	--service cart \
	--module github.com/A1sca/Douyin-Mall-Go/rpc_gen \
	--idl ../idl/cart.proto

gen-cart-server:
	@cd app/cart && \
	cwgo server \
	-I ../../idl \
	--type RPC \
	--service cart \
	--module github.com/A1sca/Douyin-Mall-Go/app/cart \
	--pass "-use github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen" \
	--idl ../../idl/cart.proto

gen-cart-http:
	@echo "gen-cart-http 待实现"

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

# ===============  checkout 模块 ===================
.PHONY: gen-checkout-client gen-checkout-server gen-checkout-http

gen-checkout-client:
	@cd rpc_gen && \
	cwgo client \
	-I ../idl \
	--type RPC \
	--service checkout \
	--module github.com/A1sca/Douyin-Mall-Go/rpc_gen \
	--idl ../idl/checkout.proto

gen-checkout-server:
	@cd app/checkout && \
	cwgo server \
	-I ../../idl \
	--type RPC \
	--service checkout \
	--module github.com/A1sca/Douyin-Mall-Go/app/checkout \
	--pass "-use github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen" \
	--idl ../../idl/checkout.proto

gen-checkout-http:
	@echo "gen-checkout-http 待实现"

# ===============  payment 模块 ===================
.PHONY: gen-payment-client gen-payment-server gen-payment-http

gen-payment-client:
	@cd rpc_gen && \
	cwgo client \
	-I ../idl \
	--type RPC \
	--service payment \
	--module github.com/A1sca/Douyin-Mall-Go/rpc_gen \
	--idl ../idl/payment.proto

gen-payment-server:
	@cd app/payment && \
	cwgo server \
	-I ../../idl \
	--type RPC \
	--service payment \
	--module github.com/A1sca/Douyin-Mall-Go/app/payment \
	--pass "-use github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen" \
	--idl ../../idl/payment.proto

gen-payment-http:
	@echo "gen-payment-http 待实现"

# ============= 生成所有模版代码 ==============
.PHONY: gen-client gen-server gen-http

gen-client: gen-auth-client gen-user-client gen-product-client gen-cart-client gen-order-client gen-checkout-client gen-payment-client 

gen-server: gen-auth-server gen-user-server gen-product-server gen-cart-server gen-order-server gen-checkout-server gen-payment-server 

gen-http: gen-auth-http gen-user-http gen-product-http gen-cart-http gen-order-http gen-checkout-http gen-payment-http 
