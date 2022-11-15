process.env.NETWORK='testnet';
import {Address} from 'bsv';
import fs from 'fs';
import { io } from 'socket.io-client';
import { AuthService } from '@fyxgaming/lib/dist/auth-service';
import { SignedMessage } from '@fyxgaming/lib/dist/signed-message';
import {Owner } from './lib/owner';
import { emit } from 'process';

const API = 'http://bitcoin-dev.aws.kronoverse.io:8081';
const AUTH = 'https://dev.api.cryptofights.io'
const network = 'test';
const UN = 'shruggr1';
const PW = 'test1234';

async function main() {
    const auth = new AuthService(AUTH, network);
    const keyPair = await auth.generateKeyPair(UN, PW);
    const bip32 = await auth.recover(UN, keyPair);
    const address = Address.fromPrivKey(bip32.derive('m/6715768').privKey);
    const socket = io(API, {
        auth: async (cb) => {
            const payload = JSON.stringify(new SignedMessage({
                subject: 'Auth'
            }, UN, keyPair));
            cb({payload});
        }
    });

    async function emit(event: string, body: any): Promise<any> {
        let resp = await new Promise<any>(r => {
            socket.emit(event, body, r)
        });
        const {status, data} = JSON.parse(resp);
        if(status >= 300) {
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
        let {uploads:[upload]} = JSON.parse(data);
        // console.log('RESP:', resp);

        data = await emit('cryptofights/deploy', upload.outpoint);
        const ctx = JSON.parse(data);

        data = await emit(`cryptofights/install`, JSON.stringify({
            name,
            contract: ctx.ctxs[0].instance.outpoint,
        }));

        return JSON.parse(data);
    }

    socket.on('connect', async () => {
        let data = await deploy('../factory/fyx.wasm.gz', 'factory');
        console.log('FACTORY:', data);

        data = await deploy('../token/fyx.wasm.gz', 'token');
        console.log('TOKEN:', data);

        
    })

    socket.onAny((event, payload) => {
        console.log('Received', event, payload, "\n");
    });
    
    await new Promise(r => setTimeout(() => r(true), 600000));
}

main().catch((e) => {console.error(e); process.exit(1)})