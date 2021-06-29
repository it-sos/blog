<template>
  <el-container v-infinite-scroll="load" class="infinite-list">
    <el-main>
      <div class="box" key="index" v-for="art in article">
        <div class="title">
          <el-link :href="'/a/'+art.article.title"><h2>{{ art.article.title }}</h2></el-link>
          <span>{{ art.duration }}</span>
        </div>
        <div class="description">{{ art.article.intro }}</div>
        <el-row class="link">
          <el-col :span="12">
            <p v-for="tag in art.tags">来自：{{ tag }} 专题</p>
          </el-col>
          <el-col :span="12" align="right" class="tag">
            <el-tag v-for="topic in art.topics">{{ topic }}</el-tag>
          </el-col>
        </el-row>
        <!-- 分割线 -->
        <el-divider></el-divider>
      </div>
      <p class="load" v-if="loading">加载中...</p>
      <el-empty v-if="noMore" description="没有更多了"></el-empty>
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
        <el-link :href="'/a/'+r.title">{{ r.title }}（{{ r.access_times }}）</el-link>
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
      loading: false
    }
  },
  mounted() {
    this.load()
    this.ranks()
  },
  computed: {
    disabled() {
      return this.loading || this.noMore
    }
  },
  methods: {
    ranks() {
      this.$http.get('/article/rank').then((response) => {
        this.rank = response.data
      }).catch((error) => {
        this.loading = false
        console.log(error)
      })
    },
    load() {
      if (this.noMore) {
        return
      }
      this.loading = true
      this.page++
      this.$http.get('/article/list', {params: {"page": this.page, "size": this.size}}).then((response) => {
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
        console.log(error)
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
}

.title h2 {
  display: inline;
}

.title span {
  vertical-align: bottom;
  margin-left: 30px;
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