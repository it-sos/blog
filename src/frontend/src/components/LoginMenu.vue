<!--<template>-->
<!--  <el-dropdown :hide-on-click="false">-->
<!--    <span class="el-dropdown-link">-->
<!--      Dropdown List<el-icon class="el-icon&#45;&#45;right"><arrow-down /></el-icon>-->
<!--    </span>-->
<!--    <template #dropdown>-->
<!--      <el-dropdown-menu>-->
<!--        <el-dropdown-item>Action 1</el-dropdown-item>-->
<!--        <el-dropdown-item>Action 2</el-dropdown-item>-->
<!--        <el-dropdown-item>Action 3</el-dropdown-item>-->
<!--        <el-dropdown-item disabled>Action 4</el-dropdown-item>-->
<!--        <el-dropdown-item divided>Action 5</el-dropdown-item>-->
<!--      </el-dropdown-menu>-->
<!--    </template>-->
<!--  </el-dropdown>-->
<!--</template>-->
<template>
  <el-dropdown @visible-change="visibleChange">
  <span class="el-dropdown-link">
      <el-icon class="el-icon--right" style="font-size:32px;"><Fold /></el-icon>
    <el-icon style="vertical-align: middle;">
    </el-icon>
  </span>
    <template #dropdown>
      <el-dropdown-menu class="login-menu">
        <el-dropdown-item :icon="Plus">
          <router-link to="/e/">创建文章</router-link>
        </el-dropdown-item>
        <el-dropdown-item :icon="el-icon-edit" v-if="article.id > 0">
          <router-link :to="'/e/'+article.id">编辑文章</router-link>
        </el-dropdown-item>
        <el-dropdown-item :icon="el-icon-lock" disabled>内容加密（未激活）</el-dropdown-item>
        <el-dropdown-item divided>小鹿 [退出]</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script lang="ts">
import {defineComponent, inject, ref} from 'vue'
import {Fold, Plus} from "@element-plus/icons-vue"

export default defineComponent({
  name: "LoginMenu",
  components: {
    Fold,
    Plus,
  },
  setup() {
    let article = inject("article-id", {id: ref<number>()})
    return {
      article,
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
