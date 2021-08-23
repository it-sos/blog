import {App} from "vue";

declare const version = '0.0.1'
declare const install: (app: App, opt: any) => void;
export { version, install };
declare const _default: {
    version: string;
    install: (app: App<any>, opt: any) => void;
};
export default _default;