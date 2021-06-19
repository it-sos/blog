<template>
  <el-container>
    <el-main>
      <div class="box">
        <div class="title">
          <a href="#"><h2 style="display: inline;">java笔记</h2></a>
          <span style="vertical-align: bottom;">3天前</span>
        </div>
        <div class="description">
          xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文
        </div>
        <el-row class="link">
          <el-col :span="12">
            <p><el-link href="#" target="_blank">来自：xxoo 专题</el-link></p>
          </el-col>
          <el-col :span="12" align="right" class="tag">
            <el-link href="#" type="primary" target="_blank"><el-tag>标签一</el-tag></el-link>
            <el-link href="#" type="primary" target="_blank"><el-tag>标签二</el-tag></el-link>
          </el-col>
        </el-row>
      </div>
      <div class="box">
        <div class="title">
          <a href="#"><h2 style="display: inline;">golang笔记 - 通道论据</h2></a>
          <span style="vertical-align: bottom;">3天前</span>
        </div>
        <div class="description">
          xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文
          xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文 xxooxoxo中文
        </div>
        <el-row class="link">
          <el-col :span="12">
            <p><el-link href="#" target="_blank">来自：xxoo 专题</el-link></p>
          </el-col>
          <el-col :span="12" align="right" class="tag">
            <el-link href="#" type="primary" target="_blank"><el-tag>标签一</el-tag></el-link>
          </el-col>
        </el-row>
      </div>
    </el-main>
  </el-container>
  <el-aside width="200px">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>访问前50文章</span>
        </div>
      </template>
      <div v-for="o in 4" :key="o" class="text item">
        {{'列表内容 ' + o }}
      </div>
    </el-card>
  </el-aside>
</template>

<script>
export default {
  name: "Home"
}
</script>

<style scoped>
.title a {
  text-decoration: none;
  color: #303133;
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
.tag a {
  margin-left: 6px;

}
.box {
  padding-bottom: 30px;
}
</style>