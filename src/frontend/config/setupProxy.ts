import { ProxyOptions } from 'vite';
// 更多请参看：https://cn.vitejs.dev/config/#server-proxy
const proxy: Record<string, string | ProxyOptions> = {
    // 选项写法
    '/api': {
        target: 'http://localhost:8091',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
    },
};
export default proxy;