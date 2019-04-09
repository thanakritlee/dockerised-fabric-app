FROM hyperledger/fabric-ca:amd64-1.4.0

COPY ./crypto-config/peerOrganizations/org1.example.com/ca/ /etc/hyperledger/fabric-ca-server-config

ENV FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
ENV FABRIC_CA_SERVER_CA_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.org1.example.com-cert.pem
ENV FABRIC_CA_SERVER_CA_KEYFILE=/etc/hyperledger/fabric-ca-server-config/d1522eac6db1a54b89e812764f43c70b9d878167a43db403bd77024aa3609ec4_sk

EXPOSE 7054

CMD sh -c 'fabric-ca-server start -b admin:adminpw'