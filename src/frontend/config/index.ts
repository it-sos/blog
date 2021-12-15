export default {
    "development": {
        baseUrl: '/api' // 测试接口域名
    },
    "production": {
        baseUrl: '//www.demo.com/api' // 正式接口域名
    }
}

/**
 * @description 是否在打包环境下，开启打包的分析可视化图
 */
export const VITE_APP_VISUALIZER = false;

/**
 * @description 是否在打包环境下，去除console.log
 */
export const VITE_APP_CONSOLE = true;

/**
 * @description 打包环境下，删除debugger
 */
export const VITE_APP_DEBUGGER = true;

/**
 * @description 打包环境下是否生成source map 文件
 */
export const VITE_APP_SOURCEMAP = false;

/**
 * @description 开发端口
 */
export const VITE_APP_PORT = 3000;

/**
 * @description 公共基础路径
 */
export const VITE_APP_BASE = '/';

/**
 * @description 是否自动在浏览器中打开应用程序
 */
export const VITE_APP_OPEN = true;