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

// import {Editor, EditorContent, Extensions, VueNodeViewRenderer} from '@tiptap/vue-3'
// import Document from '@tiptap/extension-document'
// import StarterKit from '@tiptap/starter-kit'
// import Highlight from '@tiptap/extension-highlight'
// import Paragraph from '@tiptap/extension-paragraph'
// import Text from '@tiptap/extension-text'
// import Typography from '@tiptap/extension-typography'
// import TaskList from '@tiptap/extension-task-list'
// import TaskItem from '@tiptap/extension-task-item'
// import TextAlign from '@tiptap/extension-text-align'
// import CodeBlockLowlight from '@tiptap/extension-code-block-lowlight'
// import Underline from '@tiptap/extension-underline'
// import Subscript from '@tiptap/extension-subscript'
// import Superscript from '@tiptap/extension-superscript'
// import CodeBlockComponent from './editor/CodeBlockComponent.vue'

import {defineComponent, reactive, ref, toRefs} from "vue";
import axios from "axios";
import {router} from "@/routes";

export default defineComponent({

  setup() {
    let state = reactive({
      prev_title: ref<string>(),
      next_title: ref<string>(),
      prev_title_link: ref<string>(),
      next_title_link: ref<string>(),
      title: ref<string>(),
      duration: ref<string>(),
      article_content: ref<string>(),
      topics: ref<any[]>(),
      tags: ref<any[]>(),
      tagList: ref<any[]>(),
      topicList: ref<any[]>(),
    })

    document.title = "详情：" + decodeURIComponent(router.currentRoute.value.params.title.toString())

    axios.get('/article/content', {params: {title: decodeURIComponent(router.currentRoute.value.params.title.toString())}}).then((response) => {
      let article = response.data
      state.prev_title = "已经是顶部了"
      state.next_title = "已经是底部了"
      state.prev_title_link = "javascript:void(0);"
      state.next_title_link = "javascript:void(0);"
      if (article.navigations.prev_title) {
        state.prev_title = article.navigations.prev_title
        state.prev_title_link = '/a/' + encodeURIComponent(article.navigations.prev_title)
      }
      if (article.navigations.next_title) {
        state.next_title = article.navigations.next_title
        state.next_title_link = '/a/' + encodeURIComponent(article.navigations.next_title)
      }
      state.article_content = article.article_content.data
      state.topics = article.article.topics
      state.tags = article.article.tags
      state.title = article.article.article.title
      state.duration = article.article.duration
      state.topicList = article.topics;
      state.tagList = article.tags;
    }).catch((error) => {
      console.log(error)
    })
    return {
      ...toRefs(state)
    }
  }
})
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
