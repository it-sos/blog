<template>
  <el-container>
    <el-aside width="200px" class="hidden-xs-only">
      <el-tabs type="border-card">
        <el-tab-pane label="文章">
          <template #header>
            <div class="card-header">
              <el-input
                  placeholder="请输入内容"
                  v-model="keyword"
                  class="input-with-select"
              >
                <template #prefix>
                  <i class="el-input__icon el-icon-search"></i>
                </template>
              </el-input>
            </div>
          </template>
          <el-dropdown trigger="contextmenu" @command="handleCommand">
            <div class="infinite-list-wrapper" style="height:400px;overflow-y:auto;overflow-x: hidden;"
                 v-infinite-scroll="load">
              <div v-bind:key="idx" v-for="(art, idx) in article" class="text item dirs">
                <el-link href="javascript:void(0);"><span
                    v-html="art.article.title_match ? art.article.title_match : art.article.title"></span></el-link>
                <el-tag style="margin-left:0.3rem;" effect="plain" type="danger" size="mini">{{ art.duration }}</el-tag>
              </div>
            </div>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item icon="el-icon-plus">黄金糕</el-dropdown-item>
                <el-dropdown-item icon="el-icon-circle-plus">狮子头</el-dropdown-item>
                <el-dropdown-item icon="el-icon-circle-plus-outline">螺蛳粉</el-dropdown-item>
                <el-dropdown-item icon="el-icon-check">双皮奶</el-dropdown-item>
                <el-dropdown-item icon="el-icon-circle-check">蚵仔煎</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </el-tab-pane>
        <el-tab-pane label="专题">配置管理</el-tab-pane>
        <el-tab-pane label="标签">角色管理</el-tab-pane>
      </el-tabs>
    </el-aside>
    <el-main>
      <tool-editor/>
    </el-main>
  </el-container>
</template>
<script lang="ts">

import ToolEditor from '../components/ToolEditor.vue'

import {defineComponent, reactive, toRefs, watch} from 'vue'
import {ElMessage} from "element-plus";

export default defineComponent({
  components: {
    ToolEditor
  },

  setup() {
    document.title = "editing"

    const article: any[] = []
    const state = reactive({
      keyword: "",
      page: 0,
      size: 10,
      noMore: false,
      article: article,
    });

    const defaults = () => {
      state.page = 0
      state.noMore = false
    }

    const handleCommand = (command: any) => {
      ElMessage(`click on item ${command}`);
    };

    return {
      ...toRefs(state),
      handleCommand,
      defaults,
    }
  },

  created() {
    let timer: any = null
    watch(
        () => this.keyword,
        () => {
          clearTimeout(timer)
          timer = setTimeout(() => {
            this.defaults()
            this.load()
          }, 1000)
        }
    )
  },

  methods: {
    rightMenu(): void {
      alert("right click");
    },
    load(): void {
      if (this.noMore) {
        return
      }
      this.page++
      this.$http.get('/article/list', {
        params: {
          "page": this.page,
          "size": this.size,
          // "keyword": this.keyword
        }
      }).then((response) => {
        if (response.data.length === 0) {
          this.noMore = true
          return
        }
        response.data.forEach((v: any) => {
          this.article.push(v)
        })
      }).catch((error) => {
        console.log(error)
      })
    }
  }
})
</script>

<style lang="scss" scoped>
.hidden-xs-only {
  padding-top: 20px;
}

.dirs {
  margin-bottom: 8px;
  display: flex;
  justify-content: space-between;

span {
  font-size: 12px;
}

}
</style>
