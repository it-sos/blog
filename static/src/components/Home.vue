<template>
  <el-container>
    <el-main>
      <div>
        <div class="box" v-infinite-scroll="load" infinite-scroll-disabled="disabled" v-for="i in count">
          <div class="title">
            <el-link href="/a/java笔记"><h2>java笔记</h2></el-link>
            <span style="vertical-align: bottom;">3天前</span>
          </div>
          <div class="description">xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文</div>
          <el-row class="link">
            <el-col :span="12">
              <p>来自：xxoo 专题</p>
            </el-col>
            <el-col :span="12" align="right" class="tag">
              <el-tag>标签一</el-tag>
              <el-tag>标签二</el-tag>
            </el-col>
          </el-row>
          <el-divider></el-divider>
        </div>
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
      <div v-for="o in 4" :key="o" class="text item">
        <el-link href="/a/列表">{{'列表内容这篇文章，内容很多，非常多，躲到爆炸 ' + o }}</el-link>
      </div>
    </el-card>
  </el-aside>
</template>

<script>
export default {
  name: "Home",
  data () {
    return {
      count: 10,
      loading: false
    }
  },
  computed: {
    noMore () {
      return this.count >= 20
    },
    disabled () {
      return this.loading || this.noMore
    }
  },
  methods: {
    load () {
      this.loading = true
      setTimeout(() => {
        this.count += 5
        this.loading = false
      }, 2000)
    }
  }
}
</script>

<style scoped>
.load {
  padding:0px;
  margin:0px;
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
  margin-left: 30px;
}
.description {
  padding-top:30px;
  color: #606266;
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
  padding-bottom: 30px;
}
</style>