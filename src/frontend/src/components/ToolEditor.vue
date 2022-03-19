<template>
  <div class="editor" v-if="editor">
    <el-card class="box-card">
      <template #header>
        <div class="card-header menu-bar-user">
          <menu-bar @save="save" :editor="editor"/>
          <el-tooltip class="item" effect="dark" :content="saveStatus.message" placement="left">
            <el-icon :class="saveStatus.icon" :color="saveStatus.color" style="font-size:26px;">
              <loading v-if="saveStatus.unsaved"/>
              <circle-check-filled v-if="saveStatus.icon==='success'"/>
              <circle-close-filled v-if="saveStatus.icon==='error'"/>
            </el-icon>
          </el-tooltip>
        </div>
      </template>
      <bubble-menu :editor="editor" v-show="editor.isActive('image')" v-if="editor">
        <button @click="resizeImage(editor)">
          设置图像大小
        </button>
      </bubble-menu>
      <editor-content :editor="editor"/>
    </el-card>
  </div>
</template>

<script lang="ts">
import {BubbleMenu, Editor, EditorContent} from '@tiptap/vue-3'
import MenuBar from '@/components/editor/MenuBar.vue'
import {defineComponent, inject, onMounted, onUnmounted, provide, reactive, ref, toRefs, watch} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'
import utils from '../common/utils'
import {backendExtensions} from "../common/tiptap/tiptap-extensions";
import {CircleCheckFilled, CircleCloseFilled, Loading} from "@element-plus/icons-vue";
import {useRoute, useRouter} from "vue-router";
import {useStore} from "../store/store";

