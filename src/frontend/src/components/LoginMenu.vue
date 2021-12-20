<template>
  <el-dropdown>
  <span class="el-dropdown-link">
    <el-icon class="el-icon--right" style="font-size:32px;"><Operation/></el-icon>
    <el-icon style="vertical-align: middle;"></el-icon>
  </span>
    <template #dropdown>
      <el-dropdown-menu class="login-menu">
        <el-dropdown-item :icon="CirclePlus">
          <router-link to="/e/">创建文章</router-link>
        </el-dropdown-item>
        <el-dropdown-item :icon="Edit" v-if="article.id > 0">
          <router-link :to="'/e/'+article.id">编辑文章</router-link>
        </el-dropdown-item>
        <el-dropdown-item :icon="Key" disabled>内容加密（未激活）</el-dropdown-item>
        <el-dropdown-item divided>小鹿 [退出]</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script lang="ts">
import {defineComponent, inject, ref} from 'vue'
import {CirclePlus, Edit, Key, Operation} from "@element-plus/icons-vue"


export default defineComponent({
  name: "LoginMenu",
  components: {
    Operation,
  },
  setup() {
    let article = inject("article-id", {id: ref<number>()})
    return {
      article,
      Key,
      Edit,
      CirclePlus,
    }
  }
})
</script>

<style scoped>
.login-menu a {
  color: #606266;
  text-decoration: none;
}
</style>
