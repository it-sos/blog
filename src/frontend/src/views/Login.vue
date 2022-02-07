<template>
</template>

<script lang="ts">

import {defineComponent, inject, onMounted, reactive, ref, toRefs, watch} from 'vue'
import 'element-plus/theme-chalk/display.css'
import router from "../routes";

export default defineComponent({
  setup() {
    const $axios: any = inject('$axios')
    let article = inject("article-id", {id: ref<number>()})
    article.id = ref(0)
    let state = reactive({
      page: 0,
      size: 10,
      article: [],
      rank: [],
      noMore: false,
      loading: false,
      errorMsg: false,
    })

    const defaults = () => {
      document.title = "IT.SOS 技术笔记"
      state.errorMsg = false
      state.noMore = false
      state.page = 0
      state.article = []
    }

    const load = () => {
      if (state.noMore || state.errorMsg) {
        return
      }
      state.loading = true
      state.page++

      let keyword = router.currentRoute.value.params.keyword;
      if (keyword) {
        keyword = decodeURIComponent(keyword.toString())
      }

      $axios.get('/article/list', {
        params: {
          "page": state.page,
          "size": state.size,
          "keyword": keyword
        }
      }).then((response: any) => {
        state.loading = false
        if (response.data.length === 0) {
          state.noMore = true
          return
        }
        response.data.forEach((v: any) => {
          state.article.push(v)
        })
      }).catch((error: any) => {
        console.log(error)
        state.loading = false
        state.errorMsg = true
      })
    }

    watch(() => router.currentRoute.value.params.keyword, () => {
      defaults()
      load()
      ranks()
    })

    let ranks = () => {
      $axios.get('/article/rank').then((response: any) => {
        state.rank = response.data
      }).catch((error: any) => {
        state.loading = false
        console.log(error)
      })
    }

    onMounted(() => {
      defaults()
      load()
      ranks()
    })

    return {
      ...toRefs(state),
      load,
    }
  }
})
</script>

<style scoped>
.load {
  padding: 0px;
  margin: 0px;
  text-align: center;
}

.title a {
  text-decoration: none;
  color: #303133;
  margin-right: 30px;
}

.title h2 {
  display: inline;
}

.title span {
  vertical-align: bottom;
}

.description {
  word-break: break-word;
  padding-top: 30px;
  color: #606266;
}

.link {
  padding-top: 20px;
}

.link p {
  margin: 0;
}

.tag span {
  margin-left: 6px;

}

.box {
  padding-bottom: 30px;
}

.hidden-xs-only {
  padding-top: 20px;
}
</style>
