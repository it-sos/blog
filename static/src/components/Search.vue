<template>
  <el-input v-model="input" @keyup.delete="deletes()" placeholder="请输入内容">
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
      timer: null,
    }
  },
  methods: {
    deletes() {
      if (this.input == "") {
        this.search()
      }
    },
    search() {
      this.$emit("search", this.input)
    }
  }
})
</script>

<style>

</style>