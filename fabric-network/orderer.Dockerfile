FROM hyperledger/fabric-orderer:amd64-1.4.0

WORKDIR /opt/gopath/src/github.com/hyperledger/fabric/orderer

COPY ./config/ /etc/hyperledger/configtx
COPY ./crypto-config/ordererOrganizations/example.com/orderers/orderer0.example.com/ /etc/hyperledger/msp/orderer
COPY ./crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/ /etc/hyperledger/msp/peerOrg1

ENV ORDERER_GENERAL_GENESISFILE=/etc/hyperledger/configtx/genesis.block
ENV ORDERER_GENERAL_LOCALMSPDIR=/etc/hyperledger/msp/orderer/msp

EXPOSE 7050

CMD orderer