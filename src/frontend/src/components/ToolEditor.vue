<template>
  <div class="editor" v-if="editor">
    <menu-bar class="" :editor="editor" />
    <editor-content :editor="editor" />
  </div>
</template>

<script lang="ts">
import {Editor, EditorContent} from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import Highlight from '@tiptap/extension-highlight'
import Typography from '@tiptap/extension-typography'
import MenuBar from './editor/MenuBar.vue'

import { Vue, Options } from 'vue-class-component';

@Options({
  components: {
    MenuBar,
    EditorContent,
  },
  inject: []
})
export default class ToolEditor extends Vue {

  editor: Editor | null = null;

  mounted () {
    this.editor = new Editor({
      content: '<p>I’m running tiptap with Vue.js. 🎉</p>',
      extensions: [
        StarterKit,
        Highlight,
        Typography
      ],
    })
  }

  beforeDestroy () {
    if (this.editor) this.editor.destroy();
  }
}
</script>
<style lang="scss">
</style>

