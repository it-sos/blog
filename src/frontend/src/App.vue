<template>
  <div class="common-layout">
    <el-container style="min-height: 600px">
      <el-header>
        <el-row
          type="flex"
          class="row-bg"
          justify="space-between"
          align="middle"
        >
          <el-col :span="6">
            <router-link to="/">
              <el-image
                style="width: 105px; height: 30px; border-radius: 4px"
                src="/logo.png"
                fit="fill"
              ></el-image>
            </router-link>
          </el-col>
          <el-col :span="6" align="center">
            <Search />
          </el-col>
          <el-col :span="6" align="right">
            <LoginMenu />
          </el-col>
        </el-row>
      </el-header>
      <el-container>
        <router-view></router-view>
      </el-container>
      <el-footer
        style="text-align: center; font-size: 12px; padding: 30px 0 60px 0"
        >Copyright &copy; IT.SOS 技术笔记 ICP证:
        <a style="text-decoration:none;color:#e6a23c;" href="https://beian.miit.gov.cn/" target="_blank"
          >京ICP备17025733号-8</a
        >
      </el-footer>
    </el-container>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

import Search from "./components/Search.vue";
import LoginMenu from "./components/LoginMenu.vue";

export default defineComponent({
  components: { Search, LoginMenu },
  setup() {},
});
</script>
<style>
/* 屏幕小于768px */
@media screen and (max-width: 768px) {
  .el-input-group__append {
    display: none;
  }
}
</style>
