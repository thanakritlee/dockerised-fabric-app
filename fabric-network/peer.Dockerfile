FROM hyperledger/fabric-peer:amd64-1.4.0

WORKDIR /opt/gopath/src/github.com/hyperledger/fabric

COPY ./crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp /etc/hyperledger/msp/peer
COPY ./crypto-config/peerOrganizations/org1.example.com/users /etc/hyperledger/msp/users
COPY ./config /etc/hyperledger/configtx

ENV CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
ENV CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/peer/

EXPOSE 7051
EXPOSE 7053

CMD peer node start