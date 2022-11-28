import { Bn, KeyPair, Script, Sig, TxOut } from 'bsv';

export class Owner {
    public keyPairs = new Map<string, any>();
    public address: string;
    public pubkey: string;
    public scripthash: string;

    constructor(private bip32) {}


    // async sign(rawtx: string, parents: { satoshis: number, script: string }[], locks: any[]): Promise<string> {
    //     const tx = Tx.fromHex(rawtx);

    //     await Promise.all(tx.txIns.map(async (txIn, i) => {
    //         const lockScript = Script.fromHex(parents[i].script);
    //         const txOut = TxOut.fromProperties(new Bn(parents[i].satoshis), lockScript);
    //         const keyPair = this.keyPairs.get(txOut.script.toHex());
    //         if (!keyPair) return;
    //         const sig = await tx.asyncSign(keyPair, Sig.SIGHASH_ALL | Sig.SIGHASH_FORKID, i, txOut.script, txOut.valueBn);
    //         txIn.setScript(new Script().writeBuffer(sig.toTxFormat()).writeBuffer(keyPair.pubKey.toBuffer()));
    //     }));

    //     return tx.toHex();
    // }

    async signVin(tx, vin, script, satoshis, path) {
        const lockScript = Script.fromBuffer(Buffer.from(script, 'base64'));
        const txOut = TxOut.fromProperties(new Bn(satoshis), lockScript);
        const keyPair = KeyPair.fromPrivKey(this.bip32.derive(path).privKey);
        const sig = tx.sign(keyPair, Sig.SIGHASH_ALL | Sig.SIGHASH_FORKID, vin, txOut.script, txOut.valueBn);
        tx.txIns[vin].setScript(new Script().writeBuffer(sig.toTxFormat()).writeBuffer(keyPair.pubKey.toBuffer()));
    }
}