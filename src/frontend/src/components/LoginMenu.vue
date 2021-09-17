<template>
  <el-dropdown @visible-change="visibleChange">
  <span class="el-dropdown-link">
    <i :class="[active ? 'el-icon-s-unfold' : 'el-icon-s-fold']" style="font-size: 32px;"></i>
  </span>
    <template #dropdown>
      <el-dropdown-menu class="login-menu">
        <el-dropdown-item icon="el-icon-document-add"><router-link to="/e">创建文章</router-link></el-dropdown-item>
        <el-dropdown-item icon="el-icon-edit" disabled>编辑文章</el-dropdown-item>
        <el-dropdown-item icon="el-icon-lock" disabled>内容加密（未激活）</el-dropdown-item>
        <el-dropdown-item divided>小鹿 [退出]</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
export default defineComponent({
  name: "LoginMenu",
  setup() {
    let active :boolean = false
    const visibleChange = (isFold: boolean) => {
      active = isFold
    }
    return {
      active,
      visibleChange,
    }
  }
})
</script>

<style scoped>
.login-menu a {
  color:#606266;
  text-decoration: none;
}
</style>