export default defineComponent({

  components: {
    Loading,
    CircleCheckFilled,
    CircleCloseFilled,
    MenuBar,
    EditorContent,
    BubbleMenu,
  },

  setup(prop, context) {

    const $axios: any = inject('$axios')
    const store = useStore()

    const switchStatus = reactive({
      publish: ref(false),
      encrypt: ref(false),
    })
    provide('switch-status', switchStatus)

    const selectValue = reactive({
      topic: [],
      tag: [],
      topicOld: [],
      tagOld: [],
    })
    provide('select-value', selectValue)

    const state = reactive({
      saveStatus: {
        icon: "success",
        message: "已自动保存",
        color: "#67C23A",
        unsaved: false,
      },
      id: 0,
    });

    const stateUnsaved = () => {
      store.commit('setSaved', false)
      state.saveStatus = {
        icon: "is-loading",
        message: "停止编辑 5s 后将自动保存",
        color: "#67C23A",
        unsaved: true,
      }
    };

    const stateSaved = () => {
      store.commit('setSaved', true)
      state.saveStatus = {
        icon: "success",
        message: "已自动保存",
        color: "#67C23A",
        unsaved: false,
      }
    };

    const stateSaveFail = (message?: string, type: 'success' | 'warning' | 'info' | 'error' | '' = "error", duration: number = 3000) => {
      state.saveStatus = {
        icon: "error",
        message: "保存失败." + message,
        color: "#F56C6C",
        unsaved: false,
      }

      ElMessage({
        duration: duration,
        showClose: true,
        message: message,
        // @ts-ignored
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

    const router = useRouter()
    const route = useRoute()

    const save = () => {
      stateUnsaved()
      let json = editor.getJSON()
      if (json.content[0].type != 'heading' || json.content[0].attrs.level != 2) {
        stateSaveFail("请设置h2开头的标题名称，可通过点击【H】图标进行设置！", 'warning')
        return
      }
      if (!json.content[0].hasOwnProperty("content")) {
        stateSaveFail("标题不能为空！", 'warning')
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
        stateSaveFail(`第二行是简介超过最大长度255，当前长度：${intro.length}！`, 'warning')
        return
      }
      let id: number = utils.getArticleId()
      let is_state = switchStatus.publish ? SWITCH_PUBLISH_STATUS.Public : SWITCH_PUBLISH_STATUS.Private
      $axios('/admin/article', {
        method: "post",
        responseType: "json",
        data: {
          "id": id,
          "title": title,
          "intro": intro,
          "content": content,
          "is_encrypt": switchStatus.encrypt ? SWITCH_ENCRYPT_STATUS.Encrypt : SWITCH_ENCRYPT_STATUS.Plaintext,
          "is_state": is_state,
        }
      }).then((response: any) => {
        store.commit("setArticleId", id || response.data)
        if (id == 0) {
          router.push('/e/' + response.data)
          context.emit("syncArticleList", 'add', response.data, title, is_state)
        } else {
          context.emit("syncArticleList", 'update', response.data, title, is_state)
        }
        stateSaved()
      }).catch((error: any) => {
        stateSaveFail(error.response.data.message)
      })
    }

    const loadArticle = () => {
      stateSaved()
      let id: number = utils.getArticleId()
      if (id > 0) {
        $axios('/admin/article', {
          method: "get",
          responseType: "json",
          params: {
            "id": id,
          }
        }).then((response: any) => {
          store.commit("setArticleId", id)
          editor.commands.setContent(`<h2>${response.data.title}</h2>\n${response.data.content}`)
          switchStatus.publish = response.data.is_state == SWITCH_PUBLISH_STATUS.Public
          switchStatus.encrypt = response.data.is_encrypt == SWITCH_ENCRYPT_STATUS.Encrypt
          selectValue.topic = selectValue.topicOld = response.data.topics
          selectValue.tag = selectValue.tagOld = response.data.tags
        }).catch((error: any) => {
          stateSaveFail(error.response.data.message)
        })
      } else {
        editor.commands.setContent('<h2>此行为标题，固定样式为H2</h2>\r<p>此行为简介，可自定义文本格式</p>')
        switchStatus.publish = false
        switchStatus.encrypt = false
        selectValue.topic = selectValue.topicOld = []
        selectValue.tag = selectValue.tagOld = []
      }
    }

    onMounted(() => {
      loadArticle()
    })

    watch(() => route.params.id, () => {
      loadArticle()
    })

    let timer: any = null;
    let editor: any = new Editor({
      injectCSS: true,
      autofocus: true,
      extensions: backendExtensions,
      onUpdate() {
        stateUnsaved()
        clearTimeout(timer)
        timer = setTimeout(() => {
          save()
        }, 2000)
      },
    })

    const resizeImage = (editors: any) => {
      if (!editors.isActive('image')) {
        return
      }
      let src = editors.getAttributes("image").src
      let size = ""
      if (src.indexOf('?') > -1) {
        let srcs = src.split('?')
        src = srcs[0]
        size = srcs[1].split('=')[1]
      }

      ElMessageBox.prompt('请输入图片尺寸，为空则为原图尺寸，若800x0以高度自适应', '输入尺寸', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputValue: size,
        inputPattern: /^(\d+x\d+)?$/,
        inputErrorMessage: '无效的尺寸格式，如800x800，800x0',
      }).then(({value}) => {
        if (value == null) {
          editors.commands.updateAttributes('image', {src: `${src}`})
        } else {
          editors.commands.updateAttributes('image', {src: `${src}?size=${value}`})
        }
      })
          .catch(() => {
            ElMessage({
              type: 'info',
              message: '取消',
            })
          })
    }

    // let selectionUpdate = (props: any) => {
    //   if (props.editor.isActive("image")) {
    //     resizeImage(props)
    //   }
    //   console.log(1)
    // }
    // editor.on('selectionUpdate', selectionUpdate)

    onUnmounted(() => {
      if (editor) editor.destroy();
    })


    let tippyOptions = {}

    return {
      tippyOptions,
      ...toRefs(state),
      editor,
      resizeImage,
      save,
    }
  },
})
</script>
<style lang="scss">
.global-drag-handle {
  position: absolute;

&
::after {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 1rem;
  height: 1.25rem;
  content: '⠿';
  font-weight: 700;
  cursor: grab;
  background: #0D0D0D10;
  color: #0D0D0D50;
  border-radius: 0.25rem;
}

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

&
.ProseMirror-selectednode {
  outline: 3px solid #68CEF8;
}

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

a {
  color: #68CEF8;
}

}
</style>

