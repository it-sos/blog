<template>
  <div>
    <template v-for="(item, index) in items">
      <div class="divider" v-if="item.type === 'divider'" :key="index"/>
      <menu-item :key="item.title" v-else v-bind="item"/>
    </template>
    <div class="menu-opt">
      <el-switch
          class="menu-opt-item"
          v-model="publish"
          active-text="公开"
          :loading="publishLoading"
          :beforeChange="publishBeforeChange"></el-switch>
      <el-switch
          class="menu-opt-item"
          v-model="encrypt"
          active-text="加密"
          :loading="encryptLoading"
          :beforeChange="encryptBeforeChange"></el-switch>
      <el-select
          class="menu-opt-item"
          v-model="topicValue"
          multiple
          allow-create
          filterable
          default-first-option
          size="mini"
          placeholder="选择专题">
        <el-option
            v-for="item in topicOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
      </el-select>
      <el-select
          class="menu-opt-item"
          v-model="tagValue"
          multiple
          allow-create
          filterable
          default-first-option
          size="mini"
          placeholder="选择标签">
        <el-option
            v-for="item in tagOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            v-right-click="rightMenu(1, 'tag')">
        </el-option>
      </el-select>
    </div>
  </div>
  <right-menu ref="rightMenuRefs" @trigger="trigger" />
</template>

<script lang="ts">
import '@fortawesome/fontawesome-free/css/all.css'
import MenuItem from './MenuItem.vue'
import {defineComponent, reactive, ref, toRefs} from 'vue'
import {ElMessage} from 'element-plus'
import {Position, default as RightMenu} from "@/components/RightMenu.vue";

