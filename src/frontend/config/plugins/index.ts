import legacy from '@vitejs/plugin-legacy';
import vue from '@vitejs/plugin-vue';
import * as path from 'path';
import AutoImport from 'unplugin-auto-import/vite';
import { Plugin } from 'vite';

import {
    ElementPlusResolver
} from 'unplugin-vue-components/resolvers';
import Components from 'unplugin-vue-components/vite';

import { viteVConsole } from 'vite-plugin-vconsole';
import { VITE_APP_VISUALIZER } from '../index';
import configVisualizerConfig from './visualizer';
export default function createVitePlugins() {
    const vitePlugins: (Plugin | Plugin[])[] = [
        vue(),
        AutoImport({
            // targets to transform
            include: [
                /\.[tj]sx?$/, // .ts, .tsx, .js, .jsx
                /\.vue$/, /\.vue\?vue/, // .vue
                /\.md$/, // .md
            ],

            // global imports to register
            imports: [
                // presets
                'vue',
                'vue-router',
                // custom
                {
                    '@vueuse/core': [
                        // named imports
                        'useMouse', // import { useMouse } from '@vueuse/core',
                        // alias
                        // ['useFetch', 'useMyFetch'], // import { useFetch as useMyFetch } from '@vueuse/core',
                    ],
                    // 'axios': [
                    //     // default imports
                    //     ['default', 'axios'], // import { default as axios } from 'axios',
                    // ],
                    // '[package-name]': [
                    //     '[import-names]',
                    //     // alias
                    //     ['[from]', '[alias]'],
                    // ],
                },
                // example type import
                {
                    from: 'vue-router',
                    imports: ['RouteLocationRaw'],
                    type: true,
                },
            ],
            // Enable auto import by filename for default module exports under directories
            defaultExportByFilename: false,

            // Auto import for module exports under directories
            // by default it only scan one level of modules under the directory
            dirs: [
                // './hooks',
                // './composables' // only root modules
                'src/components/**', // all nested modules
                'src/types/**', // all nested modules
                'src/common/**', // all nested modules
                'src/routes/**', // all nested modules
                'src/store/**', // all nested modules
                'src/utils/**', // all nested modules
                'src/global/**', // all nested modules
            ],

            // Filepath to generate corresponding .d.ts file.
            // Defaults to './auto-imports.d.ts' when `typescript` is installed locally.
            // Set `false` to disable.
            dts: './auto-imports.d.ts',

            // Cache the result of resolving, across multiple vite builds.
            // A custom path is supported.
            // When set to `true`, the cache will be stored in `node_modules/.cache/unplugin-auto-import.json`.
            cache: false,

            // Auto import inside Vue template
            // see https://github.com/unjs/unimport/pull/15 and https://github.com/unjs/unimport/pull/72
            vueTemplate: true,

            // Custom resolvers, compatible with `unplugin-vue-components`
            // see https://github.com/antfu/unplugin-auto-import/pull/23/
            resolvers: [ElementPlusResolver()],

            // Inject the imports at the end of other imports
            injectAtEnd: true,

            // Generate corresponding .eslintrc-auto-import.json file.
            // eslint globals Docs - https://eslint.org/docs/user-guide/configuring/language-options#specifying-globals
            eslintrc: {
                enabled: false, // Default `false`
                filepath: './.eslintrc-auto-import.json', // Default `./.eslintrc-auto-import.json`
                globalsPropValue: true, // Default `true`, (true | false | 'readonly' | 'readable' | 'writable' | 'writeable')
            },
        }),
        Components({
            // ui库解析器，也可以自定义
            resolvers: [
                ElementPlusResolver(),
                // AntDesignVueResolver(),
                // VantResolver(),
                // HeadlessUiResolver(),
                // ElementUiResolver()
            ]
        }),
        viteVConsole({
            entry: path.resolve('src/main.ts'),
            localEnabled: true,
            enabled: false,
            config: {
                maxLogNumber: 1000,
                theme: 'dark'
            }
        }),
        legacy({
            // 不考虑ie的兼容性和老的vivo/锤子/荣耀等国产机型，则可以使用下面
            // targets: ['defaults', 'not IE 11', '> 0.25%, not dead'],
            // 如果考虑上面说的兼容性问题，使用下面配置
            targets: ['ie >= 11'],
            additionalLegacyPolyfills: ['regenerator-runtime/runtime'],

            // 下面可以根据情况选用
            renderLegacyChunks: true,
            polyfills: [
                'es.symbol',
                'es.array.filter',
                'es.promise',
                'es.promise.finally',
                'es/map',
                'es/set',
                'es.array.for-each',
                'es.object.define-properties',
                'es.object.define-property',
                'es.object.get-own-property-descriptor',
                'es.object.get-own-property-descriptors',
                'es.object.keys',
                'es.object.to-string',
                'web.dom-collections.for-each',
                'esnext.global-this',
                'esnext.string.match-all'
            ],
            modernPolyfills: ['es.promise.finally']
        }),
        // importToCDN({
        //     modules: [
        //         //https://unpkg.com/vue@3.2.29/dist/vue.global.js
        //         // autoComplete('vue'),
        //         // autoComplete('axios'),
        //         // autoComplete('crypto-js'),
        //         // {
        //         //     name: 'vue',
        //         //     var: 'Vue',
        //         //     path: `https://unpkg.com/vue@3.2.26/dist/vue.global.prod.js`,
        //         // },
        //         // {
        //         //     name: 'axios',
        //         //     var: 'axios',
        //         //     path: `https://unpkg.com/axios@0.27.2/dist/axios.min.js`,
        //         // },
        //         // {
        //         //     name: 'crypto-js',
        //         //     var: 'CryptoJs',
        //         //     path: `https://unpkg.com/crypto-js@4.1.1/crypto-js.js`,
        //         // },

        //         // autoComplete('@vueuse/shared'),
        //         // autoComplete('@vueuse/core'),
        //         // {
        //         //     name: 'element-plus',
        //         //     var: 'ElementPlus',
        //         //     path: `https://unpkg.com/element-plus@1.2.0-beta.6/dist/index.full.min.js`,
        //         //     css: `https://unpkg.com/element-plus@1.2.0-beta.6/dist/index.css`
        //         // },
        //         // {
        //         //     name: 'vue-router',
        //         //     var: 'VueRouter',
        //         //     path: `https://unpkg.com/vue-router@4.0.12`,
        //         // },
        //         // {
        //         //     name: 'vuex',
        //         //     var: 'Vuex',
        //         //     path: `https://unpkg.com/vuex@4.0.0/dist/vuex.global.js`,
        //         // },
        //     ]
        // }),
    ];
    VITE_APP_VISUALIZER && vitePlugins.push(configVisualizerConfig());
    return vitePlugins;
}