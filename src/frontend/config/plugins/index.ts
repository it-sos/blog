import { Plugin } from 'vite';
import vue from '@vitejs/plugin-vue'
import { VITE_APP_VISUALIZER } from '../index';
import configVisualizerConfig from './visualizer';
export default function createVitePlugins() {
    const vitePlugins: (Plugin | Plugin[])[] = [
        vue()
    ];
    VITE_APP_VISUALIZER && vitePlugins.push(configVisualizerConfig());
    return vitePlugins;
}