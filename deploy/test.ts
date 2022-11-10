process.env.NETWORK='testnet';
import fs from 'fs';
import { io } from 'socket.io-client';
import { AuthService } from '@fyxgaming/lib/dist/auth-service';
import { SignedMessage } from '@fyxgaming/lib/dist/signed-message';
import {Owner } from './lib/owner';
import { emit } from 'process';

const API = 'http://bitcoin-dev.aws.kronoverse.io:8081';
const network = 'test';
const UN = 'shruggr';
const PW = 'test1234';

async function main() {
    const auth = new AuthService(API, network);
    const keyPair = await auth.generateKeyPair(UN, PW);
    const socket = io(API, {
        auth: async (cb) => {
            const payload = JSON.stringify(new SignedMessage({
                subject: 'Auth'
            }, UN, keyPair));
            cb({payload});
        }
    });

    async function emit(event, body): Promise<any> {
        let resp = await new Promise<any>(r => {
            socket.emit(event, JSON.stringify(body), r)
        });
        const {status, data} = JSON.parse(resp);
        if(status >= 300) {
            console.error(resp);
            throw new Error("bad-response");
        }
        return data;
    }

    socket.on('connect', async () => {
        let code = fs.readFileSync('../factory/fyx.wasm.gz')
        let data = await emit('cryptofights/upload',{name:'factory.wasm.gz', data: code.toString('base64')});
        console.log('RESP:', data);
        // let {txid} = JSON.parse(data)
        

    })

    socket.onAny((event, payload) => {
        console.log('Received', event, payload, "\n");
    });
    
    await new Promise(r => setTimeout(() => r(true), 600000));
}

main().catch((e) => {console.error(e); process.exit(1)})