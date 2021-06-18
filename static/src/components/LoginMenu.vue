<template>
  <el-dropdown @visible-change="visibleChange">
  <span class="el-dropdown-link">
    <i :class="[active ? 'el-icon-s-unfold' : 'el-icon-s-fold']" style="font-size: 32px;"></i>
  </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item icon="el-icon-plus">黄金糕</el-dropdown-item>
        <el-dropdown-item>狮子头</el-dropdown-item>
        <el-dropdown-item>螺蛳粉</el-dropdown-item>
        <el-dropdown-item disabled>双皮奶</el-dropdown-item>
        <el-dropdown-item divided>蚵仔煎</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script>
export default {
  name: "LoginMenu",
  data() {
    return {
      active: false
    }
  },
  methods: {
    visibleChange(isFold) {
      if (isFold) {
        this.active = true
      } else {
        this.active = false
      }
    }
  }
}
</script>

<style scoped>

</style>