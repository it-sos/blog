// @ts-ignore
import md5 from "js-md5"
import { JSEncrypt } from 'jsencrypt'

export function randomString(e: number) {
    e = e || 32;
    var t = "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz012345678",
        a = t.length,
        n = "";
    for (var i = 0; i < e; i++) n += t.charAt(Math.floor(Math.random() * a));
    return n
}

export function sign(ts: string, data: any, token: string) {
    let p, r, a, e, n, d, b;
    p = "-----BEGIN PUBLIC KEY-----\n" +
        "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDEnH1ObJ90TdO+mdQLAwsBlehH\n" +
        "E8/Cgl6AO/ojo5FB8t0AkN+gcjkkdmD5j8QnEBWxdsZjqGRJkVcwp23lkOMtA88N\n" +
        "qUQ4CncuGYe1um/Xm1xW5PwaJOq+FbjHBoNFI01uKYpVJG+k3MaTT6f84J5vYWMd\n" +
        "X+lAQZjNvN8ooPW0JwIDAQAB\n" +
        "-----END PUBLIC KEY-----"
    r = /[\":{},\[\]]/g
    a = md5(randomString(32) + token)
    e = new JSEncrypt();
    e.setPublicKey(p);
    n = e.encrypt(a);
    d = JSON.stringify(data)
    data = Array.from(d.replaceAll(r, "")).sort().join("")
    b = md5(token + ts + a + data + n)
    return [n, b]
}
