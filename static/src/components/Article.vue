<template>
  <el-container>
    <el-main>
      <el-row class="nav">
        <el-col :span="12"><el-link href="#">上一篇：xxoo</el-link></el-col>
        <el-col :span="12" align="right"><el-link href="#">下一篇：xxoo</el-link></el-col>
      </el-row>
      <div class="box">
        <div class="title">
          <el-link href="/a/java笔记"><h2>java笔记</h2></el-link>
          <span style="vertical-align: bottom;">3天前</span>
        </div>
        <div class="description">xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文</div>
        <el-row class="link">
          <el-col :span="12">
            <p>来自：xxoo 专题</p>
            <p>来自：xxoo 专题</p>
          </el-col>
          <el-col :span="12" align="right" class="tag">
            <el-tag>标签一</el-tag>
            <el-tag>标签二</el-tag>
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
  name: "Article"
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
  margin-left: 30px;
}
.description {
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
</style>
