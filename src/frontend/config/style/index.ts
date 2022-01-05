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
    postcss: {
        plugins: [
            {
                postcssPlugin: 'internal:charset-removal',
                AtRule: {
                    charset: (atRule) => {
                        if (atRule.name === 'charset') {
                            atRule.remove();
                        }
                    }
                }
            }
        ]
    }
};
export default cssOption;