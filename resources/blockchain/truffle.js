let HDWalletProvider = require("truffle-hdwallet-provider");
let Web3 = require("web3");

let provider = (endpoint) => {
    if (process.env.HDWALLET_MNEMONIC) {
        return new HDWalletProvider(process.env.HDWALLET_MNEMONIC, endpoint);
    } else {
        return new Web3.providers.HttpProvider(endpoint);
    }
};

const ganachePort = parseInt(process.env.DAEMON_GANACHE_PORT || 8545);

module.exports = {
    networks: {
        local: {
            host: "127.0.0.1",
            port: ganachePort,
            network_id: "*" // Any network ID
        },
        localhd: {
            provider: () => provider("http://127.0.0.1:" + ganachePort),
            network_id: "*" // Any network ID
        },
        sepolia: {
            gasPrice: 500000000,
            provider: () => provider("https://sepolia.infura.io/v3"),
            network_id: "11155111" // Sepolia network ID
        },
    },
};
