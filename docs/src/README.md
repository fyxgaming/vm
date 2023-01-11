# Introduction

## How did we get here

When we set out to build [Cryptofights](https://whitepaper.cryptofights.io/), one of the things uppermost in our minds was whether the game would scale, not just to thousands but to millions of users.

Our company was one of the first adopters of the web3 technologies. Having finally settled on using [BitcoinSV](https://bitcoinsv.com/) as the blockchain where our games' transactional data would live, the next challenge was to find a solution that would enable us to create assets that would live on the blockchain. These assets entailed both code, that would be run on demand, as well as non-code assets such as plaintext data, images, videos and other multi-media assets.

We finally chose the [RUN SDK](https://run.network/) to interact with the BitcoinSV blockchain. However, RUN SDK came with its own slew of challenges, not the least of which was the fact that RUN was written in JavaScript. While the platform functionally worked well, it wasn't long that we began encountering performance & scaling bottlenecks.

In our own performance tests, we had hit over 10M+ blockchain transactions in a single day and there was nothing publicly available that could support the kind of volume we were looking to generate on the blockchain. Our Chief Architect soon realized that we needed to create something ourselves that could support the performance and throughput we had benchmarked. 

## Welcome to FYX VM!

If you are aware and/or familiar with the Ethereum Virtual Machine (EVM), you would know that an EVM is a layer on top of the core Ethereum blockchain that enables interaction between artifacts in this layer and the core blockchain. 

FYX VM is somewhat similar in that it comprises the following two core components:

- Fyx Protocol - This is a layer that interacts with the core Bitcoin SV blockchain. This is the layer that is responsible for creating transactions on the blockchain, and for querying its state. 
- Fyx Runtime - These are the custom components (call them smart contracts if you will) that invoke the Fyx Protocol and that ultimately update the state of the blockchain e.g. a component (smart contract code) that saves an image to the blockchain.

Fyx Protocol + Fyx Runtime = Fyx VM
