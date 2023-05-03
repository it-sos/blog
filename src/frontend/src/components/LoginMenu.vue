<template>
  <el-dropdown>
  <span class="el-dropdown-link">
    <el-icon class="el-icon--right" style="font-size:32px;"><Operation/></el-icon>
    <el-icon style="vertical-align: middle;"></el-icon>
  </span>
    <template #dropdown>
      <el-dropdown-menu class="login-menu" v-if="account">
        <el-dropdown-item :icon="CirclePlus">
          <router-link to="/e/">创建文章</router-link>
        </el-dropdown-item>
        <el-dropdown-item :icon="Edit" v-if="articleId > 0 && !isBackend">
          <router-link :to="'/e/'+articleId">编辑文章</router-link>
        </el-dropdown-item>
        <el-dropdown-item :icon="Delete" v-if="articleId > 0 && isBackend">
          <span @click="delConfirm(articleId)">删除文章</span>
        </el-dropdown-item>
        <el-dropdown-item :icon="Key" disabled>内容加密（未激活）</el-dropdown-item>
        <el-tooltip content="退出登陆" effect="dark" placement="left">
          <el-dropdown-item :icon="CircleClose" @click="logout" divided>你好，{{ account }}</el-dropdown-item>
        </el-tooltip>
      </el-dropdown-menu>
      <el-dropdown-menu class="login-menu" v-else>
        <el-dropdown-item :icon="Lock">
          <router-link to="/login">登录</router-link>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>
<!--<el-icon><circle-close-filled /></el-icon>-->
<script lang="ts" setup>

import { CircleClose, CirclePlus, Delete, Edit, Key, Lock, Operation } from "@element-plus/icons-vue";
import { ElMessageBox } from "element-plus";
import { computed } from 'vue';
import { useRouter } from "vue-router";
import { useStore } from "../store/store";

const store = useStore()
const router = useRouter()

let logout = () => {
  ElMessageBox.confirm('确认退出登陆吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
      .then(() => {
        store.commit('logout')
        router.push({path: '/login'})
      })
      .catch(() => {
      });
}

let delConfirm = (id: number) => {
  store.commit('article/remove', {id: id})
}

let account = computed(() => store.getters.getAccount())
let articleId = computed(() => store.getters.getArticleId())
let isBackend = computed(() => store.getters.getIsBackend())

</script>

<style scoped>
.login-menu a {
  color: #606266;
  text-decoration: none;
}
</style>
