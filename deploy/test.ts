process.env.NETWORK = 'testnet';
import { Address, Tx } from 'bsv';
import fs from 'fs';
import { io } from 'socket.io-client';
import { AuthService } from '@fyxgaming/lib/dist/auth-service';
import { SignedMessage } from '@fyxgaming/lib/dist/signed-message';
import { Owner } from './lib/owner';
import { emit } from 'process';

const API = 'https://dev.api.fyxgaming.com';
const AUTH = 'https://dev.api.cryptofights.io'
const network = 'test';
const UN = 'shruggr1';
const PW = 'test1234';

async function main() {
    const auth = new AuthService(AUTH, network);
    const keyPair = await auth.generateKeyPair(UN, PW);
    const bip32 = await auth.recover(UN, keyPair);
    const address = Address.fromPrivKey(bip32.derive('m/6715768').privKey);
    const owner = new Owner(bip32);
    const socket = io(API, {
        auth: async (cb) => {
            const payload = JSON.stringify(new SignedMessage({
                subject: 'Auth'
            }, UN, keyPair));
            cb({ payload });
        }
    });

    async function emit(event: string, body: any): Promise<any> {
        let resp = await new Promise<any>(r => {
            socket.emit(event, body, r)
        });
        const { status, data } = JSON.parse(resp);
        if (status >= 300) {
            console.error(resp);
            throw new Error("bad-response");
        }
        return data;
    }

    async function deploy(filePath, name) {
        let code = fs.readFileSync(filePath)
        let data = await emit('cryptofights/upload', JSON.stringify({
            // name:'fyx.wasm.gz', 
            data: code.toString('base64'),
            type: 'application/wasm',
            enc: 'gzip'
        }));
        let { uploads: [upload] } = JSON.parse(data);
        // console.log('RESP:', resp);

        data = await emit('cryptofights/deploy', upload.outpoint);
        const ctx = JSON.parse(data);

        await install(name, ctx.ctxs[0].instance.outpoint);
    }

    async function install(name, contract) {
        await emit(`cryptofights/install`, JSON.stringify({
            name,
            contract
        }));
    }

    function sign(signTxn): string {
        const tx = Tx.fromBuffer(Buffer.from(signTxn.rawtx, 'base64'));
        for(let input of signTxn.inputs) {
            owner.signVin(tx, input.vin, input.script, input.sats, 'm/6715768')
        }
        console.log('SIGNED:', tx.toHex());
        return tx.toBuffer().toString('base64');
    }

    socket.on('connect', async () => {
        let data: any;
        const lock = address.toTxOutScript().toBuffer().toString('hex');
        console.log("LOCK", lock);
        // await deploy('../factory/fyx.wasm.gz', 'factory');
        // await deploy('../token/fyx.wasm.gz', 'token');

        let services = JSON.parse(await emit('cryptofights/services', '{}'));
        console.log('SERVICES:', services);

        // let actions = [{
        //     action: 2,
        //     outpoint: services.token.contract,
        //     service: 'factory',
        //     method: 'Init',
        //     callData: Buffer.from(JSON.stringify({
        //         supply: 1000000000,
        //         lock,
        //     }), 'utf8').toString('base64')
        // }];

        // console.log("REQ:", actions);

        // data = await emit('cryptofights/actions', JSON.stringify(actions));
        // console.log("RESP:", data);

        // const signTx = JSON.parse(data);
        // // console.log("SignTxn:", signTx);
        // const signedTx = sign(signTx);
        // // console.log("Signed Txn:", signedTx);

        // data = await emit('cryptofights/broadcast', signedTx);
        // console.log('Broadcast:', data);
        
        data = await emit('cryptofights/instances', JSON.stringify({
            lock,
            kind: Buffer.from(services.token.contract, 'base64').toString('hex'),
        }));
        console.log('Tokens:', data);
    });

    socket.onAny((event, payload) => {
        console.log('Received', event, payload, "\n");
    });

    await new Promise(r => setTimeout(() => r(true), 600000));
}

main().catch((e) => { console.error(e.message); process.exit(1) })