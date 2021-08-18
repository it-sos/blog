<template>
  <el-container>
    <el-main>
      <el-row class="nav">
        <el-col :span="12">
          <el-link :href="prev_title_link">上一篇：{{ prev_title }}</el-link>
        </el-col>
        <el-col :span="12" align="right">
          <el-link :href="next_title_link">下一篇：{{ next_title }}</el-link>
        </el-col>
      </el-row>
      <div class="box">
        <div class="title">
          <el-link :href="'/a/'+encodeURIComponent(title)"><h2>{{ title }}</h2></el-link>
          <el-tag effect="plain" type="danger" size="small">{{ duration }}</el-tag>
        </div>
        <div class="description" id="showHtml" v-html="article_content"></div>
        <el-row class="link">
          <el-col :span="12">
            <p v-bind:key="idx" v-for="(topic,idx) in topics">来自：{{ topic }} 专题</p>
          </el-col>
          <el-col :span="12" align="right" class="tag">
            <el-tag effect="plain" v-bind:key="idx" v-for="(tag, idx) in tags">{{ tag }}</el-tag>
          </el-col>
        </el-row>
      </div>
    </el-main>
  </el-container>
  <el-aside class="hidden-xs-only">
    <el-card class="box-card" v-bind:key="idx" v-for="(topic,idx) in topicList">
      <template #header>
        <div class="card-header">
          <span>{{ topic.title }} 专题</span>
        </div>
      </template>
      <div v-bind:key="idx" v-for="(art, idx) in topic.article" class="text item">
        <el-link :href="'/a/'+encodeURIComponent(art.title)">{{ art.title + ' （' + art.access_times + '）' }}</el-link>
      </div>
    </el-card>
    <el-card class="box-card" v-bind:key="idx" v-for="(tag, idx) in tagList">
      <template #header>
        <div class="card-header">
          <span>{{ tag.title }} 标签</span>
        </div>
      </template>
      <div v-bind:key="idx" v-for="(art, idx) in tag.article" class="text item">
        <el-link :href="'/a/'+encodeURIComponent(art.title)">{{ art.title + ' （' + art.access_times + '）' }}</el-link>
      </div>
    </el-card>
  </el-aside>
</template>

<script lang="ts">

import {Vue} from "vue-class-component";

export default class Article extends Vue {
  prev_title: string = ""
  next_title: string = ""
  prev_title_link: string = ""
  next_title_link: string = ""
  title: string = ""
  duration: string = ""
  article_content: string = ""
  topics: any[] = []
  tags: any[] = []
  article: any = {}
  tagList: any[] = []
  topicList: any[] = []

  mounted() {
    let title: string = this.$route.params.title.toString();
    document.title = "详情：" + decodeURIComponent(title)
    this.content()
  }

  content() {
    this.$http.get('/article/content', {params: {title: decodeURIComponent(this.$route.params.title.toString())}}).then((response) => {
      this.article = response.data
      this.prev_title = "已经是顶部了"
      this.next_title = "已经是底部了"
      this.prev_title_link = "javascript:void(0);"
      this.next_title_link = "javascript:void(0);"
      if (this.article.navigations.prev_title) {
        this.prev_title = this.article.navigations.prev_title
        this.prev_title_link = '/a/' + encodeURIComponent(this.prev_title)
      }
      if (this.article.navigations.next_title) {
        this.next_title = this.article.navigations.next_title
        this.next_title_link = '/a/' + encodeURIComponent(this.next_title)
      }
      this.article_content = this.article.article_content.data
      this.topics = this.article.article.topics
      this.tags = this.article.article.tags
      this.title = this.article.article.article.title
      this.duration = this.article.article.duration
      this.topicList = this.article.topics;
      this.tagList = this.article.tags;
    }).catch((error) => {
      console.log(error)
    })
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
  margin-right: 30px;
}

.title h2 {
  display: inline;
  font-size: larger;
  font-size: 24px;
}

.title span {
  vertical-align: bottom;
}

.description {
  word-break: break-all;
  word-wrap: break-word;
  padding-top: 30px;
  min-height: 300px;
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
  padding-top: 30px;
  padding-bottom: 30px;
}

.box-card {
  margin-bottom: 30px;
}

.hidden-xs-only {
  padding-top: 20px;
}
</style>
