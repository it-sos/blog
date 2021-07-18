<template>
  <el-input v-model="input" placeholder="请输入内容">
    <template #append>
    <el-button type="primary" icon="el-icon-search" :loading="false">搜索</el-button>
    </template>
  </el-input>
</template>

<script>
import { defineComponent, ref } from 'vue'

export default defineComponent ({
  setup() {
    return {
      input: ref(''),
    }
  }
})
</script>

<style>

</style>