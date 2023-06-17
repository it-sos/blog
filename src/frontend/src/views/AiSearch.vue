<template>
  <el-container>
    <el-main>
      <el-row
        type="flex"
        justify="space-between"
        align="middle"
      >
        <el-col :span="6">
        </el-col>
        <el-col :span="6" align="center">
          <el-alert v-for="(msg, msgIndex) in chatList" v-bind:key="msg+msgIndex" :title="msg" type="info" @close="removeMessage(msgIndex)" />
          <el-input
            v-model="chatMsg"
            placeholder="请输入内容"
            @keydown.enter="chatMessage()" 
            class="input-with-select"
            style="margin-top:30px;
          ">
            <template #append>
              <el-button :icon="Position" @click="chatMessage()" />
            </template>
          </el-input>
        </el-col>
        <el-col :span="6" align="left">
        </el-col>
      </el-row>
      <div class="box">
        <div class="title">
          <!-- <el-link :href="'/a/'+encodeURIComponent(title)"><h2>{{ title }}</h2></el-link> -->
          <el-link><h2>{{ title }}</h2></el-link>
        </div>
        <div class="description">
          <el-skeleton :rows="10" :loading="loading" animated>
            <template #default>
              <editor-content :editor="editor"/>
            </template>
          </el-skeleton>
        </div>
        <div style="padding-top: 60px;text-align: center;">
          <el-button style="width: 100px;" size="small" @click="back" type="info">返回</el-button>
          <el-button style="width: 100px;" size="small" @click="onlySave" type="primary">保存</el-button>
          <el-button style="width: 100px;" size="small" @click="saveEdit" type="warning">保存&编辑</el-button>
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<script lang="ts" setup>
import { Position } from '@element-plus/icons-vue';
import { Editor, EditorContent } from '@tiptap/vue-3';
import { ElMessage } from 'element-plus';
import 'element-plus/theme-chalk/display.css';
import MarkdownIt from 'markdown-it';
import { frontendExtensions } from "../common/tiptap/tiptap-extensions";
const chatMsg = ref<string>("")

let router = useRouter()
let store = useStore()
store.commit('setIsBackend', false)

const route = useRoute()
let loading = ref<boolean>(true)

let editor: any = new Editor({
  injectCSS: true,
  extensions: frontendExtensions,
  editable: false,
  content: 'test'
})

let chatList = computed(() => store.getters['chat/list'])

let chatMessage = () => {
  if (chatMsg.value != '') {
    // ElMessage({
    //   duration: 1000,
    //   showClose: false,
    //   message: '请输入内容',
    //   type: "error",
    // });
    // return;
    store.commit('chat/send', chatMsg.value)
    chatMsg.value = ''
  }
  searchAi()
}

let removeMessage = (msgIndex: number) => {
  store.commit('chat/remove', msgIndex)
}

// 返回
let back = () => {
  router.back()
}

// 只保存
let onlySave = () => {
  save(false)
}

// 保存&编辑
let saveEdit = () => {
  let id = store.getters.getArticleId()
  if (id > 0) {
    router.push('/e/' + id)
  } else {
    save(true)
  }
}
let titleEncode = (txt: string) => { return decodeURIComponent(txt) }
let title = ref<string>(titleEncode(route.params.keyword.toString()))
chatMsg.value = title.value

let searchAi = () => {
  console.log(chatList.value)
  let data = new FormData()
  for (let i in chatList.value) {
    data.append('keyword', chatList.value[i])
  }
  loading.value = true
  axios('/admin/chat/completion', {
    method: "post",
    responseType: "json",
    data: data
  }).then((response: any) => {
    store.commit('setArticleId', 0)
    const md = new MarkdownIt()
    let html = md.render(response.data)
    editor.commands.setContent(html)
    loading.value = false
  }).catch((error: any) => {
    console.error(error)
    // ElMessage({
    //   duration: 500,
    //   showClose: true,
    //   message: error.response.data.message,
    //   type: "error",
    // });
  })
}

watch(() => route.params.keyword, () => {
  if (route.params.keyword != null) {
    title.value = titleEncode(route.params.keyword.toString())
    chatMsg.value = title.value
    chatMessage()
  }
})

onMounted(() => {
  chatMessage()
})

onUnmounted(() => {
  store.commit('chat/reset')
})


let save = (is_edit: boolean) => {
  axios('/admin/article', {
    method: "post",
    responseType: "json",
    data: {
      "title": title.value,
      "intro": "",
      "content": editor.getHTML(),
      "is_encrypt": SWITCH_ENCRYPT_STATUS.Plaintext,
      "is_state": SWITCH_PUBLISH_STATUS.Private,
    }
  }).then((response: any) => {
    store.commit("setArticleId", response.data)
    if (is_edit) {
      router.push('/e/' + response.data)
    } else {
      ElMessage({
        duration: 500,
        showClose: true,
        message: '保存成功',
        type: "success",
      });
    }
  }).catch((error: any) => {
    ElMessage({
      duration: 500,
      showClose: true,
      message: error.response.data.message,
      type: "error",
    });
  })
}

</script>

<style scoped>
.el-alert {
  margin: 20px 0 0;
}
.el-alert:first-child {
  margin: 0;
}
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