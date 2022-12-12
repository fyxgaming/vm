process.env.NETWORK = 'testnet';
import { Address, Bip32, Tx } from 'bsv';
import fs from 'fs';
import { io } from 'socket.io-client';
import { AuthService } from '@fyxgaming/lib/dist/auth-service';
import { SignedMessage } from '@fyxgaming/lib/dist/signed-message';
import { Owner } from './lib/owner';
import { emit } from 'process';

const API = 'https://dev.api.fyxgaming.com';
const AUTH = 'https://dev.api.fyxgaming.com'
// const AUTH = 'https://dev.api.cryptofights.io'
const network = 'test';
const UN = 'shruggr1';
const PW = 'test1234';
const PATH = 'm/6715768';

const actions = {
    Auth: 0,
    Mint: 1,
    Call: 2,
    Spawn: 3,
    Deploy: 4
};

let services: any;
let lock: string;
async function main() {
    const auth = new AuthService(AUTH, network);
    const keyPair = await auth.generateKeyPair(UN, PW);
    const bip32 = await auth.recover(UN, keyPair);
    const address = Address.fromPrivKey(bip32.derive(PATH).privKey);
    const owner = new Owner(bip32);
    lock = address.toTxOutScript().toBuffer().toString('base64');
    console.log("LOCK", lock);
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
        for (let input of signTxn.inputs) {
            owner.signVin(tx, input.vin, input.script, input.sats, 'm/6715768')
        }
        console.log('SIGNED:', tx.toHex());
        return tx.toBuffer().toString('base64');
    }

    async function doActions(actions: any[]) {
        console.log("REQ:", actions);

        let data = await emit('cryptofights/actions', JSON.stringify(actions));

        const signTx = JSON.parse(data);
        console.log("SignTxn:", signTx);
        const signedTx = sign(signTx);
        // console.log("Signed Txn:", signedTx);

        return JSON.parse(await emit('cryptofights/broadcast', signedTx));
    }

    async function getInstances(kind: string) {
        const data = await emit('cryptofights/instances', JSON.stringify({
            lock,
            kind,
        }));
        return JSON.parse(data);
    }

    async function load(outpoint: string) {
        const data = await emit('cryptofights/instance', outpoint);
        return JSON.parse(data);
    }

    async function getTokens() {
        let txos = await getInstances(services.token.contract);

        return Promise.all(txos.map(t => load(t.outpoint)));
    }

    socket.on('connect', async () => {
        let data: any;
        // await deploy('../factory/fyx.wasm.gz', 'factory');
        // await deploy('../token/fyx.wasm.gz', 'token');
        // await deploy('../notes/fyx.wasm.gz', 'notes');

        services = JSON.parse(await emit('cryptofights/services', '{}'));
        console.log('SERVICES:', services);

        // const result = await doActions([{
        //     action: actions.Mint,
        //     service: 'notes',
        //     method: '',
        //     callData: Buffer.from(JSON.stringify({
        //         note: 'This is a note.',
        //         lock
        //     }), 'utf8').toString('base64')
        // }]);
        // console.log('RESULT:', result)

        // const notes = await getInstances(services.notes.contract);
        // console.log("Notes TXOs:", notes);

        // await doActions([{
        //     action: actions.Call,
        //     outpoint: services.token.contract,
        //     service: 'factory',
        //     method: 'Init',
        //     callData: Buffer.from(JSON.stringify({
        //         supply: 1000000000,
        //         lock
        //     }), 'utf8').toString('base64')
        // }]);


        // let tokens = await getTokens(lock);
        // tokens.forEach(t => console.log('Token:', t.outpoint, parseInt(Buffer.from(t.store, 'base64').toString('hex'), 16)))

        // const shruggr = JSON.parse(await emit('cryptofights/user/shruggr', '{}'));
        // const bip32 = Bip32.fromString(shruggr.xpub);
        // const dest = Address.fromPubKey(bip32.derive(PATH).pubKey).toTxOutScript().toBuffer().toString('base64')

        // const [token] = tokens;

        // await doActions( [{
        //     action: actions.Call,
        //     outpoint: token.outpoint,
        //     service: 'token',
        //     method: 'Send',
        //     callData: Buffer.from(JSON.stringify({
        //         sends: [{
        //             to: lock,
        //             amount: 500000000
        //         }]
        //     }), 'utf8').toString('base64')
        // }]);


        // tokens = await getTokens(lock);
        // tokens.forEach(t => console.log('Token:', t.outpoint, parseInt(Buffer.from(t.store, 'base64').toString('hex'), 16)))
    });

    socket.onAny((event, payload) => {
        console.log('Received', event, payload, "\n");
    });

    await new Promise(r => setTimeout(() => r(true), 600000));
}

main().catch((e) => { console.error(e.message); process.exit(1) })