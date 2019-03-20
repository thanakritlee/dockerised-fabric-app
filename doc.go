// Package dockerisedfabricapp shows an example Hyperledger Fabric application structure.
//
// Basic workflow
//
//      1) Bring up the Fabric network using docker-compose in fabric-network/.
//      2) Start the Go API server in web/.
//      3) Send request to the API server.
//      4) Stop the server, and delete user key store at /tmp/dockerised-fabric-app*.
//      5) Down the Fabric docker-compose network.
package dockerisedfabricapp