export default defineComponent({
  components: {
    MenuItem,
    RightMenu,
  },

  props: {
    editor: {
      type: Object,
      required: true,
    },
  },

  directives: {
    rightClick: {
      mounted(el, binding) {
        el.addEventListener("contextmenu", function (ev: any) {
          ev.preventDefault();
          const position: Position = {
            x: ev.x,
            y: ev.y,
          }
          binding.value(position)
          return false;
        }, false)
      }
    }
  },

  setup() {

    let triggerState = {
      id: 0,
      type: ""
    }
    const rightMenuRefs = ref()
    const rightMenu = (id: number, type: string) => {
      return (position: Position) => {
        triggerState = {id, type}
        rightMenuRefs.value.showMenu()
        rightMenuRefs.value.positionMenu(position)
        document.onclick = () => {
          rightMenuRefs.value.hideMenu()
        }
      }
    }
    const trigger = (type: string) => {
      if (type == "delete") {
        open()
      }
      console.log(type, triggerState)
    }

    const publishStatus = reactive({
      publish: false,
      publishLoading: false,
    })

    const encryptStatus = reactive({
      encrypt: false,
      encryptLoading: false,
    })

    const publishBeforeChange = () => {
      publishStatus.publishLoading = true
      return new Promise(resolve => {
        setTimeout(() => {
          publishStatus.publishLoading = false
          ElMessage.success('切换成功')
          return resolve(true)
        }, 1000)
      })
    }

    const encryptBeforeChange = () => {
      encryptStatus.encryptLoading = true
      return new Promise(resolve => {
        setTimeout(() => {
          encryptStatus.encryptLoading = false
          ElMessage.success('切换成功')
          return resolve(true)
        }, 1000)
      })
    }

    const selectStatus = reactive({
      topicValue: ['选项1'],
      topicOptions: [{
        value: '选项1',
        label: 'mac 下以 root 角色开机启动执行脚本或命令'
      }, {
        value: '选项2',
        label: '双皮奶'
      }, {
        value: '选项3',
        label: '蚵仔煎'
      }, {
        value: '选项4',
        label: '龙须面'
      }, {
        value: '选项5',
        label: '北京烤鸭'
      }],
      tagValue: false,
      tagOptions: [{lable:"php", value:"php"}],
    })

    return {
      ...toRefs(publishStatus),
      ...toRefs(encryptStatus),
      ...toRefs(selectStatus),
      publishBeforeChange,
      encryptBeforeChange,
      rightMenu,
      triggerState,
      trigger,
      rightMenuRefs,
    }

  },

  data() {
    return {
      items: [
        {
          icon: 'fas fa-bold',
          title: 'Bold',
          action: () => this.editor.chain().focus().toggleBold().run(),
          isActive: () => this.editor.isActive('bold'),
        },
        {
          icon: 'fas fa-italic',
          title: 'Italic',
          action: () => this.editor.chain().focus().toggleItalic().run(),
          isActive: () => this.editor.isActive('italic'),
        },
        {
          icon: 'fas fa-strikethrough',
          title: 'Strike',
          action: () => this.editor.chain().focus().toggleStrike().run(),
          isActive: () => this.editor.isActive('strike'),
        },
        {
          icon: 'fas fa-underline',
          title: 'Underline',
          action: () => this.editor.chain().focus().toggleUnderline().run(),
          isActive: () => this.editor.isActive('underline'),
        },
        {
          icon: 'fas fa-subscript',
          title: 'Subscript',
          action: () => this.editor.chain().focus().toggleSubscript().run(),
          isActive: () => this.editor.isActive('subscript'),
        },
        {
          icon: 'fas fa-superscript',
          title: 'Superscript',
          action: () => this.editor.chain().focus().toggleSuperscript().run(),
          isActive: () => this.editor.isActive('superscript'),
        },
        {
          icon: 'fas fa-lightbulb',
          title: 'Highlight',
          action: () => this.editor.chain().focus().toggleHighlight().run(),
          isActive: () => this.editor.isActive('highlight'),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-align-left',
          title: 'left',
          action: () => this.editor.chain().focus().setTextAlign('left').run(),
          isActive: () => this.editor.isActive({textAlign: 'left'}),
        },
        {
          icon: 'fas fa-align-center',
          title: 'left',
          action: () => this.editor.chain().focus().setTextAlign('center').run(),
          isActive: () => this.editor.isActive({textAlign: 'center'}),
        },
        {
          icon: 'fas fa-align-right',
          title: 'right',
          action: () => this.editor.chain().focus().setTextAlign('right').run(),
          isActive: () => this.editor.isActive({textAlign: 'right'}),
        },
        {
          icon: 'fas fa-align-justify',
          title: 'justify',
          action: () => this.editor.chain().focus().setTextAlign('justify').run(),
          isActive: () => this.editor.isActive({textAlign: 'justify'}),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-heading',
          title: 'Heading',
          action: () => this.editor.chain().focus().toggleHeading({level: 2}).run(),
          isActive: () => this.editor.isActive('heading', {level: 2}),
        },
        {
          icon: 'fas fa-paragraph',
          title: 'Paragraph',
          action: () => this.editor.chain().focus().setParagraph().run(),
          isActive: () => this.editor.isActive('paragraph'),
        },
        {
          icon: 'fas fa-list-ul',
          title: 'Bullet List',
          action: () => this.editor.chain().focus().toggleBulletList().run(),
          isActive: () => this.editor.isActive('bulletList'),
        },
        {
          icon: 'fas fa-list-ol',
          title: 'Ordered List',
          action: () => this.editor.chain().focus().toggleOrderedList().run(),
          isActive: () => this.editor.isActive('orderedList'),
        },
        {
          icon: 'fas fa-tasks',
          title: 'Task List',
          action: () => this.editor.chain().focus().toggleTaskList().run(),
          isActive: () => this.editor.isActive('taskList'),
        },
        {
          icon: 'fas fa-terminal',
          title: 'Code Block',
          action: () => this.editor.chain().focus().toggleCodeBlock().run(),
          isActive: () => this.editor.isActive('codeBlock'),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-quote-right',
          title: 'Blockquote',
          action: () => this.editor.chain().focus().toggleBlockquote().run(),
          isActive: () => this.editor.isActive('blockquote'),
        },
        {
          icon: 'fas fa-minus',
          title: 'Horizontal Rule',
          action: () => this.editor.chain().focus().setHorizontalRule().run(),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-level-down-alt',
          title: 'Hard Break',
          action: () => this.editor.chain().focus().setHardBreak().run(),
        },
        {
          icon: 'fas fa-eraser',
          title: 'Clear Format',
          action: () => this.editor.chain()
              .focus()
              .clearNodes()
              .unsetAllMarks()
              .run(),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-undo',
          title: 'Undo',
          action: () => this.editor.chain().focus().undo().run(),
        },
        {
          icon: 'fas fa-redo',
          title: 'Redo',
          action: () => this.editor.chain().focus().redo().run(),
        },
      ],
    }
  },
})
</script>

<style lang="scss" scoped>
.divider {
  width: 1px;
  margin-left: 1rem;
  display: inline-block;
}

.menu-opt {
  padding-top: 0.5rem;
  line-height: 2rem;
  .menu-opt-item {
    margin-right: 12px;
  }
}
</style>
