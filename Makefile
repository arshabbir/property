build:
	go build -o property
run:
	go build -o property
	CORE_CHAINCODE_ID_NAME=prop:0 CORE_PEER_TLS_ENABLED=false ./property -peer.address localhost:7051