<template>
  <el-container v-infinite-scroll="load" class="infinite-list" infinite-scroll-immediate="false">
    <el-main>
      <div class="box" v-bind:key="idx" v-for="(art, idx) in article">
        <div class="title">
          <el-link :href="'/a/'+encodeURIComponent(art.article.title)"><h2
              v-html="art.article.title_match ? art.article.title_match : art.article.title"></h2></el-link>
          <el-tag effect="plain" type="danger" size="mini">{{ art.duration }}</el-tag>
        </div>
        <div class="description" v-html="art.article.intro_match ? art.article.intro_match : art.article.intro"></div>
        <el-row class="link">
          <el-col :span="12">
            <p v-bind:key="idx" v-for="(topic, idx) in art.topics">来自：{{ topic }} 专题</p>
          </el-col>
          <el-col :span="12" align="right" class="tag">
            <el-tag effect="plain" v-bind:key="idx" v-for="(tag, idx) in art.tags">{{ tag }}</el-tag>
          </el-col>
        </el-row>
        <!-- 分割线 -->
        <el-divider></el-divider>
      </div>
      <p class="load" v-if="loading">加载中...</p>
      <el-empty v-if="noMore" description="木有了"></el-empty>
      <el-result v-if="errorMsg" icon="error" title="错误提示" subTitle="内部错误，稍后再试试">
      </el-result>
    </el-main>
  </el-container>
  <el-aside class="hidden-xs-only">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>访问前50文章</span>
        </div>
      </template>
      <div v-bind:key="idx" v-for="(r, idx) in rank" class="text item">
        <el-link :href="'/a/'+encodeURIComponent(r.title)">{{ r.title }}（{{ r.access_times }}）</el-link>
      </div>
    </el-card>
  </el-aside>
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
      document.title = "YU鹏sir笔记"
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
