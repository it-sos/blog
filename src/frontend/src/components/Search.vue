<template>
  <el-input v-model="inputSearch" @keyup.delete="deletes()" @keydown.enter="search()" placeholder="请输入内容">
    <template #append>
      <el-button type="primary" icon="el-icon-search" v-on:click="search()" :loading="false">搜索</el-button>
    </template>
  </el-input>
</template>

<script lang="ts">
import {defineComponent, ref} from 'vue'

export default defineComponent({
  setup() {
    return {
      inputSearch: ref("")
    }
  },
  created() {
    this.$watch(
        () => this.$route.params.keyword,
        () => {
          if (typeof this.$route.params.keyword == "undefined") {
            this.inputSearch = ""
          } else {
            this.inputSearch = this.$route.params.keyword.toString()
          }
        },
        // 组件创建完后获取数据，
        // 此时 data 已经被 observed 了
        { immediate: true }
    )
  },
  methods: {
    deletes() {
      if (this.inputSearch === '') {
        document.title = "YuPengSir Blog"
        this.$router.push('/')
      }
    },
    search() {
      document.title = "搜索结果：" + this.inputSearch
      this.$router.push('/'+encodeURIComponent(this.inputSearch))
    }
  }
})
</script>

<style>
</style>