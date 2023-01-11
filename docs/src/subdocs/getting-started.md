# Getting Started

- Sign up for a developer account. 
    - has an L&P that will bill the users account
- Deploy a contract that points to that code (from our library/template) (puts a contract out on chain)
    - The user would select a contract from the list of available public templates or choose to create one from scratch.
    - Service gets set up for the user (AWS VMs would run many binaries that users want as a service).
    - This can be used as an on-chain file storage service for apps/games. A CDN cache for resources on chain.
    - "deploy [file location]" using a terminal app.
    - Creating a CLI to interact.
        - Would keep track of what files are uploaded.
- Contract installed (getting contract up and running on the server) (file location is onchain).
    - Could be a CLI command. (Abel might be working on this)
    - User permissions (admin true: false) ➝ install proceeds
    - The contract can be installed and uninstalled.
- Mint token supply (contract interaction)
- What are the tools a dev has in writing a contract?
    - Raise Events
    - Ownership
    - Data (binary blob local variable)
    - Spawn Children


All state changes are stored within the FYX database.

Limitations of contract: have to be deterministic.

