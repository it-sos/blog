module.exports = {
    runtimeCompiler: true,
    alias: {
        'vue': "vue/dist/vue.esm-bundler.js"
    },
    cssPreprocessOptions: {
        scss: {
            // additionalData: '@import "./src/assets/scss/style.scss";' // 添加公共样式
        }
    }
}