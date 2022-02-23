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
        <el-dropdown-item :icon="Edit" v-if="article.id > 0">
          <router-link :to="'/e/'+article.id">编辑文章</router-link>
        </el-dropdown-item>
        <el-dropdown-item :icon="Key" disabled>内容加密（未激活）</el-dropdown-item>
        <el-dropdown-item divided>{{account}} [退出]</el-dropdown-item>
      </el-dropdown-menu>
      <el-dropdown-menu class="login-menu" v-else>
        <el-dropdown-item :icon="Lock">
          <router-link to="/login">登录</router-link>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script lang="ts">
import {computed, defineComponent, inject, ref} from 'vue'
import {CirclePlus, Edit, Lock, Operation} from "@element-plus/icons-vue"
import {useStore} from "../store/store";

export default defineComponent({
  name: "LoginMenu",
  components: {
    Operation,
  },
  setup() {
    const store = useStore()
    let article = inject("article-id", {id: ref<number>()})
    return {
      article,
      Lock,
      Edit,
      CirclePlus,
      account: computed(() => store.getters.getAccount())
    }
  },
})
</script>

<style scoped>
.login-menu a {
  color: #606266;
  text-decoration: none;
}
</style>
