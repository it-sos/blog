<template>
  <div class="editor" v-if="editor">
    <menu-bar class="editor__header" :editor="editor" />
    <editor-content :editor="editor" />
  </div>
</template>

<script>
import { Editor, EditorContent } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import Highlight from '@tiptap/extension-highlight'
import Typography from '@tiptap/extension-typography'
import MenuBar from './editor/MenuBar.vue'

export default {
  components: {
    EditorContent,
    MenuBar,
  },

  data() {
    return {
      editor: null,
    }
  },

  mounted() {
    this.editor = new Editor({
      content: '<p>I’m running tiptap with Vue.js. 🎉</p>',
      extensions: [
        StarterKit,
        Highlight,
        Typography
      ],
    })
  },

  beforeUnmount() {
    this.editor.destroy()
  },

}
</script>
<style lang="scss">
/* Basic editor styles */
.ProseMirror {
> * + * {
  margin-top: 0.75em;
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
}

img {
  max-width: 100%;
  height: auto;
}

hr {
  margin: 1rem 0;
}

blockquote {
  padding-left: 1rem;
  border-left: 2px solid rgba(#0D0D0D, 0.1);
}
}
</style>

