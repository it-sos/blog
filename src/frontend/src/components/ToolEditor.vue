<template>
  <div class="editor" v-if="editor">
    <el-card class="box-card">
      <template #header>
        <div class="card-header menu-bar-user">
          <menu-bar @save="save" :editor="editor"/>
          <el-tooltip class="item" effect="dark" :content="saveStatus.message" placement="left">
            <el-icon :class="saveStatus.icon" :color="saveStatus.color" :size="26">
              <loading v-if="saveStatus.unsaved"/>
            </el-icon>
          </el-tooltip>
        </div>
      </template>
      <editor-content :editor="editor"/>
    </el-card>
  </div>
</template>

<script lang="ts">
import {Editor, EditorContent} from '@tiptap/vue-3'
import MenuBar from '@/components/editor/MenuBar.vue'
import {defineComponent, onMounted, onUnmounted, provide, reactive, ref, toRefs} from 'vue'
import {ElMessage} from 'element-plus'
import axios from "axios";
import {router} from '@/routes'
import {Loading} from '@element-plus/icons'
import utils from '@/common/utils'
import extensions from "@/common/tiptap_extensions";

export default defineComponent({

  components: {
    Loading,
    MenuBar,
    EditorContent,
  },

  setup(prop, context) {

    const switchStatus = reactive({
      publish: ref(false),
      encrypt: ref(false),
    })
    provide('switch-status', switchStatus)

    const selectValue = reactive({
      topic: [],
      tag: [],
    })
    provide('select-value', selectValue)

    const state = reactive({
      saveStatus: {
        icon: "el-icon-success",
        message: "已自动保存",
        color: "#67C23A",
        unsaved: false,
      },
      id: 0,
    });

    const stateUnsaved = () => {
      state.saveStatus = {
        icon: "is-loading",
        message: "停止编辑 5s 后将自动保存",
        color: "#67C23A",
        unsaved: true,
      }
    };

    const stateSaved = () => {
      state.saveStatus = {
        icon: "el-icon-success",
        message: "已自动保存",
        color: "#67C23A",
        unsaved: false,
      }
    };

    const stateSaveFail = (message?: string, type: 'success' | 'warning' | 'info' | 'error' | '' = "error", duration: number = 3000) => {
      state.saveStatus = {
        icon: "el-icon-error",
        message: "保存失败." + message,
        color: "#F56C6C",
        unsaved: false,
      }

      ElMessage({
        duration: duration,
        showClose: true,
        message: message,
        type: type,
      });
    };

    // 文章加密状态
    // eslint-disable-next-line
    const enum SWITCH_ENCRYPT_STATUS {
      // 密文
      // eslint-disable-next-line
      Encrypt = 1,
      // 明文
      // eslint-disable-next-line
      Plaintext = 2,
    }

    // 文章发布状态
    // eslint-disable-next-line
    const enum SWITCH_PUBLISH_STATUS {
      // 私有
      // eslint-disable-next-line
      Private = 1,
      // 公开
      // eslint-disable-next-line
      Public = 2,
    }


    const save = () => {
      stateUnsaved()
      let json = editor.getJSON()
      if (json.content[0].type != 'heading' || json.content[0].attrs.level != 2) {
        stateSaveFail("请设置h2开头的标题名称，可通过点击【H】图标进行设置！", 'warning')
        return
      }
      let title: string = json.content[0].content[0].text
      let contents: string = editor.getHTML()
      let content: string = contents.substring(contents.indexOf("</h2>") + 5)
      let intro: string = ""
      if (typeof json.content[1].content != "undefined" &&
          json.content[1].content[0].text != "") {
        intro = json.content[1].content[0].text
      }
      if (intro.length > 255) {
        stateSaveFail(`第二行简介超过最大长度255，当前长度：${intro.length}！`, 'warning')
        return
      }
      let id: number = utils.getArticleId()
      axios('/admin/article', {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        },
        method: "post",
        responseType: "json",
        data: {
          "id": id,
          "title": title,
          "intro": intro,
          "content": content,
          "is_encrypt": switchStatus.encrypt ? SWITCH_ENCRYPT_STATUS.Encrypt : SWITCH_ENCRYPT_STATUS.Plaintext,
          "is_state": switchStatus.publish ? SWITCH_PUBLISH_STATUS.Public : SWITCH_PUBLISH_STATUS.Private,
        }
      }).then((response: any) => {
        if (id == 0) {
          router.push('/e/' + response.data)
          context.emit("syncArticleList", 'add', response.data, title)
        } else {
          context.emit("syncArticleList", 'update', response.data, title)
        }
        stateSaved()
      }).catch((error: any) => {
        stateSaveFail(error.response.data.message)
      })
    }

    const loadArticle = () => {
      let id: number = utils.getArticleId()
      if (id > 0) {
        axios('/admin/article', {
          method: "get",
          responseType: "json",
          params: {
            "id": id,
          }
        }).then((response: any) => {
          editor.commands.setContent(`<h2>${response.data.title}</h2>\n${response.data.content}`)
          switchStatus.publish = response.data.is_state == SWITCH_PUBLISH_STATUS.Public
          switchStatus.encrypt = response.data.is_encrypt == SWITCH_ENCRYPT_STATUS.Encrypt
          selectValue.topic = response.data.topics
          selectValue.tag = response.data.tags
        }).catch((error: any) => {
          stateSaveFail(error.response.data.message)
        })
      } else {
        editor.commands.setContent('<h2>此行为标题，固定样式为H2</h2>\r<p>此行为简介，可自定义文本格式</p>')
        switchStatus.publish = false
        switchStatus.encrypt = false
        selectValue.topic = []
        selectValue.tag = []
      }
    }

    onMounted(() => {
      loadArticle()
    })

    let timer: any = null;
    let editor: any = new Editor({
      injectCSS: true,
      content: '<h2>此行为标题，固定样式为H2</h2>\r<p>此行为简介，可自定义文本格式</p>',
      extensions: extensions,
      onUpdate() {
        stateUnsaved()
        clearTimeout(timer)
        timer = setTimeout(() => {
          save()
        }, 2000)
      },
    })

    onUnmounted(() => {
      if (editor) editor.destroy();
    })

    return {
      ...toRefs(state),
      editor,
      loadArticle,
      save,
    }
  },
})
</script>
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

