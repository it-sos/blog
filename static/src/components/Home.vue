<template>
  <el-container v-infinite-scroll="load" class="infinite-list">
    <el-main>
      <div class="box" key="index" v-for="art in article">
        <div class="title">
          <el-link :href="'/a/'+escape(art.article.title)"><h2 v-html="art.article.title_match ? art.article.title_match : art.article.title"></h2></el-link>
          <el-tag effect="plain" type="danger" size="mini">{{ art.duration }}</el-tag>
        </div>
        <div class="description" v-html="art.article.intro_match ? art.article.intro_match : art.article.intro"></div>
        <el-row class="link">
          <el-col :span="12">
            <p v-for="topic in art.topics">来自：{{ topic }} 专题</p>
          </el-col>
          <el-col :span="12" align="right" class="tag">
            <el-tag effect="plain" v-for="tag in art.tags">{{ tag }}</el-tag>
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
      <div v-for="r in rank" class="text item">
        <el-link :href="'/a/'+escape(r.title)">{{ r.title }}（{{ r.access_times }}）</el-link>
      </div>
    </el-card>
  </el-aside>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      page: 0,
      size: 10,
      article: [],
      rank: [],
      noMore: false,
      loading: false,
      errorMsg: false
    }
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
        { immediate: true }
    )
  },
  mounted() {
    document.title = "YuPengSir Blog"
    this.ranks()
  },
  methods: {
    defaults() {
      this.errorMsg = false
      this.noMore = false
      this.page = 0
      this.article = []
    },
    escape(str) {
      return encodeURIComponent(str)
    },
    unscape(str) {
      return decodeURIComponent(str)
    },
    ranks() {
      this.$http.get('/article/rank').then((response) => {
        this.rank = response.data
      }).catch((error) => {
        this.loading = false
        console.log(error)
      })
    },
    load() {
      if (this.noMore || this.errorMsg) {
        return
      }
      this.loading = true
      this.page++

      var keyword = this.$route.params.keyword
      if (keyword) {
        keyword = this.unscape(keyword)
      }

      this.$http.get('/article/list', {params: {"page": this.page, "size": this.size, "keyword": keyword}}).then((response) => {
        this.loading = false
        if (response.data.length === 0) {
          this.noMore = true
          return
        }
        response.data.forEach(v => {
          this.article.push(v)
        })
      }).catch((error) => {
        this.loading = false
        this.errorMsg = true
      })
    }
  }
}
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