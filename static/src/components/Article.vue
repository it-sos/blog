<template>
  <el-container>
    <el-main>
      <el-row class="nav">
        <el-col :span="12"><el-link :href="prev_title_link">上一篇：{{ prev_title }}</el-link></el-col>
        <el-col :span="12" align="right"><el-link :href="next_title_link">下一篇：{{ next_title }}</el-link></el-col>
      </el-row>
      <div class="box">
        <div class="title">
          <el-link :href="'/a/'+title"><h2>{{ title }}</h2></el-link>
          <span>{{ duration }}</span>
        </div>
        <div class="description" id="showHtml" v-html="article_content"></div>
        <el-row class="link">
          <el-col :span="12">
            <p v-for="topic in topics">来自：{{ topic }} 专题</p>
          </el-col>
          <el-col :span="12" align="right" class="tag">
            <el-tag v-for="tag in tags">{{ tag }}</el-tag>
          </el-col>
        </el-row>
      </div>
    </el-main>
  </el-container>
  <el-aside class="hidden-xs-only">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>xxoo 专题</span>
        </div>
      </template>
      <div v-for="o in 4" :key="o" class="text item">
        <el-link href="/a/列表">{{'列表内容这篇文章，内容很多，非常多，躲到爆炸 （' + o + '）' }}</el-link>
      </div>
    </el-card>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>xxoo 专题</span>
        </div>
      </template>
      <div v-for="o in 4" :key="o" class="text item">
        <el-link href="/a/列表">{{'列表内容这篇文章，内容很多，非常多，躲到爆炸 （' + o + '）' }}</el-link>
      </div>
    </el-card>
  </el-aside>
</template>

<script>

export default {
  name: "Article",
  created() {
    // 更多 UEditor 配置，参考 http://fex.baidu.com/ueditor/#start-config
    this.editorConfig = {
      UEDITOR_HOME_URL: '/ueditor/', // 访问 UEditor 静态资源的根路径，可参考常见问题1
      serverUrl: '//demo.com/cos', // 服务端接口（这个地址是我为了方便各位体验文件上传功能搭建的临时接口，请勿在生产环境使用！！！）
    };
  },
  data() {
    return {
      prev_title: "",
      next_title: "",
      prev_title_link: "",
      next_title_link: "",
      title: "",
      duration: "",
      article_content: "",
      topics: [],
      tags: [],
      article: {},
    }
  },
  mounted() {
    this.content()
  },
  methods: {
    content() {
      this.$http.get('/article/content', {params: {title: this.$route.params.title}}).then((response) => {
        this.article = response.data
        console.log(this.article.article)
        this.prev_title = "已经是顶部了"
        this.next_title = "已经是底部了"
        this.prev_title_link = "javascript:void(0);"
        this.next_title_link = "javascript:void(0);"
        if (this.article.navigations.prev_title) {
          this.prev_title = this.article.navigations.prev_title
          this.prev_title_link = '/a/' + this.prev_title
        }
        if (this.article.navigations.next_title) {
          this.next_title = this.article.navigations.next_title
          this.next_title_link = '/a/' + this.next_title
        }
        this.article_content = this.article.article_content.data
        this.topics = this.article.article.topics
        this.tags = this.article.article.tags
        this.title = this.article.article.article.title
        this.duration = this.article.article.duration
      }).catch((error) => {
        console.log(error)
      })
    },
  }
}
</script>

<style scoped>
.nav {
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  padding: 15px 20px 20px 20px;
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
  word-break: break-all;
  word-wrap: break-word;
  padding-top:30px;
  min-height:300px;
}
.link {
  padding-top: 20px;
}
.link p {
  margin:0;
}
.tag span {
  margin-left: 6px;

}
.box {
  padding-top:30px;
  padding-bottom: 30px;
}
.box-card {
  margin-bottom: 30px;
}
.hidden-xs-only {
  padding-top:20px;
}
</style>
