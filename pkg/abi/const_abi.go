package abi

// Trc20AbiFragment for unit test
const Trc20AbiFragment = `[
    {
        "inputs": [{"name":"to","type":"address"},{"name":"value","type":"uint256"}],
        "name": "transfer",
        "outputs": [{"name":"","type":"bool"}],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"value","type":"uint256"}],
        "name": "transferFrom",
        "outputs": [{"name":"","type":"bool"}],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [{"name":"spender","type":"address"},{"name":"value","type":"uint256"}],
        "name": "approve",
        "outputs": [{"name":"","type":"bool"}],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [{"name":"spender","type":"address"},{"name":"addedValue","type":"uint256"}],
        "name": "increaseAllowance",
        "outputs": [{"name":"","type":"bool"}],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [{"name":"spender","type":"address"},{"name":"subtractedValue","type":"uint256"}],
        "name": "decreaseAllowance",
        "outputs": [{"name":"","type":"bool"}],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [{"name":"account","type":"address"}],
        "name": "balanceOf",
        "outputs": [{"name":"","type":"uint256"}],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [{"name":"owner","type":"address"},{"name":"spender","type":"address"}],
        "name": "allowance",
        "outputs": [{"name":"","type":"uint256"}],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "totalSupply",
        "outputs": [{"name":"","type":"uint256"}],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "name",
        "outputs": [{"name":"","type":"string"}],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "symbol",
        "outputs": [{"name":"","type":"string"}],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "decimals",
        "outputs": [{"name":"","type":"uint8"}],
        "stateMutability": "view",
        "type": "function"
    }
]`
