# FyxVM

FyxVM takes lessons learned from building on EVM (Ethereum Virtual Machine) and Run (Run on Bitcoin) to produce a new UTXO-based virtual state machine for running smart contract functionality on bitcoin-compatible blockchains. 

FyxVM contracts are stored on the blockchain as gzipped web assembly modules. Contracts provide methods for creating and modifying the state of on-chain object instances. Instances are encoded into Bitcoin UTXO script, such that 
