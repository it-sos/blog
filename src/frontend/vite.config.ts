import {defineConfig} from 'vite'
import createVitePlugins from "./config/plugins";
import {VITE_APP_BASE, VITE_APP_OPEN, VITE_APP_PORT} from "./config";
import proxy from "./config/setupProxy";
import build from "./config/build";
// @ts-ignore
import path from 'path'
import cssOption from "./config/style";

// https://vitejs.dev/config/
// https://cn.vitejs.dev/config/#conditional-config
export default defineConfig(({command, mode}) => {
    return {
        base: VITE_APP_BASE,
        // optimizeDeps: {
        //     include: [
        //         '@/common/tiptap/tiptap-extensions',
        //     ]
        // },
        plugins: createVitePlugins(),
        css: cssOption,
        resolve: {
            alias: {
                // @ts-ignore
                '~': path.resolve(__dirname, './'),
                // @ts-ignore
                '@': path.resolve(__dirname, 'src')
            }
        },
        server: {
            host: true,
            port: VITE_APP_PORT,
            open: VITE_APP_OPEN,
            proxy,
        },
        build
    }
})
