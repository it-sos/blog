<template>
  <div class="editor" v-if="editor">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <menu-bar class="" :editor="editor"/>
        </div>
      </template>
      <editor-content :editor="editor" style="min-height: 436px;"/>
    </el-card>
  </div>
</template>

<script lang="ts">
import {Editor, EditorContent, VueNodeViewRenderer} from '@tiptap/vue-3'
import Document from '@tiptap/extension-document'
import StarterKit from '@tiptap/starter-kit'
import Highlight from '@tiptap/extension-highlight'
import Paragraph from '@tiptap/extension-paragraph'
import Text from '@tiptap/extension-text'
import Typography from '@tiptap/extension-typography'
import TaskList from '@tiptap/extension-task-list'
import TaskItem from '@tiptap/extension-task-item'
import TextAlign from '@tiptap/extension-text-align'
import CodeBlockLowlight from '@tiptap/extension-code-block-lowlight'
import CodeBlockComponent from './editor/CodeBlockComponent.vue'
import MenuBar from './editor/MenuBar.vue'

import lowlight from 'lowlight'

const CodeBlock = CodeBlockLowlight
    .extend({
      addNodeView() {
        return VueNodeViewRenderer(CodeBlockComponent)
      },
    })
    .configure({lowlight})

import {defineComponent} from 'vue'

export default defineComponent({

  components: {
    MenuBar,
    EditorContent,
  },

  setup() {
    let timer: any = null;
    let editor: Editor = new Editor({
      injectCSS: false,
      content: '<h3>首行为标题</h3>',
      extensions: [
        Document,
        Paragraph,
        Text,
        CodeBlock,
        TaskList,
        TaskItem,
        StarterKit,
        TextAlign.configure({
          types: ['heading', 'paragraph'],
        }),
        Highlight,
        Typography,
      ],
      onUpdate({ editor }) {
        clearTimeout(timer)
        timer = setTimeout(() => {
          // console.log("set time out.")
        }, 5000)
        let json :Record<string, any> = editor.getJSON()
        if (json.content[0].type == 'heading') {
          console.log(json.content[0].content[0].text)
        }
        // console.log(json.content[0].attrs.level)
        // console.log(json.content[0].type)
        // console.log(json.content[0].content[0].text)
        // The content has changed.
      },
    })
    return {
      editor: editor
    }
  },

  beforeUnmount() {
    if (this.editor) this.editor.destroy();
  }
})
</script>
<style lang="scss">
/* Basic editor styles */

.ProseMirror {
  blockquote {
    padding-left: 1rem;
    border-left: 2px solid rgba(#0D0D0D, 0.1);
  }

  hr {
    border: none;
    border-top: 2px solid rgba(#0D0D0D, 0.1);
    margin: 2rem 0;
  }

  min-height: 430px;

  padding-left: 2px;

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
</style>

