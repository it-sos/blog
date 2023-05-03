<template>
  <el-input v-model="inputSearch" @keyup.delete="deletes()" class="input-with-select" @keydown.enter="search()" placeholder="请输入内容">
    <template #prepend>
      <el-select v-model="searchType" style="width: 110px" placeholder="搜索类型">
        <el-option label="全站搜索" value="1" />
        <el-option v-show="isLogin" label="智能搜索" value="2" />
      </el-select>
    </template>
    <template #append>
      <el-button :icon="Search" v-on:click="search()" :loading="false" />
    </template>
  </el-input>
</template>

<script lang="ts" setup>
import { Search } from '@element-plus/icons-vue';

const store = useStore()
const inputSearch = ref<string>("")
const searchType = ref<string>(SearchSite)
const router = useRouter()
const route = useRoute()
watch(() => route.params.keyword, () => {
  if (typeof route.params.keyword == "undefined") {
    inputSearch.value = ""
  } else {
    inputSearch.value = route.params.keyword.toString()
  }
})

const deletes = () => {
  if (inputSearch.value === '') {
    router.push('/')
  }
}

const search = () => {
  document.title = "搜索结果：" + inputSearch.value
  let keyword = encodeURIComponent(inputSearch.value)
  let url = ''
  if (searchType.value == SearchSite) {
    url = '/'
  } else {
    url = '/ai/'
  }
  router.push(url + keyword)
}

const isLogin = computed(() => store.getters.getAccount() != null)

</script>

<style>
.input-with-select .el-input-group__prepend {
  background-color: var(--el-fill-color-blank);
}
</style>