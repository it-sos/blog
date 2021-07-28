<template>
  <div class="common-layout">
    <el-container style="min-height:700px;">
      <el-header>
        <el-row type="flex" class="row-bg" justify="space-between"  align="middle">
          <el-col :span="6">
            <router-link to="/">
              <el-image style="width: 104px; height: 24px; border-radius: 4px;" src="/logo.png" fit="fill"></el-image>
            </router-link>
          </el-col>
          <el-col :span="6" align="center"><Search @search="search" /></el-col>
          <el-col :span="6" align="right"><LoginMenu/></el-col>
        </el-row>
      </el-header>
      <el-container>
        <router-view></router-view>
      </el-container>
      <el-footer style="text-align: center;font-size: 12px; padding:30px 0 60px 0;">Copyright &copy; 鹏sir笔记</el-footer>
    </el-container>
  </div>
<!--  <p>-->
<!--    <router-link to="/">首页</router-link>-->
<!--    <router-link to="/about">关于</router-link>-->
<!--  </p>-->
<!--  <router-view></router-view>-->
</template>

<script>

import Search from "./components/Search.vue"
import LoginMenu from "./components/LoginMenu.vue"

export default {
  name: 'App',
  components: {Search, LoginMenu},
  methods: {
    search(txt) {
      console.log(txt)
    }
  }
}
</script>
