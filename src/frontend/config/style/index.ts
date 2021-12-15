import { CSSOptions } from 'vite';
const cssOption: CSSOptions = {
    preprocessorOptions: {
        less: {
            javascriptEnabled: true,
        },
        scss: {
            additionalData: '@import "./src/assets/scss/varible.scss";',
        },
    },
};
export default cssOption;