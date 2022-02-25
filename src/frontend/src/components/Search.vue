<template>
  <el-input v-model="inputSearch" @keyup.delete="deletes()" @keydown.enter="search()" placeholder="请输入内容">
    <template #append>
      <el-button type="primary" icon="el-icon-search" v-on:click="search()" :loading="false">搜索</el-button>
    </template>
  </el-input>
</template>

<script lang="ts">
import {defineComponent, ref, watch} from 'vue'
import {useRoute, useRouter} from "vue-router";

export default defineComponent({
  setup() {
    let inputSearch = ref<string>("")
    const router = useRouter()
    const route = useRoute()
    watch(() => route.params.keyword, () => {
      if (typeof route.params.keyword == "undefined") {
        inputSearch.value = ""
      } else {
        inputSearch.value = route.params.keyword.toString()
      }
    })

    const deletes = () => {
      if (inputSearch.value === '') {
        router.push('/')
      }
    }

    const search = () => {
      document.title = "搜索结果：" + inputSearch.value
      router.push('/' + encodeURIComponent(inputSearch.value))
    }

    return {
      deletes,
      search,
      inputSearch
    }
  }
})
</script>

<style>
</style>