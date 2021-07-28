<template>
  <el-input v-model="input" @keyup.delete="deletes()" @keydown.enter="search()" placeholder="请输入内容">
    <template #append>
      <el-button type="primary" icon="el-icon-search" v-on:click="search()" :loading="false">搜索</el-button>
    </template>
  </el-input>
</template>

<script>
import {defineComponent, ref} from 'vue'

export default defineComponent({
  setup() {
    return {
      input: ref(''),
    }
  },
  created() {
    this.$watch(
        () => this.$route.params.keyword,
        () => {
          var keyword = this.$route.params.keyword
          if (keyword) {
            this.input = keyword
          }
        },
        // 组件创建完后获取数据，
        // 此时 data 已经被 observed 了
        { immediate: true }
    )
  },
  methods: {
    deletes() {
      if (this.input == "") {
        document.title = "yupengSir 首页"
        this.$router.push('/')
      }
    },
    search() {
      document.title = "搜索结果：" + this.input
      this.$router.push('/'+encodeURIComponent(this.input))
    }
  }
})
</script>

<style>

</style>