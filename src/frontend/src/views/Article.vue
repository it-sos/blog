<template>
  <el-skeleton :rows="10" :loading="loading" animated>
    <template #default>
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
            <div class="description">
              <editor-content :editor="editor"/>
            </div>
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
            <el-link :href="'/a/'+encodeURIComponent(art.title)">{{
                art.title + ' （' + art.access_times + '）'
              }}
            </el-link>
          </div>
        </el-card>
        <el-card class="box-card" v-bind:key="idx" v-for="(tag, idx) in tagList">
          <template #header>
            <div class="card-header">
              <span>{{ tag.title }} 标签</span>
            </div>
          </template>
          <div v-bind:key="idx" v-for="(art, idx) in tag.article" class="text item">
            <el-link :href="'/a/'+encodeURIComponent(art.title)">{{
                art.title + ' （' + art.access_times + '）'
              }}
            </el-link>
          </div>
        </el-card>
      </el-aside>
    </template>
  </el-skeleton>
</template>

<script lang="ts">

import {Editor, EditorContent} from '@tiptap/vue-3'
import {defineComponent, inject, onUnmounted, reactive, ref, toRefs} from "vue";
import axios from "axios";
import router from "../routes";
import {frontendExtensions} from "../common/tiptap/tiptap-extensions";
import Document from "@tiptap/extension-document";
import Paragraph from "@tiptap/extension-paragraph";
import {Link} from "@tiptap/extension-link";
import Text from "@tiptap/extension-text";


export default defineComponent({
  components: {
    EditorContent,
  },
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
      loading: true,
    })

    let article_state = inject("article-id", {id: ref<number>()})

    document.title = "详情：" + decodeURIComponent(router.currentRoute.value.params.title.toString())

    let editor: any = new Editor({
      injectCSS: true,
      extensions: frontendExtensions,
      editable: false,
    })

    onUnmounted(() => {
      if (editor) editor.destroy();
    })
    axios.get('/article/content', {params: {title: decodeURIComponent(router.currentRoute.value.params.title.toString())}}).then((response) => {
      let article = response.data
      article_state.id = article.article.article.id
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
      editor.commands.setContent(article.article_content.data)
      state.topics = article.article.topics
      state.tags = article.article.tags
      state.title = article.article.article.title
      state.duration = article.article.duration
      state.topicList = article.topics
      state.tagList = article.tags
      state.loading = false
    }).catch((error) => {
      console.log(error)
    })
    return {
      ...toRefs(state),
      editor,
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

<style lang="scss">
/* Basic editor styles */

.ProseMirror {

  min-height: 335px;
  padding: 5px;

*, :after, :before {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
  box-shadow: none;
  outline: none;
}

> * + * {
  margin-top: 0.75em;
}

blockquote {
  padding-left: 1rem;
  border-left: 2px solid rgba(#0D0D0D, 0.1);
}

hr {
  border: none;
  border-top: 2px solid rgba(#0D0D0D, 0.1);
  margin: 2rem 0;
}

ul,
ol {
  padding: 0 1rem;
}

h1,
h2,
h3,
h4,
h5,
h6 {
  line-height: 1.1;
}

code {
  background-color: rgba(#616161, 0.1);
  color: #616161;
}

mark {
  background-color: #FAF594;
}

img {
  max-width: 100%;
  height: auto;
}

pre {
  background: #0D0D0D;
  color: #FFF;
  font-family: 'JetBrainsMono', monospace;
  padding: 0.75rem 1rem;
  border-radius: 0.5rem;

code {
  color: inherit;
  padding: 0;
  background: none;
  font-size: 0.8rem;
}

.hljs-comment,
.hljs-quote {
  color: #616161;
}

.hljs-variable,
.hljs-template-variable,
.hljs-attribute,
.hljs-tag,
.hljs-name,
.hljs-regexp,
.hljs-link,
.hljs-name,
.hljs-selector-id,
.hljs-selector-class {
  color: #F98181;
}

.hljs-number,
.hljs-meta,
.hljs-built_in,
.hljs-builtin-name,
.hljs-literal,
.hljs-type,
.hljs-params {
  color: #FBBC88;
}

.hljs-string,
.hljs-symbol,
.hljs-bullet {
  color: #B9F18D;
}

.hljs-title,
.hljs-section {
  color: #FAF594;
}

.hljs-keyword,
.hljs-selector-tag {
  color: #70CFF8;
}

.hljs-emphasis {
  font-style: italic;
}

.hljs-strong {
  font-weight: 700;
}

}

ul[data-type="taskList"] {
  list-style: none;
  padding: 0;

li {
  display: flex;
  align-items: center;

> label {
  flex: 0 0 auto;
  margin-right: 0.5rem;
}

}

input[type="checkbox"] {
  cursor: pointer;
}

}
}

.menu-bar-user {
  display: flex;
  justify-content: space-between;
}

/* Color swatches */
.color {
  white-space: nowrap;

&
::before {
  content: ' ';
  display: inline-block;
  width: 1em;
  height: 1em;
  border: 1px solid rgba(128, 128, 128, 0.3);
  vertical-align: middle;
  margin-right: 0.1em;
  margin-bottom: 0.15em;
  border-radius: 2px;
  background-color: var(--color);
}

}
</style>

