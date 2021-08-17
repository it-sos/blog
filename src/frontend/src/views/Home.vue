<template>
  <el-container v-infinite-scroll="load" class="infinite-list" infinite-scroll-immediate="false">
    <el-main>
      <div class="box" v-bind:key="idx" v-for="(art, idx) in article">
        <div class="title">
          <el-link :href="'/a/'+escape(art.article.title)"><h2
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
        <el-link :href="'/a/'+escape(r.title)">{{ r.title }}（{{ r.access_times }}）</el-link>
      </div>
    </el-card>
  </el-aside>
</template>

<script lang="ts">

import {defineComponent} from 'vue'

export default defineComponent({
  data() {
    let page: number = 0
    let size: number = 10
    let article: any = []
    let rank: any = []
    let noMore: boolean = false
    let loading: boolean = false
    let errorMsg: boolean = false
    return {page, size, article, rank, noMore, loading, errorMsg}
  },

  created() {
    this.$watch(
        () => this.$route.params.keyword,
        () => {
          this.defaults()
          this.load()
        },
        // 组件创建完后获取数据，
        // 此时 data 已经被 observed 了
        {immediate: true}
    )
  },

  mounted(): void {
    document.title = "yupeng-sir-blog"
    this.ranks()
  },

  methods: {
    ranks(): void {
      this.$http.get('/article/rank').then((response) => {
        this.rank = response.data
      }).catch((error) => {
        this.loading = false
        console.log(error)
      })
    },

    defaults(): void {
      this.errorMsg = false
      this.noMore = false
      this.page = 0
      this.article = []
    },

    escape(str: string): string {
      return encodeURIComponent(str)
    },

    unscape(str: string): string {
      return decodeURIComponent(str)
    },

    load(): void {
      if (this.noMore || this.errorMsg) {
        return
      }
      this.loading = true
      this.page++

      var keyword = this.$route.params.keyword
      if (keyword) {
        keyword = this.unscape(keyword.toString())
      }

      this.$http.get('/article/list', {
        params: {
          "page": this.page,
          "size": this.size,
          "keyword": keyword
        }
      }).then((response) => {
        this.loading = false
        if (response.data.length === 0) {
          this.noMore = true
          return
        }
        response.data.forEach((v: any) => {
          this.article.push(v)
        })
      }).catch((error) => {
        console.log(error)
        this.loading = false
        this.errorMsg = true
      })
    }
  },
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