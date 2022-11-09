process.env.NETWORK='testnet';
import { io } from 'socket.io-client';
import { AuthService } from '@fyxgaming/lib/dist/auth-service';
import { SignedMessage } from '@fyxgaming/lib/dist/signed-message';
import {Owner } from './lib/owner';

const API = 'https://dev.api.cryptofights.io';
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

    socket.on('connect', () => {
        console.log('Yo');
    })

    socket.onAny((event, payload) => {
        console.log('Received', event, payload, "\n");
    });
    
    await new Promise(r => setTimeout(() => r(true), 600000));
}

main().catch((e) => {console.error(e); process.exit(1)})